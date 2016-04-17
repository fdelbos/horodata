package billing

import (
	"encoding/json"
	"time"

	"dev.hyperboloide.com/fred/horodata/middlewares"
	"dev.hyperboloide.com/fred/horodata/models/billing"
	"dev.hyperboloide.com/fred/horodata/models/errors"
	"dev.hyperboloide.com/fred/horodata/models/user"
	"dev.hyperboloide.com/fred/horodata/www/api/jsend"
	"github.com/gin-gonic/gin"
)

func ChangePlan(c *gin.Context) {
	u := middlewares.GetUser(c)

	var data struct {
		Plan string `json:"plan"`
	}

	if err := json.NewDecoder(c.Request.Body).Decode(&data); err != nil {
		jsend.ErrorJson(c)
		return
	} else {
		switch data.Plan {
		case user.QuotaFree:
		case user.QuotaSmall:
		case user.QuotaMedium:
		case user.QuotaLarge:
		default:
			jsend.BadRequest(c, map[string]string{"plan": "Plan invalide"})
			return
		}
	}

	cust, err := billing.CustomerByUserId(u.Id)
	if err == errors.NotFound {
		jsend.Forbidden(c)
		return
	} else if err != nil {
		jsend.Error(c, err)
		return
	}

	sub, err := cust.Subscription()
	if err == errors.NotFound {
		sub = &billing.Subscription{
			UserId: u.Id,
		}
	} else if err != nil {
		jsend.Error(c, err)
		return
	}

	if err := sub.Update(data.Plan); err != nil {
		jsend.Error(c, err)
	} else {
		jsend.Ok(c, data)
	}
}

func GetPlan(c *gin.Context) {
	u := middlewares.GetUser(c)

	if s, err := billing.SubscriptionByUserId(u.Id); err == errors.NotFound {
		jsend.Ok(c, struct {
			Plan string `json:"plan"`
		}{"free"})
	} else if err != nil {
		jsend.Error(c, err)
	} else {
		jsend.Ok(c, struct {
			Plan string     `json:"plan"`
			End  *time.Time `json:"end,omitempty"`
		}{s.Plan, s.End})
	}
}

func GetEndPeriod(c *gin.Context) {
	u := middlewares.GetUser(c)

	if sub, err := billing.SubscriptionByUserId(u.Id); err == errors.NotFound {
		jsend.Error(c, err)
	} else if end, err := sub.EndPeriod(); err != nil {
		jsend.Error(c, err)
	} else {
		jsend.Ok(c, struct {
			End *time.Time `json:"end"`
		}{end})
	}
}
