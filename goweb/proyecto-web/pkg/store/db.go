package store

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var storageDb *sql.DB

func NewDb() *sql.DB {
	dataSource := "meli_sprint_user:Meli_Sprint#123@tcp(localhost:3306)/storage"

	// Patrón Singleton
	if storageDb != nil {
		return storageDb
	}
	var err error
	storageDb, err = sql.Open("mysql", dataSource)

	if err != nil {
		panic(err)
	}

	if err := storageDb.Ping(); err != nil {
		panic(err)
	}

	return storageDb
}
