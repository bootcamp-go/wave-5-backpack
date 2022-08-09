package products

import (
	"database/sql"
	"testing"

	"github.com/bootcamp-go/wave-5-backpack/tree/Ramos_Andres/goweb/practica/internal/domain"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func TestStore(t *testing.T) {
	db, err := sql.Open("mysql", "root@/storage")
	assert.Nil(t, err)
	repo := NewRepository(db)

	mockProduct := domain.Product{
		Name:         "test",
		Color:        "Red",
		Price:        10.99,
		Stock:        10,
		Code:         "JH7BU998G",
		Published:    true,
		Warehouse_id: 1,
	}

	res, err := repo.Store(mockProduct)
	mockProduct.Id = res.Id
	assert.Nil(t, err)
	assert.Equal(t, mockProduct, res)
}

func TestGetByName(t *testing.T) {
	db, err := sql.Open("mysql", "root@/storage")
	assert.Nil(t, err)
	repo := NewRepository(db)

	mockProduct := domain.Product{
		Id:           1,
		Name:         "product 1",
		Color:        "red",
		Price:        10.99,
		Stock:        100,
		Code:         "HJ988BH",
		Published:    true,
		Created_at:   "2022-08-09",
		Warehouse_id: 1,
	}
	assert.Nil(t, err)
	res, err := repo.GetByName("product 1")
	assert.Nil(t, err)
	assert.Equal(t, mockProduct, res)
}

func TestGetAll(t *testing.T) {
	db, err := sql.Open("mysql", "root@/storage")
	assert.Nil(t, err)
	repo := NewRepository(db)
	assert.Nil(t, err)
	_, err = repo.GetAll()
	assert.Nil(t, err)
}
