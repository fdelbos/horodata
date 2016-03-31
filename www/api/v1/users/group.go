package users

import (
	"dev.hyperboloide.com/fred/horodata/middlewares"
	"dev.hyperboloide.com/fred/horodata/www/api/jsend"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Group(r *gin.RouterGroup) {
	users := r.Group("/users")
	{
		users.GET("/me", GetMe)
	}
}

func GetMe(c *gin.Context) {
	u := middlewares.GetUser(c)
	jsend.Success(c, http.StatusOK, u)
}
