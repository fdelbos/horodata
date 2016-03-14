package postgres

import (
	"database/sql"
	"fmt"
	log "github.com/Sirupsen/logrus"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

var db *sql.DB

func Configure() {
	sslMode := "verify-full"
	if !viper.GetBool("pgssl") {
		sslMode = "disable"
	}
	connectionString := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s sslmode=%s",
		viper.GetString("pg_user"),
		viper.GetString("pg_password"),
		viper.GetString("pg_dbname"),
		viper.GetString("pg_host"),
		sslMode)

	if tmp, err := sql.Open("postgres", connectionString); err != nil {
		log.Fatal(err)
	} else {
		db = tmp
	}

	db.SetMaxOpenConns(viper.GetInt("pg_pool_max"))
	db.SetMaxIdleConns(viper.GetInt("pg_pool_idle"))

	if err := Ping(); err != nil {
		log.WithField("error", err).Fatal("Failed to ping PostgreSQL.")
	}

	log.WithFields(log.Fields{
		"host": viper.GetString("pg_host"),
		"db":   viper.GetString("pg_dbname"),
	}).Info("Connected to PostgreSQL.")
}
