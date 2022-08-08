package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	StorageDB *sql.DB
)

func ViendoqueOnda() error {
	return nil
}

func init() {
	dataSource := "root:12345678@tcp(localhost:3306)/storage"

	var err error
	StorageDB, err := sql.Open("mysql", dataSource)
	if err != nil {
		log.Fatal(err)
	}
	if err = StorageDB.Ping(); err != nil {
		log.Fatal(err)
	}
	log.Println("database configured")

}
