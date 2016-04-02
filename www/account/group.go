package account

import (
	"net/http"

	"dev.hyperboloide.com/fred/horodata/models/user"
	"dev.hyperboloide.com/fred/horodata/services/cookies"
	"dev.hyperboloide.com/fred/horodata/services/urls"
	"github.com/gin-gonic/gin"
)

func Group(r *gin.RouterGroup) {

	account := r.Group("/account")
	{
		account.GET("/register", GetRegister)
		account.POST("/register", PostRegister)

		account.GET("/login", GetLogin)
		account.POST("/login", PostLogin)

		account.GET("/logout", GetLogout)

		account.GET("/password_reset", GetResetStart)
		account.POST("/password_reset", PostResetStart)
		account.GET("/password_reset/:url", GetResetInput)
		account.POST("/password_reset/:url", PostResetInput)

		account.GET("provider/connect/:provider", ProviderConnect)
		account.GET("provider/callback/:provider", ProviderCallback)
		account.GET("/complete_registration", ProviderComplete)
	}
}

func LogUser(u *user.User, c *gin.Context) {
	if id, err := user.NewSession(u, c.Request.RemoteAddr); err != nil {
		GetError(c, err)
	} else if err := cookies.NewSession(id, c); err != nil {
		GetError(c, err)
	} else {
		// the hashtag is necessary to remove fb crap...
		c.Redirect(http.StatusTemporaryRedirect, urls.WWWApp+"#")
	}
}
