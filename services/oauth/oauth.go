package oauth

import (
	"fmt"
	"net/http"
	"strings"

	"dev.hyperboloide.com/fred/horodata/services/cookies"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/facebook"
	"github.com/markbates/goth/providers/gplus"
	"github.com/spf13/viper"
)

func Configure() {

	goth.UseProviders(
		gplus.New(
			viper.GetString("oauth_provider_google_key"),
			viper.GetString("oauth_provider_google_secret"),
			fmt.Sprintf("%s/account/provider/callback/gplus", viper.GetString("www_root"))),
		facebook.New(
			viper.GetString("oauth_provider_facebook_key"),
			viper.GetString("oauth_provider_facebook_secret"),
			fmt.Sprintf("%s/account/provider/callback/facebook", viper.GetString("www_root")),
		),
	)

	gothic.Store = cookies.GetStore()

	gothic.GetProviderName = func(req *http.Request) (string, error) {
		path := strings.Split(req.URL.Path, "/")
		return path[len(path)-1], nil
	}
}

func IsVerified(user goth.User) bool {
	if verified, ok := user.RawData["verified_email"].(bool); ok && !verified {
		return false
	}
	return true
}
