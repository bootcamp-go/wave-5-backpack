package products

import (
	"Clase1_Parte1/productos/internal/domain"
	"context"
	"database/sql"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/DATA-DOG/go-txdb"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	_ "github.com/go-sql-driver/mysql"
)

// CLASE 2 PARTE 1

func TestGetAll(t *testing.T) {
	dataSource := "root@tcp(localhost:3306)/storage"
	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		t.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		t.Fatal(err)
	}
	repo := NewRepository(db)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	products, err := repo.GetAll(ctx)
	expectedProducts := []domain.Product{
		{Id: 4, Nombre: "Carro", Color: "Blanco", Precio: 10000, Stock: 10, Codigo: "A001", Publicado: true, FechaCreacion: "01-01-2020"},
		{Id: 7, Nombre: "Moto", Color: "Azul", Precio: 2000, Stock: 30, Codigo: "B001", Publicado: true, FechaCreacion: "02-02-2021"},
		{Id: 8, Nombre: "Carro", Color: "Negro", Precio: 12000, Stock: 20, Codigo: "C001", Publicado: true, FechaCreacion: "03-03-2022"},
	}
	assert.Equal(t, expectedProducts, products)
	assert.Nil(t, err)
}

func TestGetByIdWithContext(t *testing.T) {
	id := 4
	dataSource := "root@tcp(localhost:3306)/storage"
	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		t.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		t.Fatal(err)
	}
	repo := NewRepository(db)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	product, err := repo.GetByID(ctx, id)
	expectedProduct := domain.Product{Id: 4, Nombre: "Carro", Color: "Blanco", Precio: 10000, Stock: 10, Codigo: "A001", Publicado: true, FechaCreacion: "01-01-2020"}
	assert.Equal(t, expectedProduct, product)
	assert.Nil(t, err)
}

// USANDO SQL MOCK

func TestRepositoryGetById(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	productId := 4
	columns := []string{"id", "nombre", "color", "precio", "stock", "codigo", "publicado", "fecha_creacion"}
	rows := sqlmock.NewRows(columns)
	rows.AddRow(productId, "Carro", "Blanco", 10000, 10, "A001", true, "01-01-2020")

	mock.ExpectQuery(regexp.QuoteMeta(GetProdyctByID)).WillReturnRows(rows)

	repo := NewRepository(db)
	expectedProduct := domain.Product{
		Id:            productId,
		Nombre:        "Carro",
		Color:         "Blanco",
		Precio:        10000,
		Stock:         10,
		Codigo:        "A001",
		Publicado:     true,
		FechaCreacion: "01-01-2020",
	}

	ctx := context.TODO()
	p, err := repo.GetByID(ctx, productId)
	assert.NoError(t, err)
	assert.NotZero(t, p)
	assert.Equal(t, expectedProduct, p)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRepositoryStore(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectPrepare("INSERT INTO products")
	mock.ExpectExec("INSERT INTO products").WillReturnResult(sqlmock.NewResult(1, 1))

	productId := 1

	repo := NewRepository(db)
	product := domain.Product{
		Id:            productId,
		Nombre:        "Bicicleta",
		Color:         "Azul",
		Precio:        100,
		Stock:         50,
		Codigo:        "G000",
		Publicado:     true,
		FechaCreacion: "03-06-2015",
	}

	ctx := context.TODO()
	p, err := repo.Store(ctx, product.Nombre, product.Color, product.Precio, product.Stock, product.Codigo, product.Publicado, product.FechaCreacion)
	assert.NoError(t, err)
	assert.NotZero(t, p)
	assert.Equal(t, product.Id, p.Id)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRepositoryUpdate(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	productId := 1
	columns := []string{"id", "nombre", "color", "precio", "stock", "codigo", "publicado", "fecha_creacion"}
	rows := sqlmock.NewRows(columns)
	rows.AddRow(productId, "Carro", "Blanco", 10000, 10, "A001", true, "01-01-2020")

	mock.ExpectPrepare(regexp.QuoteMeta(UpdateProduct))
	mock.ExpectExec(regexp.QuoteMeta(UpdateProduct)).WillReturnResult(sqlmock.NewResult(1, 1))

	repo := NewRepository(db)
	updatedProduct := domain.Product{
		Id:            productId,
		Nombre:        "Carro1",
		Color:         "Blanco1",
		Precio:        11000,
		Stock:         11,
		Codigo:        "A002",
		Publicado:     false,
		FechaCreacion: "02-02-2022",
	}

	ctx := context.TODO()
	p, err := repo.Update(ctx, updatedProduct.Id, updatedProduct.Nombre, updatedProduct.Color, updatedProduct.Precio, updatedProduct.Stock, updatedProduct.Codigo, updatedProduct.Publicado, updatedProduct.FechaCreacion)
	assert.NoError(t, err)
	assert.NotZero(t, p)
	assert.Equal(t, updatedProduct, p)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRepositoryDelete(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	productId := 1
	columns := []string{"id", "nombre", "color", "precio", "stock", "codigo", "publicado", "fecha_creacion"}
	rows := sqlmock.NewRows(columns)
	rows.AddRow(productId, "Carro", "Blanco", 10000, 10, "A001", true, "01-01-2020")

	mock.ExpectPrepare(regexp.QuoteMeta(DeleteProduct))
	mock.ExpectQuery(regexp.QuoteMeta(GetProdyctByID)).WillReturnRows(rows)
	mock.ExpectExec(regexp.QuoteMeta(DeleteProduct)).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectQuery(regexp.QuoteMeta(GetProdyctByID)).WillReturnRows(rows)
	mock.ExpectQuery(GetAllProducts).WillReturnRows(rows)

	repo := NewRepository(db)
	deletedProduct := domain.Product{
		Id:            productId,
		Nombre:        "Carro",
		Color:         "Blanco",
		Precio:        10000,
		Stock:         10,
		Codigo:        "A001",
		Publicado:     true,
		FechaCreacion: "01-01-2020",
	}

	ctx := context.TODO()
	p, err := repo.Delete(ctx, deletedProduct.Id)
	p2, err2 := repo.GetByID(ctx, deletedProduct.Id)
	p3, err3 := repo.GetAll(ctx)
	assert.NoError(t, err)
	assert.NotZero(t, p)
	assert.Equal(t, deletedProduct, p)
	assert.NoError(t, err2)
	assert.Zero(t, p2)
	assert.NoError(t, err3)
	assert.NotContains(t, p3, p)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRepositoryGetByIdFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	productId := 1
	columns := []string{"id", "nombre", "color", "precio", "stock", "codigo", "publicado", "fecha_creacion"}
	rows := sqlmock.NewRows(columns)

	mock.ExpectQuery(regexp.QuoteMeta(GetProdyctByID)).WillReturnRows(rows)

	repo := NewRepository(db)

	ctx := context.TODO()
	p, err := repo.GetByID(ctx, productId)
	assert.NoError(t, err)
	assert.Zero(t, p)
	assert.Equal(t, domain.Product{}, p)
	assert.NoError(t, mock.ExpectationsWereMet())
}

// GO TXDB

func TestRepositoryGetByIdTXDB(t *testing.T) {
	dataSource := "root@tcp(localhost:3306)/storage"
	txdb.Register("txdb", "mysql", dataSource)
	db, err := sql.Open("txdb", uuid.New().String())
	assert.NoError(t, err)

	repo := NewRepository(db)
	productId := 4
	expectedProduct := domain.Product{
		Id:            productId,
		Nombre:        "Carro",
		Color:         "Blanco",
		Precio:        10000,
		Stock:         10,
		Codigo:        "A001",
		Publicado:     true,
		FechaCreacion: "01-01-2020",
	}

	ctx := context.TODO()
	p, err := repo.GetByID(ctx, productId)
	assert.NoError(t, err)
	assert.NotZero(t, p)
	assert.Equal(t, expectedProduct, p)
}

func TestRepositoryStoreTXDB(t *testing.T) {
	dataSource := "root@tcp(localhost:3306)/storage"
	txdb.Register("txdb_store", "mysql", dataSource)
	db, err := sql.Open("txdb_store", uuid.New().String())
	assert.NoError(t, err)

	repo := NewRepository(db)
	product := domain.Product{
		Nombre:        "Bicicleta",
		Color:         "Azul",
		Precio:        100,
		Stock:         50,
		Codigo:        "G000",
		Publicado:     true,
		FechaCreacion: "03-06-2015",
	}

	ctx := context.TODO()
	p, err := repo.Store(ctx, product.Nombre, product.Color, product.Precio, product.Stock, product.Codigo, product.Publicado, product.FechaCreacion)
	assert.NoError(t, err)
	assert.NotZero(t, p)
}

func TestRepositoryUpdateTXDB(t *testing.T) {
	dataSource := "root@tcp(localhost:3306)/storage"
	txdb.Register("txdb_update", "mysql", dataSource)
	db, err := sql.Open("txdb_update", uuid.New().String())
	assert.NoError(t, err)

	repo := NewRepository(db)
	updatedProduct := domain.Product{
		Id:            4,
		Nombre:        "Carro1",
		Color:         "Blanco1",
		Precio:        11000,
		Stock:         11,
		Codigo:        "A002",
		Publicado:     false,
		FechaCreacion: "02-02-2022",
	}

	ctx := context.TODO()
	p, err := repo.Update(ctx, updatedProduct.Id, updatedProduct.Nombre, updatedProduct.Color, updatedProduct.Precio, updatedProduct.Stock, updatedProduct.Codigo, updatedProduct.Publicado, updatedProduct.FechaCreacion)
	assert.NoError(t, err)
	assert.NotZero(t, p)
	assert.Equal(t, updatedProduct, p)
}

func TestRepositoryDeleteTXDB(t *testing.T) {
	dataSource := "root@tcp(localhost:3306)/storage"
	txdb.Register("txdb_delete", "mysql", dataSource)
	db, err := sql.Open("txdb_delete", uuid.New().String())
	assert.NoError(t, err)

	repo := NewRepository(db)
	deletedProduct := domain.Product{
		Id:            4,
		Nombre:        "Carro",
		Color:         "Blanco",
		Precio:        10000,
		Stock:         10,
		Codigo:        "A001",
		Publicado:     true,
		FechaCreacion: "01-01-2020",
	}

	ctx := context.TODO()
	p, err := repo.Delete(ctx, deletedProduct.Id)
	p2, err2 := repo.GetByID(ctx, deletedProduct.Id)
	p3, err3 := repo.GetAll(ctx)
	assert.NoError(t, err)
	assert.NotZero(t, p)
	assert.Equal(t, deletedProduct, p)
	assert.NoError(t, err2)
	assert.Zero(t, p2)
	assert.NoError(t, err3)
	assert.NotContains(t, p3, p)
}

func TestRepositoryGetByIdFailTXDB(t *testing.T) {
	dataSource := "root@tcp(localhost:3306)/storage"
	txdb.Register("txdb_id_fail", "mysql", dataSource)
	db, err := sql.Open("txdb_id_fail", uuid.New().String())
	assert.NoError(t, err)

	repo := NewRepository(db)
	productId := 0

	ctx := context.TODO()
	p, err := repo.GetByID(ctx, productId)
	assert.NoError(t, err)
	assert.Zero(t, p)
	assert.Equal(t, domain.Product{}, p)
}
