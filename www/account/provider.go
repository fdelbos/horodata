package account

import (
	"bitbucket.com/hyperboloide/horo/models/errors"
	sqlerrors "bitbucket.com/hyperboloide/horo/models/errors"
	"bitbucket.com/hyperboloide/horo/models/user"
	"bitbucket.com/hyperboloide/horo/services/cookies"
	"bitbucket.com/hyperboloide/horo/services/oauth"
	"bitbucket.com/hyperboloide/horo/services/urls"
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"net/http"
)

func Provider(c *gin.Context) {
	gothic.BeginAuthHandler(c.Writer, c.Request)
}

func ProviderCallback(c *gin.Context) {
	ru, err := gothic.CompleteUserAuth(c.Writer, c.Request)
	if err != nil {
		GetError(c, err)
		return
	} else if !oauth.IsVerified(ru) {
		UserNotVerified(c)
		return
	}

	if err := cookies.Clear("_gothic_session", c); err != nil {
		GetError(c, err)
		return
	}

	u, err := user.ByEmail(ru.Email)
	if err == errors.NotFound {
		if err := cookies.Set("session", "provider", ru, c); err != nil {
			GetError(c, err)
		} else {
			c.Redirect(http.StatusTemporaryRedirect, urls.WWWComplete)
		}
	} else if err != nil {
		GetError(c, err)
	} else if !u.Active {
		c.Redirect(http.StatusTemporaryRedirect, urls.WWWLogin)
	} else {
		LogUser(u, c)
	}
}

func ProviderComplete(c *gin.Context) {
	tmp, err := cookies.Get("session", "provider", c)
	if err != nil {
		if err == sqlerrors.NotFound {
			c.Redirect(http.StatusTemporaryRedirect, urls.WWWLogin)
		} else {
			GetError(c, err)
		}
		return
	}
	guser := tmp.(goth.User)
	u := &user.User{}
	u.Active = true
	u.Email = guser.Email
	u.FullName = guser.Name

	if err := u.Insert(); err != nil {
		GetError(c, err)
	} else if err := u.SendWelcome(); err != nil {
		GetError(c, err)
	} else if err := cookies.Delete("session", "provider", c); err != nil {
		GetError(c, err)
	} else {
		LogUser(u, c)
	}
}
