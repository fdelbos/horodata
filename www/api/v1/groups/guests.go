package groups

import (
	"bitbucket.com/hyperboloide/horo/middlewares"
	sqlerrors "bitbucket.com/hyperboloide/horo/models/errors"
	"bitbucket.com/hyperboloide/horo/www/api/jsend"
	"encoding/json"
	valid "github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GuestAdd(c *gin.Context) {
	g := middlewares.GetGroup(c)

	var data struct {
		Email string `json:"email"`
		Admin bool   `json:"admin"`
		Rate  int    `json:"rate"`
	}
	if err := json.NewDecoder(c.Request.Body).Decode(&data); err != nil {
		jsend.ErrorJson(c)
		return
	}

	errors := map[string]string{}

	if data.Email == "" {
		errors["email"] = "Ce champs est obligatoire."
	} else if len(data.Email) > 100 {
		errors["email"] = "Ce champ ne doit pas dépasser 100 caractères."
	} else if valid.IsEmail(data.Email) == false {
		errors["email"] = "Cette adresse email n'est pas valide."
	}

	if data.Rate < 0 {
		errors["rate"] = "Ce champ doit être supérieur ou égal a 0."
	}

	if len(errors) > 0 {
		jsend.BadRequest(c, errors)
		return
	}

	if err := g.GuestAdd(data.Email, data.Rate, data.Admin, false); err != nil {
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
	if err == sqlerrors.NotFound {
		jsend.NotFound(c)
	} else if err != nil {
		jsend.Error(c, err)
	}

	var data struct {
		Admin bool `json:"admin"`
		Rate  int  `json:"admin"`
	}
	if err := json.NewDecoder(c.Request.Body).Decode(&data); err != nil {
		jsend.ErrorJson(c)
		return
	}

	errors := map[string]string{}
	if data.Rate < 0 {
		errors["rate"] = "Ce champ doit être supérieur ou égal a 0."
	}

	if len(errors) > 0 {
		jsend.BadRequest(c, errors)
	} else {
		// IMPORTANT
		if *guest.UserId != g.OwnerId {
			guest.Admin = data.Admin
		}
		guest.Rate = data.Rate
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
