package store

import (
	"database/sql"
	"log"
	"os"

	"github.com/DATA-DOG/go-txdb"
	"github.com/google/uuid"
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

func RunTxdb() (*sql.DB, error) {
	txdb.Register("txdb", "mysql", "root@tcp(localhost:3306)/storage")
	db, err := sql.Open("txdb", uuid.New().String())

	if err != nil {
		return nil, err
	}

	return db, nil
}
