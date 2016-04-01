package user

import (
	"database/sql"
	"fmt"
	"time"

	"dev.hyperboloide.com/fred/horodata/services/cache"
	"dev.hyperboloide.com/fred/horodata/services/mail"
	"dev.hyperboloide.com/fred/horodata/services/postgres"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       int64     `json:"id,omitempty"`
	Created  time.Time `json:"created"`
	Active   bool      `json:"active"`
	Email    string    `json:"email,omitempty"`
	FullName string    `json:"name"`
}

const (
	currentHashVersion = 1
	cacheUserId        = "models.users.id"
)

func (u *User) Scan(scanFn func(dest ...interface{}) error) error {
	return scanFn(
		&u.Id,
		&u.Created,
		&u.Active,
		&u.Email,
		&u.FullName)
}

func (u User) saveInCache() error {
	id := fmt.Sprintf("%d", u.Id)
	return cache.SetPackage(cacheUserId, id, u, time.Hour)
}

func (u User) removeFromCache() error {
	id := fmt.Sprintf("%d", u.Id)
	return cache.DelPackage(cacheUserId, id)
}

func (u *User) SendWelcome() error {
	m := &mail.Mail{
		Dest:     u.Email,
		Subject:  "Bienvenue sur Horodata",
		Template: "welcome",
		Data:     map[string]interface{}{},
	}
	return m.Send()
}

func (u *User) Update() error {
	const query = `
	update guests
	set active = $2, full_name = $3
	where id = $1;`
	return postgres.Exec(query, u.Id, u.Active, u.FullName)
}

func (u *User) Insert() error {
	const query = `SELECT * from "user_new"($1, $2);`

	if err := postgres.Exec(query, u.Email, u.FullName); err != nil {
		return err
	} else if nu, err := ByEmail(u.Email); err != nil {
		return err
	} else {
		*u = *nu
	}
	return u.removeFromCache()
}

func (u User) CheckPassword(password string) (bool, error) {
	var tmp struct {
		hash        []byte
		hashVersion uint32
	}
	const query = `
	select hash, hash_version
    from users
    where active = true and id = $1;`

	err := postgres.DB().QueryRow(query, u.Id).Scan(&tmp.hash, &tmp.hashVersion)
	if err != nil && err == sql.ErrNoRows {
		return false, nil
	} else if err != nil {
		return false, err
	}
	ok := bcrypt.CompareHashAndPassword(tmp.hash, []byte(password)) == nil
	return ok, nil
}

func (u User) UpdatePassword(password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	const query = `update users set hash = $1, hash_version = $2 where id = $3;`
	return postgres.Exec(query, hash, currentHashVersion, u.Id)
}

func ByEmail(email string) (*User, error) {
	user := &User{}
	const query = `select * from  users_active where email = $1;`
	return user, postgres.QueryRow(user, query, email)
}

func ById(id int64) (*User, error) {
	user := &User{}
	if err := cache.GetPackage(cacheUserId, fmt.Sprintf("%d", id), user); err == nil {
		return user, nil
	}

	const query = `select * from  users_active where id = $1;`
	if err := postgres.QueryRow(user, query, id); err != nil {
		return nil, err
	}
	return user, user.saveInCache()
}
