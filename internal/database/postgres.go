package database

import (
	"fmt"
	"net/url"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

const (
	maxIdleConnections = 2
	maxOpenConnections = 10
)

func BuildDSN(user string, password string, host string, port int, dbName string) string {
	dsn := url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(user, password),
		Host:   fmt.Sprintf("%s:%d", host, port),
		Path:   dbName,
	}

	return dsn.String()
}

func ConnectDB(dsn string) *sqlx.DB {
	db, err := sqlx.Connect("pgx", dsn)

	if err != nil {
		log.Fatal().Err(err).Msg("Unable to connect to database")
	}

	db.SetMaxIdleConns(maxIdleConnections)
	db.SetMaxOpenConns(maxOpenConnections)

	return db
}
