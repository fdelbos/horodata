package config

import (
	"bitbucket.com/hyperboloide/horo/services/cache"
	"bitbucket.com/hyperboloide/horo/services/captcha"
	"bitbucket.com/hyperboloide/horo/services/cookies"
	"bitbucket.com/hyperboloide/horo/services/oauth"
	"bitbucket.com/hyperboloide/horo/services/payment"
	"bitbucket.com/hyperboloide/horo/services/postgres"
	"bitbucket.com/hyperboloide/horo/services/urls"
	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var started bool

func init() {
	viper.SetEnvPrefix("horo")

	viper.BindEnv("dev_mode")
	viper.SetDefault("dev_mode", "true")

	viper.BindEnv("port")
	viper.SetDefault("port", "3000")

	viper.BindEnv("host")
	viper.SetDefault("host", "localhost")

	viper.BindEnv("www_root")
	viper.SetDefault("www_root", "http://localhost:3000/www")

	viper.BindEnv("api_root")
	viper.SetDefault("api_root", "http://localhost:3000/www/api/v1")

	viper.BindEnv("static_root")
	viper.SetDefault("static_root", "http://localhost:3000/static")

	viper.BindEnv("www_angular_base")
	viper.SetDefault("www_angular_base", "/www/app/")

	//
	// Captcha
	//

	viper.BindEnv("captcha_pub_key")
	viper.SetDefault("captcha_pub_key", "6LdCPAATAAAAAPEx6JGdu8TbbVz-QNHA3LrOkO7W")

	viper.BindEnv("captcha_priv_key")
	viper.SetDefault("captcha_priv_key", "6LdCPAATAAAAABnKEfGG-KmuZiTXmrWk5KrPDslj")

	//
	// Cookies
	//

	viper.BindEnv("session_auth_b64")
	viper.SetDefault(
		"session_auth_b64",
		"CVfR/RqhdXK/HCGCC4WjYdkUg9B2hZlUY729THvc89+73L4TfSaNjgYfFaEHcg3KkVUgO67TsGy1Bg5nAiHqeQ==")

	viper.BindEnv("session_encryption_b64")
	viper.SetDefault(
		"session_encryption_b64",
		"oPmNI/RVzMJzN/s5m7odRTQBLZXByE6M4fOjtdViUlU=")

	viper.BindEnv("session_csrf_b64")
	viper.SetDefault(
		"session_csrf_b64",
		"QzxWLmCyobdIFH2np5UCArFPLK7T9O6d/gDK+8J9gtQ=")

	//
	// Redis
	//

	viper.BindEnv("redis_host")
	viper.SetDefault("redis_host", "localhost")

	viper.BindEnv("redis_port")
	viper.SetDefault("redis_port", "6379")

	viper.BindEnv("redis_pool")
	viper.SetDefault("redis_pool", "10")

	//
	// PostgreSQL
	//

	viper.BindEnv("pg_host")
	viper.SetDefault("pg_host", "localhost")

	viper.BindEnv("pg_dbname")
	viper.SetDefault("pg_dbname", "horodata")

	viper.BindEnv("pg_user")
	viper.SetDefault("pg_user", "postgres")

	viper.BindEnv("pg_password")
	viper.SetDefault("pg_password", "password")

	viper.BindEnv("pg_ssl")
	viper.SetDefault("pg_ssl", "false")

	viper.BindEnv("pg_pool_max")
	viper.SetDefault("pg_pool_max", "30")

	viper.BindEnv("pg_pool_idle")
	viper.SetDefault("pg_pool_idle", "10")

	//
	// MailGun
	//

	viper.BindEnv("mail_key")
	viper.SetDefault(
		"mail_key",
		"key-5nz98-8e6edsuw3gr9jphc1x8l2vpri4")

	viper.BindEnv("mail_domain")
	viper.SetDefault(
		"mail_domain",
		"sandboxf4743199e9ba4a069d656b6e4fe40b19.mailgun.org")

	viper.BindEnv("mail_sender")
	viper.SetDefault(
		"mail_sender",
		"Test User <me@sandboxf4743199e9ba4a069d656b6e4fe40b19.mailgun.org>")

	//
	// oauth
	//

	viper.BindEnv("oauth_provider_google_key")
	viper.SetDefault("oauth_provider_google_key", "763925452628-6ujmpo0aa5gvrtnnjn7rleaok7tqvdi2.apps.googleusercontent.com")

	viper.BindEnv("oauth_provider_google_secret")
	viper.SetDefault("oauth_provider_google_secret", "yS2NQ0SgU8_FZ4fy39gm4Dht")

	//
	// log
	//

	viper.BindEnv("log_format")
	viper.SetDefault("log_format", "text")

	viper.BindEnv("log_level")
	viper.SetDefault("log_level", "debug")

	//
	// payment
	//

	viper.BindEnv("payment_publishable_key")
	viper.SetDefault("payment_publishable_key", "pk_test_nTWT6X97pIQHuQk8k8XYkfL1")

	viper.BindEnv("payment_secret_key")
	viper.SetDefault("payment_secret_key", "sk_test_ANrSNC5Yy2xAenilHrdGW9Lw")

}

func Configure() {
	if started {
		log.Fatal("System already started.")
	}
	started = true

	if viper.GetBool("dev_mode") == false {
		gin.SetMode(gin.ReleaseMode)
	}

	switch viper.GetString("log_format") {
	case "text":
		log.SetFormatter(&log.TextFormatter{})
	case "json":
		log.SetFormatter(&log.JSONFormatter{})
	default:
		log.WithField("format", viper.GetString("log_format")).Fatal("Unknow log format")
	}

	switch viper.GetString("log_level") {
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "warning":
		log.SetLevel(log.WarnLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	case "fatal":
		log.SetLevel(log.FatalLevel)
	case "panic":
		log.SetLevel(log.PanicLevel)
	default:
		log.WithField("level", viper.GetString("log_level")).Fatal("Unknow log level")
	}

	cache.Configure()
	captcha.Configure()
	cookies.Configure()
	oauth.Configure()
	payment.Configure()
	postgres.Configure()
	urls.Configure()
}
