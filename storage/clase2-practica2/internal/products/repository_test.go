package products

import (
	"context"
	"database/sql"
	"fmt"
	"practica2-clase2/internal/domain"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/DATA-DOG/go-txdb"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var dataSource = "root@tcp(localhost:3306)/storage"

// -----------------------------------------------
// Clase 1 - Práctica 1
// -----------------------------------------------

func TestStore(t *testing.T) {
	product := domain.Product{
		Name:  "Producto 1",
		Type:  "Tipo 1",
		Count: 1,
		Price: 1.0,
	}
	StorageDB, err := sql.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	myRepo := NewRepository(StorageDB)

	productResult, err := myRepo.Store(context.TODO(), product)
	if err != nil {
		t.Errorf("Error al guardar el producto: %v", err)
	}
	assert.Equal(t, product.Name, productResult.Name)
}

func TestGetByName(t *testing.T) {
	StorageDB, err := sql.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	myRepo := NewRepository(StorageDB)

	product, err := myRepo.GetByName("Producto 1")
	assert.Nil(t, err)
	assert.Equal(t, "Producto 1", product.Name)
}

func TestGetAll(t *testing.T) {
	StorageDB, err := sql.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	myRepo := NewRepository(StorageDB)

	products, err := myRepo.GetAll(context.TODO())
	assert.Nil(t, err)
	assert.Equal(t, 1, len(products))
}

// -----------------------------------------------
// Clase 2 - Práctica 2
// -----------------------------------------------

// Ejercicios 1

func TestStoreTXDB(t *testing.T) {
	txdb.Register("txdb", "mysql", dataSource)
	db, err := sql.Open("txdb", uuid.New().String())
	assert.NoError(t, err)
	repo := NewRepository(db)

	ctx := context.TODO()
	product := domain.Product{
		Name:  "Producto 1",
		Type:  "Tipo 1",
		Count: 1,
		Price: 1.0,
	}

	p, err := repo.Store(ctx, product)
	product.ID = p.ID

	assert.NoError(t, err)
	assert.NotZero(t, p)
	getResult, err := repo.GetOne(ctx, p.ID)
	assert.NoError(t, err)
	assert.Equal(t, product.Name, getResult.Name)
	assert.Equal(t, product.ID, getResult.ID)
}

// Ejercicios 2
func TestUpdateTXDB(t *testing.T) {
	txdb.Register("txdb", "mysql", dataSource)
	db, err := sql.Open("txdb", uuid.New().String())
	if err != nil {
		panic(err)
	}
	repo := NewRepository(db)

	ctx := context.TODO()
	idProduct := 1
	product := domain.Product{
		ID:    idProduct,
		Name:  "Producto 1",
		Type:  "Tipo 1",
		Count: 1,
		Price: 1.0,
	}

	p, err := repo.Update(ctx, product)

	assert.NoError(t, err)
	assert.NotZero(t, p)
	getResult, err := repo.GetOne(ctx, p.ID)
	assert.NoError(t, err)
	assert.Equal(t, product.Name, getResult.Name)
	assert.Equal(t, product.ID, getResult.ID)
}

func TestDeleteTXDB(t *testing.T) {
	txdb.Register("txdb", "mysql", dataSource)
	db, err := sql.Open("txdb", uuid.New().String())
	if err != nil {
		panic(err)
	}
	repo := NewRepository(db)

	ctx := context.TODO()
	idProduct := 1

	err = repo.Delete(ctx, idProduct)
	assert.NoError(t, err)

	getResult, err := repo.GetOne(ctx, idProduct)
	assert.Error(t, err)
	assert.Zero(t, getResult)
}

// Ejercicio 3

func TestUpdateSQLMock(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	//mock.ExpectPrepare("UPDATE products SET")
	mock.ExpectPrepare(regexp.QuoteMeta("UPDATE products SET name = ?, type = ?, count = ?, price = ? WHERE id = ?"))
	mock.ExpectExec("UPDATE products SET").WillReturnResult(sqlmock.NewResult(1, 1))

	repo := NewRepository(db)

	product := domain.Product{
		ID:    1,
		Name:  "Producto 1",
		Type:  "Tipo 1",
		Count: 1,
		Price: 1.0,
	}

	p, err := repo.Update(context.TODO(), product)
	assert.NoError(t, err)
	assert.Equal(t, product.ID, p.ID)
	assert.Equal(t, product, p)
}

func TestDeleteSQLMock(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	//mock.ExpectPrepare("DELETE FROM products WHERE")
	mock.ExpectPrepare(regexp.QuoteMeta("DELETE FROM products WHERE id = ?"))
	mock.ExpectExec("DELETE FROM products WHERE").WillReturnResult(sqlmock.NewResult(1, 1))

	repo := NewRepository(db)

	err = repo.Delete(context.TODO(), 1)
	assert.NoError(t, err)
}

// Ejercicio 4

func TestUpdateSQLMockError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	//mock.ExpectPrepare("UPDATE products SET")
	mock.ExpectPrepare(regexp.QuoteMeta("UPDATE products SET name = ?, type = ?, count = ?, price = ? WHERE id = ?"))
	mock.ExpectExec("UPDATE products SET").WillReturnError(fmt.Errorf("error updating"))

	repo := NewRepository(db)

	product := domain.Product{
		ID:    1,
		Name:  "Producto 1",
		Type:  "Tipo 1",
		Count: 1,
		Price: 1.0,
	}
	errorExpected := fmt.Errorf("error updating")

	_, err = repo.Update(context.TODO(), product)
	assert.Error(t, errorExpected, err)
}
