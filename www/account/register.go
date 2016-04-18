package account

import (
	"net/http"

	"dev.hyperboloide.com/fred/horodata/html"
	sqlerrors "dev.hyperboloide.com/fred/horodata/models/errors"
	"dev.hyperboloide.com/fred/horodata/models/user"
	"dev.hyperboloide.com/fred/horodata/services/captcha"
	valid "github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

func GetRegister(c *gin.Context) {
	data := map[string]interface{}{"CaptchaPubKey": captcha.PubKey}
	html.Render("account/register.html", c, data, http.StatusOK)
}

func PostRegister(c *gin.Context) {
	errors := map[string]string{}

	email := c.PostForm("email")
	if email == "" {
		errors["email"] = "Ce champ est obligatoire."
	} else if len(email) > 100 {
		errors["email"] = "Ce champ ne doit pas dépasser 100 caractères."
	} else if valid.IsEmail(email) == false {
		errors["email"] = "Cette adresse email n'est pas valide."
	} else if _, err := user.ByEmail(email); err == nil {
		errors["email"] = "Cette adresse email est déjà utilisée par un autre compte."
	} else if err != sqlerrors.NotFound {
		GetError(c, err)
		return
	}

	password := c.PostForm("password")
	if password == "" {
		errors["password"] = "Ce champ est obligatoire."
	} else if len(password) < 6 {
		errors["password"] = "Ce champ doit faire au moins 6 caractères."
	} else if len(password) > 100 {
		errors["password"] = "Ce champ ne doit pas dépasser 100 caractères."
	}

	fullName := c.PostForm("full_name")
	if fullName == "" {
		errors["full_name"] = "Ce champ est obligatoire."
	} else if len(fullName) < 4 {
		errors["full_name"] = "Ce champ doit faire au moins 4 caractères."
	} else if len(fullName) > 50 {
		errors["full_name"] = "Ce champ ne doit pas dépasser 50 caractères."
	}

	accept := c.PostForm("accept")
	if accept != "yes" {
		errors["accept"] = "Ce champ est obligatoire."
	}

	recaptcha := c.PostForm("g-recaptcha-response")
	if recaptcha == "" {
		errors["recaptcha"] = "Ce champ est obligatoire."
	} else if ok, err := captcha.Validate(recaptcha); err != nil {
		GetError(c, err)
		return
	} else if !ok {
		errors["recaptcha"] = "Ce champ est n'est pas valide."
	}

	if len(errors) != 0 {
		data := map[string]interface{}{
			"Title":         "Inscription",
			"errors":        errors,
			"full_name":     fullName,
			"email":         email,
			"CaptchaPubKey": captcha.PubKey,
		}
		html.Render("account/register.html", c, data, http.StatusBadRequest)
	} else {
		u := &user.User{}
		u.Active = true
		u.FullName = fullName
		u.Email = email

		if err := u.Insert(); err != nil {
			GetError(c, err)
		} else if err := u.UpdatePassword(password); err != nil {
			GetError(c, err)
		} else if err := u.SendWelcome(); err != nil {
			GetError(c, err)
		} else {
			LogUser(u, c)
		}
	}
}
