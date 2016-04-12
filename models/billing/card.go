package billing

import (
	"time"

	"dev.hyperboloide.com/fred/horodata/services/postgres"
)

type Card struct {
	UserId     int64     `json:"user_id"`
	Created    time.Time `json:"created"`
	StripeId   string    `json:"stripe_id"`
	Last4      string    `json:"last4"`
	Brand      string    `json:"brand"`
	Expiration time.Time `json:"expiration"`
}

func (c *Card) insert() error {
	const query = `
	INSERT INTO cards (
		user_id,
        stripe_id,
        last4,
        brand,
        expiration
	)
	VALUES ($1, $2, $3, $4, $5);`
	return postgres.Exec(
		query,
		c.UserId,
		c.StripeId,
		c.Last4,
		c.Brand,
		c.Expiration,
	)
}

func (c *Card) Scan(scanFn func(dest ...interface{}) error) error {
	return scanFn(
		&c.UserId,
		&c.Created,
		&c.StripeId,
		&c.Last4,
		&c.Brand,
		&c.Expiration,
	)
}

func (c *Customer) Card() (*Card, error) {
	card := &Card{}
	const query = `select * from cards where user_id = $1`
	return card, postgres.QueryRow(card, query, c.UserId)
}
