package cache

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/spf13/viper"
	"gopkg.in/redis.v3"
)

var (
	client *redis.Client
)

func Configure() {
	addr := fmt.Sprintf(
		"%s:%d",
		viper.GetString("redis_host"),
		viper.GetInt("redis_port"))

	client = redis.NewClient(&redis.Options{
		Addr:     addr,
		PoolSize: viper.GetInt("redis_port"),
	})
	if err := Ping(); err != nil {
		log.WithField("error", err).Fatal("Cannot ping redis.")
	}

	log.WithField("host", viper.GetString("redis_host")).Info("Connected to Redis.")
}
