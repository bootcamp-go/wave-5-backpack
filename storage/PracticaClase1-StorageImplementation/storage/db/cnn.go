package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func MySQLConnection() *sql.DB {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", os.Getenv("USERNAME"), os.Getenv("PASSWORD"), os.Getenv("DATABASE")))
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	return db
}
