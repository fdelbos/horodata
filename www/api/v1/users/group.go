package users

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"dev.hyperboloide.com/fred/horodata/middlewares"
	"dev.hyperboloide.com/fred/horodata/models/user"
	"dev.hyperboloide.com/fred/horodata/services/mail"
	"dev.hyperboloide.com/fred/horodata/www/api/jsend"
	"github.com/gin-gonic/gin"
	"github.com/hyperboloide/qmail/client"
)

func Group(r *gin.RouterGroup) {
	users := r.Group("/users")
	{
		users.GET("/me", Get)
		users.PUT("/me", Update)
		users.GET("/me/quotas", Quota)
		users.POST("/contact_message", ContactMessage)
	}
}

func Get(c *gin.Context) {
	u := middlewares.GetUser(c)
	jsend.Success(c, http.StatusOK, u)
}

func Update(c *gin.Context) {
	u := middlewares.GetUser(c)

	name := c.PostForm("name")

	errors := map[string]string{}
	name = strings.TrimSpace(name)
	if name == "" {
		errors["name"] = "Ce champ est obligatoire."
	} else if len(name) < 4 {
		errors["name"] = "Ce champ doit faire au moins 4 caractères."
	} else if len(name) > 50 {
		errors["name"] = "Ce champ ne doit pas dépasser 50 caractères."
	} else {
		u.FullName = name
	}
	if len(errors) > 0 {
		jsend.BadRequest(c, errors)
		return
	}

	file, header, err := c.Request.FormFile("file")
	if err == nil {
		if err := u.PictureSetFromRequest(file, header.Header.Get("Content-Type")); err != nil {
			jsend.Error(c, err)
			return
		}
	}

	if err := u.Update(); err != nil {
		jsend.Error(c, err)
	} else {
		jsend.Ok(c, nil)
	}
}

func Quota(c *gin.Context) {
	u := middlewares.GetUser(c)

	if quotas, err := u.Quotas(); err != nil {
		jsend.Error(c, err)
	} else if uGroups, err := u.UsageGroups(); err != nil {
		jsend.Error(c, err)
	} else if uGests, err := u.UsageGuests(); err != nil {
		jsend.Error(c, err)
	} else if uJobs, err := u.UsageJobs(); err != nil {
		jsend.Error(c, err)
	} else {
		res := struct {
			Quotas *user.Quotas `json:"quotas"`
			Usage  user.Limits  `json:"usage"`
		}{
			quotas,
			user.Limits{
				Jobs:   uJobs,
				Guests: uGests,
				Groups: uGroups,
			},
		}
		jsend.Ok(c, res)
	}
}

func ContactMessage(c *gin.Context) {
	u := middlewares.GetUser(c)

	var data struct {
		Content string `json:"content"`
	}
	if err := json.NewDecoder(c.Request.Body).Decode(&data); err != nil {
		jsend.ErrorJson(c)
		return
	}

	errors := map[string]string{}
	data.Content = strings.TrimSpace(data.Content)
	if data.Content == "" {
		errors["content"] = "Ce champ est obligatoire."
	} else if len(data.Content) > 5000 {
		errors["content"] = "Ce champ ne doit pas dépasser 5000 caractères."
	} else if err := mail.Mailer().Send(client.Mail{
		Dests:    []string{"contact@hyperboloide.com"},
		Subject:  "Nouveau message Horodata",
		Template: "message",
		Data: map[string]interface{}{
			"name":    u.FullName,
			"email":   u.Email,
			"message": data.Content,
			"created": time.Now().Format("01/02/2006 15h04"),
		},
	}); err != nil {
		jsend.Error(c, err)
	} else {
		jsend.Ok(c, nil)
	}
}
