package users

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/bootcamp-go/wave-5-backpack/tree/flood_patricio/goweb/internal/domain"
	"github.com/stretchr/testify/assert"
)

type fileStubStore struct {
	db interface{}
}

func (fs *fileStubStore) Write(data interface{}) error {
	fs.db = data
	return nil
}

func (fs *fileStubStore) Read(data interface{}) error {
	rv := reflect.ValueOf(data)
	rv = reflect.Indirect(rv)
	rv.Set(reflect.ValueOf(fs.db))
	return nil
}

func TestGetAll(t *testing.T) {
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

	db := &fileStubStore{
		db: testUsers,
	}
	fmt.Println(db.db)
	repo := NewRepository(db)

	emptyFilter := make(map[string]interface{})
	resultado, _ := repo.GetAll(emptyFilter)
	fmt.Println(resultado)
	assert.Equal(t, testUsers, resultado)
}
