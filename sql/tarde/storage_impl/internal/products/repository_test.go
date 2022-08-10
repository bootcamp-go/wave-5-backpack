package products

import (
	"context"
	"testing"

	cnn "github.com/bootcamp-go/storage/db"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func init() {
	err := godotenv.Load("./../../.env")
	if err != nil {
		panic("can't connect to database")
	}
}

/* Ejercicio 2 - Implementar TestGetAll()
Diseñar un test que permita comprobar el método GetAll.
    1. Dentro del archivo repository_test.go desarrollar el método de prueba.
    2. Comprobar el correcto funcionamiento. */
func TestGetAll(t *testing.T) {
	db := cnn.MySQLConnection()
	repo := NewRepository(db)

	// recordar tener el producto con id 1 y tener un warehouse asociado a este producto
	products, err := repo.GetAll(context.Background(), 1)

	assert.NoError(t, err)
	assert.True(t, len(products) > 0)
}
