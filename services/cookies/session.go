package cookies

import (
	"github.com/codahale/charlie"
	"github.com/dchest/uniuri"
	"github.com/gin-gonic/gin"
	"time"
)

const (
	SessionStoreName = "session"
	CSRFTokenSize    = 32
	CSRFField        = "csrf"
	CSRFMaxAge       = time.Hour * 4
)

var (
	csrfKey []byte
)

func NewSession(id string, c *gin.Context) error {
	csrfId := uniuri.NewLen(32)

	if err := SessionSet("id", id, c); err != nil {
		return err
	} else if err := SessionSet(CSRFField, csrfId, c); err != nil {
		return err
	}
	return nil
}

func NewCSRFToken(c *gin.Context) (string, error) {
	params := charlie.New(csrfKey)
	params.MaxAge = CSRFMaxAge
	if csrfId, err := SessionGet("csrf", c); err != nil {
		return "", err
	} else {
		return params.Generate(csrfId.(string)), nil
	}
}

func ValidateCSRF(token string, c *gin.Context) (bool, error) {
	params := charlie.New(csrfKey)
	params.MaxAge = CSRFMaxAge
	if csrfId, err := SessionGet(CSRFField, c); err != nil {
		return false, err
	} else {
		return params.Validate(csrfId.(string), token) == nil, nil
	}
}

func SessionGet(key string, c *gin.Context) (interface{}, error) {
	return Get(SessionStoreName, key, c)
}

func SessionSet(key string, value interface{}, c *gin.Context) error {
	return Set(SessionStoreName, key, value, c)
}

func SessionClear(c *gin.Context) error {
	return Clear(SessionStoreName, c)
}

func SessionDelete(key string, c *gin.Context) error {
	return Delete(SessionStoreName, key, c)
}
