package transactions

import (
	"clase2-storage-implementation-tt/internal/domain"
	"context"
	"database/sql"
	"log"
	"testing"
	"time"

	"github.com/DATA-DOG/go-txdb"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

// TestEcommerce ...
func TestEcommerce(t *testing.T) {
	var (
		StorageDB *sql.DB
		err       error
	)
	transaction := domain.Transaction{}

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

	myRepo := NewRepository(StorageDB)
	productResult, err := myRepo.Ecommerce(transaction)
	if err != nil {
		log.Println(err)
	}
	assert.Equal(t, transaction.CodigoTransaccion, productResult.CodigoTransaccion)
}

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

func TestGetOneWithContextFail(t *testing.T) {
	id := 10
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
	// Validation
	assert.Error(t, err)
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

func TestUpdate(t *testing.T) {
	idSelected := 1

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

	req := domain.Transaction{
		CodigoTransaccion: "cba", Moneda: "USD", Monto: 233, Emisor: "Citigroup Inc",
		Receptor: "HSBC", Fecha: "2021-02-23",
	}

	beforeUpdate, err := repo.GetOne(idSelected)
	if err != nil {
		t.Errorf("err must be nil, but got %v", err)
	}

	result, err := repo.Update(ctx, idSelected, req.CodigoTransaccion,
		req.Moneda, req.Monto, req.Emisor, req.Receptor, req.Fecha)

	if err != nil {
		t.Errorf("err must be nil, but got %v", err)
	}

	// Validation
	assert.NotEqual(t, beforeUpdate, result)
}

func TestDelete(t *testing.T) {
	idSelected := 6

	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/storage")
	if err != nil {
		t.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		t.Fatal(err)
	}

	repo := NewRepository(db)
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	beforeDelete, err := repo.Delete(idSelected)
	if err != nil {
		t.Errorf("err must be nil, but got %v", err)
	}

	result, err := repo.Delete(idSelected)
	if err != nil {
		t.Errorf("err must be nil, but got %v", err)
	}

	// Validation
	assert.Empty(t, result)
	assert.NotEqual(t, beforeDelete, result)
}

func TestEcommerceTXDB(t *testing.T) {
	txdb.Register("txdb", "mysql", "root@tcp(localhost:3306)/storage") //nos conectamos a nuestra base de datos
	db, err := sql.Open("txdb", uuid.New().String())                   //sql.Open recibe el driver de base de datos y un string de conexion
	repo := NewRepository(db)                                          //Generamos nuestro repository
	transaction := domain.Transaction{
		CodigoTransaccion: "cde", Moneda: "EUR", Monto: 983.07,
		Emisor: "Lloyds Banking", Receptor: "Deutsche Bank-Rg", Fecha: "2019-05-114",
	}
	p, err := repo.Ecommerce(transaction) //consulta el repo.
	// Validation
	assert.NoError(t, err)
	assert.NotZero(t, p)
}

func TestGetOneTXDB(t *testing.T) {
	idSelected := 1
	txdb.Register("txdb", "mysql", "root@tcp(localhost:3306)/storage") //nos conectamos a nuestra base de datos
	db, err := sql.Open("txdb", uuid.New().String())                   //sql.Open recibe el driver de base de datos y un string de conexion
	repo := NewRepository(db)                                          //Generamos nuestro repository
	p, err := repo.GetOne(idSelected)                                  //consulta el repo.
	// Validation
	assert.NoError(t, err)
	assert.NotZero(t, p)
}
