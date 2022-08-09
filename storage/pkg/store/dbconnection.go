package store

import (
	"database/sql"
	"log"
	"os"
)

func DBConnection() *sql.DB {
	dataSource := os.Getenv("DATASOURCE")

	storageDB, err := sql.Open("mysql", dataSource)
	if err != nil {
		log.Fatal("ERROR")
	}

	if err = storageDB.Ping(); err != nil {
		log.Fatal("ERROR PING")
	}

	return storageDB
}
