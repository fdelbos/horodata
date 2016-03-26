package urls

import (
	"fmt"
	"github.com/spf13/viper"
)

var (
	StaticRoot string

	ApiRoot  string
	ApiUsers string
	ApiRoles string
	ApiVisit string

	WWWRoot     string
	WWWApp      string
	WWWAccount  string
	WWWLogin    string
	WWWComplete string

	AngularBase string
)

func Configure() {
	StaticRoot = viper.GetString("static_root")

	ApiRoot = viper.GetString("api_root")
	ApiUsers = ApiRoot + "/users"
	ApiRoles = ApiRoot + "/roles"
	ApiVisit = ApiRoot + "/visit"

	WWWRoot = viper.GetString("www_root")
	WWWApp = WWWRoot + "/app"
	WWWAccount = WWWRoot + "/account"
	WWWLogin = WWWAccount + "/login"
	WWWComplete = WWWRoot + "/account/complete_registration"

	AngularBase = viper.GetString("www_angular_base")
}

func ApiGroup(url string) string {
	return fmt.Sprintf("%s/groups/%s", ApiRoot, url)
}
