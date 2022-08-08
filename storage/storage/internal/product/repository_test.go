package product

import (
	"github.com/stretchr/testify/assert"
	"storage/internal/domain"
	"storage/pkg/store"
	"testing"
)

func getRepo() Repository {
	store.Init()
	return NewRepository(store.StorageDB)
}

func TestGetByName(t *testing.T) {
	//arrange
	repo := getRepo()

	//act
	result, err := repo.GetByName("monitor")

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

	//act
	result, err := repo.Store(prod)

	//assert
	assert.Nil(t, err)
	assert.NotEqual(t, 0, result.ID)
	assert.Equal(t, "testing", result.Name)
}
