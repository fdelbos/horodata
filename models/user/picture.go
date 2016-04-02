package user

import (
	"time"

	sqlerrors "dev.hyperboloide.com/fred/horodata/models/errors"
	"dev.hyperboloide.com/fred/horodata/services/picture"
	"dev.hyperboloide.com/fred/horodata/services/postgres"
	log "github.com/Sirupsen/logrus"
	"github.com/dchest/uniuri"
)

type Picture struct {
	Id      string    `json:"id"`
	Created time.Time `json:"created"`
	UserId  int64     `json:"user_id"`
	Origin  string    `json:"origin"`
}

func (p *Picture) Scan(scanFn func(dest ...interface{}) error) error {
	return scanFn(
		&p.Id,
		&p.Created,
		&p.UserId,
		&p.Origin)
}

func (p *Picture) Insert() error {
	const query = `
    insert into user_pictures (id, user_id, origin)
	values ($1, $2, $3);`

	return postgres.Exec(query, p.Id, p.UserId, p.Origin)
}

func (p *Picture) Update() error {
	const query = `
    update user_pictures set origin = $2
	where id = $1;`

	return postgres.Exec(query, p.Id, p.Origin)
}

func (u User) PictureGet() (*Picture, error) {
	const query = `
    select * from user_pictures
    where user_id = $1;`

	picture := &Picture{}
	return picture, postgres.QueryRow(picture, query, u.Id)
}

func (u User) PictureSetFromOrigin(origin string) error {

	isNew := false
	current, err := u.PictureGet()
	if err == sqlerrors.NotFound {
		isNew = true
		current = &Picture{
			Id:     uniuri.NewLen(32),
			UserId: u.Id,
			Origin: origin,
		}
	} else if err != nil {
		return err
	} else if current.Origin == "" || current.Origin == origin {
		return nil
	}
	current.Origin = origin

	if err := picture.ProfileFromUrl(origin, current.Id); err != nil {
		return err
	} else if isNew {
		log.WithFields(map[string]interface{}{
			"user": u.Id,
			"id":   current.Id,
		}).Info("Insert picture.")
		return current.Insert()
	} else {
		log.WithFields(map[string]interface{}{
			"user": u.Id,
			"id":   current.Id,
		}).Info("Update picture.")
		return current.Update()
	}
}
