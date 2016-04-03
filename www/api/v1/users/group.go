package users

import (
	"net/http"

	"dev.hyperboloide.com/fred/horodata/middlewares"
	"dev.hyperboloide.com/fred/horodata/models/user"
	"dev.hyperboloide.com/fred/horodata/www/api/jsend"
	"github.com/gin-gonic/gin"
)

func Group(r *gin.RouterGroup) {
	users := r.Group("/users")
	{
		users.GET("/me", Me)
		users.GET("/me/quotas", Quota)
	}
}

func Me(c *gin.Context) {
	u := middlewares.GetUser(c)
	jsend.Success(c, http.StatusOK, u)
}

func Quota(c *gin.Context) {
	u := middlewares.GetUser(c)

	if quotas, err := u.Quotas(); err != nil {
		jsend.Error(c, err)
	} else if uGroups, err := u.UsageGroups(); err != nil {
		jsend.Error(c, err)
	} else if uGests, err := u.UsageGuests(); err != nil {
		jsend.Error(c, err)
	} else if uJobs, err := u.UsageJobs(); err != nil {
		jsend.Error(c, err)
	} else {
		res := struct {
			Quotas *user.Quotas `json:"quotas"`
			Usage  user.Limits  `json:"usage"`
		}{
			quotas,
			user.Limits{
				Jobs:   uJobs,
				Guests: uGests,
				Groups: uGroups,
			},
		}
		jsend.Ok(c, res)
	}
}
