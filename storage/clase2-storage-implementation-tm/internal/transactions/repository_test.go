package transactions

import (
	"clase2-storage-implementation-tm/internal/domain"
	"context"
	"database/sql"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func TestGetOneWithContext(t *testing.T) {
	id := 6
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/storage")
	if err != nil {
		t.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		t.Fatal(err)
	}
	repo := NewRepository(db)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err = repo.GetOneWithContext(ctx, id)
	if err != nil {
		t.Errorf("err must be nil, but got %v", err)
	}
}

func TestGetAll(t *testing.T) {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/storage")
	if err != nil {
		t.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		t.Fatal(err)
	}
	repository := NewRepository(db)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	transactions, err := repository.GetAll(ctx)
	expectedTransactions := []domain.Transaction{
		{ID: 1, CodigoTransaccion: "abc", Moneda: "JPY", Monto: 1012.76, Emisor: "SMFG", Receptor: "Mitsubishi UFJ", Fecha: "2018-06-24"},
		{ID: 2, CodigoTransaccion: "cde", Moneda: "EUR", Monto: 983.07, Emisor: "Lloyds Banking", Receptor: "Deutsche Bank-Rg", Fecha: "2019-05-11"},
		{ID: 3, CodigoTransaccion: "efg", Moneda: "MXN", Monto: 2302.75, Emisor: "BBVA", Receptor: "Banorte", Fecha: "2017-12-24"},
		{ID: 4, CodigoTransaccion: "hij", Moneda: "USD", Monto: 1012.76, Emisor: "Bank Of America", Receptor: "Morgan Stanley", Fecha: "2016-04-07"},
		{ID: 5, CodigoTransaccion: "klm", Moneda: "KPW", Monto: 1012.76, Emisor: "Shang Pudong", Receptor: "Kb Financial Gro", Fecha: "2015-07-23"},
		{ID: 6, CodigoTransaccion: "nop", Moneda: "CHF", Monto: 1012.76, Emisor: "Bankcomm-H", Receptor: "Bank Of Ningbo", Fecha: "2018-06-24"},
	}

	//Validation
	assert.Nil(t, err)
	assert.Equal(t, transactions, expectedTransactions)
}
