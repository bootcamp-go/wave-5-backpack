package util

import (
	"database/sql"

	"github.com/DATA-DOG/go-txdb"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

func init() {
	dataSource := "root@tcp(localhost:3306)/storage"
	txdb.Register("txdb", "mysql", dataSource)
}

func InitDB() (*sql.DB, error) {
	db, err := sql.Open("txdb", uuid.New().String())

	if err == nil {
		return db, db.Ping()
	}
	return db, err
}
