package group

import (
	"time"

	"dev.hyperboloide.com/fred/horodata/models/errors"
	"dev.hyperboloide.com/fred/horodata/models/user"
	"dev.hyperboloide.com/fred/horodata/services/mail"
	"dev.hyperboloide.com/fred/horodata/services/postgres"
)

type Guest struct {
	Id      int64     `json:"id"`
	Created time.Time `json:"-"`
	Active  bool      `json:"-"`
	GroupId int64     `json:"-"`
	UserId  *int64    `json:"-"`
	Rate    int       `json:"rate"`
	Admin   bool      `json:"admin"`
	Email   string    `json:"email"`
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
		&g.Email)
}

func (g *Guest) Update() error {
	const query = `
	update guests
	set active = $2, rate = $3, admin = $4, user_id = $5
	where id = $1;`
	return postgres.Exec(query, g.Id, g.Active, g.Rate, g.Admin, g.UserId)
}

func (g *Group) GuestAdd(email string, rate int, admin, sendMail bool) error {

	u, err := user.ByEmail(email)
	if err == errors.NotFound {
		u = nil
	} else if err != nil {
		return err
	}
	owner, err := g.GetOwner()
	if err != nil {
		return err
	}

	guest := &Guest{}
	const findQuery = `
    select * from guests where group_id = $1 and email = $2;`

	if err := postgres.QueryRow(guest, findQuery, g.Id, email); err == nil {
		wasActive := guest.Active
		if guest.UserId == nil || *guest.UserId != g.OwnerId {
			guest.Admin = admin
		}
		guest.Active = true
		guest.Rate = rate

		if u != nil {
			guest.UserId = &u.Id
		}
		if err := guest.Update(); err != nil {
			return err
		}
		if wasActive {
			return nil
		}
	} else if err != errors.NotFound {
		return err
	} else if u == nil {
		const insertQuery = `
		insert into guests (group_id, rate, admin, email)
		values ($1, $2, $3, $4);`
		if err := postgres.Exec(insertQuery, g.Id, rate, admin, email); err != nil {
			return err
		}
	} else if u != nil {
		const insertQuery = `
		insert into guests (group_id, rate, admin, email, user_id)
		values ($1, $2, $3, $4, $5);`
		if err := postgres.Exec(insertQuery, g.Id, rate, admin, email, u.Id); err != nil {
			return err
		}
	}

	if !sendMail {
		return nil
	}
	m := mail.Mail{
		Dest:     email,
		Subject:  "Nouvelle invitation sur Horodata",
		Template: "invitation",
		Data: map[string]interface{}{
			"ownerName": owner.FullName,
			"groupName": g.Name,
			"groupUrl":  g.Url,
		},
	}
	return m.Send()

}

func (g *Group) GuestGetByEmail(email string) (*Guest, error) {
	guest := &Guest{}
	const query = `
    select * from guests
    where 	group_id = $1
		and active = true
		and email = $2`
	return guest, postgres.QueryRow(guest, query, g.Id, email)
}

func (g *Group) GuestGetByUserId(id int64) (*Guest, error) {
	guest := &Guest{}
	const query = `
    select *
    from guests
    where   group_id = $1
        and active = true
        and user_id = $2`
	return guest, postgres.QueryRow(guest, query, g.Id, id)
}

func (g *Group) GuestGetById(id int64) (*Guest, error) {
	guest := &Guest{}
	const query = `
    select *
    from guests
    where   group_id = $1
        and active = true
        and id = $2`
	return guest, postgres.QueryRow(guest, query, g.Id, id)
}

type ApiGuest struct {
	Id       int64   `json:"id"`
	Active   bool    `json:"active"`
	Rate     int     `json:"rate"`
	Admin    bool    `json:"admin"`
	Email    string  `json:"email"`
	FullName *string `json:"full_name,omitempty"`
}

func (g *ApiGuest) Scan(scanFn func(dest ...interface{}) error) error {
	return scanFn(
		&g.Id,
		&g.Active,
		&g.Rate,
		&g.Admin,
		&g.Email,
		&g.FullName)
}

func (g *Group) ApiGuests() ([]ApiGuest, error) {
	const query = `
    select
		g.id, g.active, g.rate, g.admin, g.email,
		u.full_name
	from
		guests g left outer join users u on (g.user_id = u.id)
    where
		g.group_id = $1`

	rows, err := postgres.DB().Query(query, g.Id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	results := []ApiGuest{}
	for rows.Next() {
		i := &ApiGuest{}
		if err := i.Scan(rows.Scan); err != nil {
			return nil, err
		}
		results = append(results, *i)
	}
	return results, rows.Err()
}
