package billing

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Group(r *gin.RouterGroup) {
	billing := r.Group("/billing")
	{
		billing.GET("/address", GetAddress)
		billing.POST("/address", NewAddress)
		billing.GET("/card", GetCard)
		billing.POST("/card", NewCard)
		billing.GET("/stripe_key", StripeKey)
		billing.GET("/plan", GetPlan)
		billing.GET("/end_period", GetEndPeriod)
		billing.POST("/change_plan", ChangePlan)

		endpoint := fmt.Sprintf("/%s", viper.GetString("payment_endpoint"))
		billing.POST(endpoint, StripeEndpoint)
	}
}

func StripeEndpoint(c *gin.Context) {}
