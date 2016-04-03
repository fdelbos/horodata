package group

import (
	"database/sql"
	"fmt"
	"time"

	"dev.hyperboloide.com/fred/horodata/models/types/listing"
	"dev.hyperboloide.com/fred/horodata/services/postgres"
)

type Job struct {
	Id         int64      `json:"id"`
	Created    time.Time  `json:"created"`
	GroupId    int64      `json:"-"`
	TaskId     int64      `json:"task_id"`
	CustomerId int64      `json:"customer_id"`
	CreatorId  int64      `json:"creator_id"`
	Duration   int64      `json:"duration"`
	Comment    string     `json:"comment"`
	Updated    *time.Time `json:"updated,omitempty"`
	UpdaterId  *int64     `json:"updater_id,omitempty"`
}

func (j *Job) Scan(scanFn func(dest ...interface{}) error) error {
	return scanFn(
		&j.Id,
		&j.Created,
		&j.GroupId,
		&j.TaskId,
		&j.CustomerId,
		&j.CreatorId,
		&j.Duration,
		&j.Comment,
		&j.Updated,
		&j.UpdaterId)
}

func (j *Job) Update(updaterId int64) error {
	const query = `
	update jobs
	set
        task_id = $2,
        customer_id = $3,
        duration = $4,
        comment = $5,
        updated = now(),
        updater_id = $6
	where id = $1;`
	return postgres.Exec(
		query,
		j.Id,
		j.TaskId,
		j.CustomerId,
		j.Duration,
		j.Comment,
		updaterId)
}

func (j *Job) Remove() error {
	const query = `
    delete from jobs where id = $1`
	return postgres.Exec(query, j.Id)
}

func (g *Group) JobGet(id int64) (*Job, error) {
	j := &Job{}
	const query = `
	select * from jobs where group_id = $1 and id = $2`
	return j, postgres.QueryRow(j, query, g.Id, id)
}

func (g *Group) JobAdd(task, customer, creator, duration int64, comment string) error {
	const query = `
	insert into jobs
        (group_id, task_id, customer_id, creator_id, duration, comment)
	values
        ($1, $2, $3, $4, $5, $6);`
	return postgres.Exec(query, g.Id, task, customer, creator, duration, comment)
}

func (g *Group) JobRemove(id int64) error {
	const query = `
    delete from jobs where group_id = $1 and id = $2`
	return postgres.Exec(query, g.Id, id)
}

type JobApi struct {
	Id         int64     `json:"id"`
	Created    time.Time `json:"created"`
	TaskId     int64     `json:"task_id"`
	CustomerId int64     `json:"customer_id"`
	CreatorId  int64     `json:"creator_id"`
	Duration   int64     `json:"duration"`
	Comment    string    `json:"comment,omitempty"`
}

func (j *JobApi) Scan(scanFn func(dest ...interface{}) error) error {
	return scanFn(
		&j.Id,
		&j.Created,
		&j.TaskId,
		&j.CustomerId,
		&j.Duration,
		&j.Comment,
		&j.CreatorId)
}

func (g *Group) JobApiGet(id int64) (*JobApi, error) {
	j := &JobApi{}
	const query = `
    select
		id, created, task_id, customer_id, duration, comment, creator_id
    from
		jobs
    where
		group_id = $1 and id = $2;`
	return j, postgres.QueryRow(j, query, g.Id, id)
}

func (g *Group) JobApiList(begin, end time.Time, customer, creator *int64, request *listing.Request) (*listing.Result, error) {
	result := &listing.Result{}
	result.Offset = request.Offset

	query := jobApiGenQuery(customer, creator)
	var rows *sql.Rows
	var err error
	if customer != nil && creator != nil {
		rows, err = postgres.DB().Query(query, g.Id, begin, end, request.Size, request.Offset, *customer, *creator)
	} else if customer != nil {
		rows, err = postgres.DB().Query(query, g.Id, begin, end, request.Size, request.Offset, *customer)
	} else if creator != nil {
		rows, err = postgres.DB().Query(query, g.Id, begin, end, request.Size, request.Offset, *creator)
	} else {
		rows, err = postgres.DB().Query(query, g.Id, begin, end, request.Size, request.Offset)
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		job := &JobApi{}
		if err := job.Scan(rows.Scan); err != nil {
			return nil, err
		}
		result.Results = append(result.Results, job)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	result.Size = len(result.Results)

	queryCount := jobApiGenQueryCount(customer, creator)
	if customer != nil && creator != nil {
		err = postgres.DB().QueryRow(queryCount, g.Id, begin, end, *customer, creator).Scan(&result.Total)
	} else if customer != nil {
		err = postgres.DB().QueryRow(queryCount, g.Id, begin, end, *customer).Scan(&result.Total)
	} else if creator != nil {
		err = postgres.DB().QueryRow(queryCount, g.Id, begin, end, *creator).Scan(&result.Total)
	} else {
		err = postgres.DB().QueryRow(queryCount, g.Id, begin, end).Scan(&result.Total)
	}
	return result, err
}

func jobApiGenQuery(customer, creator *int64) string {
	const query = `
	select
		id, created, task_id, customer_id, duration, comment, creator_id
	from jobs
	where
			group_id = $1
		and created > $2
		and created < $3
		%s
	order by id desc
	limit $4 offset $5;`

	cond := ""
	if customer != nil && creator != nil {
		cond = "and customer_id = $6 and creator_id = $7"
	} else if customer != nil {
		cond = "and customer_id = $6"
	} else if creator != nil {
		cond = "and creator_id = $6"
	}
	return fmt.Sprintf(query, cond)
}

func jobApiGenQueryCount(customer, creator *int64) string {
	const query = `
		select count(id)
		from jobs
		where
				group_id = $1
			and created > $2
			and created < $3
			%s;`

	cond := ""
	if customer != nil && creator != nil {
		cond = "and customer_id = $4 and creator_id = $5"
	} else if customer != nil {
		cond = "and customer_id = $4"
	} else if creator != nil {
		cond = "and creator_id = $4"
	}
	return fmt.Sprintf(query, cond)
}
