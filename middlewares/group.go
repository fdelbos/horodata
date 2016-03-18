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

		if url := c.Params.ByName("group"); url == "" {
			log.WithFields(map[string]interface{}{
				"context": c,
			}).Fatal("no url definted for GroupFilter")
		} else if g, err := group.ByUrl(url); err == errors.NotFound {
			c.JSON(404, http.StatusText(404))
		} else if err != nil {
			log.Error(err)
			c.JSON(500, http.StatusText(500))
		} else if !g.Active {
			c.JSON(404, http.StatusText(404))
		} else if guest, err := g.GuestGetByUserId(u.Id); err == errors.NotFound {
			c.JSON(404, http.StatusText(404))
		} else if err != nil {
			log.Error(err)
			c.JSON(500, http.StatusText(500))
		} else {
			c.Set("guest", guest)
			c.Set("group", g)
			c.Next()
			return
		}
		c.Abort()
	}
}

func GroupAdminFilter() gin.HandlerFunc {
	return func(c *gin.Context) {
		guest := GetGuest(c)
		if !guest.Admin {
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

func GetGuest(c *gin.Context) *group.Guest {
	return c.MustGet("guest").(*group.Guest)
}
