package www

import (
	"encoding/json"
	"fmt"
	"net/http"

	"dev.hyperboloide.com/fred/horodata/html"
	"dev.hyperboloide.com/fred/horodata/middlewares"
	"dev.hyperboloide.com/fred/horodata/services/payment"
	"dev.hyperboloide.com/fred/horodata/services/urls"
	"dev.hyperboloide.com/fred/horodata/www/account"
	"dev.hyperboloide.com/fred/horodata/www/api"
	"dev.hyperboloide.com/fred/horodata/www/api/jsend"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Group(r *gin.RouterGroup) {
	www := r.Group("/www")
	{
		account.Group(www)
		api.Group(www)
		www.Any("/app/*all", middlewares.UserFilter(), GetApp)

		if gin.IsDebugging() {
			www.Static("/profiles", viper.GetString("profile_pictures"))
		}

		endpoint := fmt.Sprintf("/%s", viper.GetString("payment_endpoint"))
		www.POST(endpoint, StripeEndpoint)
	}
}

func GetApp(c *gin.Context) {
	data := map[string]interface{}{
		"base": urls.AngularBase,
		"api":  urls.ApiRoot,
	}
	html.Render("app.html", c, data, http.StatusOK)
}

func StripeEndpoint(c *gin.Context) {
	data := &payment.StripeEvent{}
	if err := json.NewDecoder(c.Request.Body).Decode(&data); err != nil {
		jsend.ErrorJson(c)
	} else if err := payment.NewEvent(data.Id); err != nil {
		jsend.Error(c, err)
	} else {
		jsend.Ok(c, nil)
	}
}
