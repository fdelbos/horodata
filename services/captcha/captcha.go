package captcha

import (
	"encoding/json"
	"github.com/spf13/viper"
	"net/http"
	"net/url"
)

var (
	PubKey     string
	privateKey string
)

func Configure() {
	PubKey = viper.GetString("captcha_pub_key")
	privateKey = viper.GetString("captcha_priv_key")
}

func Validate(response string) (bool, error) {
	resp, err := http.PostForm("https://www.google.com/recaptcha/api/siteverify",
		url.Values{"secret": {privateKey}, "response": {response}})
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	var data struct {
		Success bool `json:"success"`
	}
	err = json.NewDecoder(resp.Body).Decode(&data)
	return data.Success, err
}
