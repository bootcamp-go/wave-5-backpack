package products

import (
	"context"
	"database/sql"
	"testing"
	"time"

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
	ctx, close := context.WithTimeout(context.Background(), time.Second*5)
	defer close()
	res, err := repo.Store(ctx, mockProduct)
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
	ctx, close := context.WithTimeout(context.Background(), time.Second*5)
	defer close()
	res, err := repo.GetByName(ctx, "product 1")
	assert.Nil(t, err)
	assert.Equal(t, mockProduct, res)
}

func TestGetAll(t *testing.T) {
	db, err := sql.Open("mysql", "root@/storage")
	assert.Nil(t, err)
	repo := NewRepository(db)
	assert.Nil(t, err)
	ctx, close := context.WithTimeout(context.Background(), time.Second*5)
	defer close()
	_, err = repo.GetAll(ctx)
	assert.Nil(t, err)
}
