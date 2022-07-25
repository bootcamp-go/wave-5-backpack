package transaction

import (
	"encoding/json"
	"fmt"
	"proyecto-web/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

type StubStorage struct {
	ReadWasCalled bool
}

func TestUpdateParcial(t *testing.T) {
	//arrange
	StorageMock := &StubStorage{}
	repo := NewRepository(StorageMock)

	// act
	previusUpdate := repo.GetAll()
	updatedData, _ := repo.UpdateParcial(0, "AFTER UPDATE", 5.0)

	fmt.Println("Updated data: ", updatedData)

	// assert
	assert.Equal(t, true, StorageMock.ReadWasCalled)
	assert.Equal(t, "BEFORE UPDATE", previusUpdate[0].CodigoTransaccion)
	assert.Equal(t, "AFTER UPDATE", updatedData.CodigoTransaccion)
}

func (s *StubStorage) Read(data interface{}) error {
	s.ReadWasCalled = true

	transacciones := []domain.Transaction{
		{
			Id:                0,
			CodigoTransaccion: "BEFORE UPDATE",
			Moneda:            "PESOS",
			Monto:             5.0,
			Emisor:            "ARCOR",
			Receptor:          "AFIP",
			FechaTransaccion:  "12-01-2022",
		},
	}

	byteData, _ := json.Marshal(transacciones)
	json.Unmarshal(byteData, data)
	return nil
}

func (s *StubStorage) Write(data interface{}) error {
	return nil
}
