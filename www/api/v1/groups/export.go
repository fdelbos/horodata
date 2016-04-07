package groups

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"dev.hyperboloide.com/fred/horodata/middlewares"
	sqlerrors "dev.hyperboloide.com/fred/horodata/models/errors"
	"dev.hyperboloide.com/fred/horodata/services/urls"
	"dev.hyperboloide.com/fred/horodata/www/api/jsend"
	"github.com/gin-gonic/gin"
)

// http://localhost:3000/www/api/v1/groups/kiki/export?begin=2016-02-23&end=2016-03-23

type exportParams struct {
	begin      time.Time
	end        time.Time
	guestId    *int64
	customerId *int64
}

func validateParams(c *gin.Context) (*exportParams, error, map[string]string) {
	g := middlewares.GetGroup(c)

	begin, end, errors, err := extractTime(c)
	if err != nil {
		return nil, err, nil
	} else if len(errors) > 0 {
		return nil, nil, errors
	}

	guestId, errors, err := extractGuestId(c)
	if err != nil {
		return nil, err, nil
	} else if len(errors) > 0 {
		return nil, nil, errors
	}

	var customerId *int64
	if str := c.Query("customer"); str == "" {
		customerId = nil
	} else if i, err := strconv.ParseInt(str, 10, 64); err != nil {
		errors["customer"] = "Ce champ n'est pas valide."
	} else if _, err := g.CustomerGet(i); err == sqlerrors.NotFound {
		errors["customer"] = "Ce champ n'est pas valide."
	} else if err != nil {
		return nil, err, nil
	} else {
		customerId = &i
	}

	if len(errors) > 0 {
		return nil, nil, errors
	}
	return &exportParams{begin, end, guestId, customerId}, nil, nil
}

func ExportCSV(c *gin.Context) {
	g := middlewares.GetGroup(c)

	c.Writer.Header().Set("Content-Type", "text/csv;charset=UTF-8")
	c.Writer.Header().Set(
		"Content-Disposition",
		fmt.Sprintf("inline; filename=\"export_%s.csv\"", g.Name))

	if p, err, errors := validateParams(c); err != nil {
		jsend.Error(c, err)
	} else if errors != nil {
		jsend.BadRequest(c, errors)
	} else if err := g.ExportCSV(c.Writer, p.begin, p.end, p.customerId, p.guestId); err != nil {
		jsend.Error(c, err)
	}
}

func ExportXLSX(c *gin.Context) {
	g := middlewares.GetGroup(c)

	begin, end, errors, err := extractTime(c)
	if err != nil {
		jsend.Error(c, err)
	} else if len(errors) > 0 {
		jsend.BadRequest(c, errors)
	} else if data, err := g.ExportStruct(begin, end); err != nil {
		jsend.Error(c, err)
	} else {
		c.Writer.Header().Set(
			"Content-Type",
			"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
		c.Writer.Header().Set(
			"Content-Disposition",
			fmt.Sprintf("inline; filename=\"export_%s.xlsx\"", g.Name))

		if json, err := json.Marshal(data); err != nil {
			jsend.Error(c, err)
		} else if resp, err := http.Post(urls.ExportService, "application/json", bytes.NewBuffer(json)); err != nil {
			jsend.Error(c, err)
		} else if _, err := io.Copy(c.Writer, resp.Body); err != nil {
			jsend.Error(c, err)
		}
	}
}
