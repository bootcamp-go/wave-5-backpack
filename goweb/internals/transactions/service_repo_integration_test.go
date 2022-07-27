package transactions

import (
	"errors"
	"fmt"
	"goweb/internals/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockIntegration struct {
	dataMock      []domain.Transaction
	readWasCalled bool
	errWrite      string
	errRead       string
}

func (m *MockIntegration) Read(data interface{}) error {
	if m.errRead != "" {
		return fmt.Errorf(m.errRead)
	}
	trans := data.(*[]domain.Transaction)
	*trans = m.dataMock
	m.readWasCalled = true
	return nil
}

func (m *MockIntegration) Write(data interface{}) error {
	if m.errWrite != "" {
		return fmt.Errorf(m.errWrite)
	}
	a := data.([]domain.Transaction)
	m.dataMock = append(m.dataMock, a...)
	return nil
}

func (m *MockIntegration) Ping() error {
	return nil
}

func TestIntegrationUpdate(t *testing.T) {
	//Arrange
	database := []domain.Transaction{
		{
			Id: 1, Codigo: "ASKDD423423NJK", Moneda: "MXN", Monto: 1500, Emisor: "Yvonne", Receptor: "Fernando", Fecha: "27-06-2022",
		},
		{
			Id: 2, Codigo: "567KDDHESDFNJK", Moneda: "ARG", Monto: 2000, Emisor: "Cristian", Receptor: "Francisco", Fecha: "28-07-2022",
		},
		{
			Id: 3, Codigo: "ZF442G23423NYT", Moneda: "COL", Monto: 700, Emisor: "Abelardo", Receptor: "Franco", Fecha: "30-05-2022",
		},
	}
	newTrans := domain.Transaction{
		Id: 1, Codigo: "HYED4DD423NJ57", Moneda: "CHL", Monto: 5000, Emisor: "Toni", Receptor: "Mati", Fecha: "10-04-2022",
	}

	MockInte := MockIntegration{
		dataMock: database,
		errWrite: "",
		errRead:  "",
	}
	//Act
	repo := NewRepository(&MockInte)
	service := NewService(repo)
	trans, err := service.Update(
		newTrans.Id,
		newTrans.Codigo,
		newTrans.Moneda,
		newTrans.Monto,
		newTrans.Emisor,
		newTrans.Receptor,
	)

	//Assert
	assert.Nil(t, err)
	assert.Equal(t, MockInte.dataMock[0], trans)
	assert.True(t, MockInte.readWasCalled)
}

func TestIntegrationUpdateFailRead(t *testing.T) {
	//Arrange
	expectedError := errors.New("No se puede leer la base de datos")
	MockInte := MockIntegration{
		dataMock: nil,
		errWrite: "",
		errRead:  "No se puede leer la base de datos",
	}
	//Act
	repo := NewRepository(&MockInte)
	service := NewService(repo)
	_, err := service.Update(0, "", "", 0, "", "")
	//Assert
	assert.Equal(t, expectedError, err)
}

func TestIntegrationStoreFailWrite(t *testing.T) {
	//Arrange
	expectedError := errors.New("No se pudo escribir en la base de datos")
	database := []domain.Transaction{
		{
			Id: 1, Codigo: "ZF442G23423NYT", Moneda: "COL", Monto: 700, Emisor: "Abelardo", Receptor: "Franco", Fecha: "30-05-2022",
		},
	}
	mockInte := MockIntegration{
		dataMock: database,
		errRead:  "",
		errWrite: "No se pudo escribir en la base de datos",
	}
	//Act
	repo := NewRepository(&mockInte)
	service := NewService(repo)
	_, err := service.Store("ZF442G23423NYT", "COL", 700, "Abelardo", "Franco") //Como err trae errores dentro de errores, hay que aplicar la siguiente línea.
	finalError := errors.Unwrap(err)
	//Assert
	assert.Equal(t, finalError, expectedError)
}

func TestIntegrationDelete(t *testing.T) {
	//Arrange
	database := []domain.Transaction{
		{
			Id: 1, Codigo: "ASKDD423423NJK", Moneda: "MXN", Monto: 1500, Emisor: "Yvonne", Receptor: "Fernando", Fecha: "27-06-2022",
		},
		{
			Id: 2, Codigo: "567KDDHESDFNJK", Moneda: "ARG", Monto: 2000, Emisor: "Cristian", Receptor: "Francisco", Fecha: "28-07-2022",
		},
		{
			Id: 3, Codigo: "ZF442G23423NYT", Moneda: "COL", Monto: 700, Emisor: "Abelardo", Receptor: "Franco", Fecha: "30-05-2022",
		},
	}

	mockInte := MockIntegration{
		dataMock: database,
		errWrite: "",
		errRead:  "",
	}
	//Act
	repo := NewRepository(&mockInte)
	service := NewService(repo)
	err := service.Delete(1)
	//Assert
	assert.Nil(t, err)
	assert.NotEqual(t, mockInte.dataMock[0].Id, 1)
}

func TestIntegrationDeleteError(t *testing.T) {
	//Arrange
	database := []domain.Transaction{
		{
			Id: 1, Codigo: "ASKDD423423NJK", Moneda: "MXN", Monto: 1500, Emisor: "Yvonne", Receptor: "Fernando", Fecha: "27-06-2022",
		},
		{
			Id: 2, Codigo: "567KDDHESDFNJK", Moneda: "ARG", Monto: 2000, Emisor: "Cristian", Receptor: "Francisco", Fecha: "28-07-2022",
		},
		{
			Id: 3, Codigo: "ZF442G23423NYT", Moneda: "COL", Monto: 700, Emisor: "Abelardo", Receptor: "Franco", Fecha: "30-05-2022",
		},
	}

	mockInte := MockIntegration{
		dataMock: database,
		errWrite: "",
		errRead:  "",
	}
	falseID := 69
	transNotFound := fmt.Errorf("Transacción %d no encontrada", falseID)
	//Act
	repo := NewRepository(&mockInte)
	service := NewService(repo)
	errNotFound := service.Delete(falseID)
	//Assert
	assert.ErrorContains(t, errNotFound, transNotFound.Error())
}
