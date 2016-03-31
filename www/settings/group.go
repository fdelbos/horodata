package settings

import (
	"dev.hyperboloide.com/fred/horodata/middlewares"
	"github.com/gin-gonic/gin"
)

func Group(r *gin.RouterGroup) {
	settings := r.Group("/settings")
	// settings.Use(middlewares.CSRFFilter)
	{
		settings.GET("/profile", GetProfile)
		settings.POST("/profile", middlewares.PostCSRFFilter, PostProfile)
		settings.POST("/language", middlewares.PostCSRFFilter, PostLanguage)
		settings.GET("/usage", GetUsage)

		settings.GET("/billing", GetBilling)
		settings.POST("/billing", middlewares.PostCSRFFilter, PostBilling)
	}
}
