package group

import (
	"dev.hyperboloide.com/fred/horodata/models/errors"
	"dev.hyperboloide.com/fred/horodata/services/postgres"
	"time"
)

type Customer struct {
	Id      int64     `json:"id"`
	Created time.Time `json:"-"`
	Active  bool      `json:"-"`
	GroupId int64     `json:"-"`
	Name    string    `json:"name"`
}

func (c *Customer) Scan(scanFn func(dest ...interface{}) error) error {
	return scanFn(
		&c.Id,
		&c.Created,
		&c.Active,
		&c.GroupId,
		&c.Name)
}

func (c *Customer) Update() error {
	const query = `
	update customers
	set active = $2, name = $3
	where id = $1;`
	return postgres.Exec(query, c.Id, c.Active, c.Name)
}

func (g *Group) CustomerGet(id int64) (*Customer, error) {
	c := &Customer{}
	const query = `
	select * from customers where active = true and group_id = $1 and id = $2`
	return c, postgres.QueryRow(c, query, g.Id, id)
}

func (g *Group) CustomerAdd(name string) error {
	customer := &Customer{}
	const findQuery = `
    select * from customers where group_id = $1 and name = $2;`

	if err := postgres.QueryRow(customer, findQuery, g.Id, name); err == nil {
		customer.Active = true
		return customer.Update()
	} else if err != errors.NotFound {
		return err
	}

	const insertQuery = `
	insert into customers (group_id, name)
	values ($1, $2);`

	return postgres.Exec(insertQuery, g.Id, name)
}

func (g *Group) Customers() ([]Customer, error) {
	const query = `
    select * from customers
    where group_id = $1 and active = true
    order by name, id;`

	rows, err := postgres.DB().Query(query, g.Id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	results := []Customer{}
	for rows.Next() {
		i := Customer{}
		if err := i.Scan(rows.Scan); err != nil {
			return nil, err
		}
		results = append(results, i)
	}
	return results, rows.Err()
}
