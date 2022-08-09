package products

import (
	"context"
	"database/sql"
	"products_project/internal/domain"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/DATA-DOG/go-txdb"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	dataSource := "root@tcp(localhost:3306)/storage"
	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		t.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		t.Fatal(err)
	}
	repository := NewRepository(db)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	products, err := repository.GetAll(ctx)
	expectedProducts := []domain.Product{
		{Id: 6, Nombre: "Carro", Color: "Azul", Precio: 300, Stock: 6, Codigo: "C3456", Publicado: true, Fecha: "23-04-12"},
		{Id: 8, Nombre: "Moto", Color: "Negro", Precio: 244, Stock: 5, Codigo: "M456", Publicado: true, Fecha: "11-05-21"},
	}
	assert.Nil(t, err)
	assert.Equal(t, expectedProducts, products)
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
		Id:        productId,
		Nombre:    "Avion",
		Color:     "Blanco",
		Precio:    788,
		Stock:     8,
		Codigo:    "AF456",
		Publicado: true,
		Fecha:     "02-09-22",
	}

	ctx := context.TODO()
	p, err := repo.Store(ctx, product.Nombre, product.Color, product.Precio, product.Stock, product.Codigo, product.Publicado, product.Fecha)
	assert.NoError(t, err)
	assert.NotZero(t, p)
	assert.Equal(t, product, p)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRepositoryStoreTXDB(t *testing.T) {
	dataSource := "root@tcp(localhost:3306)/storage"
	txdb.Register("txdb_store", "mysql", dataSource)
	db, err := sql.Open("txdb_store", uuid.New().String())
	assert.NoError(t, err)

	repo := NewRepository(db)
	product := domain.Product{
		Nombre:    "Avion",
		Color:     "Blanco",
		Precio:    788,
		Stock:     8,
		Codigo:    "AF456",
		Publicado: true,
		Fecha:     "02-09-22",
	}

	ctx := context.TODO()
	p, err := repo.Store(ctx, product.Nombre, product.Color, product.Precio, product.Stock, product.Codigo, product.Publicado, product.Fecha)
	assert.NoError(t, err)
	assert.NotZero(t, p)
}

func TestRepositoryGetById(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	productId := 6
	columns := []string{"id", "nombre", "color", "precio", "stock", "codigo", "publicado", "fecha"}
	rows := sqlmock.NewRows(columns)
	rows.AddRow(productId, "Carro", "Azul", 300, 6, "C3456", true, "23-04-12")
	mock.ExpectQuery(regexp.QuoteMeta(GetById)).WillReturnRows(rows)
	repo := NewRepository(db)
	expectedProduct := domain.Product{
		Id:        productId,
		Nombre:    "Carro",
		Color:     "Azul",
		Precio:    300,
		Stock:     6,
		Codigo:    "C3456",
		Publicado: true,
		Fecha:     "23-04-12",
	}

	ctx := context.TODO()
	p, err := repo.GetById(ctx, productId)
	assert.NoError(t, err)
	assert.NotZero(t, p)
	assert.Equal(t, expectedProduct, p)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRepositoryGetByIdTXDB(t *testing.T) {
	dataSource := "root@tcp(localhost:3306)/storage"
	txdb.Register("txdb_getbyid", "mysql", dataSource)
	db, err := sql.Open("txdb_getbyid", uuid.New().String())
	assert.NoError(t, err)

	repo := NewRepository(db)
	productId := 6
	expectedProduct := domain.Product{
		Id:        productId,
		Nombre:    "Carro",
		Color:     "Azul",
		Precio:    300,
		Stock:     6,
		Codigo:    "C3456",
		Publicado: true,
		Fecha:     "23-04-12",
	}

	ctx := context.TODO()
	p, err := repo.GetById(ctx, productId)
	assert.NoError(t, err)
	assert.NotZero(t, p)
	assert.Equal(t, expectedProduct, p)
}

func TestRepositoryUpdate(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	productId := 1
	columns := []string{"id", "nombre", "color", "precio", "stock", "codigo", "publicado", "fecha"}
	rows := sqlmock.NewRows(columns)
	rows.AddRow(productId, "Avion", "Blanco", 788, 8, "AF456", true, "02-09-22")
	mock.ExpectPrepare(regexp.QuoteMeta(Update))
	mock.ExpectExec(regexp.QuoteMeta(Update)).WillReturnResult(sqlmock.NewResult(1, 1))

	repo := NewRepository(db)
	updatedProduct := domain.Product{
		Id:        productId,
		Nombre:    "Avioneta",
		Color:     "Rojo",
		Precio:    800,
		Stock:     10,
		Codigo:    "Af41",
		Publicado: true,
		Fecha:     "05-06-20",
	}

	ctx := context.TODO()
	p, err := repo.Update(ctx, updatedProduct.Id, updatedProduct.Nombre, updatedProduct.Color, updatedProduct.Precio, updatedProduct.Stock, updatedProduct.Codigo, updatedProduct.Publicado, updatedProduct.Fecha)
	assert.NoError(t, err)
	assert.NotZero(t, p)
	assert.Equal(t, updatedProduct, p)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRepositoryUpdateTXDB(t *testing.T) {
	dataSource := "root@tcp(localhost:3306)/storage"
	txdb.Register("txdb_update", "mysql", dataSource)
	db, err := sql.Open("txdb_update", uuid.New().String())
	assert.NoError(t, err)

	repo := NewRepository(db)
	productId := 6
	updatedProduct := domain.Product{
		Id:        productId,
		Nombre:    "Carro",
		Color:     "Verde",
		Precio:    300,
		Stock:     6,
		Codigo:    "C3456",
		Publicado: true,
		Fecha:     "23-04-12",
	}

	ctx := context.TODO()
	p, err := repo.Update(ctx, updatedProduct.Id, updatedProduct.Nombre, updatedProduct.Color, updatedProduct.Precio, updatedProduct.Stock, updatedProduct.Codigo, updatedProduct.Publicado, updatedProduct.Fecha)
	assert.NoError(t, err)
	assert.NotZero(t, p)
	assert.Equal(t, updatedProduct, p)
}

func TestRepositoryDelete(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	productId := 1
	columns := []string{"id", "nombre", "color", "precio", "stock", "codigo", "publicado", "fecha_creacion"}
	rows := sqlmock.NewRows(columns)
	rows.AddRow(productId, "Carro", "Verde", 300, 10, "C3456", true, "23-04-12")

	mock.ExpectPrepare(regexp.QuoteMeta(Delete))
	mock.ExpectQuery(regexp.QuoteMeta(GetById)).WillReturnRows(rows)
	mock.ExpectExec(regexp.QuoteMeta(Delete)).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectQuery(regexp.QuoteMeta(GetById)).WillReturnRows(rows)
	mock.ExpectQuery(GetAllProducts).WillReturnRows(rows)

	repo := NewRepository(db)
	deletedProduct := domain.Product{
		Id:        productId,
		Nombre:    "Carro",
		Color:     "Verde",
		Precio:    300,
		Stock:     10,
		Codigo:    "C3456",
		Publicado: true,
		Fecha:     "23-04-12",
	}

	ctx := context.TODO()
	p, err := repo.Delete(ctx, deletedProduct.Id)
	assert.NoError(t, err)
	assert.NotZero(t, p)
	assert.Equal(t, deletedProduct, p)

	p2, err2 := repo.GetById(ctx, deletedProduct.Id)
	assert.NoError(t, err2)
	assert.Zero(t, p2)

	p3, err3 := repo.GetAll(ctx)
	assert.NoError(t, err3)
	assert.NotContains(t, p3, p)
	assert.NoError(t, mock.ExpectationsWereMet())
}
func TestRepositoryDeleteTXDB(t *testing.T) {
	dataSource := "root@tcp(localhost:3306)/storage"
	txdb.Register("txdb_delete", "mysql", dataSource)
	db, err := sql.Open("txdb_delete", uuid.New().String())
	assert.NoError(t, err)

	repo := NewRepository(db)
	deletedProduct := domain.Product{
		Id:        6,
		Nombre:    "Carro",
		Color:     "Azul",
		Precio:    300,
		Stock:     6,
		Codigo:    "C3456",
		Publicado: true,
		Fecha:     "23-04-12",
	}

	ctx := context.TODO()
	p, err := repo.Delete(ctx, deletedProduct.Id)
	assert.NoError(t, err)
	assert.NotZero(t, p)
	assert.Equal(t, deletedProduct, p)

	p2, err2 := repo.GetById(ctx, deletedProduct.Id)
	assert.NoError(t, err2)
	assert.Zero(t, p2)

	p3, err3 := repo.GetAll(ctx)
	assert.NoError(t, err3)
	assert.NotContains(t, p3, p)
}

func TestRepositoryGetByIdFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	productId := 1
	columns := []string{"id", "nombre", "color", "precio", "stock", "codigo", "publicado", "fecha_creacion"}
	rows := sqlmock.NewRows(columns)

	mock.ExpectQuery(regexp.QuoteMeta(GetById)).WillReturnRows(rows)

	repo := NewRepository(db)

	ctx := context.TODO()
	p, err := repo.GetById(ctx, productId)
	assert.NoError(t, err)
	assert.Zero(t, p)
	assert.Equal(t, domain.Product{}, p)
	assert.NoError(t, mock.ExpectationsWereMet())
}
func TestRepositoryGetByIdFailTXDB(t *testing.T) {
	dataSource := "root@tcp(localhost:3306)/storage"
	txdb.Register("txdb_getbyid_fail", "mysql", dataSource)
	db, err := sql.Open("txdb_getbyid_fail", uuid.New().String())
	assert.NoError(t, err)

	repo := NewRepository(db)
	productId := 1

	ctx := context.TODO()
	p, err := repo.GetById(ctx, productId)
	assert.NoError(t, err)
	assert.Zero(t, p)
	assert.Equal(t, domain.Product{}, p)
}
