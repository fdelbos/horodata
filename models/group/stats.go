package group

import (
	"database/sql"
	"fmt"
	"time"

	"dev.hyperboloide.com/fred/horodata/services/postgres"
)

type CustomerTime struct {
	CustomerId int64 `json:"customer_id"`
	Duration   int64 `json:"duration"`
}

func (ct *CustomerTime) Scan(scanFn func(dest ...interface{}) error) error {
	return scanFn(
		&ct.CustomerId,
		&ct.Duration)
}

func (g *Group) StatsCustomerTime(begin, end time.Time, creator *int64) ([]CustomerTime, error) {
	const query = `
    select customer_id, sum(duration)
    from jobs
    where
            group_id = $1
        and created > $2
    	and created < $3 %s
    group by customer_id`

	var rows *sql.Rows
	var err error
	if creator == nil {
		rows, err = postgres.DB().Query(fmt.Sprintf(query, ""), g.Id, begin, end)
	} else {
		rows, err = postgres.DB().Query(fmt.Sprintf(query, "and creator_id = $4"), g.Id, begin, end, *creator)
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	results := []CustomerTime{}
	for rows.Next() {
		i := CustomerTime{}
		if err := i.Scan(rows.Scan); err != nil {
			return nil, err
		}
		results = append(results, i)
	}
	return results, rows.Err()
}

type TaskTime struct {
	TaskId   int64 `json:"task_id"`
	Duration int64 `json:"duration"`
}

func (st *TaskTime) Scan(scanFn func(dest ...interface{}) error) error {
	return scanFn(
		&st.TaskId,
		&st.Duration)
}

func (g *Group) StatsTaskTime(begin, end time.Time, creator *int64) ([]TaskTime, error) {
	const query = `
    select task_id, sum(duration)
    from jobs
    where
            group_id = $1
        and created > $2
    	and created < $3 %s
    group by task_id`

	var rows *sql.Rows
	var err error
	if creator == nil {
		rows, err = postgres.DB().Query(fmt.Sprintf(query, ""), g.Id, begin, end)
	} else {
		rows, err = postgres.DB().Query(fmt.Sprintf(query, "and creator_id = $4"), g.Id, begin, end, *creator)
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	results := []TaskTime{}
	for rows.Next() {
		i := TaskTime{}
		if err := i.Scan(rows.Scan); err != nil {
			return nil, err
		}
		results = append(results, i)
	}
	return results, rows.Err()
}

type GuestTime struct {
	GuestId  int64 `json:"guest_id"`
	Duration int64 `json:"duration"`
}

func (st *GuestTime) Scan(scanFn func(dest ...interface{}) error) error {
	return scanFn(
		&st.GuestId,
		&st.Duration)
}

func (g *Group) StatsGuestTime(begin, end time.Time, creator *int64) ([]GuestTime, error) {
	const query = `
    select creator_id, sum(duration)
    from jobs
    where
            group_id = $1
        and created > $2
    	and created < $3
    group by creator_id`

	rows, err := postgres.DB().Query(query, g.Id, begin, end)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	results := []GuestTime{}
	for rows.Next() {
		i := GuestTime{}
		if err := i.Scan(rows.Scan); err != nil {
			return nil, err
		}
		results = append(results, i)
	}
	return results, rows.Err()
}
