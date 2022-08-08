package product

import (
	"database/sql"
	"log"
	"testing"

	"github.com/bootcamp-go/wave-5-backpack/tree/ospina_christian/storage/storageprueba/internal/domain"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

var (
	StorageDB *sql.DB
	err       error
)

func TestStore(t *testing.T) {
	producto := domain.Productos{
		Nombre:        "Maradona",
		Color:         "Verde",
		Precio:        124333,
		Stock:         11,
		Codigo:        "12asc",
		Publicado:     true,
		FechaCreación: "2021-02-15",
	}
	StorageDB, err = sql.Open("mysql", "meli_sprint_user:Meli_Sprint#123@/storage")

	assert.Nil(t, err)

	myrepo := NewRepo(StorageDB)
	productResult, err := myrepo.Store(producto)
	if err != nil {
		log.Println(err)
	}
	assert.Equal(t, producto.Nombre, productResult.Nombre)
}

func TestGetOne(t *testing.T) {
	producto := domain.Productos{
		Nombre:        "Maradona",
		Color:         "Verde",
		Precio:        124333,
		Stock:         11,
		Codigo:        "12asc",
		Publicado:     true,
		FechaCreación: "2021-02-15",
	}

	StorageDB, err = sql.Open("mysql", "meli_sprint_user:Meli_Sprint#123@/storage")

	assert.Nil(t, err)

	myrepo := NewRepo(StorageDB)
	productResult, err := myrepo.GetOne(1)
	if err != nil {
		log.Println(err)
	}
	assert.Equal(t, producto.Nombre, productResult.Nombre)
}

func TestUpdate(t *testing.T) {
	producto := domain.Productos{
		Nombre:        "Maradona",
		Color:         "Rojo",
		Precio:        2500,
		Stock:         1,
		Codigo:        "12ac",
		Publicado:     true,
		FechaCreación: "2021-02-16",
	}

	StorageDB, err = sql.Open("mysql", "root:@/storage")
	assert.Nil(t, err)
	idtest := 2

	myrepo := NewRepo(StorageDB)
	productResult, err := myrepo.Update(idtest, producto.Nombre, producto.Color, producto.Precio, producto.Stock, producto.Codigo, producto.Publicado, producto.FechaCreación)
	if err != nil {
		log.Println(err)
	}
	producto.Id = idtest
	assert.Equal(t, producto, productResult)
}
