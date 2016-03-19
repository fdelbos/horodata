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

	}
}
