package group

import (
	"bitbucket.com/hyperboloide/horo/services/postgres"
	"time"
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
