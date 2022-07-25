package directorio

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type StubDB struct{}

func (d StubDB) BuscarPorNombre(nombre string) string {
	return "12345678"
}

func (d StubDB) BuscarPorTelefono(telefono string) string {
	return ""
}

func (d StubDB) AgregarEntrada(nombre, telefono string) error {
	return nil
}

func TestFindByName(t *testing.T) {
	//arrange

	myStubDB := StubDB{}
	motor := NewEngine(myStubDB)
	telefonoEsperado := "12345678"

	//act
	resultado := motor.FindByName("Nacho")

	//assert
	assert.Equal(t, telefonoEsperado, resultado)
}

func TestFindByNameWithNameShorterThan3(t *testing.T) {
	//arrange

	myStubDB := StubDB{}
	motor := NewEngine(myStubDB)
	telefonoEsperado := ""

	//act
	resultado := motor.FindByName("Ah")

	//assert
	assert.Equal(t, telefonoEsperado, resultado)
}
