package groups

import (
	"dev.hyperboloide.com/fred/horodata/middlewares"
	"dev.hyperboloide.com/fred/horodata/www/api/jsend"
	"github.com/gin-gonic/gin"
)

func StatsCustomerTime(c *gin.Context) {
	g := middlewares.GetGroup(c)

	begin, end, errors, err := extractTime(c)
	if err != nil {
		jsend.Error(c, err)
		return
	} else if len(errors) > 0 {
		jsend.BadRequest(c, errors)
		return
	}

	guestId, errors, err := extractGuestId(c)
	if err != nil {
		jsend.Error(c, err)
		return
	} else if len(errors) > 0 {
		jsend.BadRequest(c, errors)
		return
	}

	if res, err := g.StatsCustomerTime(begin, end, guestId); err != nil {
		jsend.Error(c, err)
	} else {
		jsend.Ok(c, res)
	}
}
