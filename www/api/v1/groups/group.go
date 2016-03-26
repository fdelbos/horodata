package groups

import (
	"bitbucket.com/hyperboloide/horo/middlewares"
	"github.com/gin-gonic/gin"
)

func Group(r *gin.RouterGroup) {
	groups := r.Group("/groups")
	{
		groups.GET("", Listing)
		groups.POST("", Create)

<<<<<<< HEAD
		gr := groups.Group("/:group")
		gr.Use(middlewares.GroupFilter())
		{
			gr.GET("", Get)
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
			}
		}
=======
		groups.GET("/:group",
			middlewares.GroupFilter(),
			Get)

		groups.POST("/:group/tasks",
			middlewares.GroupFilter(),
			middlewares.GroupAdminFilter(),
			TaskAdd)

		groups.PUT("/:group/tasks/:taskId",
			middlewares.GroupFilter(),
			middlewares.GroupAdminFilter(),
			TaskUpdate)

		groups.DELETE("/:group/tasks/:taskId",
			middlewares.GroupFilter(),
			middlewares.GroupAdminFilter(),
			TaskDelete)

		groups.POST("/:group/customers",
			middlewares.GroupFilter(),
			middlewares.GroupAdminFilter(),
			CustomerAdd)

		groups.PUT("/:group/customers/:customerId",
			middlewares.GroupFilter(),
			middlewares.GroupAdminFilter(),
			CustomerUpdate)

		groups.DELETE("/:group/customers/:customerId",
			middlewares.GroupFilter(),
			middlewares.GroupAdminFilter(),
			CustomerDelete)

		groups.POST("/:group/guests",
			middlewares.GroupFilter(),
			middlewares.GroupAdminFilter(),
			GuestAdd)

		groups.PUT("/:group/guests/:guestId",
			middlewares.GroupFilter(),
			middlewares.GroupAdminFilter(),
			GuestUpdate)

		groups.DELETE("/:group/guests/:guestId",
			middlewares.GroupFilter(),
			middlewares.GroupAdminFilter(),
			GuestDelete)

		groups.GET("/:group/jobs",
			middlewares.GroupFilter(),
			JobListing)

		groups.POST("/:group/jobs",
			middlewares.GroupFilter(),
			JobAdd)

>>>>>>> fc157b92e6aaffac4e506b943790d083b62a8426
	}
}
