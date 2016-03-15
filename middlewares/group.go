package middlewares

import (
	"bitbucket.com/hyperboloide/horo/models/errors"
	"bitbucket.com/hyperboloide/horo/models/group"
	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GroupFilter() gin.HandlerFunc {

	return func(c *gin.Context) {
		u := GetUser(c)

		url := c.Params.ByName("group")
		if url == "" {
			log.WithFields(map[string]interface{}{
				"context": c,
			}).Fatal("no url definted for GroupFilter")
		}

		g, err := group.ByUrl(url)
		if err != nil {
			if err == errors.NotFound {
				c.JSON(404, http.StatusText(404))
			} else {
				c.JSON(500, http.StatusText(500))
			}
			c.Abort()
			return
		} else if u.Id != g.OwnerId {
			if _, err := g.GuestGetByUserId(u.Id); err != nil {
				if err == errors.NotFound {
					c.JSON(404, http.StatusText(404))
				} else {
					c.JSON(500, http.StatusText(500))
				}
				c.Abort()
				return
			}
		}
		c.Set("group", g)
		c.Next()
	}
}

func GroupOwnerFilter() gin.HandlerFunc {
	return func(c *gin.Context) {
		u := GetUser(c)
		g := GetGroup(c)
		if u.Id != g.OwnerId {
			c.JSON(403, http.StatusText(403))
			c.Abort()
		} else {
			c.Next()
		}
	}
}

func GetGroup(c *gin.Context) *group.Group {
	return c.MustGet("group").(*group.Group)
}
