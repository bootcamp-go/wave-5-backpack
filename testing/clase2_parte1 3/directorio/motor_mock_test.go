package directorio

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockDB struct {
	BuscarPorNombreWasCalled bool
}

func (m *MockDB) BuscarPorNombre(nombre string) string {
	m.BuscarPorNombreWasCalled = true
	return "12345678"
}

func (m *MockDB) BuscarPorTelefono(telefono string) string {
	return ""
}

func (m *MockDB) AgregarEntrada(nombre, telefono string) error {
	return nil
}

func TestFindByNameMocked(t *testing.T) {
	//arrange

	myMockDB := MockDB{}
	motor := NewEngine(&myMockDB)
	telefonoEsperado := "12345678"

	//act
	resultado := motor.FindByName("Nacho")

	//assert
	assert.Equal(t, telefonoEsperado, resultado)
	assert.True(t, myMockDB.BuscarPorNombreWasCalled)
}

func TestFindByNameMockedNotCalled(t *testing.T) {
	//arrange

	myMockDB := MockDB{}
	motor := NewEngine(&myMockDB)
	telefonoEsperado := ""

	//act
	resultado := motor.FindByName("Nac")

	//assert
	assert.Equal(t, telefonoEsperado, resultado)
	assert.False(t, myMockDB.BuscarPorNombreWasCalled)
}
