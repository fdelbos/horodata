package middlewares

import (
	"net/http"

	"dev.hyperboloide.com/fred/horodata/services/cookies"
	"dev.hyperboloide.com/fred/horodata/services/urls"
	"github.com/gin-gonic/gin"
)

func PostCSRFFilter(c *gin.Context) {
	token := c.PostForm(cookies.CSRFField)
	if ok, err := cookies.ValidateCSRF(token, c); err != nil || !ok {
		c.String(400, http.StatusText(400))
		c.Abort()
	} else {
		c.Next()
	}
}

func AjaxCSRFFilter(c *gin.Context) {
	origin := c.Request.Header.Get("Origin")
	if origin == "" {
		switch c.Request.Method {
		case "POST", "PUT", "DELETE":
			if c.Request.Header.Get("X-Requested-With") != "XMLHttpRequest" {
				c.String(400, http.StatusText(400))
				c.Abort()
			} else {
				c.Next()
			}
		default:
			c.Next()
		}
	} else if origin != urls.WWWRoot {
		c.String(400, http.StatusText(400))
		c.Abort()
	} else {
		c.Next()
	}

}
