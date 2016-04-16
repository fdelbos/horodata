package groups

import (
	"encoding/json"
	"strconv"
	"time"

	"dev.hyperboloide.com/fred/horodata/middlewares"
	sqlerrors "dev.hyperboloide.com/fred/horodata/models/errors"
	"dev.hyperboloide.com/fred/horodata/models/types/listing"
	"dev.hyperboloide.com/fred/horodata/www/api/jsend"
	"github.com/gin-gonic/gin"
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
		errors["size"] = "La valeur de ce champ ne peut dépasser 100."
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

type jobData struct {
	Task     int64  `json:"task"`
	Customer int64  `json:"customer"`
	Duration int64  `json:"duration"`
	Comment  string `json:"comment"`
}

func validateJob(c *gin.Context) (*jobData, map[string]string, error) {
	g := middlewares.GetGroup(c)

	data := &jobData{}
	if err := json.NewDecoder(c.Request.Body).Decode(data); err != nil {
		return nil, nil, err
	}

	errors := map[string]string{}

	commentMandatory := true
	if data.Task == 0 {
		errors["task"] = "Ce champ est obligatoire."
	} else if task, err := g.TaskGet(data.Task); err == sqlerrors.NotFound {
		errors["task"] = "Ce champ n'est pas valide."
	} else if err != nil {
		return nil, nil, err
	} else {
		commentMandatory = task.CommentMandatory
	}

	if data.Customer == 0 {
		errors["customer"] = "Ce champ est obligatoire."
	} else if _, err := g.CustomerGet(data.Customer); err == sqlerrors.NotFound {
		errors["customer"] = "Ce champ n'est pas valide."
	} else if err != nil {
		return nil, nil, err
	}

	if data.Duration == 0 {
		errors["duration"] = "La durée ne peut être égale à 0."
	} else if data.Duration > 13*3600 { // more than 13 hours...
		errors["duration"] = "Ce champ n'est pas valide."
	} else if data.Duration < 0 {
		errors["duration"] = "Ce champ n'est pas valide."
	}

	if data.Comment == "" && commentMandatory {
		errors["comment"] = "Ce champ est obligatoire."
	} else if len(data.Comment) > 1500 {
		errors["comment"] = "Ce champ ne doit pas dépasser 1500 caractères."
	}

	if len(errors) != 0 {
		return nil, errors, nil
	} else {
		return data, nil, nil
	}
}

func JobAdd(c *gin.Context) {
	g := middlewares.GetGroup(c)

	u, err := g.GetOwner()
	if err != nil {
		jsend.ErrorJson(c)
		return
	} else if ok, err := u.QuotaCanAddJob(); err != nil {
		jsend.ErrorJson(c)
		return
	} else if !ok {
		jsend.Quota(c, "jobs")
		return
	}

	guest := middlewares.GetGuest(c)

	if data, errors, err := validateJob(c); err != nil {
		jsend.Error(c, err)
	} else if errors != nil {
		jsend.BadRequest(c, errors)
	} else if err := g.JobAdd(data.Task, data.Customer, guest.Id, data.Duration, data.Comment); err != nil {
		jsend.Error(c, err)
	} else if err := u.UsageJobsIncr(); err != nil {
		jsend.Error(c, err)
	} else {
		jsend.Ok(c, nil)
	}
}

func JobUpdate(c *gin.Context) {
	g := middlewares.GetGroup(c)
	guest := middlewares.GetGuest(c)

	if data, errors, err := validateJob(c); err != nil {
		jsend.Error(c, err)
	} else if errors != nil {
		jsend.BadRequest(c, errors)
	} else if id, err := strconv.ParseInt(c.Param("jobId"), 10, 64); err != nil {
		jsend.BadRequest(c, nil)
	} else if j, err := g.JobGet(id); err == sqlerrors.NotFound {
		jsend.NotFound(c)
	} else if err != nil {
		jsend.Error(c, err)
	} else if !guest.Admin && j.CreatorId != guest.Id {
		jsend.Forbidden(c)
	} else if !guest.Admin && !isSameDay(j.Created, time.Now()) {
		jsend.Forbidden(c)
	} else {
		j.TaskId = data.Task
		j.CustomerId = data.Customer
		j.Duration = data.Duration
		j.Comment = data.Comment
		if err := j.Update(g.Id); err != nil {
			jsend.Error(c, err)
		} else {
			jsend.Ok(c, nil)
		}
	}
}

func isSameDay(a, b time.Time) bool {
	aY, aM, aD := a.Date()
	bY, bM, bD := b.Date()
	return aY == bY && aM == bM && aD == bD
}

func JobDelete(c *gin.Context) {
	g := middlewares.GetGroup(c)
	guest := middlewares.GetGuest(c)

	if id, err := strconv.ParseInt(c.Param("jobId"), 10, 64); err != nil {
		jsend.BadRequest(c, nil)
	} else if j, err := g.JobGet(id); err == sqlerrors.NotFound {
		jsend.NotFound(c)
	} else if err != nil {
		jsend.Error(c, err)
	} else if !guest.Admin && j.CreatorId != guest.Id {
		jsend.Forbidden(c)
	} else if !guest.Admin && !isSameDay(j.Created, time.Now()) {
		jsend.Forbidden(c)
	} else if err := g.JobRemove(id); err != nil {
		jsend.Error(c, err)
	} else if owner, err := g.GetOwner(); err != nil {
		jsend.Error(c, err)
	} else if err := owner.UsageJobsDecr(); err != nil {
		jsend.Error(c, err)
	} else {
		jsend.Ok(c, nil)
	}
}
