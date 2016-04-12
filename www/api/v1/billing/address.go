package billing

import (
	"encoding/json"
	"strings"

	"dev.hyperboloide.com/fred/horodata/middlewares"
	"dev.hyperboloide.com/fred/horodata/models/billing"
	sqlerrors "dev.hyperboloide.com/fred/horodata/models/errors"
	"dev.hyperboloide.com/fred/horodata/www/api/jsend"
	valid "github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

func GetAddress(c *gin.Context) {
	user := middlewares.GetUser(c)

	if addr, err := billing.CurrentAddress(user.Id); err == sqlerrors.NotFound {
		jsend.NotFound(c)
	} else if err != nil {
		jsend.Error(c, err)
	} else {
		jsend.Ok(c, addr)
	}
}

func NewAddress(c *gin.Context) {
	user := middlewares.GetUser(c)

	data := &billing.Address{}
	if err := json.NewDecoder(c.Request.Body).Decode(&data); err != nil {
		jsend.ErrorJson(c)
		return
	}
	data.UserId = user.Id

	errors := map[string]string{}

	data.Name = strings.TrimSpace(data.Name)
	if data.Name == "" {
		errors["name"] = "Ce champ est obligatoire."
	} else if len(data.Name) < 4 {
		errors["name"] = "Ce champ doit faire au moins 4 caractères."
	} else if len(data.Name) > 50 {
		errors["name"] = "Ce champ ne doit pas dépasser 50 caractères."
	}

	data.Email = strings.TrimSpace(data.Email)
	if data.Email == "" {
		errors["email"] = "Ce champ est obligatoire."
	} else if len(data.Email) > 100 {
		errors["email"] = "Ce champ ne doit pas dépasser 100 caractères."
	} else if valid.IsEmail(data.Email) == false {
		errors["email"] = "Cette adresse email n'est pas valide."
	}

	data.Company = strings.TrimSpace(data.Company)
	if len(data.Company) > 100 {
		errors["company"] = "Ce champ ne doit pas dépasser 100 caractères."
	}

	data.VAT = strings.TrimSpace(data.VAT)
	if len(data.VAT) > 25 {
		errors["vat"] = "Ce champ ne doit pas dépasser 25 caractères."
	}

	data.Address1 = strings.TrimSpace(data.Address1)
	if data.Address1 == "" {
		errors["addr1"] = "Ce champ est obligatoire."
	} else if len(data.Address1) > 150 {
		errors["addr1"] = "Ce champ ne doit pas dépasser 150 caractères."
	}

	data.Address2 = strings.TrimSpace(data.Address2)
	if len(data.Address2) > 150 {
		errors["addr2"] = "Ce champ ne doit pas dépasser 150 caractères."
	}

	data.City = strings.TrimSpace(data.City)
	if data.City == "" {
		errors["city"] = "Ce champ est obligatoire."
	} else if len(data.City) > 100 {
		errors["city"] = "Ce champ ne doit pas dépasser 100 caractères."
	}

	data.Zip = strings.TrimSpace(data.Zip)
	if data.Zip == "" {
		errors["zip"] = "Ce champ est obligatoire."
	} else if len(data.Zip) > 100 {
		errors["zip"] = "Ce champ ne doit pas dépasser 6 caractères."

	}

	if len(errors) > 0 {
		jsend.BadRequest(c, errors)
	} else if err := data.Insert(); err != nil {
		jsend.Error(c, err)
	} else if addr, err := billing.CurrentAddress(user.Id); err != nil {
		jsend.Error(c, err)
	} else {
		jsend.Ok(c, addr)
	}
}
