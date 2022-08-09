package products

import (
	"context"
	"database/sql"
	"fmt"
	"regexp"
	"storage/internal/domain"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/DATA-DOG/go-txdb"
	"github.com/stretchr/testify/assert"

	_ "github.com/go-sql-driver/mysql"
)

func TestUpdateWithContext(t *testing.T) {

	productTest := domain.Products{
		Nombre: "Sandia",
		Color:  "verdoso",
		Precio: 2132222,
		Stock:  2,
	}

	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/storage")
	if err != nil {
		t.Fatal(err)
	}
	repo := InitRepository(db)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = repo.Update(ctx, productTest)
	if err != nil {
		t.Errorf("error: %v", err)
	}
}

func TestGetAll(t *testing.T) {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/storage")
	if err != nil {
		t.Fatal(err)
	}

	expect := []domain.Products{
		{
			Id:            1,
			Nombre:        "Sandia",
			Color:         "verdoso",
			Precio:        2132220,
			Stock:         2,
			Codigo:        "21fe2",
			Publicado:     true,
			FechaCreacion: "23/10/2022",
		},
		{
			Id:            2,
			Nombre:        "Guayaba",
			Color:         "Amarillo",
			Precio:        323123,
			Stock:         24,
			Codigo:        "3efe2",
			Publicado:     true,
			FechaCreacion: "30/10/2022",
		},
		{
			Id:            5,
			Nombre:        "Pera",
			Color:         "Verde",
			Precio:        3213430,
			Stock:         43,
			Codigo:        "y2h76",
			Publicado:     true,
			FechaCreacion: "14/11/2022",
		},
	}

	repo := InitRepository(db)

	r, err := repo.GetAll()

	assert.Equal(t, err, nil)
	assert.Equal(t, expect, r)
}

func TestSqlRepositoryCreateMock(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectPrepare(regexp.QuoteMeta("INSERT INTO products(name, color, price, stock, code, publish, creation_date) VALUES(?, ?, ?, ?, ?, ?, ?)"))
	mock.ExpectExec("INSERT INTO products").WillReturnResult(sqlmock.NewResult(1, 1))

	productId := 1

	repo := InitRepository(db)
	product := domain.Products{
		Id:            productId,
		Nombre:        "Prueba",
		Color:         "negro",
		Precio:        12412412,
		Stock:         2,
		Codigo:        "23sd2",
		Publicado:     true,
		FechaCreacion: "23/10/2020",
	}

	p, err := repo.CreateProduct(product)
	assert.NoError(t, err)
	assert.NotZero(t, p)
	assert.Equal(t, productId, p.Id)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRepositoryCreateTXDB(t *testing.T) {

	txdb.Register("txdb", "mysql", "root@tcp(localhost:3306)/storage")

	db, err := sql.Open("txdb", "124123123123")
	assert.NoError(t, err)

	repo := InitRepository(db)
	product := domain.Products{
		Nombre:        "Prueba",
		Color:         "negro",
		Precio:        12412412,
		Stock:         2,
		Codigo:        "23sd2",
		Publicado:     true,
		FechaCreacion: "23/10/2020",
	}
	p, err := repo.CreateProduct(product)
	assert.NoError(t, err)
	assert.NotZero(t, p)

}

func TestRepositoryGetOneTXDB(t *testing.T) {
	txdb.Register("txdb", "mysql", "root@tcp(localhost:3306)/storage")

	db, err := sql.Open("txdb", "124123123123")
	assert.NoError(t, err)
	repo := InitRepository(db)
	idGet := 1

	p, err := repo.GetById(idGet)
	assert.NoError(t, err)
	assert.NotZero(t, p)
}

func TestRepositoryUpdateTXDB(t *testing.T) {
	txdb.Register("txdb", "mysql", "root@tcp(localhost:3306)/storage")
	ctx := context.TODO()
	db, err := sql.Open("txdb", "124123123123")
	assert.NoError(t, err)
	product := domain.Products{
		Nombre: "Prueba2",
		Color:  "Azul",
		Precio: 342112,
		Stock:  20,
	}

	repo := InitRepository(db)
	er := repo.Update(ctx, product)

	assert.Equal(t, nil, er)
}

func TestRepositoryDeleteTXDB(t *testing.T) {
	txdb.Register("txdb", "mysql", "root@tcp(localhost:3306)/storage")
	db, err := sql.Open("txdb", "124123123123")
	assert.NoError(t, err)
	idProduct := 2

	repo := InitRepository(db)
	er := repo.Delete(idProduct)
	fmt.Println(er)

	assert.Equal(t, nil, er)

}
