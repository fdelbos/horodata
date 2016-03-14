package groups

import (
	"github.com/gin-gonic/gin"
)

func Group(r *gin.RouterGroup) {
	roles := r.Group("/groups")
	{
		roles.GET("", Listing)
		roles.POST("", Create)
		roles.GET("/:url", Get)
		roles.PUT("/:url", Update)
		roles.GET("/:url/users", UserListing)
	}
}
