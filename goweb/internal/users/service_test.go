package users

import (
	"testing"

	"github.com/bootcamp-go/wave-5-backpack/tree/flood_patricio/goweb/internal/domain"
	"github.com/bootcamp-go/wave-5-backpack/tree/flood_patricio/goweb/pkg/store"
	"github.com/stretchr/testify/assert"
)

func TestServiceIntegrationUpdate(t *testing.T) {
	testUsers := []domain.User{
		{
			Id:        1,
			Age:       22,
			FirstName: "Patricio",
			LastName:  "Flood",
			Email:     "patricio.flood@mercadolibre.com",
			Height:    1.72,
			Active:    true,
		},
	}

	usersAfterUpdate := []domain.User{
		{
			Id:        1,
			Age:       10,
			FirstName: "Patricio",
			LastName:  "Flores",
			Email:     "patricio.flood@mercadolibre.com",
			Height:    1.72,
			Active:    true,
		},
	}

	db := &store.MockStorage{
		DataMock:      testUsers,
		ReadWasCalled: false,
	}
	repo := NewRepository(db)
	service := NewService(repo)

	resultado, _ := service.Update(1, 10, "Patricio", "Flores", "patricio.flood@mercadolibre.com", "", 1.72, true)

	//Test de actualización de la base de datos
	assert.Equal(t, usersAfterUpdate, db.DataMock)

	//Test de respuesta del Update
	assert.Equal(t, usersAfterUpdate[0], resultado)

	//Test de método Read ejecutado
	assert.True(t, db.ReadWasCalled)
}

func TestServiceIntegrationDelete(t *testing.T) {
	testUsers := []domain.User{
		{
			Id:        1,
			Age:       22,
			FirstName: "Patricio",
			LastName:  "Flood",
			Email:     "patricio.flood@mercadolibre.com",
			Height:    1.72,
			Active:    true,
		},
	}

	db := &store.MockStorage{
		DataMock:      testUsers,
		ReadWasCalled: false,
	}
	repo := NewRepository(db)
	service := NewService(repo)

	service.Delete(1)
	//Test de borrado en base de datos
	assert.Equal(t, []domain.User{}, db.DataMock)

	err := service.Delete(2)
	//Test de error en el borrado
	assert.ErrorContains(t, err, "usuario no encontrado")
}
