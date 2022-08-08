package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	StorageDB *sql.DB
)

func NewConnection(source string) (*sql.DB, error) {
	StorageDB, err := sql.Open("mysql", source)
	if err != nil {
		return nil, err
	}

	if err := StorageDB.Ping(); err != nil {
		return nil, err
	}
	log.Println("Open DB")

	return StorageDB, nil
}
