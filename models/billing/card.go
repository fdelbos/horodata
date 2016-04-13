package billing

import (
	"time"

	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/card"

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

func (c *Customer) UpdateCard(token string) error {
	resp, err := card.New(&stripe.CardParams{
		Customer: c.StripeId,
		Token:    token,
	})
	if err != nil {
		return err
	}

	expiration := time.Date(
		int(resp.Year),
		time.Month(int(resp.Month)),
		1, 0, 0, 0, 0, time.UTC)

	const query = `
    insert into
        cards (user_id, stripe_id, last4, brand, expiration)
    values
        ($1, $2, $3, $4, $5)
    on conflict (user_id) do update set
        stripe_id = $2, last4 = $3, brand = $4, expiration = $5`

	return postgres.Exec(
		query,
		c.UserId,
		resp.ID,
		resp.LastFour,
		string(resp.Brand),
		expiration,
	)
}
