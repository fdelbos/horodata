package groups

import (
	"encoding/json"

	"bitbucket.com/hyperboloide/horo/middlewares"
	"bitbucket.com/hyperboloide/horo/models/group"
	"bitbucket.com/hyperboloide/horo/models/types/listing"
	"bitbucket.com/hyperboloide/horo/www/api/jsend"
	"github.com/gin-gonic/gin"
)

// Listing lists groups available to the user
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

// Create a new group
func Create(c *gin.Context) {
	u := middlewares.GetUser(c)
	var data struct {
		Name string `json:"name"`
	}
	if err := json.NewDecoder(c.Request.Body).Decode(&data); err != nil {
		jsend.ErrorJson(c)
		return
	}
	name := data.Name
	errors := map[string]string{}
	if name == "" {
		errors["name"] = "Ce champ est obligatoire."
	} else if len(name) > 30 {
		errors["name"] = "Ce champ ne doit pas depasser plus de 30 caractères."
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
	} else if newG, err := group.ByUrl(g.Url); err != nil {
		jsend.Error(c, err)
	} else if err := newG.GuestAdd(u.Email, 0, true, false); err != nil {
		jsend.Error(c, err)
	} else {
		jsend.Ok(c, newG)
	}
}

// Get the group from url
func Get(c *gin.Context) {
	guest := middlewares.GetGuest(c)
	group := middlewares.GetGroup(c)

	if detail, err := group.ApiDetail(guest.Admin); err != nil {
		jsend.Error(c, err)
	} else {
		jsend.Ok(c, detail)
	}
}

// Delete the group from url
func Delete(c *gin.Context) {
	group := middlewares.GetGroup(c)

	if err := group.Delete(); err != nil {
		jsend.Error(c, err)
	} else {
		jsend.Ok(c, nil)
	}
}
