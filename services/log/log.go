package log

import (
	log "github.com/Sirupsen/logrus"
	"github.com/spf13/viper"
)

func Configure() {
	switch viper.GetString("log_format") {
	case "text":
		log.SetFormatter(&log.TextFormatter{
			DisableColors: true,
		})
	case "json":
		log.SetFormatter(&log.JSONFormatter{})
	default:
		log.WithField("format", viper.GetString("log_format")).Fatal("Unknow log format")
	}

	switch viper.GetString("log_level") {
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "warning":
		log.SetLevel(log.WarnLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	case "fatal":
		log.SetLevel(log.FatalLevel)
	case "panic":
		log.SetLevel(log.PanicLevel)
	default:
		log.WithField("level", viper.GetString("log_level")).Fatal("Unknow log level")
	}
}
