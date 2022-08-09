package products

import (
	"context"
	"storage/internal/domain"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	prod := domain.Product{
		Name:  "test",
		Type:  "tipe1",
		Count: 1,
		Price: 10.5,
	}
	repo := NewRepo()
	res, err := repo.Store(prod)

	assert.Nil(t, err)
	assert.Equal(t, prod.Name, res.Name)
}
func TestGetByName(t *testing.T) {
	ExpectedProd := domain.Product{
		ID:    1,
		Name:  "test",
		Type:  "tipe1",
		Count: 1,
		Price: 10.5,
	}
	repo := NewRepo()
	res, err := repo.GetByName("test")

	assert.Nil(t, err)
	assert.Equal(t, ExpectedProd, res)
}
func TestGetAll(t *testing.T) {
	ExpectedProds := []domain.Product{
		{
			ID:    1,
			Name:  "test",
			Type:  "tipe1",
			Count: 1,
			Price: 10.5,
		}, {
			ID:    2,
			Name:  "test2",
			Type:  "tipe2",
			Count: 2,
			Price: 20.5,
		},
	}
	repo := NewRepo()
	res, err := repo.GetAll()

	assert.Nil(t, err)
	assert.Equal(t, ExpectedProds, res)
}
func TestUpdate(t *testing.T) {
	prod := domain.Product{
		ID:    1,
		Name:  "test10",
		Type:  "tipe10",
		Count: 10,
		Price: 100.5,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	repo := NewRepo()
	res, err := repo.Update(ctx, prod)

	assert.Nil(t, err)
	assert.Equal(t, prod.Name, res.Name)
}
