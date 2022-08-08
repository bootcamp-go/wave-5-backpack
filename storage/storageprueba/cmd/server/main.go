package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var StorageDB *sql.DB

func main() {

	var err error
	StorageDB, err = sql.Open("mysql", "meli_sprint_user:Meli_Sprint#123@/storage")
	if err != nil {
		panic(err)
	}
	if err = StorageDB.Ping(); err != nil {
		panic(err)
	}
	log.Println("database configured")
}
