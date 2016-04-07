package groups

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"dev.hyperboloide.com/fred/horodata/middlewares"
	"dev.hyperboloide.com/fred/horodata/services/urls"
	"dev.hyperboloide.com/fred/horodata/www/api/jsend"
	"github.com/gin-gonic/gin"
)

func ExportCSV(c *gin.Context) {
	g := middlewares.GetGroup(c)

	begin, end, errors, err := extractTime(c)
	if err != nil {
		jsend.Error(c, err)
	} else if len(errors) > 0 {
		jsend.BadRequest(c, errors)
	} else {
		c.Writer.Header().Set("Content-Type", "text/csv;charset=UTF-8")
		c.Writer.Header().Set(
			"Content-Disposition",
			fmt.Sprintf("inline; filename=\"export_%s.csv\"", g.Name))
		if err := g.ExportCSV(c.Writer, begin, end); err != nil {
			jsend.Error(c, err)
		}
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
