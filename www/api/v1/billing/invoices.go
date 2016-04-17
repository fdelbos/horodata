package billing

import (
	"fmt"
	"strconv"

	"dev.hyperboloide.com/fred/horodata/middlewares"
	"dev.hyperboloide.com/fred/horodata/models/billing"
	sqlerrors "dev.hyperboloide.com/fred/horodata/models/errors"
	"dev.hyperboloide.com/fred/horodata/www/api/jsend"
	"github.com/gin-gonic/gin"
)

func Invoices(c *gin.Context) {
	u := middlewares.GetUser(c)

	if pr, err := billing.InvoicePreviewsByUserId(u.Id); err != nil {
		jsend.Error(c, err)
	} else {
		jsend.Ok(c, pr)
	}
}

func GetInvoice(c *gin.Context) {
	u := middlewares.GetUser(c)

	if id, err := strconv.ParseInt(c.Param("id"), 10, 64); err != nil {
		jsend.BadRequest(c, nil)
	} else if invoice, err := billing.InvoiceById(id); err == sqlerrors.NotFound {
		jsend.NotFound(c)
	} else if err != nil {
		jsend.Error(c, err)
	} else if u.Id != invoice.UserId {
		jsend.Forbidden(c)
	} else {
		c.Writer.Header().Set("Content-Type", "application/pdf")
		c.Writer.Header().Set(
			"Content-Disposition",
			fmt.Sprintf("inline; filename=\"%s.pdf\"", invoice.FileId()))
		if err := invoice.Write(c.Writer); err != nil {
			jsend.Error(c, err)
		}
	}
}
