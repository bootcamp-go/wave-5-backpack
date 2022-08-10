package tests

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/DATA-DOG/go-txdb"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load("./../../.env")
	if err != nil {
		panic("can't connect to database")
	}

	txdb.Register("txdb", "mysql", fmt.Sprintf("%s:%s@/%s", os.Getenv("USERNAME"), os.Getenv("PASSWORD"), os.Getenv("DATABASE")))

}

func NewDBMock(t *testing.T) (*sql.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	return db, mock, nil
}

func InitDb() (*sql.DB, error) {
	db, err := sql.Open("txdb", uuid.New().String())

	if err == nil {
		return db, db.Ping()
	}

	return db, err
}
