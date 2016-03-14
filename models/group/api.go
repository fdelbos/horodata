package group

import (
	"bitbucket.com/hyperboloide/horo/models/types/listing"
	"bitbucket.com/hyperboloide/horo/models/user"
	"bitbucket.com/hyperboloide/horo/services/postgres"
	"bitbucket.com/hyperboloide/horo/services/urls"
	"encoding/json"
	"time"
)

type GroupApi struct {
	Url     string        `json:"url"`
	Owner   user.UserLink `json:"owner,omitempty"`
	Created time.Time     `json:"created,omitempty"`
	Name    string        `json:"name"`
}

func (ga *GroupApi) MarshalJSON() ([]byte, error) {
	type alias GroupApi

	return json.Marshal(&struct {
		Link string `json:"_link"`
		*alias
	}{urls.ApiGroup(ga.Url), (*alias)(ga)})
}

func (ag *GroupApi) Scan(scanFn func(dest ...interface{}) error) error {
	return scanFn(
		&ag.Url,
		&ag.Owner.Login,
		&ag.Created,
		&ag.Name)
}

func ApiByUser(user_id int64, request *listing.Request) (*listing.Result, error) {
	result := &listing.Result{}
	result.Offset = request.Offset

	const query = `
    select g.url, u.login, g.created, g.name
    from
		groups g
		left outer join users u on g.owner_id = u.id
    where
			g.active = true
		and (	g.owner_id = $1
			or 	g.id in (
					select distinct group_id
					from guests
					where user_id = $1 and active = true
				)
			)
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
		and (	g.owner_id = $1
			or	g.id in (
				select distinct group_id
				from guests
				where user_id = $1 and active = true
				)
			)`
	err = postgres.DB().QueryRow(queryCount, user_id).Scan(&result.Total)
	return result, err
}

func ApiByUrl(url string) (*GroupApi, error) {
	ga := &GroupApi{}
	query := `
    select g.url, u.login, g.created, g.name
    from users u, groups g
    where u.id = g.owner_id and g.url = $1;`
	return ga, postgres.QueryRow(ga, query, url)
}
