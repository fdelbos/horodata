package groups

import (
	"dev.hyperboloide.com/fred/horodata/middlewares"
	"github.com/gin-gonic/gin"
)

// Group for Groups package urls
func Group(r *gin.RouterGroup) {
	groups := r.Group("/groups")
	{
		groups.GET("", Listing)
		groups.POST("", Create)

		gr := groups.Group("/:group")
		gr.Use(middlewares.GroupFilter())
		{
			gr.GET("", Get)
			gr.DELETE("", middlewares.GroupOwnerFilter(), Delete)
			gr.POST("/tasks", middlewares.GroupAdminFilter(), TaskAdd)
			gr.PUT("/tasks/:taskId", middlewares.GroupAdminFilter(), TaskUpdate)
			gr.DELETE("/tasks/:taskId", middlewares.GroupAdminFilter(), TaskDelete)
			gr.POST("/customers", middlewares.GroupAdminFilter(), CustomerAdd)
			gr.PUT("/customers/:customerId", middlewares.GroupAdminFilter(), CustomerUpdate)
			gr.DELETE("/customers/:customerId", middlewares.GroupAdminFilter(), CustomerDelete)
			gr.POST("/guests", middlewares.GroupAdminFilter(), GuestAdd)
			gr.PUT("/guests/:guestId", middlewares.GroupAdminFilter(), GuestUpdate)
			gr.DELETE("/guests/:guestId", middlewares.GroupAdminFilter(), GuestDelete)
			gr.GET("/jobs", JobListing)
			gr.POST("/jobs", JobAdd)
			gr.GET("/export_csv", ExportCSV)
			gr.GET("/export_xlsx", ExportXLSX)

			stats := gr.Group("/stats")
			{
				stats.GET("/customer_time", StatsCustomerTime)
				stats.GET("/task_time", StatsTaskTime)
				stats.GET("/guest_time", StatsGuestTime)
			}
		}
	}
}
