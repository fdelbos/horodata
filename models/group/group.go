package group

import (
	"bitbucket.com/hyperboloide/horo/helpers"
	"bitbucket.com/hyperboloide/horo/models/user"
	"bitbucket.com/hyperboloide/horo/services/postgres"
	"time"
)

type Group struct {
	Id      int64     `json:"id"`
	Created time.Time `json:"created"`
	Active  bool      `json:"active"`
	OwnerId int64     `json:"owner_id"`
	Name    string    `json:"name"`
	Url     string    `json:"url"`
}

func (g *Group) Scan(scanFn func(dest ...interface{}) error) error {
	return scanFn(
		&g.Id,
		&g.Created,
		&g.Active,
		&g.OwnerId,
		&g.Name,
		&g.Url)
}

func (g *Group) GetOwner() (*user.User, error) {
	return user.ById(g.OwnerId)
}

func (g *Group) Insert() error {
	urlTest := func(url string) (interface{}, error) {
		return ByUrl(url)
	}

	if tmp, err := helpers.Gen(g.Name, urlTest); err != nil {
		return err
	} else {
		g.Url = tmp
	}

	const query = `
	insert into groups (owner_id, name, url)
	values ($1, $2, $3);`
	return postgres.Exec(query, g.OwnerId, g.Name, g.Url)
}

func ByUrl(url string) (*Group, error) {
	g := &Group{}
	query := `
    select * from groups where url = $1 and active = true
	order by name, id;`
	return g, postgres.QueryRow(g, query, url)
}
