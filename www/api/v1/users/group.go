package users

import (
	"bitbucket.com/hyperboloide/horo/middlewares"
	"bitbucket.com/hyperboloide/horo/www/api/jsend"
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
