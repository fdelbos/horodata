package groups

import (
	"encoding/json"
	"strconv"
	"strings"

	"dev.hyperboloide.com/fred/horodata/middlewares"
	sqlerrors "dev.hyperboloide.com/fred/horodata/models/errors"
	"dev.hyperboloide.com/fred/horodata/www/api/jsend"
	valid "github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

func stringDecimal2ToInt(str string) (int, bool) {
	chk := strings.Split(str, ".")
	if len(chk) > 2 {
		return 0, false
	} else if len(chk) == 1 {
		nb, err := strconv.Atoi(chk[0])
		return nb * 100, err == nil
	}

	if len(chk[1]) == 1 {
		chk[1] = chk[1] + "0"
	}

	if begin, err := strconv.Atoi(chk[0]); err != nil {
		return 0, false
	} else if end, err := strconv.Atoi(chk[1]); err != nil {
		return 0, false
	} else if end >= 100 {
		return 0, false
	} else {
		return (begin*100 + end), true
	}
}

func GuestAdd(c *gin.Context) {
	g := middlewares.GetGroup(c)

	if u, err := g.GetOwner(); err != nil {
		jsend.ErrorJson(c)
		return
	} else if ok, err := u.QuotaCanAddGuest(); err != nil {
		jsend.ErrorJson(c)
		return
	} else if !ok {
		jsend.Quota(c, "guests")
		return
	}

	var data struct {
		Email string `json:"email"`
		Admin bool   `json:"admin"`
		Rate  string `json:"rate"`
	}
	if err := json.NewDecoder(c.Request.Body).Decode(&data); err != nil {
		jsend.ErrorJson(c)
		return
	}

	errors := map[string]string{}

	if data.Email == "" {
		errors["email"] = "Ce champ est obligatoire."
	} else if len(data.Email) > 100 {
		errors["email"] = "Ce champ ne doit pas dépasser 100 caractères."
	} else if valid.IsEmail(data.Email) == false {
		errors["email"] = "Cette adresse email n'est pas valide."
	}

	rate, ok := stringDecimal2ToInt(data.Rate)
	if !ok {
		errors["rate"] = "Ce champ n'est pas valide."
	} else if rate < 0 {
		errors["rate"] = "Ce champ doit être supérieur ou égal à 0."
	} else if rate >= 1000000 {
		errors["rate"] = "Ce champ doit être inférieur 10000.00 ."
	}

	if len(errors) > 0 {
		jsend.BadRequest(c, errors)
		return
	}

	if err := g.GuestAdd(data.Email, rate, data.Admin, true); err != nil {
		jsend.Error(c, err)
	} else {
		jsend.Ok(c, nil)
	}
}

func GuestUpdate(c *gin.Context) {
	g := middlewares.GetGroup(c)
	id, err := strconv.ParseInt(c.Param("guestId"), 10, 64)
	if err != nil {
		jsend.BadRequest(c, nil)
		return
	}

	guest, err := g.GuestGetById(id)
	if err != nil {
		if err == sqlerrors.NotFound {
			jsend.NotFound(c)
		} else if err != nil {
			jsend.Error(c, err)
		}
		return
	}

	var data struct {
		Admin bool   `json:"admin"`
		Rate  string `json:"rate"`
	}
	if err := json.NewDecoder(c.Request.Body).Decode(&data); err != nil {
		jsend.ErrorJson(c)
		return
	}

	errors := map[string]string{}
	rate, ok := stringDecimal2ToInt(data.Rate)
	if !ok {
		errors["rate"] = "Ce champ n'est pas valide."
	} else if rate < 0 {
		errors["rate"] = "Ce champ doit être supérieur ou égal à 0."
	} else if rate >= 1000000 {
		errors["rate"] = "Ce champ doit être inférieur 10000.00 ."
	}

	if len(errors) > 0 {
		jsend.BadRequest(c, errors)
	} else {
		// IMPORTANT
		if guest.UserId != nil && *guest.UserId != g.OwnerId {
			guest.Admin = data.Admin
		}
		guest.Rate = rate
		if err := guest.Update(); err != nil {
			jsend.Error(c, err)
		} else {
			jsend.Ok(c, guest)
		}
	}
}

func GuestDelete(c *gin.Context) {
	g := middlewares.GetGroup(c)

	if id, err := strconv.ParseInt(c.Param("guestId"), 10, 64); err != nil {
		jsend.BadRequest(c, nil)
	} else if guest, err := g.GuestGetById(id); err == sqlerrors.NotFound {
		jsend.NotFound(c)
	} else if err != nil {
		jsend.Error(c, err)
	} else if guest.UserId != nil && g.OwnerId == *guest.UserId {
		jsend.Ok(c, nil)
	} else {
		guest.Active = false
		if err := guest.Update(); err != nil {
			jsend.Error(c, err)
		} else {
			jsend.Ok(c, nil)
		}
	}
}
