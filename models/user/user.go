package user

import (
	"bitbucket.com/hyperboloide/horo/services/cache"
	"bitbucket.com/hyperboloide/horo/services/mail"
	"bitbucket.com/hyperboloide/horo/services/postgres"
	"bitbucket.com/hyperboloide/horo/services/urls"
	"database/sql"
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	Id           int64     `json:"id,omitempty"`
	Created      time.Time `json:"created"`
	Active       bool      `json:"active"`
	Login        string    `json:"login"`
	Email        string    `json:"email,omitempty"`
	FullName     string    `json:"full_name,omitempty"`
	Organization string    `json:"organization,omitempty"`
	Website      string    `json:"website,omitempty"`
	About        string    `json:"about,omitempty"`
}

const (
	currentHashVersion = 1
	cacheUserId        = "models.users.id"
)

func (u *User) Scan(scanFn func(dest ...interface{}) error) error {
	var fullName, organization, website, about sql.NullString
	err := scanFn(
		&u.Id,
		&u.Created,
		&u.Active,
		&u.Login,
		&u.Email,
		&fullName,
		&organization,
		&website,
		&about)
	if err != nil {
		return err
	}
	if fullName.Valid {
		u.FullName = fullName.String
	}
	if organization.Valid {
		u.Organization = organization.String
	}
	if website.Valid {
		u.Website = website.String
	}
	if about.Valid {
		u.About = about.String
	}
	return nil
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
		Dests:    []string{u.Email},
		Subject:  "Bienvenue sur Horo Data",
		Template: "welcome",
		Data: map[string]interface{}{
			"login": u.Login,
		},
	}
	return m.Send()
}

func (u *User) Insert() error {
	const query = `SELECT * from "user_new"($1, $2);`

	if err := postgres.Exec(query, u.Login, u.Email); err != nil {
		return err
	} else if nu, err := ByLogin(u.Login); err != nil {
		return err
	} else {
		*u = *nu
	}
	return u.removeFromCache()
}

func (u User) UpdateProfile() error {
	const query = `
	update users
	set full_name = $2, organization = $3, website = $4, about = $5
	where id = $1;`

	if err := postgres.Exec(query, u.Id, u.FullName, u.Organization, u.Website, u.About); err != nil {
		return err
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

func ByLogin(login string) (*User, error) {
	user := &User{}
	const query = `select * from  users_active where login = $1;`
	return user, postgres.QueryRow(user, query, login)
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

type UserLink struct {
	Login string `json:"login"`
}

func (ul *UserLink) MarshalJSON() ([]byte, error) {
	type alias UserLink
	return json.Marshal(&struct {
		Link string `json:"_link"`
		*alias
	}{urls.ApiUsers + "/" + ul.Login, (*alias)(ul)})
}
