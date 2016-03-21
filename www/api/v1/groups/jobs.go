package groups

import (
	"bitbucket.com/hyperboloide/horo/middlewares"
	sqlerrors "bitbucket.com/hyperboloide/horo/models/errors"
	"bitbucket.com/hyperboloide/horo/models/group"
	"bitbucket.com/hyperboloide/horo/models/types/listing"
	"bitbucket.com/hyperboloide/horo/www/api/jsend"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func JobListing(c *gin.Context) {
	g := middlewares.GetGroup(c)
	guest := middlewares.GetGuest(c)

	const datefmt = "2006-01-02"

	errors := map[string]string{}

	var end time.Time
	if c.Query("end") == "" {
		errors["end"] = "Ce champ est obligatoire."
	} else if t, err := time.Parse(datefmt, c.Query("end")); err != nil {
		errors["end"] = "Ce champ est invalide."
	} else if t.After(time.Now()) {
		errors["end"] = "Ce champ ne peut être supérieur a la date du jour."
	} else {
		end = t.Add(24 * time.Hour)
	}

	var begin time.Time
	if len(errors) > 0 {
		jsend.BadRequest(c, errors)
		return
	} else if c.Query("begin") == "" {
		errors["begin"] = "Ce champ est obligatoire."
	} else if t, err := time.Parse(datefmt, c.Query("begin")); err != nil {
		errors["begin"] = "Ce champ est invalide."
	} else if t.After(time.Now()) {
		errors["begin"] = "Ce champ ne peut être supérieur a la date du jour."
	} else if t.After(end) {
		errors["begin"] = "Ce champs ne peut être supérieur a la date de fin."
	} else {
		begin = t
	}

	var customerId *int64
	if str := c.Query("customer"); str == "" {
		customerId = nil
	} else if i, err := strconv.ParseInt(str, 10, 64); err != nil {
		errors["customer"] = "Ce champ est invalide."
	} else if _, err := g.CustomerGet(i); err == sqlerrors.NotFound {
		errors["customer"] = "Ce champ est invalide."
	} else if err != nil {
		jsend.Error(c, err)
		return
	} else {
		customerId = &i
	}

	var guestUserId *int64
	if !guest.Admin {
		guestUserId = guest.UserId
	} else if str := c.Query("guest"); str == "" {
		guestUserId = nil
	} else if i, err := strconv.ParseInt(str, 10, 64); err != nil {
		errors["guest"] = "Ce champ est invalide."
	} else if guestObj, err := g.GuestGetById(i); err == sqlerrors.NotFound {
		errors["guest"] = "Ce champ est invalide."
	} else if err != nil {
		jsend.Error(c, err)
		return
	} else {
		guestUserId = guestObj.UserId
	}

	offset := 0
	if str := c.Query("offset"); str == "" {
		offset = 0
	} else if i, err := strconv.ParseInt(str, 10, 32); err != nil {
		errors["offset"] = "Ce champ est invalide."
	} else if i < 0 {
		errors["offset"] = "Ce champ est invalide."
	} else {
		offset = int(i)
	}

	size := 10
	if str := c.Query("size"); str == "" {
		size = 10
	} else if i, err := strconv.ParseInt(str, 10, 32); err != nil {
		errors["size"] = "Ce champ est invalide."
	} else if i < 0 {
		errors["size"] = "Ce champ est invalide."
	} else if i > 100 {
		errors["size"] = "La valeure de ce champ ne peut depasser 100."
	} else {
		size = int(i)
	}

	if len(errors) > 0 {
		jsend.BadRequest(c, errors)
		return
	}

	request := &listing.Request{offset, size}
	if res, err := g.JobApiList(begin, end, customerId, guestUserId, request); err != nil {
		jsend.Error(c, err)
	} else {
		jsend.Ok(c, res)
	}

}

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
		errors["task"] = "Ce champ est obligatoire."
	} else if task, err = g.TaskGet(data.Task); err == sqlerrors.NotFound {
		errors["task"] = "Ce champ est invalide."
	} else if err != nil {
		jsend.Error(c, err)
		return
	}

	if data.Customer == 0 {
		errors["customer"] = "Ce champ est obligatoire."
	} else if _, err := g.CustomerGet(data.Customer); err == sqlerrors.NotFound {
		errors["customer"] = "Ce champ est invalide."
	} else if err != nil {
		jsend.Error(c, err)
		return
	}

	if data.Duration == 0 {
		errors["duration"] = "La duree de la tâche ne peut être nulle."
	} else if data.Duration > 13*3600 { // more than 13 hours...
		errors["duration"] = "Ce champ est invalide."
	} else if data.Duration < 0 {
		errors["duration"] = "Ce champ est invalide."
	}

	if len(errors) > 0 {
		jsend.BadRequest(c, errors)
		return
	} else if data.Comment == "" && task.CommentMandatory {
		errors["comment"] = "Ce champ est obligatoire."
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
