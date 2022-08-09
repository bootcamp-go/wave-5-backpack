package product

import (
	"context"
	"github.com/stretchr/testify/assert"
	"storage/internal/domain"
	"storage/pkg/store"
	"testing"
	"time"
)

func getRepo() Repository {
	store.Init()
	return NewRepository(store.StorageDB)
}

func TestGetAll(t *testing.T) {
	//arrage
	repo := getRepo()
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	//act
	result, err := repo.GetAll(ctx)

	//assert
	assert.Nil(t, err)
	assert.Greater(t, len(result), 0)
}

func TestGetByName(t *testing.T) {
	//arrange
	repo := getRepo()
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	//act
	result, err := repo.GetByName(ctx, "monitor")

	//assert
	assert.Nil(t, err)
	assert.Equal(t, 1, result.ID)
}

func TestStore(t *testing.T) {
	//arrange
	repo := getRepo()
	prod := domain.Product{
		Name:  "testing",
		Type:  "Prueba",
		Count: 1,
		Price: 10,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	//act
	result, err := repo.Store(ctx, prod)

	//assert
	assert.Nil(t, err)
	assert.NotEqual(t, 0, result.ID)
	assert.Equal(t, "testing", result.Name)
}

func TestUpdate(t *testing.T) {
	//arrage
	repo := getRepo()
	prod := domain.Product{
		ID:    5,
		Name:  "prueba",
		Type:  "Test",
		Count: 1,
		Price: 10,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	//act
	result, err := repo.Update(ctx, prod)

	//assemble
	assert.Nil(t, err)
	assert.NotEqual(t, 0, result.ID)
	assert.Equal(t, "prueba", result.Name)
}

func TestDelete(t *testing.T) {
	//arrage
	repo := getRepo()
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	//act
	err := repo.Delete(ctx, 7)

	//assert
	assert.Nil(t, err)
}
