package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

func MySQLConnection() *sql.DB {
	cnxDb, err := sql.Open("mysql", fmt.Sprintf("%s:@/%s", os.Getenv("USERNAME"), os.Getenv("DATABASE")))
	if err != nil {
		panic(err)
	}
	if err := cnxDb.Ping(); err != nil {
		panic(err)
	}

	return cnxDb
}
