package billing

import (
	"fmt"
	"time"

	"dev.hyperboloide.com/fred/horodata/models/user"
	"dev.hyperboloide.com/fred/horodata/services/postgres"
	log "github.com/Sirupsen/logrus"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/card"
	"github.com/stripe/stripe-go/customer"
)

type Customer struct {
	UserId   int64     `json:"user_id"`
	Created  time.Time `json:"created"`
	StripeId string    `json:"stripe_id"`
}

func (c *Customer) insert() error {
	const query = `
	insert into subscribers (user_id, stripe_id)
	values ($1, $2);`
	err := postgres.Exec(query, c.UserId, c.StripeId)
	if err != nil {
		log.WithFields(map[string]interface{}{
			"package":  "horodata.models.billing",
			"function": "func (c *Customer) insert() error",
			"step":     "postgres save",
		}).Error(err)
	}
	return err
}

func (c *Customer) Scan(scanFn func(dest ...interface{}) error) error {
	return scanFn(&c.UserId, &c.Created, &c.StripeId)
}

func NewCustomer(userId int64, token string) error {
	u, err := user.ById(userId)
	if err != nil {
		return err
	}

	cp := &stripe.CustomerParams{
		Desc:  fmt.Sprintf("%s - %s", u.FullName, u.Email),
		Email: u.Email,
	}

	cp.SetSource(token)
	sc, err := customer.New(cp)
	if err != nil {
		log.WithFields(map[string]interface{}{
			"package":  "horodata.models.billing",
			"function": "func NewCustomer(userId int64, token string) error",
			"step":     "Stripe new customer",
		}).Error(err)
		return err
	}
	cus := &Customer{
		UserId:   u.Id,
		StripeId: sc.ID,
	}
	if err := cus.insert(); err != nil {
		return err
	}

	cardObj, err := card.Get(
		sc.DefaultSource.ID,
		&stripe.CardParams{Customer: sc.ID})
	if err != nil {
		log.WithFields(map[string]interface{}{
			"package":  "horodata.models.billing",
			"function": "func NewCustomer(userId int64, token string) error",
			"step":     "Stripe get card",
		}).Error(err)
		return err
	}

	newCard := &Card{
		UserId:     u.Id,
		StripeId:   cardObj.ID,
		Last4:      cardObj.LastFour,
		Brand:      string(cardObj.Brand),
		Expiration: time.Date(int(cardObj.Year), time.Month(int(cardObj.Month)), 1, 0, 0, 0, 0, time.UTC),
	}
	return newCard.insert()
}

func CustomerByUserId(userId int64) (*Customer, error) {
	c := &Customer{}
	const query = `select * from subscribers where user_id = $1`
	return c, postgres.QueryRow(c, query, userId)
}

func CustomerByStripeId(id string) (*Customer, error) {
	c := &Customer{}
	const query = `select * from subscribers where stripe_id = $1`
	return c, postgres.QueryRow(c, query, id)
}

func (c *Customer) User() (*user.User, error) {
	return user.ById(c.UserId)
}

func (c *Customer) Subscription() (*Subscription, error) {
	s := &Subscription{}
	const query = `select * from stripe_subscriptions where user_id = $1`
	return s, postgres.QueryRow(s, query, c.UserId)
}
