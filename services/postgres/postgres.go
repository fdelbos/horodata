package postgres

import (
	"bitbucket.com/hyperboloide/horo/models/errors"
	"database/sql"
)

type Scanner interface {
	Scan(func(dest ...interface{}) error) error
}

func DB() *sql.DB {
	return db
}

func Ping() error {
	return DB().Ping()
}

func Exec(query string, params ...interface{}) error {
	stmt, err := DB().Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(params...)
	return err
}

func QueryRow(obj Scanner, query string, params ...interface{}) error {
	err := obj.Scan(DB().QueryRow(query, params...).Scan)
	if err != nil && err == sql.ErrNoRows {
		return errors.NotFound
	}
	return err
}
