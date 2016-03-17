package tests

import (
	"bitbucket.com/hyperboloide/horo/models/user"
	"bitbucket.com/hyperboloide/horo/services/postgres"
	"github.com/dchest/uniuri"
)

func NewUser() (*user.User, error) {
	u := &user.User{
		Login: uniuri.NewLen(12),
		Email: uniuri.NewLen(12) + "@test.com",
	}
	if err := u.Insert(); err != nil {
		return nil, err
	}
	return user.ByLogin(u.Login)
}

func CleanupUser(u *user.User) error {
	_, err := postgres.DB().Exec("DELETE FROM users WHERE login = $1", u.Login)
	return err
}
