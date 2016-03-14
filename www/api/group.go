package api

import (
	"bitbucket.com/hyperboloide/horo/middlewares"
	"bitbucket.com/hyperboloide/horo/www/api/v1"
	"github.com/gin-gonic/gin"
)

func Group(r *gin.RouterGroup) {
	api := r.Group("/api")
	api.Use(middlewares.UserFilter())
	{
		v1.Group(api)
	}
}
