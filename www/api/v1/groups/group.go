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
			middlewares.GroupOwnerFilter(),
			TaskAdd)

		groups.PUT("/:group/tasks/:taskId",
			middlewares.GroupFilter(),
			middlewares.GroupOwnerFilter(),
			TaskUpdate)

		groups.DELETE("/:group/tasks/:taskId",
			middlewares.GroupFilter(),
			middlewares.GroupOwnerFilter(),
			TaskDelete)

		groups.POST("/:group/customers",
			middlewares.GroupFilter(),
			middlewares.GroupOwnerFilter(),
			CustomerAdd)

		groups.PUT("/:group/customers/:customerId",
			middlewares.GroupFilter(),
			middlewares.GroupOwnerFilter(),
			CustomerUpdate)

		groups.DELETE("/:group/customers/:customerId",
			middlewares.GroupFilter(),
			middlewares.GroupOwnerFilter(),
			CustomerDelete)

		// groups.PUT("/:url", Update, middlewares.GroupFilter(), middlewares.GroupOwnerFilter())
		// groups.GET("/:group/users", UserListing)

	}
}
