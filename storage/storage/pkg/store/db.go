package store

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var (
	StorageDB *sql.DB
)

func Init() {
	dataSource := "root@tcp(localhost:3306)/storage"
	var err error
	StorageDB, err = sql.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	if err = StorageDB.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("DATABASE OK")
}
