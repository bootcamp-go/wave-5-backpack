package util

import (
	"database/sql"

	"github.com/DATA-DOG/go-txdb"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	txdb.Register("txdb", "mysql", "root:@tcp(localhost:3306)/storage")
}

func InitDb() (*sql.DB, error) {
	db, err := sql.Open("txdb", "root:@tcp(localhost:3306)/storage")
	if err == nil {
		return db, db.Ping()
	}

	return db, err
}
