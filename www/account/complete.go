package account

import (
	"bitbucket.com/hyperboloide/horo/html"
	sqlerrors "bitbucket.com/hyperboloide/horo/models/errors"
	"bitbucket.com/hyperboloide/horo/models/user"
	"bitbucket.com/hyperboloide/horo/services/cookies"
	"bitbucket.com/hyperboloide/horo/services/urls"
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth"
	"net/http"
	"regexp"
)

func GetComplete(c *gin.Context) {
	_, err := cookies.Get("session", "provider", c)
	if err != nil {
		if err == sqlerrors.NotFound {
			c.Redirect(http.StatusTemporaryRedirect, urls.WWWLogin)
		} else {
			GetError(c, err)
		}
	} else {
		html.Render("account/choose_username.html", c, nil, http.StatusOK)
	}
}

func PostComplete(c *gin.Context) {
	errors := map[string]string{}

	tmp, err := cookies.Get("session", "provider", c)
	if err != nil {
		if err == sqlerrors.NotFound {
			c.Redirect(http.StatusTemporaryRedirect, urls.WWWLogin)
		} else {
			GetError(c, err)
		}
		return
	}
	guser := tmp.(goth.User)

	username := c.PostForm("username")
	if username == "" {
		errors["username"] = "Ce champs est obligatoire."
	} else if len(username) < 4 {
		errors["username"] = "Ce champ doit faire au moins 4 caractères."
	} else if len(username) > 30 {
		errors["username"] = "Ce champ ne doit pas dépasser 30 caractères."
	} else if ok, err := regexp.MatchString(`^[\w.-]+$`, username); err != nil {
		GetError(c, err)
		return
	} else if !ok {
		errors["username"] = "Ce champs ne peut contenir que des lettres, des chiffres et les caractères ./-/_ ."
	} else if _, err := user.ByLogin(username); err == nil {
		errors["username"] = "Ce nom d'utilisateur est déjà associé à un autre compte."
	} else if err != sqlerrors.NotFound {
		GetError(c, err)
		return
	}

	if len(errors) != 0 {
		c.Writer.WriteHeader(http.StatusBadRequest)

		data := map[string]interface{}{
			"Title":    "Inscription",
			"errors":   errors,
			"username": username,
		}
		html.Render("account/choose_username.html", c, data, http.StatusBadRequest)
	} else {
		u := &user.User{}
		u.Active = true
		u.Login = username
		u.Email = guser.Email
		u.FullName = guser.Name
		if err := u.Insert(); err != nil {
			GetError(c, err)
		} else if err := u.SendWelcome(); err != nil {
			GetError(c, err)
		} else if err := u.UpdateProfile(); err != nil {
			GetError(c, err)
		} else if err := cookies.Delete("session", "provider", c); err != nil {
			GetError(c, err)
		} else {
			LogUser(u, c)
		}
	}
}
