package group

import (
	"time"

	"dev.hyperboloide.com/fred/horodata/helpers"
	"dev.hyperboloide.com/fred/horodata/models/user"
	"dev.hyperboloide.com/fred/horodata/services/postgres"
)

type Group struct {
	Id      int64     `json:"-"`
	Created time.Time `json:"-"`
	Active  bool      `json:"-"`
	OwnerId int64     `json:"-"`
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

func (g *Group) Delete() error {
	const query = `delete from groups where id = $1`
	return postgres.Exec(query, g.Id)
}

func ByUrl(url string) (*Group, error) {
	g := &Group{}
	query := `
    select * from groups where url = $1 and active = true
	order by name, id;`
	return g, postgres.QueryRow(g, query, url)
}
