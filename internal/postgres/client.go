package postgres

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
)

type PostgresSQLDB struct {
	*sql.DB
}

func NewDBConnect() (*PostgresSQLDB, error) {

	dataSourceName, found := os.LookupEnv("SQLCONNECT")
	if !found {
		log.Panic("environment variable SQLCONNECT not found in .env")
	}

	bd, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return &PostgresSQLDB{}, err
	}

	return &PostgresSQLDB{bd}, nil
}
