package products

import (
	"encoding/json"
	"testing"

	"github.com/bootcamp-go/wave-5-backpack/tree/Ramos_Andres/goweb/practica/internal/domain"
	"github.com/stretchr/testify/assert"
)

type StubFile struct{}

func (s StubFile) Read(data interface{}) error {
	byteData, _ := json.Marshal([]domain.Product{
		{
			Id:         1,
			Name:       "Laptop",
			Color:      "black",
			Price:      999.99,
			Stock:      100,
			Code:       "SJD23RFG",
			Published:  false,
			Created_at: "2022-06-30",
		},
	})
	json.Unmarshal(byteData, data)
	return nil

}

func (s StubFile) Write(data interface{}) error {
	return nil
}

func (s StubFile) Ping() error {
	return nil
}

func TestGetAll(t *testing.T) {
	expected := []domain.Product{
		{
			Id:         1,
			Name:       "Laptop",
			Color:      "black",
			Price:      999.99,
			Stock:      100,
			Code:       "SJD23RFG",
			Published:  false,
			Created_at: "2022-06-30",
		},
	}
	stub := StubFile{}
	r := NewRepository(stub)
	prList, err := r.GetAll()
	assert.Nil(t, err)
	assert.Equal(t, expected, prList)
}
