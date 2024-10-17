package postgres

import (
	"fmt"

	"github.com/tarunngusain08/culturehub/config"
	"github.com/tarunngusain08/culturehub/pkg/log"
)

var logger = log.New("internal/db/postgres")

func dbHost() string { return config.GetString("database.host") }

func driver() string { return config.GetString("database.driver") }

func DBName() string {
	env := config.GetEnv().String()
	dbname := config.GetString("database.dbname")
	return fmt.Sprintf("%s_%s", env, dbname)
}

func dbUser() string { return config.GetString("database.user") }

func dbPort() int { return config.GetInt("database.port") }

func dbPassword() string { return config.GetString("database.password") }

func ConnectionStrWithDB() string {
	return fmt.Sprintf("%s database=%s", ConnectionStr(), DBName())
}

func ConnectionStr() string {
	str := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s",
		dbHost(),
		dbPort(),
		dbUser(),
		dbPassword(),
	)
	ssl := "sslmode=disable"
	if config.IsProd() {
		ssl = "sslmode=require"
	}
	return fmt.Sprintf("%s %s", str, ssl)
}
