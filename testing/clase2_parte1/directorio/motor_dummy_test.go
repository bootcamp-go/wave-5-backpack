package directorio

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type DummyDB struct{}

func (d DummyDB) BuscarPorNombre(nombre string) string {
	return ""
}

func (d DummyDB) BuscarPorTelefono(telefono string) string {
	return ""
}

func (d DummyDB) AgregarEntrada(nombre, telefono string) error {
	return nil
}

func TestGetVersion(t *testing.T) {
	//arrange

	myDummyDB := DummyDB{}
	motor := NewEngine(myDummyDB)
	versionEsperada := "1.0"

	//act
	resultado := motor.GetVersion()

	//assert
	assert.Equal(t, versionEsperada, resultado)

}
