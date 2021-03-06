package tests

import (
	"dev.hyperboloide.com/fred/horodata/models/user"
	"dev.hyperboloide.com/fred/horodata/services/postgres"
	"github.com/dchest/uniuri"
)

func NewUser() (*user.User, error) {
	u := &user.User{
		FullName: uniuri.NewLen(12),
		Email:    uniuri.NewLen(12) + "@test.com",
	}
	if err := u.Insert(); err != nil {
		return nil, err
	}
	return user.ByEmail(u.Email)
}

func CleanupUser(u *user.User) error {
	_, err := postgres.DB().Exec("DELETE FROM users WHERE email = $1", u.Email)
	return err
}
