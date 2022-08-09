package store

import (
	"database/sql"
	"log"
)

func DBConnection() *sql.DB {
	dataSource := "root@tcp(localhost:3306)/storage"

	storageDB, err := sql.Open("mysql", dataSource)
	if err != nil {
		log.Fatal("ERROR")
	}

	if err = storageDB.Ping(); err != nil {
		log.Fatal("ERROR PING")
	}

	return storageDB
}
