package products

import (
	"encoding/json"
	"testing"

	"github.com/bootcamp-go/wave-5-backpack/tree/Ramos_Andres/goweb/practica/internal/domain"
	"github.com/stretchr/testify/assert"
)

type MockFile struct{}

var isReadCalled bool = false

func (s MockFile) Read(data interface{}) error {
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
	isReadCalled = true
	json.Unmarshal(byteData, data)
	return nil

}

func (s MockFile) Write(data interface{}) error {
	return nil
}

func (s MockFile) Ping() error {
	return nil
}

func TestUpdateName(t *testing.T) {
	expected := domain.Product{
		Id:         1,
		Name:       "Laptop mod",
		Color:      "black",
		Price:      999.99,
		Stock:      100,
		Code:       "SJD23RFG",
		Published:  false,
		Created_at: "2022-06-30",
	}
	stub := MockFile{}
	r := NewRepository(stub)
	prMod, err := r.UpdatePartial(1, "Laptop mod", "", 0, 100, "", false, "")
	assert.Nil(t, err)
	assert.True(t, isReadCalled)
	assert.Equal(t, expected, prMod)
}
