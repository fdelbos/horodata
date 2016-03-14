package oauth

import (
	"bitbucket.com/hyperboloide/horo/services/cookies"
	"fmt"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/gplus"
	"github.com/spf13/viper"
	"net/http"
)

func Configure() {
	googleCb := fmt.Sprintf("%s/account/provider/google/callback", viper.GetString("www_root"))
	googleProv := gplus.New(
		viper.GetString("oauth_provider_google_key"),
		viper.GetString("oauth_provider_google_secret"),
		googleCb)

	goth.UseProviders(googleProv)
	gothic.Store = cookies.GetStore()

	gothic.GetProviderName = func(req *http.Request) (string, error) {
		return "gplus", nil
	}
}

func IsVerified(user goth.User) bool {
	return user.RawData["verified_email"].(bool)
}
