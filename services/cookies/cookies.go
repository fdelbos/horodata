package cookies

import (
	"bitbucket.com/hyperboloide/horo/models/errors"
	"encoding/base64"
	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/spf13/viper"
)

var (
	store *sessions.CookieStore
)

func Configure() {
	auth, err := base64.StdEncoding.DecodeString(viper.GetString("session_auth_b64"))
	if err != nil {
		log.WithField("error", err).Fatal("Cannot decode session auth.")
	}
	encrypt, err := base64.StdEncoding.DecodeString(viper.GetString("session_encryption_b64"))
	if err != nil {
		log.WithField("error", err).Fatal("Cannot decode session encrytion.")
	}
	store = sessions.NewCookieStore(auth, encrypt)
	store.MaxAge(3600 * 24 * 14)

	csrfKey, err = base64.StdEncoding.DecodeString(viper.GetString("session_csrf_b64"))
	if err != nil {
		log.WithField("error", err).Fatal("Cannot decode session auth.")
	}
}

func GetStore() sessions.Store {
	return store
}

func Set(name, key string, data interface{}, c *gin.Context) error {
	session, err := store.Get(c.Request, name)
	if err != nil {
		return err
	}
	session.Values[key] = data
	return session.Save(c.Request, c.Writer)
}

func Clear(name string, c *gin.Context) error {
	session, err := store.Get(c.Request, name)
	if err != nil {
		return err
	}
	session.Options = &sessions.Options{MaxAge: -1, Path: "/"}
	return session.Save(c.Request, c.Writer)
}

func Delete(name, key string, c *gin.Context) error {
	return Set(name, key, nil, c)
}

func Get(name, key string, c *gin.Context) (interface{}, error) {
	session, err := store.Get(c.Request, name)
	if err != nil {
		return nil, err
	}
	value, ok := session.Values[key]
	if ok == false || value == nil {
		return nil, errors.NotFound
	}
	return value, nil
}
