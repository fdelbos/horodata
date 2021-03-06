package config

import (
	"dev.hyperboloide.com/fred/horodata/services/cache"
	"dev.hyperboloide.com/fred/horodata/services/captcha"
	"dev.hyperboloide.com/fred/horodata/services/cookies"
	log_service "dev.hyperboloide.com/fred/horodata/services/log"
	"dev.hyperboloide.com/fred/horodata/services/mail"
	"dev.hyperboloide.com/fred/horodata/services/oauth"
	"dev.hyperboloide.com/fred/horodata/services/payment"
	"dev.hyperboloide.com/fred/horodata/services/picture"
	"dev.hyperboloide.com/fred/horodata/services/postgres"
	"dev.hyperboloide.com/fred/horodata/services/urls"
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
	viper.SetDefault("www_root", "http://localhost:3000")

	viper.BindEnv("api_root")
	viper.SetDefault("api_root", "http://localhost:3000/api/v1")

	viper.BindEnv("static_root")
	viper.SetDefault("static_root", "http://localhost:3000/static")

	viper.BindEnv("www_angular_base")
	viper.SetDefault("www_angular_base", "/app/")

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
	// Email
	//

	viper.BindEnv("mail_queue_name")
	viper.SetDefault("mail_queue_name", "mails")

	viper.BindEnv("mail_queue_host")
	viper.SetDefault("mail_queue_host", "amqp://guest:guest@localhost:5672/")

	//
	// oauth
	//

	viper.BindEnv("oauth_provider_google_key")
	viper.SetDefault(
		"oauth_provider_google_key",
		"682921687076-lgsos35csnfvq53n6dv95qpuf1pkg08j.apps.googleusercontent.com")

	viper.BindEnv("oauth_provider_google_secret")
	viper.SetDefault("oauth_provider_google_secret", "q-cA7pU5KENQT5ImkmfVMsEG")

	viper.BindEnv("oauth_provider_facebook_key")
	viper.SetDefault("oauth_provider_facebook_key", "583125351864412")

	viper.BindEnv("oauth_provider_facebook_secret")
	viper.SetDefault("oauth_provider_facebook_secret", "6f398a76151645884bb85a21589130db")

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
	viper.SetDefault("payment_publishable_key", "pk_test_Qepx3Si0RvnDab2jPoFq4fZw")

	viper.BindEnv("payment_secret_key")
	viper.SetDefault("payment_secret_key", "sk_test_ksm8vhIDWTyGCzDpHulwPF6l")

	viper.BindEnv("payment_endpoint")
	viper.SetDefault("payment_endpoint", "X8pwmpajDokdQsbQpw9UUu9oRb8C6Ui9Gg8s99XVAv7nwmdjVL")

	viper.BindEnv("payment_queue_name")
	viper.SetDefault("payment_queue_name", "invoicing")

	viper.BindEnv("payment_queue_host")
	viper.SetDefault("payment_queue_host", "amqp://guest:guest@localhost:5672/")

	//
	// Profile Pictures
	//

	viper.BindEnv("profile_pictures")
	viper.SetDefault("profile_pictures", "/tmp/horodata/profiles")

	//
	// Export
	//

	viper.BindEnv("export_service")
	viper.SetDefault("export_service", "http://localhost:5000")

	//
	// Invoicing
	//

	viper.BindEnv("invoicing_output")
	viper.SetDefault("invoicing_output", "/tmp/invoices.horodata/")

}

func Configure() {
	if started {
		log.Fatal("System already started.")
	}
	started = true

	if viper.GetBool("dev_mode") == false {
		gin.SetMode(gin.ReleaseMode)
	}

	log_service.Configure()
	cache.Configure()
	captcha.Configure()
	cookies.Configure()
	mail.Configure()
	oauth.Configure()
	payment.Configure()
	picture.Configure()
	postgres.Configure()
	urls.Configure()
}
