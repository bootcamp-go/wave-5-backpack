package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	StorageDB *sql.DB
)

func init() {
	dataSource := "root@tcp(localhost:3306)/storage"
	//Open inicia un pool de conexiones. Solo abrir una vez
	var err error
	StorageDB, err = sql.Open("mysql", dataSource)
	if err != nil {
		log.Fatal(err)
	}
	if err = StorageDB.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("database configured")

}
