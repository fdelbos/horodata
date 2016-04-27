package main

import (
	"time"

	log_service "dev.hyperboloide.com/fred/horodata/services/log"
	"dev.hyperboloide.com/fred/horodata/services/mail"
	"dev.hyperboloide.com/fred/horodata/services/postgres"
	log "github.com/Sirupsen/logrus"
	"github.com/hyperboloide/qmail/client"
	"github.com/spf13/viper"
)

func init() {
	viper.SetEnvPrefix("stats")

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
	viper.SetDefault("pg_pool_max", "3")

	viper.BindEnv("pg_pool_idle")
	viper.SetDefault("pg_pool_idle", "1")

	//
	// Email
	//

	viper.BindEnv("mail_queue_name")
	viper.SetDefault("mail_queue_name", "mails")

	viper.BindEnv("mail_queue_host")
	viper.SetDefault("mail_queue_host", "amqp://guest:guest@localhost:5672/")

	//
	// log
	//

	viper.BindEnv("log_format")
	viper.SetDefault("log_format", "text")

	viper.BindEnv("log_level")
	viper.SetDefault("log_level", "debug")
}

type rUser struct {
	Id       int64  `json:"id,omitempty"`
	Email    string `json:"email,omitempty"`
	FullName string `json:"name"`
}

func (u *rUser) Scan(scanFn func(dest ...interface{}) error) error {
	return scanFn(
		&u.Id,
		&u.Email,
		&u.FullName)
}

func main() {
	log_service.Configure()
	mail.Configure()
	postgres.Configure()

	date := time.Now().Format("02/01/2006")

	data := struct {
		TotalUsers  int64       `json:"total_users"`
		TotalPlans  interface{} `json:"total_plans"`
		TotalGroups int64       `json:"total_groups"`
		TotalGuests int64       `json:"total_guests"`
		TotalJobs   int64       `json:"total_jobs"`
		NewUsers    interface{} `json:"new_users"`
		Date        string      `json:"date"`
	}{
		totalUsers(),
		totalPlans(),
		totalGroups(),
		totalGuests(),
		totalJobs(),
		newUsers(),
		date,
	}
	mail.Mailer().Send(client.Mail{
		Dests:    []string{"fred@hyperboloide.com", "eric@hyperboloide.com"},
		Subject:  "Rapport du " + date,
		Template: "report",
		Data:     data,
	})
}

func totalUsers() int64 {
	const query = `select count(id) from users;`
	var res int64
	err := postgres.DB().QueryRow(query).Scan(&res)
	if err != nil {
		log.Fatal(err)
	}
	return res
}

func totalPlans() interface{} {
	const query = `select count(*) from stripe_subscriptions;`
	res := struct {
		Small  int64 `json:"small"`
		Medium int64 `json:"medium"`
		Large  int64 `json:"large"`
	}{0, 0, 0}

	err := postgres.DB().QueryRow(`select count(*) from stripe_subscriptions where plan = 'small';`).Scan(&res.Small)
	if err != nil {
		log.Fatal(err)
	}
	err = postgres.DB().QueryRow(`select count(*) from stripe_subscriptions where plan = 'medium';`).Scan(&res.Medium)
	if err != nil {
		log.Fatal(err)
	}
	err = postgres.DB().QueryRow(`select count(*) from stripe_subscriptions where plan = 'large';`).Scan(&res.Large)
	if err != nil {
		log.Fatal(err)
	}
	return res
}

func totalGroups() int64 {
	const query = `select count(*) from groups;`
	var res int64
	err := postgres.DB().QueryRow(query).Scan(&res)
	if err != nil {
		log.Fatal(err)
	}
	return res
}

func totalGuests() int64 {
	const query = `select count(*) from guests;`
	var res int64
	err := postgres.DB().QueryRow(query).Scan(&res)
	if err != nil {
		log.Fatal(err)
	}
	return res
}

func totalJobs() int64 {
	const query = `select count(*) from jobs;`
	var res int64
	err := postgres.DB().QueryRow(query).Scan(&res)
	if err != nil {
		log.Fatal(err)
	}
	return res
}

func newUsers() []rUser {
	const query = `
    select
		id, email, full_name
	from users
    where
		created > current_timestamp - interval '1 day';`

	rows, err := postgres.DB().Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	res := []rUser{}

	for rows.Next() {
		u := rUser{}
		if err := u.Scan(rows.Scan); err != nil {
			log.Fatal(err)
		}
		res = append(res, u)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return res
}
