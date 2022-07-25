package users

import (
	"reflect"
	"testing"

	"github.com/bootcamp-go/wave-5-backpack/tree/flood_patricio/goweb/internal/domain"
	"github.com/stretchr/testify/assert"
)

type fileMockStore struct {
	db            interface{}
	readWasCalled bool
}

func (fs *fileMockStore) Write(data interface{}) error {
	fs.db = data
	return nil
}

func (fs *fileMockStore) Read(data interface{}) error {
	rv := reflect.ValueOf(data)
	rv = reflect.Indirect(rv)
	rv.Set(reflect.ValueOf(fs.db))
	fs.readWasCalled = true
	return nil
}

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

	db := &fileMockStore{
		db:            testUsers,
		readWasCalled: false,
	}
	repo := NewRepository(db)

	resultado, _ := repo.UpdateAgeLastName(1, 10, "Flores")

	assert.Equal(t, usersAfterUpdate, db.db)
	assert.Equal(t, usersAfterUpdate[0], resultado)
	assert.True(t, db.readWasCalled)
}
