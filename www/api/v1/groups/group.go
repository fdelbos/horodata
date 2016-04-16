package groups

import (
	"dev.hyperboloide.com/fred/horodata/middlewares"
	"dev.hyperboloide.com/fred/horodata/www/api/jsend"
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
			gr.POST("/leave", LeaveGroup)
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
			gr.DELETE("/jobs/:jobId", JobDelete)
			gr.PUT("/jobs/:jobId", JobUpdate)
			gr.GET("/export_csv", middlewares.GroupAdminFilter(), ExportCSV)
			gr.GET("/export_xlsx", middlewares.GroupAdminFilter(), ExportXLSX)

			stats := gr.Group("/stats")
			{
				stats.GET("/customer_time", StatsCustomerTime)
				stats.GET("/task_time", StatsTaskTime)
				stats.GET("/guest_time", StatsGuestTime)
			}
		}
	}
}

func LeaveGroup(c *gin.Context) {
	g := middlewares.GetGroup(c)
	guest := middlewares.GetGuest(c)

	guest.Active = false
	if g.OwnerId == *guest.UserId {
		jsend.Forbidden(c)
	} else if err := guest.Update(); err != nil {
		jsend.Error(c, err)
	} else {
		jsend.Ok(c, nil)
	}
}
