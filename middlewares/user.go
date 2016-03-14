package middlewares

import (
	"bitbucket.com/hyperboloide/horo/models/errors"
	"bitbucket.com/hyperboloide/horo/models/user"
	"bitbucket.com/hyperboloide/horo/services/cookies"
	"bitbucket.com/hyperboloide/horo/services/urls"
	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserFilter() gin.HandlerFunc {

	unauthorizedError := func(c *gin.Context) {
		contentType := c.Request.Header.Get("Content-Type")
		if contentType == "application/json" {
			c.JSON(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
		} else {
			c.Redirect(http.StatusTemporaryRedirect, urls.WWWLogin)
		}
		c.Abort()
	}

	return func(c *gin.Context) {
		id, err := cookies.SessionGet("id", c)
		if err == errors.NotFound {
			unauthorizedError(c)
			return
		} else if err != nil {
			log.WithFields(log.Fields{
				"error": err,
			}).Error("Cannot get session id.")
			c.String(500, http.StatusText(500))
			c.Abort()
			return
		}

		session, err := user.GetSession(id.(string))
		if err == errors.NotFound {
			unauthorizedError(c)
			return
		} else if err != nil {
			log.WithFields(log.Fields{
				"error": err,
			}).Error("Cannot get session from model.")
			c.String(500, http.StatusText(500))
			c.Abort()
			return
		}

		if session.IsValid() == false {
			unauthorizedError(c)
			return
		}

		user, err := session.GetUser()
		if err == errors.NotFound {
			unauthorizedError(c)
			return
		} else if err != nil {
			log.WithFields(log.Fields{
				"error": err,
			}).Error("Cannot find user.")
			c.String(500, http.StatusText(500))
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Next()
	}
}

func GetUser(c *gin.Context) *user.User {
	return c.MustGet("user").(*user.User)
}

func GetUserMaybe(c *gin.Context) *user.User {
	if u, ok := c.Get("user"); !ok {
		return nil
	} else {
		return u.(*user.User)
	}
}
