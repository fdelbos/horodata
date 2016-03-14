package main

import (
	"bitbucket.com/hyperboloide/horo/config"
	"bitbucket.com/hyperboloide/horo/www"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	config.Configure()

	engine := gin.Default()
	r := engine.Group("/")
	{
		www.Group(r)

		if gin.IsDebugging() {
			r.Static("/static", "./static")
		}
	}

	connStr := fmt.Sprintf("%s:%d", viper.GetString("host"), viper.GetInt("port"))
	if err := engine.Run(connStr); err != nil {
		log.Fatal(err)
	}
}
