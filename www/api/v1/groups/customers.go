package groups

import (
	"bitbucket.com/hyperboloide/horo/middlewares"
	sqlerrors "bitbucket.com/hyperboloide/horo/models/errors"
	"bitbucket.com/hyperboloide/horo/www/api/jsend"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

func CustomerAdd(c *gin.Context) {
	g := middlewares.GetGroup(c)

	var data struct {
		Customers string `json:"customers"`
	}
	if err := json.NewDecoder(c.Request.Body).Decode(&data); err != nil {
		jsend.ErrorJson(c)
		return
	}

	errors := map[string]string{}
	if data.Customers == "" {
		errors["customers"] = "Ce champ est obligatoire."
	}
	if len(errors) > 0 {
		jsend.BadRequest(c, errors)
		return
	}
	customers := strings.Split(data.Customers, "\n")
	if len(customers) > 100 {
		errors["customers"] = "Vous ne pouvez pas ajouter plus de 100 dossiers à la fois."
		jsend.BadRequest(c, errors)
		return
	}
	clean := []string{}
	for _, name := range customers {
		if len(name) > 200 {
			errors["customers"] = "Le nom d'un dossier ne doit pas dépasser plus de 200 caractères."
			jsend.BadRequest(c, errors)
			return
		} else if strings.TrimSpace(name) != "" {
			clean = append(clean, name)
		}
	}

	for _, name := range clean {
		if err := g.CustomerAdd(name); err != nil {
			jsend.Error(c, err)
			return
		}
	}
	res := struct {
		Total int `json:"total"`
	}{len(clean)}
	jsend.Ok(c, res)
}

func CustomerUpdate(c *gin.Context) {
	g := middlewares.GetGroup(c)
	id, err := strconv.ParseInt(c.Param("customerId"), 10, 64)
	if err != nil {
		jsend.BadRequest(c, nil)
		return
	}

	var data struct {
		Name string `json:"name"`
	}
	if err := json.NewDecoder(c.Request.Body).Decode(&data); err != nil {
		jsend.ErrorJson(c)
		return
	}

	errors := map[string]string{}
	if data.Name == "" {
		errors["name"] = "Ce champ est obligatoire."
	} else if len(data.Name) > 200 {
		errors["name"] = "Ce champ ne doit pas dépasser plus de 200 caractères."
	}
	if len(errors) > 0 {
		jsend.BadRequest(c, errors)
		return
	}

	if cust, err := g.CustomerGet(id); err == sqlerrors.NotFound {
		jsend.NotFound(c)
	} else if err != nil {
		jsend.Error(c, err)
	} else {
		cust.Name = data.Name
		if err := cust.Update(); err != nil {
			jsend.Error(c, err)
		} else {
			jsend.Ok(c, cust)
		}
	}
}

func CustomerDelete(c *gin.Context) {
	g := middlewares.GetGroup(c)
	id, err := strconv.ParseInt(c.Param("customerId"), 10, 64)
	if err != nil {
		jsend.BadRequest(c, nil)
	} else if cust, err := g.CustomerGet(id); err == sqlerrors.NotFound {
		jsend.NotFound(c)
	} else if err != nil {
		jsend.Error(c, err)
	} else {
		cust.Active = false
		if err := cust.Update(); err != nil {
			jsend.Error(c, err)
		} else {
			jsend.Ok(c, nil)
		}
	}
}
