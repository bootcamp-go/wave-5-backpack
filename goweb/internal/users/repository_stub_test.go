package users

import (
	"fmt"
	"testing"

	"github.com/bootcamp-go/wave-5-backpack/tree/flood_patricio/goweb/internal/domain"
	"github.com/bootcamp-go/wave-5-backpack/tree/flood_patricio/goweb/pkg/store"
	"github.com/stretchr/testify/assert"
)

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

	db := &store.MockStorage{
		DataMock: testUsers,
	}
	fmt.Println(db.DataMock)
	repo := NewRepository(db)

	emptyFilter := make(map[string]interface{})
	resultado, _ := repo.GetAll(emptyFilter)
	fmt.Println(resultado)
	assert.Equal(t, testUsers, resultado)
}
