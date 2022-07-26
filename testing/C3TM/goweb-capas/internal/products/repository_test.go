package products

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type StubStore struct {
}

func (fs *StubStore) Read(data interface{}) error {
	all := data.(*[]Product)
	*all = []Product{
		{ID: 1, Name: "Cama", Type: "Madera", Count: 2, Price: 270000.0},
		{ID: 2, Name: "Silla", Type: "Madera", Count: 2, Price: 27000.0},
	}
	return nil
}

func (fs *StubStore) Write(data interface{}) error {
	return nil
}

func (fs *StubStore) Ping() error {
	return nil
}

func TestGetAll(t *testing.T) {
	stub := StubStore{}
	repo := NewRepository(&stub)
	expected := []Product{
		{ID: 1, Name: "Cama", Type: "Madera", Count: 2, Price: 270000.0},
		{ID: 2, Name: "Silla", Type: "Madera", Count: 2, Price: 27000.0},
	}
	all, err := repo.GetAll()
	assert.Nil(t, err)
	assert.Equal(t, expected, all, "No coincide la información de usuario esperada con la obtenida")
}
