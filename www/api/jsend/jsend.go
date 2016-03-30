package jsend

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data, omitempty"`
}

type ErrorResponse struct {
	Status string      `json:"status"`
	Errors interface{} `json:"errors"`
}

func Success(c *gin.Context, code int, data interface{}) {
	c.JSON(code, &Response{"success", data})
}

func Fail(c *gin.Context, code int, data interface{}) {
	c.JSON(code, &ErrorResponse{"fail", data})
}

func Error(c *gin.Context, err error) {
	log.WithFields(log.Fields{
		"error": err,
	}).Error("Api error detected.")

	c.JSON(
		http.StatusInternalServerError,
		&ErrorResponse{
			"error",
			http.StatusText(http.StatusInternalServerError),
		})
}

func ErrorJson(c *gin.Context) {
	Fail(c, http.StatusBadRequest, "invalid json.")
}

func NotFound(c *gin.Context) {
	Fail(c, http.StatusNotFound, http.StatusText(http.StatusNotFound))
}

func BadRequest(c *gin.Context, errors map[string]string) {
	Fail(c, http.StatusBadRequest, errors)
}

func Ok(c *gin.Context, data interface{}) {
	Success(c, http.StatusOK, data)
}

func Forbidden(c *gin.Context) {
	Fail(c, http.StatusForbidden, http.StatusText(http.StatusForbidden))
}

func Quota(c *gin.Context) {
	Fail(c, http.StatusForbidden, "quota exceeded.")
}

func Created(c *gin.Context, data interface{}) {
	Success(c, http.StatusCreated, data)
}
