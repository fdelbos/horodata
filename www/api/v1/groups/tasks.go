package groups

import (
	"bitbucket.com/hyperboloide/horo/middlewares"
	sqlerrors "bitbucket.com/hyperboloide/horo/models/errors"
	"bitbucket.com/hyperboloide/horo/www/api/jsend"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"strconv"
)

func TaskAdd(c *gin.Context) {
	g := middlewares.GetGroup(c)

	var data struct {
		Name             string `json:"name"`
		CommentMandatory bool   `json:"comment_mandatory"`
	}
	if err := json.NewDecoder(c.Request.Body).Decode(&data); err != nil {
		jsend.Error(c, err)
	}

	errors := map[string]string{}
	if data.Name == "" {
		errors["name"] = "Ce champs est obligatoire."
	} else if len(data.Name) > 40 {
		errors["name"] = "Ce champ ne doit pas depasser 40 caractères."
	}
	if len(errors) > 0 {
		jsend.BadRequest(c, errors)
		return
	}

	if err := g.TaskAdd(data.Name, data.CommentMandatory); err != nil {
		jsend.Error(c, err)
	} else {
		jsend.Ok(c, nil)
	}
}

func TaskUpdate(c *gin.Context) {
	g := middlewares.GetGroup(c)
	id, err := strconv.ParseInt(c.Param("taskId"), 10, 64)
	if err != nil {
		jsend.BadRequest(c, nil)
		return
	}

	var data struct {
		Name             string `json:"name"`
		CommentMandatory bool   `json:"comment_mandatory"`
	}
	if err := json.NewDecoder(c.Request.Body).Decode(&data); err != nil {
		jsend.Error(c, err)
	}

	errors := map[string]string{}
	if data.Name == "" {
		errors["name"] = "Ce champs est obligatoire."
	} else if len(data.Name) > 40 {
		errors["name"] = "Ce champ ne doit pas depasser 40 caractères."
	}
	if len(errors) > 0 {
		jsend.BadRequest(c, errors)
		return
	}

	if t, err := g.TaskGet(id); err == sqlerrors.NotFound {
		jsend.NotFound(c)
	} else if err != nil {
		jsend.Error(c, err)
	} else {
		t.Name = data.Name
		t.CommentMandatory = data.CommentMandatory
		if err := t.Update(); err != nil {
			jsend.Error(c, err)
		} else {
			jsend.Ok(c, t)
		}
	}
}

func TaskDelete(c *gin.Context) {
	g := middlewares.GetGroup(c)
	id, err := strconv.ParseInt(c.Param("taskId"), 10, 64)
	if err != nil {
		jsend.BadRequest(c, nil)
	} else if t, err := g.TaskGet(id); err == sqlerrors.NotFound {
		jsend.NotFound(c)
	} else if err != nil {
		jsend.Error(c, err)
	} else {
		t.Active = false
		if err := t.Update(); err != nil {
			jsend.Error(c, err)
		} else {
			jsend.Ok(c, nil)
		}
	}
}
