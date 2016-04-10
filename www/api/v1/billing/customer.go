package billing

import (
	"encoding/json"

	"dev.hyperboloide.com/fred/horodata/middlewares"
	"dev.hyperboloide.com/fred/horodata/models/billing"
	sqlerrors "dev.hyperboloide.com/fred/horodata/models/errors"
	"dev.hyperboloide.com/fred/horodata/services/payment"
	"dev.hyperboloide.com/fred/horodata/www/api/jsend"
	"github.com/gin-gonic/gin"
)

func StripeKey(c *gin.Context) {
	jsend.Ok(c, payment.PublishableKey())
}

func GetCard(c *gin.Context) {
	user := middlewares.GetUser(c)

	if customer, err := billing.GetCustomer(user.Id); err == sqlerrors.NotFound {
		jsend.NotFound(c)
	} else if err != nil {
		jsend.Error(c, err)
	} else if card, err := customer.GetCard(); err == sqlerrors.NotFound {
		jsend.NotFound(c)
	} else if err != nil {
		jsend.Error(c, err)
	} else {
		jsend.Ok(c, card)
	}
}

func NewCard(c *gin.Context) {
	user := middlewares.GetUser(c)

	var data struct {
		Token string `json:"token"`
	}

	if err := json.NewDecoder(c.Request.Body).Decode(&data); err != nil {
		jsend.ErrorJson(c)
	} else if customer, err := billing.GetCustomer(user.Id); err == sqlerrors.NotFound {

		if err := billing.NewCustomer(user.Id, data.Token); err != nil {
			jsend.Error(c, err)
		} else {
			GetCard(c)
		}

	} else if err != nil {
		jsend.Error(c, err)
	} else if err := customer.UpdateCard(data.Token); err != nil {
		jsend.Error(c, err)
	} else {
		GetCard(c)
	}
}
