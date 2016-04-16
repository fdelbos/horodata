package users

import (
	"encoding/json"
	"net/http"
	"strings"

	"dev.hyperboloide.com/fred/horodata/middlewares"
	"dev.hyperboloide.com/fred/horodata/models/user"
	"dev.hyperboloide.com/fred/horodata/www/api/jsend"
	"github.com/gin-gonic/gin"
)

func Group(r *gin.RouterGroup) {
	users := r.Group("/users")
	{
		users.GET("/me", Get)
		users.PUT("/me", Update)
		users.GET("/me/quotas", Quota)
	}
}

func Get(c *gin.Context) {
	u := middlewares.GetUser(c)
	jsend.Success(c, http.StatusOK, u)
}

func Update(c *gin.Context) {
	u := middlewares.GetUser(c)

	var data struct {
		Name string `json:"name"`
	}
	if err := json.NewDecoder(c.Request.Body).Decode(&data); err != nil {
		jsend.ErrorJson(c)
		return
	}

	errors := map[string]string{}
	data.Name = strings.TrimSpace(data.Name)
	if data.Name == "" {
		errors["name"] = "Ce champ est obligatoire."
	} else if len(data.Name) < 4 {
		errors["name"] = "Ce champ doit faire au moins 4 caractères."
	} else if len(data.Name) > 50 {
		errors["name"] = "Ce champ ne doit pas dépasser 50 caractères."
	} else {
		u.FullName = data.Name
	}

	if len(errors) > 0 {
		jsend.BadRequest(c, errors)
	} else if err := u.Update(); err != nil {
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
