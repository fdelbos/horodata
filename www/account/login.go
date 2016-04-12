package account

import (
	"dev.hyperboloide.com/fred/horodata/html"
	sqlerrors "dev.hyperboloide.com/fred/horodata/models/errors"
	"dev.hyperboloide.com/fred/horodata/models/user"
	valid "github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetLogin(c *gin.Context) {
	html.Render("account/login.html", c, nil, http.StatusOK)
}

func PostLogin(c *gin.Context) {
	errors := map[string]string{}

	email := c.PostForm("email")
	if email == "" {
		errors["email"] = "Ce champ est obligatoire."
	} else if len(email) > 100 {
		errors["email"] = "Ce champ ne doit pas dépasser 100 caractères."
	} else if valid.IsEmail(email) == false {
		errors["email"] = "Cette adresse email n'est pas valide."
	}

	password := c.PostForm("password")
	if password == "" {
		errors["password"] = "Ce champ est obligatoire."
	} else if len(password) > 100 {
		errors["password"] = "Ce champ ne doit pas dépasser 100 caractères."
	}

	if len(errors) == 0 {
		u, err := user.ByEmail(email)
		if err == sqlerrors.NotFound {
			errors["denied"] = "true"
		} else if err != nil {
			GetError(c, err)
			return
		} else if ok, err := u.CheckPassword(password); err != nil {
			GetError(c, err)
			return
		} else if !ok {
			errors["denied"] = "true"
		} else {
			LogUser(u, c)
			return
		}
	}

	c.Writer.WriteHeader(http.StatusBadRequest)
	data := map[string]interface{}{
		"errors": errors,
		"email":  email,
	}
	html.Render("account/login.html", c, data, http.StatusBadRequest)
}
