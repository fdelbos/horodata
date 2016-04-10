package payment

import (
	"github.com/spf13/viper"
	"github.com/stripe/stripe-go"
)

var (
	pubKey string
)

func Configure() {
	stripe.Key = viper.GetString("payment_secret_key")
	pubKey = viper.GetString("payment_publishable_key")
}

func PublishableKey() string {
	return pubKey
}
