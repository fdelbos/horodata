package middlewares

import (
	"net/http"

	"dev.hyperboloide.com/fred/horodata/models/errors"
	"dev.hyperboloide.com/fred/horodata/models/group"
	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

// GroupFilter makes sure group exists and that the user is part of this group
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

// GroupAdminFilter makes sure that the user is a group admin
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

// GroupOwnerFilter makes sure the user is the owner of the group
func GroupOwnerFilter() gin.HandlerFunc {
	return func(c *gin.Context) {
		guest := GetGuest(c)
		group := GetGroup(c)
		if *guest.UserId != group.OwnerId {
			c.JSON(403, http.StatusText(403))
			c.Abort()
		} else {
			c.Next()
		}
	}
}

// GetGroup returns the Group
func GetGroup(c *gin.Context) *group.Group {
	return c.MustGet("group").(*group.Group)
}

// GetGuest returns the Guest
func GetGuest(c *gin.Context) *group.Guest {
	return c.MustGet("guest").(*group.Guest)
}
