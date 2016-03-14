package middlewares

import (
	"bitbucket.com/hyperboloide/horo/services/cookies"
	"github.com/gin-gonic/gin"
	"net/http"
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
