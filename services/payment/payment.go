package payment

import (
	"github.com/spf13/viper"
	"github.com/stripe/stripe-go"
)

func Configure() {
	stripe.Key = viper.GetString("payment_secret_key")
}

func PublishableKey() string {
	return viper.GetString("payment_publishable_key")
}
