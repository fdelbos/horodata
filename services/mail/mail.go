package mail

import (
	log "github.com/Sirupsen/logrus"
	"github.com/hyperboloide/qmail/client"
	"github.com/spf13/viper"
)

var (
	mailer *client.Mailer
)

func Configure() {
	m, err := client.New(
		viper.GetString("mail_queue_name"),
		viper.GetString("mail_queue_host"))
	if err != nil {
		log.WithField("error", err).Fatal("Cannot connect to mail queue")
	}
	log.WithField("queue", viper.GetString("mail_queue_name")).Info("Connected to mail queue.")
	mailer = m
}

func Mailer() *client.Mailer {
	return mailer
}
