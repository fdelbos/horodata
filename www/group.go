package www

import (
	"dev.hyperboloide.com/fred/horodata/html"
	"dev.hyperboloide.com/fred/horodata/middlewares"
	"dev.hyperboloide.com/fred/horodata/services/urls"
	"dev.hyperboloide.com/fred/horodata/www/account"
	"dev.hyperboloide.com/fred/horodata/www/api"
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
