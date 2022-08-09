package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func MySQLConnection() *sql.DB {
	dataSource := "root@tcp(localhost:3306)/storageC1TT"

	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	return db
}
