package account

import (
	"dev.hyperboloide.com/fred/horodata/html"
	sqlerrors "dev.hyperboloide.com/fred/horodata/models/errors"
	"dev.hyperboloide.com/fred/horodata/models/user"
	"dev.hyperboloide.com/fred/horodata/services/captcha"
	valid "github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetResetStart(c *gin.Context) {
	data := map[string]interface{}{"CaptchaPubKey": captcha.PubKey}
	html.Render("account/password_reset_start.html", c, data, http.StatusOK)
}

func PostResetStart(c *gin.Context) {
	data := map[string]interface{}{"CaptchaPubKey": captcha.PubKey}
	errors := map[string]string{}

	email := c.PostForm("email")
	if email == "" {
		errors["email"] = "Ce champ est obligatoire."
	} else if len(email) > 100 {
		errors["email"] = "Ce champ ne doit pas dépasser plus de 100 caractères."
	} else if valid.IsEmail(email) == false {
		errors["email"] = "Cette adresse email n'est pas valide."
	}

	recaptcha := c.PostForm("g-recaptcha-response")
	if recaptcha == "" {
		errors["recaptcha"] = "Ce champ est obligatoire."
	} else if ok, err := captcha.Validate(recaptcha); err != nil {
		GetError(c, err)
		return
	} else if !ok {
		errors["recaptcha"] = "Ce champ n'est pas valide."
	}

	if len(errors) != 0 {
		data["errors"] = errors
		data["email"] = email
		html.Render("account/password_reset_start.html", c, data, http.StatusBadRequest)
	} else {
		data["done"] = "true"
		u, err := user.ByEmail(email)
		if err != nil && err == sqlerrors.NotFound {
			html.Render("account/password_reset_start.html", c, data, http.StatusOK)
		} else if err != nil {
			GetError(c, err)
		} else if err := u.NewPasswordRequest(); err != nil {
			GetError(c, err)
		} else {
			html.Render("account/password_reset_start.html", c, data, http.StatusOK)
		}
	}
}

func GetResetInput(c *gin.Context) {
	data := map[string]interface{}{"Title": "Nouveau mot de passe"}

	url := c.Param("url")
	pr, err := user.GetPasswordRequest(url)
	if err == sqlerrors.NotFound || pr.IsValid() == false {
		html.Render("account/reset_input.html", c, data, http.StatusNotFound)
	} else {
		data["found"] = "true"
		html.Render("account/password_reset_input.html", c, data, http.StatusOK)
	}
}

func PostResetInput(c *gin.Context) {
	data := map[string]interface{}{"Title": "Nouveau mot de passe"}
	errors := map[string]string{}

	url := c.Param("url")
	pr, err := user.GetPasswordRequest(url)
	if err == sqlerrors.NotFound || pr.IsValid() == false {
		html.Render("account/password_reset_input.html", c, data, http.StatusNotFound)
		return
	} else if err != nil {
		GetError(c, err)
		return
	}

	u, err := pr.GetUser()
	if err == sqlerrors.NotFound {
		html.Render("account/password_reset_input.html", c, data, http.StatusNotFound)
		return
	} else if err != nil {
		GetError(c, err)
		return
	}

	data["found"] = "true"

	password1 := c.PostForm("password1")
	if password1 == "" {
		errors["password1"] = "Ce champ est obligatoire."
	} else if len(password1) < 6 {
		errors["password1"] = "Ce champ doit faire au moins 6 caractères."
	} else if len(password1) > 100 {
		errors["password1"] = "Ce champ ne doit pas dépasser plus de 100 caractères."
	}

	password2 := c.PostForm("password2")
	if password2 == "" {
		errors["password2"] = "Ce champ est obligatoire."
	} else if len(password2) < 6 {
		errors["password2"] = "Ce champ doit faire au moins 6 caractères."
	} else if len(password2) > 100 {
		errors["password2"] = "Ce champ ne doit pas dépasser plus de 100 caractères."
	} else if password1 != password2 {
		errors["password2"] = "Ce champ ne correspond pas au nouveau mot de passe."
	}

	if len(errors) != 0 {
		data["errors"] = errors
		html.Render("account/password_reset_input.html", c, data, http.StatusBadRequest)
	} else {
		if err := u.UpdatePassword(password1); err != nil {
			GetError(c, err)
		} else if err := pr.Invalidate(); err != nil {
			GetError(c, err)
		} else {
			html.Render("account/password_reset_complete.html", c, data, http.StatusOK)
		}
	}
}
