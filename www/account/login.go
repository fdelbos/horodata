package account

import (
	"bitbucket.com/hyperboloide/horo/html"
	sqlerrors "bitbucket.com/hyperboloide/horo/models/errors"
	"bitbucket.com/hyperboloide/horo/models/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetLogin(c *gin.Context) {
	html.Render("account/login.html", c, nil, http.StatusOK)
}

func PostLogin(c *gin.Context) {
	errors := map[string]string{}

	username := c.PostForm("username")
	if username == "" {
		errors["username"] = "Ce champs est obligatoire."
	} else if len(username) > 30 {
		errors["username"] = "Ce champ ne doit pas depasser 30 caractères."
	}

	password := c.PostForm("password")
	if password == "" {
		errors["password"] = "Ce champs est obligatoire."
	} else if len(password) > 100 {
		errors["password"] = "Ce champ ne doit pas depasser 100 caractères."
	}

	if len(errors) == 0 {
		u, err := user.ByLogin(username)
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
		"errors":   errors,
		"username": username,
	}
	html.Render("account/login.html", c, data, http.StatusBadRequest)
}
