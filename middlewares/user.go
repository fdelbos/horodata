package middlewares

import (
	"net/http"

	"dev.hyperboloide.com/fred/horodata/models/errors"
	"dev.hyperboloide.com/fred/horodata/models/user"
	"dev.hyperboloide.com/fred/horodata/services/cookies"
	"dev.hyperboloide.com/fred/horodata/services/urls"
	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

func extractUser(c *gin.Context) (*user.User, error) {
	if id, err := cookies.SessionGet("id", c); err == errors.NotFound {
		return nil, nil
	} else if err != nil {
		log.WithFields(map[string]interface{}{
			"package":  "horodata.middlewares",
			"function": "func extractUser(c *gin.Context) (*user.User, error)",
			"step":     "cookies.SessionGet",
		}).Error(err)
		return nil, err
	} else if session, err := user.GetSession(id.(string)); err == errors.NotFound {
		return nil, nil
	} else if err != nil {
		log.WithFields(map[string]interface{}{
			"package":  "horodata.middlewares",
			"function": "func extractUser(c *gin.Context) (*user.User, error)",
			"step":     "user.GetSession",
		}).Error(err)
		return nil, err
	} else if session.IsValid() == false {
		return nil, nil
	} else if user, err := session.GetUser(); err == errors.NotFound {
		return nil, nil
	} else if err != nil {
		log.WithFields(map[string]interface{}{
			"package":  "horodata.middlewares",
			"function": "func extractUser(c *gin.Context) (*user.User, error)",
			"step":     "session.GetUser()",
		}).Error(err)
		return nil, err
	} else {
		return user, nil
	}
}

func UserFilter(c *gin.Context) {
	if u, err := extractUser(c); err != nil {
		c.String(500, http.StatusText(500))
		c.Abort()
	} else if u == nil {
		if contentType := c.Request.Header.Get("Content-Type"); contentType == "application/json" {
			c.JSON(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
		} else {
			c.Redirect(http.StatusTemporaryRedirect, urls.WWWLogin)
		}
		c.Abort()
	} else {
		c.Set("user", u)
		c.Next()
	}
}

func GetUser(c *gin.Context) *user.User {
	return c.MustGet("user").(*user.User)
}

func GetUserMaybe(c *gin.Context) *user.User {
	if u, ok := c.Get("user"); ok {
		return u.(*user.User)
	} else if u, _ := extractUser(c); u != nil {
		return u
	}
	return nil
}
