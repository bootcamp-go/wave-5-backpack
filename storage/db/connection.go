package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var (
	StorageDB  *sql.DB
	DataSource = "root@tcp(localhost:3306)/storage"
)

func Connection() *sql.DB {

	StorageDB, err := sql.Open("mysql", DataSource)
	if err != nil {
		panic(err)
	}

	if err := StorageDB.Ping(); err != nil {
		panic(err)
	}

	return StorageDB
}
