package users

import (
	"bitbucket.com/hyperboloide/horo/middlewares"
	sqlerrors "bitbucket.com/hyperboloide/horo/models/errors"
	"bitbucket.com/hyperboloide/horo/models/user"
	"bitbucket.com/hyperboloide/horo/services/urls"
	"bitbucket.com/hyperboloide/horo/www/api/jsend"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Group(r *gin.RouterGroup) {
	users := r.Group("/users")
	{
		users.GET("/:login", GetUser)
	}
}

type ApiUser struct {
	user.User
	Quota *user.Quota `json:"quota,omitempty"`
	Usage *user.Usage `json:"usage,omitempty"`
}

func (au ApiUser) MarshalJSON() ([]byte, error) {
	type alias ApiUser
	res := &struct {
		Link string `json:"_link"`
		alias
	}{urls.ApiUsers + "/" + au.Login, (alias)(au)}
	res.Id = 0
	return json.Marshal(res)
}

func ApiByLogin(login string) (*ApiUser, error) {
	u, err := user.ByLogin(login)
	if err != nil {
		return nil, err
	}
	apiUser := &ApiUser{}
	apiUser.User = *u
	return apiUser, nil
}

func GetUser(c *gin.Context) {
	u := middlewares.GetUser(c)
	login := c.Param("login")
	if login == "me" || login == u.Login {
		GetMe(c, u)
		return
	} else {
		if res, err := ApiByLogin(login); err == sqlerrors.NotFound {
			jsend.NotFound(c)
		} else if err != nil {
			jsend.Error(c, err)
		} else {
			res.Email = ""
			jsend.Success(c, http.StatusOK, res)
		}
	}
}

func GetMe(c *gin.Context, u *user.User) {
	if au, err := ApiByLogin(u.Login); err != nil {
		jsend.Error(c, err)
	} else if usage, err := u.GetUsage(); err != nil {
		jsend.Error(c, err)
	} else if quota, err := u.GetQuota(); err != nil {
		jsend.Error(c, err)
	} else {
		au.Usage = usage
		au.Quota = quota
		jsend.Success(c, http.StatusOK, au)
	}
}
