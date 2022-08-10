package products

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/DATA-DOG/go-txdb"
	mock "github.com/bootcamp-go/wave-5-backpack/tree/Ramos_Andres/goweb/practica/mocks"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func init() {
	txdb.Register("txdb", "mysql", "root@/storage")
}

func newDBMock(t *testing.T) (*sql.DB, sqlmock.Sqlmock) {
	product := mock.MockProduct
	// sqlRes, err := stmt.ExecContext(ctx, product.Name, product.Color, product.Price, product.Stock, product.Code, product.Published, product.Warehouse_id)
	db, mock, err := sqlmock.New()
	prep := mock.ExpectPrepare("^INSERT INTO products*")
	prep.ExpectExec().WithArgs(
		product.Name,
		product.Color,
		product.Price,
		product.Stock,
		product.Code,
		product.Published,
		product.Warehouse_id,
	).WillReturnResult(sqlmock.NewResult(1, 1))

	assert.NoError(t, err)
	return db, mock
}

func newTxDB(t *testing.T) *sql.DB {
	db, err := sql.Open("txdb", uuid.New().String())
	assert.Nil(t, err)
	assert.Nil(t, db.Ping())
	return db
}

func TestGetAll(t *testing.T) {
	db, err := sql.Open("mysql", "root@/storage")
	assert.Nil(t, err)
	repo := NewRepository(db)
	assert.Nil(t, err)
	ctx, close := context.WithTimeout(context.Background(), time.Second*5)
	defer close()
	res, err := repo.GetAll(ctx)
	assert.Nil(t, err)
	assert.Equal(t, mock.MockProductList, res)
}

func TestStore(t *testing.T) {
	db, err := sql.Open("mysql", "root@/storage")
	assert.Nil(t, err)
	repo := NewRepository(db)

	mockProduct := mock.MockProduct
	ctx, close := context.WithTimeout(context.Background(), time.Second*5)
	defer close()
	res, err := repo.Store(ctx, mockProduct)
	mockProduct.Id = res.Id
	assert.Nil(t, err)
	assert.Equal(t, mockProduct, res)
}

func TestGetByName(t *testing.T) {
	db, err := sql.Open("mysql", "root@/storage")
	assert.Nil(t, err)
	repo := NewRepository(db)
	assert.Nil(t, err)
	ctx, close := context.WithTimeout(context.Background(), time.Second*5)
	defer close()
	res, err := repo.GetByName(ctx, "product 1")
	assert.Nil(t, err)
	assert.Equal(t, mock.MockProductOne, res)
}

func TestStoreTXDB(t *testing.T) {
	db := newTxDB(t)
	repo := NewRepository(db)

	mockProduct := mock.MockProduct
	ctx, close := context.WithTimeout(context.Background(), time.Second*5)
	defer close()
	res, err := repo.Store(ctx, mockProduct)
	mockProduct.Id = res.Id
	assert.Nil(t, err)
	assert.Equal(t, mockProduct, res)
}

func TestGetOneTXDB(t *testing.T) {
	db := newTxDB(t)

	repo := NewRepository(db)
	ctx, close := context.WithTimeout(context.Background(), time.Second*5)
	defer close()
	res, err := repo.GetById(ctx, 1)
	assert.Nil(t, err)
	assert.Equal(t, mock.MockProductOne, res)
}

func TestGetOneNonExistantTXDB(t *testing.T) {
	db := newTxDB(t)

	repo := NewRepository(db)
	ctx, close := context.WithTimeout(context.Background(), time.Second*5)
	defer close()
	res, err := repo.GetById(ctx, 3)
	assert.Nil(t, err)
	assert.Equal(t, mock.MockProductEmpty, res)
}

func TestUpdateTXDB(t *testing.T) {
	db := newTxDB(t)
	repo := NewRepository(db)

	ctx, close := context.WithTimeout(context.Background(), time.Second*5)
	defer close()
	_, err := repo.Update(ctx, mock.MockProductUpdated)
	assert.Nil(t, err)
	res, err := repo.GetById(ctx, mock.MockProductUpdated.Id)
	assert.Nil(t, err)
	assert.Equal(t, mock.MockProductUpdated, res)
}

func TestDeleteTXDB(t *testing.T) {
	db := newTxDB(t)
	repo := NewRepository(db)

	ctx, close := context.WithTimeout(context.Background(), time.Second*5)
	defer close()
	err := repo.Delete(ctx, 1)
	assert.Nil(t, err)
	res, err := repo.GetById(ctx, 1)
	assert.Nil(t, err)
	assert.Equal(t, mock.MockProductEmpty, res)
}

func TestStoreMock(t *testing.T) {
	db, mockDB := newDBMock(t)
	defer db.Close()
	repo := NewRepository(db)

	mockProduct := mock.MockProduct
	ctx, close := context.WithTimeout(context.Background(), time.Second*5)
	defer close()
	res, err := repo.Store(ctx, mockProduct)
	mockProduct.Id = res.Id
	assert.Nil(t, err)
	assert.Equal(t, mockProduct, res)
	assert.NoError(t, mockDB.ExpectationsWereMet())
}
