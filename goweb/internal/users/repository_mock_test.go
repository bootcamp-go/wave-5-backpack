package users

import (
	"testing"

	"github.com/bootcamp-go/wave-5-backpack/tree/flood_patricio/goweb/internal/domain"
	"github.com/bootcamp-go/wave-5-backpack/tree/flood_patricio/goweb/pkg/store"
	"github.com/stretchr/testify/assert"
)

func TestUpdateAgeLastName(t *testing.T) {
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

	resultado, _ := repo.UpdateAgeLastName(1, 10, "Flores")

	assert.Equal(t, usersAfterUpdate, db.DataMock)
	assert.Equal(t, usersAfterUpdate[0], resultado)
	assert.True(t, db.ReadWasCalled)
}
