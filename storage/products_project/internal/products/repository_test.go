package products

import (
	"context"
	"database/sql"
	"products_project/internal/domain"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
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
