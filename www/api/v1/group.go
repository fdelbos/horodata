package v1

import (
	"dev.hyperboloide.com/fred/horodata/www/api/v1/billing"
	"dev.hyperboloide.com/fred/horodata/www/api/v1/groups"
	"dev.hyperboloide.com/fred/horodata/www/api/v1/users"
	"github.com/gin-gonic/gin"
)

func Group(r *gin.RouterGroup) {
	v1 := r.Group("/v1")
	{
		users.Group(v1)
		groups.Group(v1)
		billing.Group(v1)
	}
}
