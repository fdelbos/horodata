package group

import (
	"time"

	"dev.hyperboloide.com/fred/horodata/services/postgres"
)

type TaskTime struct {
	TaskId   int64 `json:"task_id"`
	Duration int64 `json:"duration"`
}

func (st *TaskTime) Scan(scanFn func(dest ...interface{}) error) error {
	return scanFn(
		&st.TaskId,
		&st.Duration)
}

func (g *Group) StatsTaskTime(begin, end time.Time) ([]TaskTime, error) {
	const query = `
    select task_id, sum(duration)
    from jobs
    where
            group_id = $1
        and created > $2
    	and created < $3
    group by task_id`

	if rows, err := postgres.DB().Query(query, g.Id, begin, end); err != nil {
		return nil, err
	} else {
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
}

type TaskCost struct {
	TaskId int64   `json:"task_id"`
	Cost   float64 `json:"cost"`
}

func (tc *TaskCost) Scan(scanFn func(dest ...interface{}) error) error {
	return scanFn(
		&tc.TaskId,
		&tc.Cost)
}

func (g *Group) StatsTaskCost(begin, end time.Time) ([]TaskCost, error) {
	const query = `
    select
        jobs.task_id,
        sum ((jobs.duration * guests.rate) / 360000::float)::numeric(10,2) as cost
    from
        jobs
        join guests on jobs.creator_id = guests.id
    where
            jobs.group_id = $1
        and jobs.created > $2
        and jobs.created < $3
    group by task_id
	order by task_id asc;`

	if rows, err := postgres.DB().Query(query, g.Id, begin, end); err != nil {
		return nil, err
	} else {
		defer rows.Close()
		results := []TaskCost{}
		for rows.Next() {
			i := TaskCost{}
			if err := i.Scan(rows.Scan); err != nil {
				return nil, err
			}
			results = append(results, i)
		}
		return results, rows.Err()
	}
}
