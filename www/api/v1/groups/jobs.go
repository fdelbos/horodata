package groups

import (
	"bitbucket.com/hyperboloide/horo/middlewares"
	sqlerrors "bitbucket.com/hyperboloide/horo/models/errors"
	"bitbucket.com/hyperboloide/horo/models/group"
	"bitbucket.com/hyperboloide/horo/www/api/jsend"
	"encoding/json"
	"github.com/gin-gonic/gin"
	// "strconv"
)

func JobAdd(c *gin.Context) {
	g := middlewares.GetGroup(c)
	u := middlewares.GetUser(c)

	var data struct {
		Task     int64  `json:"task"`
		Customer int64  `json:"customer"`
		Duration int64  `json:"duration"`
		Comment  string `json:"comment"`
	}
	if err := json.NewDecoder(c.Request.Body).Decode(&data); err != nil {
		jsend.ErrorJson(c)
		return
	}

	errors := map[string]string{}

	var err error
	var task *group.Task
	if data.Task == 0 {
		errors["task"] = "Ce champs est obligatoire."
	} else if task, err = g.TaskGet(data.Task); err == sqlerrors.NotFound {
		errors["task"] = "Ce champs est invalide."
	} else if err != nil {
		jsend.Error(c, err)
		return
	}

	if data.Customer == 0 {
		errors["customer"] = "Ce champs est obligatoire."
	} else if _, err := g.CustomerGet(data.Customer); err == sqlerrors.NotFound {
		errors["customer"] = "Ce champs est invalide."
	} else if err != nil {
		jsend.Error(c, err)
		return
	}

	if data.Duration == 0 {
		errors["duration"] = "La duree de la tâche ne peut être nulle."
	} else if data.Duration > 13*3600 { // more than 13 hours...
		errors["duration"] = "Ce champs est invalide."
	} else if data.Duration < 0 {
		errors["duration"] = "Ce champs est invalide."
	}

	if len(errors) > 0 {
		jsend.BadRequest(c, errors)
		return
	} else if data.Comment == "" && task.CommentMandatory {
		errors["comment"] = "Ce champs est obligatoire."
	} else if len(data.Comment) > 1500 {
		errors["comment"] = "Ce champ ne doit pas depasser 1500 caractères."
	}

	if len(errors) > 0 {
		jsend.BadRequest(c, errors)
		return
	} else if err := g.JobAdd(data.Task, data.Customer, u.Id, data.Duration, data.Comment); err != nil {
		jsend.Error(c, err)
	} else {
		jsend.Ok(c, nil)
	}
}
