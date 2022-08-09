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
		Name:      "test",
		Color:     "Red",
		Price:     10.99,
		Stock:     10,
		Code:      "JH7BU998G",
		Published: true,
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
		Id:         1,
		Name:       "test",
		Color:      "Red",
		Price:      10.99,
		Stock:      10,
		Code:       "JH7BU998G",
		Published:  true,
		Created_at: "2022-08-09",
	}

	// created, err := repo.Store(mockProduct)
	// mockProduct.Id = created.Id
	assert.Nil(t, err)
	res, err := repo.GetByName("test")
	assert.Nil(t, err)
	assert.Equal(t, mockProduct, res)
}
