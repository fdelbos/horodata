package account

import (
	"bitbucket.com/hyperboloide/horo/models/user"
	"bitbucket.com/hyperboloide/horo/services/cookies"
	"bitbucket.com/hyperboloide/horo/services/urls"
	"github.com/gin-gonic/gin"
	"net/http"
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

		account.GET("provider/:provider", Provider)
		account.GET("provider/:provider/callback", ProviderCallback)
		account.GET("/complete_registration", ProviderComplete)
	}
}

func LogUser(u *user.User, c *gin.Context) {
	if id, err := user.NewSession(u, c.Request.RemoteAddr); err != nil {
		GetError(c, err)
	} else if err := cookies.NewSession(id, c); err != nil {
		GetError(c, err)
	} else {
		c.Redirect(http.StatusTemporaryRedirect, urls.WWWApp)
	}
}
