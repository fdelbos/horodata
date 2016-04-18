package api

import (
	"dev.hyperboloide.com/fred/horodata/middlewares"
	"dev.hyperboloide.com/fred/horodata/www/api/v1"
	"github.com/gin-gonic/gin"
)

func Group(r *gin.RouterGroup) {
	api := r.Group("/api")
	api.Use(middlewares.UserFilter)
	api.Use(middlewares.AjaxCSRFFilter)
	{
		v1.Group(api)
	}
}
