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

func extractTime(c *gin.Context) (begin, end time.Time, errors map[string]string, err error) {
	errors = map[string]string{}

	const datefmt = "2006-01-02"

	loc, err := time.LoadLocation("Europe/Paris")
	if err != nil {
		return
	}

	if c.Query("end") == "" {
		errors["end"] = "Ce champ est obligatoire."
	} else if t, err := time.ParseInLocation(datefmt, c.Query("end"), loc); err != nil {
		errors["end"] = "Ce champ n'est pas valide."
	} else if t.After(time.Now()) {
		errors["end"] = "Ce champ ne peut être supérieur à la date du jour."
	} else {
		end = t.Add(24 * time.Hour)
	}

	if len(errors) > 0 {
		return
	} else if c.Query("begin") == "" {
		errors["begin"] = "Ce champ est obligatoire."
	} else if t, err := time.ParseInLocation(datefmt, c.Query("begin"), loc); err != nil {
		errors["begin"] = "Ce champ n'est pas valide."
	} else if t.After(time.Now()) {
		errors["begin"] = "Ce champ ne peut être supérieur à la date du jour."
	} else if t.After(end) {
		errors["begin"] = "Ce champ ne peut être supérieur à la date de fin."
	} else {
		begin = t
	}
	return
}

func extractGuestId(c *gin.Context) (guestId *int64, errors map[string]string, err error) {
	errors = map[string]string{}

	g := middlewares.GetGroup(c)
	guest := middlewares.GetGuest(c)

	if !guest.Admin {
		guestId = &guest.Id
	} else if str := c.Query("guest"); str == "" {
		guestId = nil
	} else if i, errConv := strconv.ParseInt(str, 10, 64); errConv != nil {
		errors["guest"] = "Ce champ n'est pas valide."
	} else if guestObj, errSql := g.GuestGetById(i); errSql == sqlerrors.NotFound {
		errors["guest"] = "Ce champ n'est pas valide."
	} else if errSql != nil {
		err = errSql
		return
	} else {
		guestId = &guestObj.Id
	}
	return
}

func JobListing(c *gin.Context) {
	g := middlewares.GetGroup(c)

	begin, end, errors, err := extractTime(c)
	if err != nil {
		jsend.Error(c, err)
		return
	} else if len(errors) > 0 {
		jsend.BadRequest(c, errors)
		return
	}

	guestId, errors, err := extractGuestId(c)
	if err != nil {
		jsend.Error(c, err)
		return
	} else if len(errors) > 0 {
		jsend.BadRequest(c, errors)
		return
	}

	var customerId *int64
	if str := c.Query("customer"); str == "" {
		customerId = nil
	} else if i, err := strconv.ParseInt(str, 10, 64); err != nil {
		errors["customer"] = "Ce champ n'est pas valide."
	} else if _, err := g.CustomerGet(i); err == sqlerrors.NotFound {
		errors["customer"] = "Ce champ n'est pas valide."
	} else if err != nil {
		jsend.Error(c, err)
		return
	} else {
		customerId = &i
	}

	offset := 0
	if str := c.Query("offset"); str == "" {
		offset = 0
	} else if i, err := strconv.ParseInt(str, 10, 32); err != nil {
		errors["offset"] = "Ce champ n'est pas valide."
	} else if i < 0 {
		errors["offset"] = "Ce champ n'est pas valide."
	} else {
		offset = int(i)
	}

	size := 10
	if str := c.Query("size"); str == "" {
		size = 10
	} else if i, err := strconv.ParseInt(str, 10, 32); err != nil {
		errors["size"] = "Ce champ n'est pas valide."
	} else if i < 0 {
		errors["size"] = "Ce champ n'est pas valide."
	} else if i > 100 {
		errors["size"] = "La valeur de ce champ ne peut depasser 100."
	} else {
		size = int(i)
	}

	if len(errors) > 0 {
		jsend.BadRequest(c, errors)
		return
	}

	request := &listing.Request{offset, size}
	if res, err := g.JobApiList(begin, end, customerId, guestId, request); err != nil {
		jsend.Error(c, err)
	} else {
		jsend.Ok(c, res)
	}

}

func JobAdd(c *gin.Context) {
	g := middlewares.GetGroup(c)
	guest := middlewares.GetGuest(c)

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
		errors["task"] = "Ce champ n'est pas valide."
	} else if err != nil {
		jsend.Error(c, err)
		return
	}

	if data.Customer == 0 {
		errors["customer"] = "Ce champ est obligatoire."
	} else if _, err := g.CustomerGet(data.Customer); err == sqlerrors.NotFound {
		errors["customer"] = "Ce champ n'est pas valide."
	} else if err != nil {
		jsend.Error(c, err)
		return
	}

	if data.Duration == 0 {
		errors["duration"] = "La durée ne peut être égale à 0."
	} else if data.Duration > 13*3600 { // more than 13 hours...
		errors["duration"] = "Ce champ n'est pas valide."
	} else if data.Duration < 0 {
		errors["duration"] = "Ce champ n'est pas valide."
	}

	if len(errors) > 0 {
		jsend.BadRequest(c, errors)
		return
	} else if data.Comment == "" && task.CommentMandatory {
		errors["comment"] = "Ce champ est obligatoire."
	} else if len(data.Comment) > 1500 {
		errors["comment"] = "Ce champ ne doit pas depasser plus de 1500 caractères."
	}

	if len(errors) > 0 {
		jsend.BadRequest(c, errors)
		return
	} else if err := g.JobAdd(data.Task, data.Customer, guest.Id, data.Duration, data.Comment); err != nil {
		jsend.Error(c, err)
	} else {
		jsend.Ok(c, nil)
	}
}
