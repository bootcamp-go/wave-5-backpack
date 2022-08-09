package product

import (
	"context"
	"database/sql"
	"log"
	"testing"
	"time"

	"github.com/bootcamp-go/wave-5-backpack/tree/ospina_christian/storage/storageprueba/internal/domain"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

var (
	StorageDB *sql.DB
	err       error
	path      string = "meli_sprint_user:Meli_Sprint#123@/storage"
)

func TestStore(t *testing.T) {
	producto := domain.Productos{
		Nombre:        "Frijoles",
		Color:         "Rojos",
		Precio:        2500,
		Stock:         15,
		Codigo:        "1f",
		Publicado:     true,
		FechaCreación: "2021-03-15",
		IdWarehouse:   1,
	}
	StorageDB, err = sql.Open("mysql", path)

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
		Nombre:        "Mandarinas Francesas",
		Color:         "Naranjas",
		Precio:        1500,
		Stock:         50,
		Codigo:        "12mand",
		Publicado:     true,
		FechaCreación: "2021-04-06",
		IdWarehouse:   1,
	}

	StorageDB, err = sql.Open("mysql", path)
	assert.Nil(t, err)

	myrepo := NewRepo(StorageDB)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	productResult, err := myrepo.GetOne(ctx, 2)
	if err != nil {
		log.Println(err)
	}
	assert.Equal(t, producto.Nombre, productResult.Nombre)
}

func TestUpdate(t *testing.T) {
	//arrange
	StorageDB, err = sql.Open("mysql", path)
	assert.Nil(t, err)
	myrepo := NewRepo(StorageDB)

	producto := domain.Productos{
		Nombre:        "Mandarinas Francesas",
		Color:         "Naranjas",
		Precio:        1500,
		Stock:         50,
		Codigo:        "12mand",
		Publicado:     true,
		FechaCreación: "2021-04-06",
		IdWarehouse:   1,
	}

	//act
	idtest := 2
	producto.Id = idtest
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	productResult, err := myrepo.Update(ctx, idtest, producto.Nombre, producto.Color, producto.Precio, producto.Stock, producto.Codigo, producto.Publicado, producto.FechaCreación, producto.IdWarehouse)

	//assert
	assert.Nil(t, err)
	assert.Equal(t, producto, productResult)
}

func TestGetAll(t *testing.T) {
	//arrange
	StorageDB, err = sql.Open("mysql", path)
	assert.Nil(t, err)
	myrepo := NewRepo(StorageDB)

	producto := domain.Productos{
		Nombre:        "Mandarinas Francesas",
		Color:         "Naranjas",
		Precio:        1500,
		Stock:         50,
		Codigo:        "12mand",
		Publicado:     true,
		FechaCreación: "2021-04-06",
		IdWarehouse:   1,
	}

	//act
	idtest := 2
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	producto.Id = idtest
	productResult, err := myrepo.GetAll(ctx)

	//assert
	assert.Nil(t, err)
	assert.Equal(t, producto, productResult[idtest-1])
}

func TestDelete(t *testing.T) {
	//arrange
	StorageDB, err = sql.Open("mysql", path)
	assert.Nil(t, err)
	myrepo := NewRepo(StorageDB)

	//act
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	idtest := 5
	err := myrepo.Delete(ctx, idtest)

	//assert
	assert.Nil(t, err)
}

func TestGetAllProdWare(t *testing.T) {
	//arrange
	StorageDB, err = sql.Open("mysql", path)
	assert.Nil(t, err)
	myrepo := NewRepo(StorageDB)

	//acts
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	idtest := 1
	prodware := myrepo.GetAllProdWare(ctx, idtest)

	//assert
	assert.Nil(t, err)
	assert.Equal(t, domain.Productos{}, prodware)
}
