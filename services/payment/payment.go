package payment

import (
	"encoding/json"

	log "github.com/Sirupsen/logrus"
	"github.com/hyperboloide/dispatch"
	"github.com/spf13/viper"
	"github.com/stripe/stripe-go"
)

var (
	pubKey string
	queue  dispatch.Queue
)

type StripeEvent struct {
	Id string `json:"id"`
}

func Configure() {
	stripe.Key = viper.GetString("payment_secret_key")
	pubKey = viper.GetString("payment_publishable_key")

	q, err := dispatch.NewAMQPQueue(
		viper.GetString("payment_queue_name"),
		viper.GetString("payment_queue_host"))
	if err != nil {
		log.WithField("error", err).Fatal("Cannot connect to  payment queue.")
	}
	queue = q
}

func PublishableKey() string {
	return pubKey
}

func NewEvent(id string) error {
	buff, err := json.Marshal(StripeEvent{id})
	if err != nil {
		return err
	}
	return queue.SendBytes(buff)
}
