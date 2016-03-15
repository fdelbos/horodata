package group

import (
	// "bitbucket.com/hyperboloide/horo/models/errors"
	// "bitbucket.com/hyperboloide/horo/models/user"
	"bitbucket.com/hyperboloide/horo/services/postgres"
	"time"
)

type Guest struct {
	Id      int64     `json:"id"`
	Created time.Time `json:"created"`
	Active  bool      `json:"active"`
	GroupId int64     `json:"group_id"`
	UserId  int64     `json:"user_id"`
	Rate    int       `json:"hour_rate"`
	Admin   bool      `json:"admin"`
	Email   string    `json:"email"`
	Message string    `json:"message"`
}

func (g *Guest) Scan(scanFn func(dest ...interface{}) error) error {
	return scanFn(
		&g.Id,
		&g.Created,
		&g.Active,
		&g.GroupId,
		&g.UserId,
		&g.Rate,
		&g.Admin,
		&g.Email,
		&g.Message)
}

func (g *Guest) Update() error {
	const query = `
	update groups
	set active = $2, rate = $3, admin = $4
	where id = $1;`
	return postgres.Exec(query, g.Id, g.Active, g.Rate, g.Admin)
}

func (g *Group) GuestAdd(email string, message, rate int, admin bool) error {
	// guest := &Guest{}
	// const existsQuery = `
	// select *
	// from guests g
	// where   group_id = $1
	//     and user_id = (
	//         select id from users where email = $2
	//     );`
	// err := postgres.QueryRow(guest, existsQuery, g.Id, email)

	// guest, err := g.GuestGetByEmail(email)
	// if err != nil && err != errors.NotFound {
	// 	return err
	// } else if err == nil {
	// 	guest.Active = true
	// 	guest.Rate = rate
	// 	guest.Admin = admin
	// 	return guest.Update()
	// }
	//
	// guest = &Guest{
	//     Email: email,
	//
	// }
	// if _, err := user.ByEmail(email); err != nil && err != errors.NotFound {
	// 	return err
	// } else if err == errors.NotFound {
	//     guest.Email =
	// }
	return nil
}

func (g *Group) GuestGetByEmail(email string) (*Guest, error) {
	guest := &Guest{}
	const query = `
    select *
    from guests g
    where   group_id = $1
        and active = true
        and user_id = (
            select id from users where email = $2
        );`
	return guest, postgres.QueryRow(guest, query, g.Id, email)
}

func (g *Group) GuestGetByUserId(id int64) (*Guest, error) {
	guest := &Guest{}
	const query = `
    select *
    from guests g
    where   group_id = $1
        and active = true
        and user_id = $2`
	return guest, postgres.QueryRow(guest, query, g.Id, id)
}

func (g *Group) Guests() ([]Guest, error) {
	const query = `
    select * from guests
    where active = true and group_id = $1`

	rows, err := postgres.DB().Query(query, g.Id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	results := []Guest{}
	for rows.Next() {
		i := &Guest{}
		if err := i.Scan(rows.Scan); err != nil {
			return nil, err
		}
		results = append(results, *i)
	}
	return results, rows.Err()
}

func (g *Group) GuestRemove(id int64) error {
	return nil
}
