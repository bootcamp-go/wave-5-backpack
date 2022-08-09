package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func NewConnection() (*sql.DB, error) {
	source := "root@/storage_bootcamp"
	dbConnection, err := sql.Open("mysql", source)
	if err != nil {
		return nil, err
	}

	if err := dbConnection.Ping(); err != nil {
		return nil, err
	}

	pool := dbConnection.Stats().OpenConnections
	log.Printf("Open DB, pools: %v\n", pool)

	return dbConnection, nil
}
