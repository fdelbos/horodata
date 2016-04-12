package billing

import (
	"time"

	"dev.hyperboloide.com/fred/horodata/models/user"
	"dev.hyperboloide.com/fred/horodata/services/postgres"
	log "github.com/Sirupsen/logrus"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/sub"
)

const TaxPercent = 20.0

type Subscription struct {
	UserId     int64
	StripeId   string
	Active     bool
	Plan       string
	TaxPercent float64
	End        *time.Time
}

func (s *Subscription) Scan(scanFn func(dest ...interface{}) error) error {
	return scanFn(
		&s.UserId,
		&s.StripeId,
		&s.Active,
		&s.Plan,
		&s.TaxPercent,
		&s.End,
	)
}

func (s *Subscription) save() error {
	const query = `
    insert into stripe_subscriptions
        (user_id, stripe_id, active, plan, tax_percent, end_date)
    values
        ($1, $2, $3, $4, $5, $6)
    on conflict (user_id) do update set
        stripe_id = $2, active = $3, plan = $4, tax_percent = $5, end_date = $6;`
	err := postgres.Exec(
		query,
		s.UserId,
		s.StripeId,
		s.Active,
		s.Plan,
		s.TaxPercent,
		s.End)
	if err != nil {
		log.WithFields(map[string]interface{}{
			"step":         "Subscribtion save to db",
			"subscription": s,
		}).Error(err)
	}
	return err
}

func (s *Subscription) BackToFree() error {
	s.Plan = "free"
	if err := s.save(); err != nil {
		return err
	}

	u, err := user.ById(s.UserId)
	if err != nil {
		return err
	}
	return u.QuotaChangePlan(s.Plan)
}

func (s *Subscription) EndPeriod() (*time.Time, error) {
	c, err := s.Customer()
	if err != nil {
		return nil, err
	}

	stripeSub, err := sub.Get(
		s.StripeId,
		&stripe.SubParams{Customer: c.StripeId})
	if err != nil {
		return nil, err
	}
	end := time.Unix(stripeSub.PeriodEnd, 0)
	return &end, nil
}

func (s *Subscription) Customer() (*Customer, error) {
	c := &Customer{}
	const query = `select * from subscribers where user_id = $1`
	return c, postgres.QueryRow(c, query, s.UserId)
}

func (s *Subscription) Unsubscribe() error {
	if s.End != nil {
		return nil
	}

	c, err := s.Customer()
	if err != nil {
		return err
	}

	stripeSub, err := sub.Cancel(
		s.StripeId,
		&stripe.SubParams{
			Customer:  c.StripeId,
			EndCancel: true,
		})
	if err != nil {
		log.WithFields(map[string]interface{}{
			"step":         "Cancel subscription / send to stripe",
			"subscription": s,
		}).Error(err)
		return err
	}
	end := time.Unix(stripeSub.PeriodEnd, 0)
	s.End = &end
	return s.save()
}

func (s *Subscription) Update(plan string) error {
	if plan == "free" {
		return s.Unsubscribe()
	}

	taxPercent := TaxPercent
	var stripeSub *stripe.Sub
	var err error
	c, err := s.Customer()
	if err != nil {
		return err
	}

	if s.StripeId == "" {
		stripeSub, err = sub.New(
			&stripe.SubParams{
				Customer:   c.StripeId,
				Plan:       plan,
				TaxPercent: taxPercent,
			})

	} else {
		stripeSub, err = sub.Update(
			s.StripeId,
			&stripe.SubParams{
				Customer:   c.StripeId,
				Plan:       plan,
				TaxPercent: taxPercent,
			})
	}
	if err != nil {
		if err != nil {
			log.WithFields(map[string]interface{}{
				"step":         "Subscribe update / send to stripe",
				"plan":         plan,
				"subscription": s,
			}).Error(err)
			return err
		}
	}

	s.StripeId = stripeSub.ID
	s.Active = true
	s.Plan = plan
	s.TaxPercent = taxPercent
	s.End = nil
	if err := s.save(); err != nil {
		return err
	} else if u, err := c.User(); err != nil {
		return err
	} else {
		return u.QuotaChangePlan(plan)
	}
}

func SubscriptionByStripeId(id string) (*Subscription, error) {
	s := &Subscription{}
	const query = `select * from stripe_subscriptions where stripe_id = $1`
	return s, postgres.QueryRow(s, query, id)
}

func SubscriptionByUserId(id int64) (*Subscription, error) {
	s := &Subscription{}
	const query = `select * from stripe_subscriptions where user_id = $1`
	return s, postgres.QueryRow(s, query, id)
}
