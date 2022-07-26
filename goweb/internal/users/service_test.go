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

	err := service.Delete(1)
	//Test de error nulo
	assert.Nil(t, err)
	//Test de borrado en base de datos
	assert.Equal(t, []domain.User{}, db.DataMock)

	err = service.Delete(2)
	//Test de error en el borrado
	assert.ErrorContains(t, err, "usuario no encontrado")
}

func TestServiceIntegrationGetById(t *testing.T) {
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
		{
			Id:        2,
			Age:       25,
			FirstName: "Julian",
			LastName:  "Sanchez",
			Email:     "julian.sanchez@invento.com",
			Height:    1.52,
			Active:    false,
		},
	}

	db := &store.MockStorage{
		DataMock:      testUsers,
		ReadWasCalled: false,
	}
	repo := NewRepository(db)
	service := NewService(repo)

	user, err := service.GetById(2)

	//Test de error nulo
	assert.Nil(t, err)

	//Test de usuario correcto
	assert.Equal(t, testUsers[1], user)

	//Test de read utilizado
	assert.True(t, db.ReadWasCalled)

	//Test de id incorrecto
	_, err = service.GetById(3)
	assert.ErrorContains(t, err, "usuario no encontrado")
}

func TestServiceIntegrationStore(t *testing.T) {
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
		{
			Id:        2,
			Age:       25,
			FirstName: "Julian",
			LastName:  "Sanchez",
			Email:     "julian.sanchez@invento.com",
			Height:    1.52,
			Active:    false,
		},
	}

	db := &store.MockStorage{
		DataMock:      []domain.User{testUsers[0]},
		ReadWasCalled: false,
	}
	repo := NewRepository(db)
	service := NewService(repo)

	user, err := service.Store(25, "Julian", "Sanchez", "julian.sanchez@invento.com", 1.52, false)
	//Ignoro la fecha de creación autogenerada
	user.CreatedAt = ""

	//Test de error nulo
	assert.Nil(t, err)

	//Test de usuario correcto
	assert.Equal(t, testUsers[1], user)

	//Ignoro la fecha de creación autogenerada
	db.DataMock.([]domain.User)[1].CreatedAt = ""

	//Test de cambio en la base de datos
	assert.Equal(t, testUsers, db.DataMock)
}
