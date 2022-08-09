package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func MySQLConnection() *sql.DB {
	//Implementacion storage
	// Open indica el pool de conexiones. Solo se debe abrir 1 VEZ.
	StorageDB, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", os.Getenv("USERNAME"), os.Getenv("PASSWORD"), os.Getenv("DATABASE")))
	if err != nil {
		log.Fatal(err)
	}
	if err = StorageDB.Ping(); err != nil {
		log.Fatal(err)
	}
	log.Println("database configured")
	return StorageDB
}