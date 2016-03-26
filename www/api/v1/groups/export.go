package groups

import (
	"bitbucket.com/hyperboloide/horo/middlewares"
	sqlerrors "bitbucket.com/hyperboloide/horo/models/errors"
	"bitbucket.com/hyperboloide/horo/www/api/jsend"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

// http://localhost:3000/www/api/v1/groups/kiki/export?begin=2016-02-23&end=2016-03-23

func ExportCSV(c *gin.Context) {
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

	var customerId *int64
	if str := c.Query("customer"); str == "" {
		customerId = nil
	} else if i, err := strconv.ParseInt(str, 10, 64); err != nil {
		errors["customer"] = "Ce champ n'est pas valide."
	} else if _, err := g.CustomerGet(i); err == sqlerrors.NotFound {
		errors["customer"] = "Ce champ n'est pas valide."
	} else if err != nil {
		jsend.Error(c, err)
		return
	} else {
		customerId = &i
	}

	if len(errors) > 0 {
		jsend.BadRequest(c, errors)
		return
	}

	c.Writer.Header().Set("Content-Type", "text/csv;charset=UTF-8")
	c.Writer.Header().Set(
		"Content-Disposition",
		fmt.Sprintf("inline; filename=\"export_%s.csv\"", g.Name))
	if err := g.ExportCSV(c.Writer, begin, end, customerId, guestId); err != nil {
		jsend.Error(c, err)
	}
}

func ExportXLSX(c *gin.Context) {
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

	var customerId *int64
	if str := c.Query("customer"); str == "" {
		customerId = nil
	} else if i, err := strconv.ParseInt(str, 10, 64); err != nil {
		errors["customer"] = "Ce champ n'est pas valide."
	} else if _, err := g.CustomerGet(i); err == sqlerrors.NotFound {
		errors["customer"] = "Ce champ n'est pas valide."
	} else if err != nil {
		jsend.Error(c, err)
		return
	} else {
		customerId = &i
	}

	if len(errors) > 0 {
		jsend.BadRequest(c, errors)
		return
	}

	c.Writer.Header().Set(
		"Content-Type",
		"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Writer.Header().Set(
		"Content-Disposition",
		fmt.Sprintf("inline; filename=\"export_%s.xlsx\"", g.Name))
	if err := g.ExportXLSX(c.Writer, begin, end, customerId, guestId); err != nil {
		jsend.Error(c, err)
	}
}
