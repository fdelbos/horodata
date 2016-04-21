package group

import (
	"time"

	"dev.hyperboloide.com/fred/horodata/services/postgres"
)

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
