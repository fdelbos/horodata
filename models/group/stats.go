package group

import (
	"bitbucket.com/hyperboloide/horo/services/postgres"
	"database/sql"
	"fmt"
	"time"
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
		rows, err = postgres.DB().Query(fmt.Sprintf(query, "and creator_id = $2"), g.Id, begin, end, *creator)
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
