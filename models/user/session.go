package user

import (
	"bitbucket.com/hyperboloide/horo/services/cache"
	"bitbucket.com/hyperboloide/horo/services/postgres"
	"fmt"
	"strconv"
	"time"
)

type Session struct {
	Id      int64     `json:"id"`
	Created time.Time `json:"created"`
	UserId  int64     `json:"user_id"`
	Active  bool      `json:"active"`
	Host    string    `json:"host"`
}

const (
	cacheSessionPkg = "models.users.sessions"
)

type SessionId int64

func (s *Session) Scan(scanFn func(dest ...interface{}) error) error {
	return scanFn(
		&s.Id,
		&s.Created,
		&s.UserId,
		&s.Active,
		&s.Host)
}

func (s *Session) saveInCache() error {
	id := fmt.Sprintf("%d", s.Id)
	return cache.SetPackage(cacheSessionPkg, id, s, time.Hour)
}

func (s *Session) deleteFromCache() error {
	id := fmt.Sprintf("%d", s.Id)
	return cache.DelPackage(cacheSessionPkg, id)
}

func (s *Session) IsValid() bool {
	// check if it's 2 weeks old
	if s.Created.Add(time.Hour * 24 * 14).Before(time.Now()) {
		return false
	}
	return s.Active
}

func (s *Session) GetUser() (*User, error) {
	return ById(s.UserId)
}

func (s *Session) Close() error {
	s.Active = false
	const query = `UPDATE sessions SET active = false WHERE id = $1;`
	if err := postgres.Exec(query, s.Id); err != nil {
		return err
	}
	return s.deleteFromCache()
}

func NewSession(u *User, host string) (string, error) {
	var id string

	const req = `
	INSERT INTO sessions (user_id, host)
	VALUES ($1, $2) RETURNING id`

	var tmp int64
	err := postgres.DB().QueryRow(req, u.Id, host).Scan(&tmp)
	if err != nil {
		return id, err
	}
	id = fmt.Sprintf("%d", tmp)
	return id, nil
}

func GetSession(id string) (*Session, error) {
	session := &Session{}

	if err := cache.GetPackage(cacheSessionPkg, id, session); err == nil {
		return session, err
	}
	key, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	query := `SELECT id, created, user_id, active, host FROM sessions WHERE id = $1;`
	return session, postgres.QueryRow(session, query, key)
}
