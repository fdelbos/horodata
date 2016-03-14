package settings

import (
	"bitbucket.com/hyperboloide/horo/html"
	"bitbucket.com/hyperboloide/horo/middlewares"
	"bitbucket.com/hyperboloide/horo/models/billing"
	"bitbucket.com/hyperboloide/horo/models/errors"
	"bitbucket.com/hyperboloide/horo/services/payment"
	valid "github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"net/http"
)

func setData(c *gin.Context, data map[string]interface{}) error {
	u := middlewares.GetUser(c)

	if addrCurrent, err := billing.CurrentAddress(u.Id); err != nil && err != errors.NotFound {
		return err
	} else if err == errors.NotFound {
		data["addr"] = billing.Address{
			FullName: u.FullName,
			Email:    u.Email,
		}
	} else {
		data["addr"] = addrCurrent
		data["addrCurrent"] = addrCurrent
		data["addrCurrentCountryCode"] = "country." + addrCurrent.CountryId
	}

	if countries, err := billing.CountriesList(middlewares.GetTranslateLanguage(c)); err != nil {
		return err
	} else {
		data["addrCountries"] = countries
	}

	customer, err := billing.GetCustomer(u.Id)
	if err != nil && err != errors.NotFound {
		return err
	} else if err == nil {
		if card, err := customer.GetCard(); err != nil {
			return err
		} else {
			data["card"] = card
		}
	}
	return nil
}

func GetBilling(c *gin.Context) {
	data := map[string]interface{}{
		"page": "billing",
		"payment_publishable_key": payment.PublishableKey(),
	}

	if err := setData(c, data); err != nil {
		html.ErrorServer(c, err)
	} else {
		html.Render("billing.html", c, data, http.StatusOK)
	}
}

func PostBilling(c *gin.Context) {
	u := middlewares.GetUser(c)
	data := map[string]interface{}{
		"page": "billing",
		"payment_publishable_key": payment.PublishableKey(),
	}

	if err := setData(c, data); err != nil {
		html.ErrorServer(c, err)
		return
	}

	if c.PostForm("action") == "address" {
		addr, errors := validateAddr(c)
		if len(errors) != 0 {
			data["addr"] = addr
			data["addrErrors"] = errors
			html.Render("billing.html", c, data, http.StatusBadRequest)
		} else if err := addr.Insert(); err != nil {
			html.ErrorServer(c, err)
		} else if addr, err := billing.CurrentAddress(u.Id); err != nil {
			html.ErrorServer(c, err)
		} else {
			data["addr"] = addr
			data["addrCurrent"] = addr
			data["addrCurrentCountryCode"] = "country." + addr.CountryId
			html.Render("billing.html", c, data, http.StatusOK)
		}
	} else {
		token := c.PostForm("token")
		customer, err := billing.GetCustomer(u.Id)
		if err == errors.NotFound {
			if err := billing.NewCustomer(u.Id, token); err != nil {
				html.ErrorServer(c, err)
				return
			}
		} else if err != nil {
			html.ErrorServer(c, err)
		} else if err := customer.UpdateCard(token); err != nil {
			html.ErrorServer(c, err)
		}
		if customer, err := billing.GetCustomer(u.Id); err != nil {
			html.ErrorServer(c, err)
		} else if card, err := customer.GetCard(); err != nil {
			html.ErrorServer(c, err)
		} else {
			data["card"] = card
			html.Render("billing.html", c, data, http.StatusOK)
		}
	}
}

func validateAddr(c *gin.Context) (*billing.Address, map[string]string) {
	T := middlewares.GetTranslate(c)
	u := middlewares.GetUser(c)
	a := &billing.Address{
		UserId: u.Id,
	}
	errors := map[string]string{}

	a.FullName = c.PostForm("full_name")
	if a.FullName == "" {
		errors["full_name"] = T("generic.form.error_mandatory")
	} else if len(a.FullName) > 100 {
		errors["full_name"] = T("generic.form.error_too_large", 100)
	}

	a.Email = c.PostForm("email")
	if a.Email == "" {
		errors["email"] = T("generic.form.error_mandatory")
	} else if len(a.Email) > 100 {
		errors["email"] = T("generic.error_too_large", 100)
	} else if valid.IsEmail(a.Email) == false {
		errors["email"] = T("generic.form.error_email_invalid")
	}

	a.Company = c.PostForm("company")
	if len(a.Company) > 100 {
		errors["company"] = T("generic.form.error_too_large", 100)
	}

	a.VAT = c.PostForm("vat")
	if len(a.VAT) > 100 {
		errors["company"] = T("generic.form.error_too_large", 100)
	}

	a.Address1 = c.PostForm("address1")
	if a.Address1 == "" {
		errors["address1"] = T("generic.form.error_mandatory")
	} else if len(a.Address1) > 200 {
		errors["address1"] = T("generic.form.error_too_large", 200)
	}

	a.Address2 = c.PostForm("address2")
	if len(a.Address1) > 200 {
		errors["address2"] = T("generic.form.error_too_large", 200)
	}

	a.City = c.PostForm("city")
	if a.City == "" {
		errors["city"] = T("generic.form.error_mandatory")
	} else if len(a.City) > 150 {
		errors["address2"] = T("generic.form.error_too_large", 150)
	}

	a.Zip = c.PostForm("zip")
	if a.Zip == "" {
		errors["zip"] = T("generic.form.error_mandatory")
	} else if len(a.Zip) > 15 {
		errors["zip"] = T("generic.form.error_too_large", 15)
	}

	a.State = c.PostForm("state")
	if len(a.State) > 100 {
		errors["state"] = T("generic.form.error_too_large", 100)
	}

	a.CountryId = c.PostForm("country")
	if len(a.CountryId) != 2 {
		errors["country"] = T("generic.form.error_invalid")
	}

	return a, errors
}
