package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/hyperboloide/dispatch"
	"github.com/spf13/viper"
)

func main() {
	Configure()

	queue, err := dispatch.NewAMQPQueue(
		viper.GetString("queue_name"),
		viper.GetString("queue_host"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.WithField("queue", viper.GetString("queue_name")).Info("Invoicing daemon started")
	if err := queue.ListenBytes(listenner); err != nil {
		log.Fatal(err)
	}
}
