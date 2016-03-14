package groups

import (
	"bitbucket.com/hyperboloide/horo/middlewares"
	"bitbucket.com/hyperboloide/horo/models/errors"
	"bitbucket.com/hyperboloide/horo/models/group"
	"bitbucket.com/hyperboloide/horo/models/types/listing"
	"bitbucket.com/hyperboloide/horo/www/api/jsend"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

func Listing(c *gin.Context) {
	u := middlewares.GetUser(c)
	if request, errs := listing.NewRequest(c); errs != nil {
		jsend.BadRequest(c, errs)
	} else if errs := request.Validate(); errs != nil {
		jsend.BadRequest(c, errs)
	} else if res, err := group.ApiByUser(u.Id, request); err != nil {
		jsend.Error(c, err)
	} else {
		jsend.Ok(c, res)
	}
}

func Create(c *gin.Context) {
	u := middlewares.GetUser(c)
	var data struct {
		Name string `json:"name"`
	}
	if err := json.NewDecoder(c.Request.Body).Decode(&data); err != nil {
		jsend.Error(c, err)
		return
	}
	name := data.Name
	errors := map[string]string{}
	if name == "" {
		errors["name"] = "Ce champs est obligatoire."
	} else if len(name) > 30 {
		errors["name"] = "Ce champ ne doit pas depasser 30 caractères."
	}
	if len(errors) > 0 {
		jsend.BadRequest(c, errors)
		return
	}

	g := &group.Group{
		Name:    data.Name,
		OwnerId: u.Id,
	}
	if err := g.Insert(); err != nil {
		jsend.Error(c, err)
	} else if g, err := group.ApiByUrl(g.Url); err != nil {
		jsend.Error(c, err)
	} else {
		jsend.Ok(c, g)
	}
}

func Get(c *gin.Context) {
	//u := middlewares.GetUser(c)
	url := c.Param("url")
	if g, err := group.ApiByUrl(url); err != nil && err != errors.NotFound {
		jsend.Error(c, err)
	} else if err == errors.NotFound {
		jsend.NotFound(c)
	} else {
		jsend.Ok(c, g)
	}
}

func Update(c *gin.Context) {}
