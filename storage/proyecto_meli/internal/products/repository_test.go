package products

import (
	"context"
	"database/sql"
	"proyecto_meli/internal/domain"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/DATA-DOG/go-txdb"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/storage")
	if err != nil {
		t.Fatal(err)
	}
	repo := NewRepository(db)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err = repo.GetAll(ctx)
	if err != nil {
		t.Errorf("err must be nil, but got %v", err)
	}
}

func TestStorageTXDB(t *testing.T) {
	//Arrange
	txdb.Register("txdb", "mysql", "root@tcp(localhost:3306)/storage")
	db, err := sql.Open("txdb", uuid.New().String())
	if err != nil {
		t.Fatal(err)
	}
	repo := NewRepository(db)
	product := domain.Product{
		Nombre:        "new product",
		Color:         "white",
		Precio:        12000,
		Stock:         10,
		Codigo:        "NEW01",
		Publicado:     true,
		FechaCreacion: "2022-01-01 00:00:00",
	}
	//Act
	p, err := repo.Store(context.TODO(), product)
	//Assert
	assert.NoError(t, err)
	assert.NotZero(t, p)
	//Act
	p2, err := repo.GetById(context.TODO(), p.Id)
	//Assert
	assert.NoError(t, err)
	assert.NotZero(t, p2)
	assert.Equal(t, p, p2)
}

func TestUpdateTXDB(t *testing.T) {
	//Arrabge
	txdb.Register("txdb", "mysql", "root@tcp(localhost:3306)/storage")
	db, err := sql.Open("txdb", uuid.New().String())
	if err != nil {
		t.Fatal(err)
	}
	repo := NewRepository(db)
	product := domain.Product{
		Nombre:        "new product",
		Color:         "white",
		Precio:        12000,
		Stock:         10,
		Codigo:        "NEW01",
		Publicado:     true,
		FechaCreacion: "2022-01-01 00:00:00",
	}
	//Act
	p, err := repo.Store(context.TODO(), product)
	//Assert
	assert.NoError(t, err)
	assert.NotZero(t, p)
	//Arrange
	updateProduct := domain.Product{
		Id:            p.Id,
		Nombre:        "update product",
		Color:         "red",
		Precio:        120,
		Stock:         1,
		Codigo:        "UP01",
		Publicado:     false,
		FechaCreacion: "2022-01-02 00:00:00",
	}
	//Act
	p2, err := repo.Update(context.TODO(), updateProduct)
	assert.NoError(t, err)
	assert.NotZero(t, p2)
	assert.Equal(t, updateProduct, p2)
	//Act
	err = repo.Delete(context.TODO(), p.Id)
	//Assert
	assert.NoError(t, err)
	//Act
	p3, err := repo.GetById(context.TODO(), p.Id)
	assert.Error(t, err)
	assert.Zero(t, p3)
}

func TestStorageMock(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	mock.ExpectPrepare("INSERT INTO products")
	mock.ExpectExec("INSERT INTO products").WillReturnResult(sqlmock.NewResult(1, 1))

	productId := 1
	repo := NewRepository(db)
	product := domain.Product{
		Id:            productId,
		Nombre:        "new product",
		Color:         "white",
		Precio:        12000,
		Stock:         10,
		Codigo:        "NEW01",
		Publicado:     true,
		FechaCreacion: "2022-01-01 00:00:00",
	}
	//Act
	p, err := repo.Store(context.TODO(), product)
	assert.NoError(t, err)
	assert.NotZero(t, p)
	assert.NoError(t, mock.ExpectationsWereMet())
}
