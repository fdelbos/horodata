package settings

import (
	"bitbucket.com/hyperboloide/horo/html"
	"bitbucket.com/hyperboloide/horo/middlewares"
	"bitbucket.com/hyperboloide/horo/models/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUsage(c *gin.Context) {
	u := middlewares.GetUser(c)

	if quota, err := u.GetQuota(); err != nil {
		html.ErrorServer(c, err)
	} else if usage, err := u.GetUsage(); err != nil {
		html.ErrorServer(c, err)
	} else {
		data := map[string]interface{}{
			"page":  "usage",
			"quota": quota,
			"usage": usage,
			"percents": user.Limits{
				Instances: usage.Instances / quota.Instances * 100,
				Forms:     usage.Forms / quota.Forms * 100,
				Roles:     usage.Roles / quota.Roles * 100,
				Files:     usage.Files / quota.Files * 100,
			},
		}
		html.Render("usage.html", c, data, http.StatusOK)
	}
}
