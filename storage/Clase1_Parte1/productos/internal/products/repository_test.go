package products

import (
	"Clase1_Parte1/productos/internal/domain"
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	_ "github.com/go-sql-driver/mysql"
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

/* func Test_sqlRepository_Store(t *testing.T) {
	db, err := util.InitDB()
	assert.NoError(t, err)
	repo := NewRepository(db)
	ctx := context.TODO()
	productId := uuid.New()
	product := domain.Product{
		UUID: productId,
	}
	new_product, err := repo.Store(ctx, &product.Nombre, &product.Color, &product.Precio, &product.Stock, &product.Codigo, &product.Publicado, &product.FechaCreacion)
	assert.NoError(t, err)
	getResult, err := repo.GetByID(ctx, uuid.New())
	assert.NoError(t, err)
	assert.Nil(t, getResult)
	getResult, err = repo.GetByID(ctx, productId)
	assert.NoError(t, err)
	assert.NotNil(t, getResult)
	assert.Equal(t, product.Id, getResult.Id)
} */
