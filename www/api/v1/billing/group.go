package billing

import "github.com/gin-gonic/gin"

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
		billing.GET("/invoices", Invoices)
		billing.GET("/invoices/:id", GetInvoice)
	}
}
