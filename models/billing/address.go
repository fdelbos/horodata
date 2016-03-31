package billing

import (
	"time"

	"dev.hyperboloide.com/fred/horodata/services/postgres"
)

type Address struct {
	Id       int64     `json:"id"`
	Created  time.Time `json:"created"`
	UserId   int64     `json:"user_id"`
	FullName string    `json:"full_name"`
	Email    string    `json:"email,omitempty"`
	Company  string    `json:"company,omitempty"`
	VAT      string    `json:"vat,omitempty"`
	Address1 string    `json:"address_1,omitempty"`
	Address2 string    `json:"address_1,omitempty"`
	City     string    `json:"city,omitempty"`
	Zip      string    `json:"zip,omitempty"`
}

func (a *Address) Insert() error {
	const query = `
	INSERT INTO addresses (
		user_id,
		full_name,
		email,
		company,
		vat,
		address1,
		address2,
		city,
		zip,
	)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);`
	return postgres.Exec(
		query,
		a.UserId,
		a.FullName,
		a.Email,
		a.Company,
		a.VAT,
		a.Address1,
		a.Address2,
		a.City,
		a.Zip,
	)
}

func (a *Address) Scan(scanFn func(dest ...interface{}) error) error {
	return scanFn(
		&a.Id,
		&a.Created,
		&a.UserId,
		&a.FullName,
		&a.Email,
		&a.Company,
		&a.VAT,
		&a.Address1,
		&a.Address2,
		&a.City,
		&a.Zip)
}

func CurrentAddress(userId int64) (*Address, error) {
	const query = `select * from addresses where id = (
		select address_id
		from address_current
		where user_id = $1);`
	a := &Address{}
	return a, postgres.QueryRow(a, query, userId)
}
