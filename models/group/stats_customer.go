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

type CustomerCost struct {
	CustomerId int64   `json:"customer_id"`
	Cost       float64 `json:"cost"`
}

func (cc *CustomerCost) Scan(scanFn func(dest ...interface{}) error) error {
	return scanFn(
		&cc.CustomerId,
		&cc.Cost)
}

func (g *Group) StatsCustomerCost(begin, end time.Time) ([]CustomerCost, error) {
	const query = `
    select
        jobs.customer_id,
        sum ((jobs.duration * guests.rate) / 360000::float)::numeric(10,2) as cost
    from
        jobs
        join guests on jobs.creator_id = guests.id
    where
            jobs.group_id = $1
        and jobs.created > $2
        and jobs.created < $3
    group by customer_id
	order by customer_id asc;`

	if rows, err := postgres.DB().Query(query, g.Id, begin, end); err != nil {
		return nil, err
	} else {
		defer rows.Close()
		results := []CustomerCost{}
		for rows.Next() {
			i := CustomerCost{}
			if err := i.Scan(rows.Scan); err != nil {
				return nil, err
			}
			results = append(results, i)
		}
		return results, rows.Err()
	}
}
