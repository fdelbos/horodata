package html

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"

	"dev.hyperboloide.com/fred/horodata/middlewares"
	"dev.hyperboloide.com/fred/horodata/models/user"
	"dev.hyperboloide.com/fred/horodata/services/cookies"
	"dev.hyperboloide.com/fred/horodata/services/urls"
	log "github.com/Sirupsen/logrus"
	"github.com/flosch/pongo2"
	_ "github.com/flosch/pongo2-addons"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var (
	tmpls *pongo2.TemplateSet
)

type Loader struct {
	AssetLoader func(name string) ([]byte, error)
}

func NewLoader(loaderFunc func(name string) ([]byte, error)) *Loader {
	return &Loader{
		AssetLoader: loaderFunc,
	}
}

func (l *Loader) Abs(base, name string) string {
	return name
}

func (l *Loader) Get(path string) (io.Reader, error) {
	b, err := l.AssetLoader(path)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}

func init() {
	if viper.GetBool("dev_mode") == true {
		tmpls = pongo2.NewSet("www", pongo2.MustNewLocalFileSystemLoader("html"))
	} else {
		tmpls = pongo2.NewSet("www", NewLoader(Asset))
	}

}

func AddContext(c *gin.Context, data map[string]interface{}) error {
	u := middlewares.GetUserMaybe(c)

	csrfFun := func() *pongo2.Value {
		csrf, err := cookies.NewCSRFToken(c)
		if err != nil {
			return pongo2.AsValue("")
		}
		csrfTag := fmt.Sprintf(`<input type="hidden" name="csrf" value="%s"/>`, csrf)
		return pongo2.AsValue(csrfTag)
	}
	data["csrf"] = csrfFun

	data["ctx"] = struct {
		WWWRoot    string
		WWWApp     string
		WWWAccount string
		StaticRoot string
		User       *user.User
		Year       int
	}{
		urls.WWWRoot,
		urls.WWWApp,
		urls.WWWAccount,
		urls.StaticRoot,
		u,
		time.Now().Year(),
	}
	return nil
}

func quickError(c *gin.Context) {
	http.Error(
		c.Writer,
		http.StatusText(http.StatusInternalServerError),
		http.StatusInternalServerError)
}

func Render(name string, c *gin.Context, ctx map[string]interface{}, status int) {
	if ctx == nil {
		ctx = map[string]interface{}{}
	}
	if err := AddContext(c, ctx); err != nil {
		log.WithFields(map[string]interface{}{
			"package":  "horodata.html",
			"function": "func Render(name string, c *gin.Context, ctx map[string]interface{}, status int)",
			"step":     "AddContext(c, ctx)",
		}).Error(err)
		quickError(c)
	} else if tmpl, err := tmpls.FromCache(name); err != nil {
		log.WithFields(map[string]interface{}{
			"package":  "horodata.html",
			"function": "func Render(name string, c *gin.Context, ctx map[string]interface{}, status int)",
			"step":     "tmpls.FromCache(name)",
		}).Error(err)
		quickError(c)
	} else {
		c.Writer.WriteHeader(status)
		if err := tmpl.ExecuteWriter(ctx, c.Writer); err != nil {
			log.WithFields(map[string]interface{}{
				"package":  "horodata.html",
				"function": "func Render(name string, c *gin.Context, ctx map[string]interface{}, status int)",
				"step":     "tmpl.ExecuteWriter(ctx, c.Writer)",
			}).Error(err)
			quickError(c)
		}
	}
}

func ErrorServer(c *gin.Context, err error) {
	log.WithFields(log.Fields{
		"error": err,
	}).Error("WWW error detected.")

	quickError(c)
}
