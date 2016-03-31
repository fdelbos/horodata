package settings

import (
	"dev.hyperboloide.com/fred/horodata/html"
	"dev.hyperboloide.com/fred/horodata/middlewares"
	validator "github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func GetProfile(c *gin.Context) {
	u := middlewares.GetUser(c)

	data := map[string]interface{}{
		"page":         "profile",
		"full_name":    u.FullName,
		"organization": u.Organization,
		"website":      u.Website,
		"about":        u.About,
	}
	html.Render("profile.html", c, data, http.StatusOK)
}

func PostProfile(c *gin.Context) {
	T := middlewares.GetTranslate(c)
	u := middlewares.GetUser(c)

	data := map[string]interface{}{"page": "profile"}
	errors := map[string]string{}

	if len(errors) == 0 {
		fullName := c.PostForm("full_name")
		if len(fullName) > 100 {
			errors["full_name"] = T("generic.form.error_too_large", 100)
		}
		organization := c.PostForm("organization")
		if len(organization) > 100 {
			errors["organization"] = T("generic.form.error_too_large", 100)
		}
		website := strings.TrimSpace(c.PostForm("website"))
		if len(website) > 500 {
			errors["website"] = T("generic.form.error_too_large", 500)
		} else if len(website) > 0 && !validator.IsURL(website) {
			errors["website"] = T("generic.form.invalid_url")
		}
		about := c.PostForm("about")
		if len(about) > 500 {
			errors["about"] = T("generic.form.error_too_large", 500)
		}
		data["full_name"] = fullName
		data["organization"] = organization
		data["website"] = website
		data["about"] = about

		if len(errors) == 0 {
			u.FullName = fullName
			u.Organization = organization
			u.Website = website
			u.About = about
			if err := u.UpdateProfile(); err != nil {
				html.ErrorServer(c, err)
			} else {
				data["success"] = true
			}
		}
	}

	if len(errors) > 0 {
		data["errors"] = errors
		html.Render("profile.html", c, data, http.StatusBadRequest)
	} else {
		html.Render("profile.html", c, data, http.StatusOK)
	}
}
