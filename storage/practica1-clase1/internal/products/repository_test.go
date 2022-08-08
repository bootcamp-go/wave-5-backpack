package products

import (
	"database/sql"
	"practica1-clase1/internal/domain"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

var dataSource = "root@tcp(localhost:3306)/storage"

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

	productResult, err := myRepo.Store(product)
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
	if err != nil {
		t.Errorf("Error al obtener el producto: %v", err)
	}
	assert.Equal(t, "Producto 1", product.Name)
}
