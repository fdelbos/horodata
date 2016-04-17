package main

import (
	"os"

	log_service "dev.hyperboloide.com/fred/horodata/services/log"
	"dev.hyperboloide.com/fred/horodata/services/mail"
	"dev.hyperboloide.com/fred/horodata/services/postgres"
	log "github.com/Sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/stripe/stripe-go"
)

var started bool

func init() {
	viper.SetEnvPrefix("inv")

	viper.BindEnv("output")
	viper.SetDefault("output", "/tmp/invoices.horodata")

	viper.BindEnv("queue_name")
	viper.SetDefault("queue_name", "invoicing")

	viper.BindEnv("queue_host")
	viper.SetDefault("queue_host", "amqp://guest:guest@localhost:5672/")

	//
	// PostgreSQL
	//

	viper.BindEnv("pg_host")
	viper.SetDefault("pg_host", "localhost")

	viper.BindEnv("pg_dbname")
	viper.SetDefault("pg_dbname", "horodata")

	viper.BindEnv("pg_user")
	viper.SetDefault("pg_user", "postgres")

	viper.BindEnv("pg_password")
	viper.SetDefault("pg_password", "password")

	viper.BindEnv("pg_ssl")
	viper.SetDefault("pg_ssl", "false")

	viper.BindEnv("pg_pool_max")
	viper.SetDefault("pg_pool_max", "30")

	viper.BindEnv("pg_pool_idle")
	viper.SetDefault("pg_pool_idle", "10")

	//
	// Email
	//

	viper.BindEnv("mail_queue_name")
	viper.SetDefault("mail_queue_name", "mails")

	viper.BindEnv("mail_queue_host")
	viper.SetDefault("mail_queue_host", "amqp://guest:guest@localhost:5672/")

	//
	// log
	//

	viper.BindEnv("log_format")
	viper.SetDefault("log_format", "text")

	viper.BindEnv("log_level")
	viper.SetDefault("log_level", "debug")

	//
	// payment
	//

	viper.BindEnv("payment_secret_key")
	viper.SetDefault("payment_secret_key", "sk_test_ksm8vhIDWTyGCzDpHulwPF6l")

	//
	// Pdf
	//

	viper.BindEnv("pdf_host")
	viper.SetDefault("pdf_host", "http://localhost:8888")

	viper.BindEnv("pdf_invoice")
	viper.SetDefault("pdf_invoice", "/invoice")
}

func Configure() {
	if started {
		log.Fatal("System already started.")
	}
	started = true

	log_service.Configure()
	mail.Configure()
	postgres.Configure()

	stripe.Key = viper.GetString("payment_secret_key")

	if err := os.MkdirAll(viper.GetString("output"), 0755); err != nil {
		log.Fatal(err)
	}
}
