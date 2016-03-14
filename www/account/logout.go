package account

import (
	"bitbucket.com/hyperboloide/horo/models/user"
	"bitbucket.com/hyperboloide/horo/services/cookies"
	"bitbucket.com/hyperboloide/horo/services/urls"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetLogout(c *gin.Context) {
	id, err := cookies.SessionGet("id", c)
	if err == nil {
		if s, err := user.GetSession(id.(string)); err != nil {
			GetError(c, err)
			return
		} else if err := s.Close(); err != nil {
			GetError(c, err)
			return
		}
	}

	_ = cookies.SessionClear(c)
	c.Redirect(http.StatusTemporaryRedirect, urls.WWWRoot)
}
