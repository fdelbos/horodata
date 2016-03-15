package www

import (
	"bitbucket.com/hyperboloide/horo/html"
	"bitbucket.com/hyperboloide/horo/middlewares"
	"bitbucket.com/hyperboloide/horo/services/urls"
	"bitbucket.com/hyperboloide/horo/www/account"
	"bitbucket.com/hyperboloide/horo/www/api"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Group(r *gin.RouterGroup) {
	www := r.Group("/www")
	{
		account.Group(www)
		api.Group(www)
		www.Any("/app/*all", middlewares.UserFilter(), GetApp)
	}
}

func GetApp(c *gin.Context) {
	data := map[string]interface{}{
		"base": urls.AngularBase,
		"api":  urls.ApiRoot,
	}
	html.Render("app.html", c, data, http.StatusOK)
}
