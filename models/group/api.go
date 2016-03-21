package group

import (
	"bitbucket.com/hyperboloide/horo/models/types/listing"
	"bitbucket.com/hyperboloide/horo/services/postgres"
	"time"
)

type GroupApi struct {
	Url     string    `json:"url"`
	OwnerId int64     `json:"owner"`
	Created time.Time `json:"created"`
	Name    string    `json:"name"`
}

func (ag *GroupApi) Scan(scanFn func(dest ...interface{}) error) error {
	return scanFn(
		&ag.Url,
		&ag.OwnerId,
		&ag.Created,
		&ag.Name)
}

func ApiByUser(user_id int64, request *listing.Request) (*listing.Result, error) {
	result := &listing.Result{}
	result.Offset = request.Offset

	const query = `
    select g.url, g.owner_id, g.created, g.name
    from groups g
    where
			g.active = true
		and g.id in (
				select distinct group_id
				from guests
				where user_id = $1 and active = true
		)
	order by g.name asc
    limit $2 offset $3;`

	rows, err := postgres.DB().Query(query, user_id, request.Size, request.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		group := &GroupApi{}
		if err := group.Scan(rows.Scan); err != nil {
			return nil, err
		}
		result.Results = append(result.Results, group)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	result.Size = len(result.Results)

	const queryCount = `
	select count(g.id)
	from groups g
	where
	 		g.active = true
		and g.id in (
				select distinct group_id
				from guests
				where user_id = $1 and active = true
		)`
	err = postgres.DB().QueryRow(queryCount, user_id).Scan(&result.Total)
	return result, err
}

func (g *Group) ApiDetail(admin bool) (interface{}, error) {
	var d struct {
		Url       string     `json:"url"`
		OwnerId   int64      `json:"owner"`
		Created   time.Time  `json:"created"`
		Name      string     `json:"name"`
		Tasks     []Task     `json:"tasks"`
		Customers []Customer `json:"customers"`
		Guests    []ApiGuest `json:"guests,omitempty"`
	}
	d.Url = g.Url
	d.Created = g.Created
	d.OwnerId = g.OwnerId
	d.Name = g.Name

	if tasks, err := g.Tasks(); err != nil {
		return nil, err
	} else {
		d.Tasks = tasks
	}

	if customers, err := g.Customers(); err != nil {
		return nil, err
	} else {
		d.Customers = customers
	}

	if admin {
		if guests, err := g.ApiGuests(); err != nil {
			return nil, err
		} else {
			d.Guests = guests
		}
	}
	return d, nil
}
