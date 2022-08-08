package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var StorageDB *sql.DB

func Init() {
	dataSource := "root:@tcp(localhost:3306)/storage"
	var err error
	StorageDB, err := sql.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}

	log.Println("database configured", &StorageDB)
}
