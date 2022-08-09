package test

import (
	"clase2-storage-implementation-tm/internal/domain"
	"clase2-storage-implementation-tm/internal/transactions"
	"database/sql"
	"log"
	"testing"

	"github.com/go-playground/assert/v2"
	_ "github.com/go-sql-driver/mysql"
)

// TestEcommerce ...
func TestEcommerce(t *testing.T) {
	var (
		StorageDB *sql.DB
		err       error
	)
	transaction := domain.Transaction{
		CodigoTransaccion: "test",
	}

	dataSource := "root@tcp(localhost:3306)/storage"
	// Open inicia un pool de conexiones. Sólo abrir una vez
	StorageDB, err = sql.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}

	if err := StorageDB.Ping(); err != nil {
		panic(err)
	}
	log.Println("database Configured")

	myRepo := transactions.NewRepository(StorageDB)
	productResult, err := myRepo.Ecommerce(transaction)
	if err != nil {
		log.Println(err)
	}
	assert.Equal(t, transaction.CodigoTransaccion, productResult.CodigoTransaccion)
}
