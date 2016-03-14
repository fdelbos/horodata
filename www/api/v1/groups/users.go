package groups

import (
	"bitbucket.com/hyperboloide/horo/www/api/jsend"
	"github.com/dchest/uniuri"
	"github.com/gin-gonic/gin"
)

func UserListing(c *gin.Context) {

	res := []interface{}{}
	for i := 0; i < 200; i++ {
		data := struct {
			Name  string `json:"name"`
			Email string `json:"email"`
			Admin bool   `json:"admin"`
		}{uniuri.New(), uniuri.New() + "@" + uniuri.New() + ".com", false}
		if i%20 == 0 {
			data.Admin = true
		}
		res = append(res, data)
	}
	jsend.Ok(c, res)
}
