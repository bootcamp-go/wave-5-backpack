package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func InitDatabase() *sql.DB {
	dataSource := "meli_sprint_user:Meli_Sprint#123@/transactions?parseTime=true"

	storage, err := sql.Open("mysql", dataSource)

	if err != nil {
		panic(err)
	}

	if err = storage.Ping(); err != nil {
		panic(err)
	}

	return storage
}
