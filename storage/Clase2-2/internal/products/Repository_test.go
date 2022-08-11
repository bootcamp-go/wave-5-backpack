package products

import (
	"context"
	"database/sql"
	"storage/internal/domain"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/DATA-DOG/go-txdb"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

//Integration test
func TestInsert(t *testing.T) {
	dataSource := "root@tcp(localhost:3306)/storage"
	StorageDB, err := sql.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	if err = StorageDB.Ping(); err != nil {
		panic(err)
	}
	prod := domain.Product{
		Name:  "test",
		Type:  "tipe1",
		Count: 1,
		Price: 10.5,
	}
	repo := NewRepo(StorageDB)
	res, err := repo.Store(prod)

	assert.Nil(t, err)
	assert.Equal(t, prod.Name, res.Name)
}
func TestGetByName(t *testing.T) {
	dataSource := "root@tcp(localhost:3306)/storage"
	StorageDB, err := sql.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	if err = StorageDB.Ping(); err != nil {
		panic(err)
	}
	ExpectedProd := domain.Product{
		ID:    1,
		Name:  "test",
		Type:  "tipe1",
		Count: 1,
		Price: 10.5,
	}
	repo := NewRepo(StorageDB)
	res, err := repo.GetByName("test")

	assert.Nil(t, err)
	assert.Equal(t, ExpectedProd, res)
}
func TestGetAll(t *testing.T) {
	dataSource := "root@tcp(localhost:3306)/storage"
	StorageDB, err := sql.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	if err = StorageDB.Ping(); err != nil {
		panic(err)
	}
	ExpectedProds := []domain.Product{
		{
			ID:    1,
			Name:  "test",
			Type:  "tipe1",
			Count: 1,
			Price: 10.5,
		},
	}
	repo := NewRepo(StorageDB)
	res, err := repo.GetAll()

	assert.Nil(t, err)
	assert.Equal(t, ExpectedProds, res)
}
func TestUpdate(t *testing.T) {
	dataSource := "root@tcp(localhost:3306)/storage"
	StorageDB, err := sql.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	if err = StorageDB.Ping(); err != nil {
		panic(err)
	}
	prod := domain.Product{
		ID:    1,
		Name:  "test10",
		Type:  "tipe10",
		Count: 10,
		Price: 100.5,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	repo := NewRepo(StorageDB)
	res, err := repo.Update(ctx, prod)

	assert.Nil(t, err)
	assert.Equal(t, prod.Name, res.Name)
}

//Clase 2-2 unit test USIN TXBD
func TestInserTxbd(t *testing.T) {
	txdb.Register("txdb", "mysql", "root@tcp(localhost:3306)/storage")
	db, err := sql.Open("txdb", uuid.New().String())
	assert.NoError(t, err)
	prod := domain.Product{
		Name:  "test2",
		Type:  "tipe2",
		Count: 2,
		Price: 20.5,
	}
	repo := NewRepo(db)
	res, err := repo.Store(prod)

	assert.Nil(t, err)
	assert.Equal(t, prod.Name, res.Name)
}
func TestGeOneTxbd(t *testing.T) {
	txdb.Register("txdb", "mysql", "root@tcp(localhost:3306)/storage")
	db, err := sql.Open("txdb", uuid.New().String())
	assert.NoError(t, err)
	ExpectedProd := domain.Product{
		ID:    1,
		Name:  "test",
		Type:  "tipe1",
		Count: 1,
		Price: 10.5,
	}
	repo := NewRepo(db)
	res, err := repo.GetOne(1)

	assert.Nil(t, err)
	assert.Equal(t, ExpectedProd.Name, res.Name)
}
func TestUpdateTxbd(t *testing.T) {
	ctx := context.TODO()
	txdb.Register("txdb", "mysql", "root@tcp(localhost:3306)/storage")
	db, err := sql.Open("txdb", uuid.New().String())
	assert.NoError(t, err)
	ExpectedProd := domain.Product{
		ID:    1,
		Name:  "test10",
		Type:  "tipe10",
		Count: 10,
		Price: 100.5,
	}

	repo := NewRepo(db)
	res, err := repo.Update(ctx, ExpectedProd)

	assert.Nil(t, err)
	assert.Equal(t, ExpectedProd, res)
}
func TestDeleteTxbd(t *testing.T) {

	txdb.Register("txdb", "mysql", "root@tcp(localhost:3306)/storage")
	db, err := sql.Open("txdb", uuid.New().String())
	assert.NoError(t, err)

	repo := NewRepo(db)
	err = repo.Delete(1)

	assert.Nil(t, err)
}

//Clase 2-2 unit test USIN sqlmock
func TestInserMock(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	//mock.ExpectPrepare("INSERT INTO products(name, type, count, price) VALUES( ?, ?, ?, ? )")
	mock.ExpectPrepare("INSERT INTO products")
	mock.ExpectExec("INSERT INTO products").WillReturnResult(sqlmock.NewResult(1, 1))
	prod := domain.Product{
		ID:    1,
		Name:  "test2",
		Type:  "tipe2",
		Count: 2,
		Price: 20.5,
	}
	repo := NewRepo(db)
	res, err := repo.Store(prod)

	assert.Nil(t, err)
	assert.Equal(t, prod.ID, res.ID)
	assert.NoError(t, mock.ExpectationsWereMet())
}
