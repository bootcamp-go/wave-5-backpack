package directorio

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type SpyBD struct{
	BuscarPorTelefonoWasCalled bool
}

func (d SpyBD) BuscarPorNombre(nombre string) string {
	return ""
}

func (d *SpyBD) BuscarPorTelefono(telefono string) string {
	d.BuscarPorTelefonoWasCalled = true
	return ""
}

func (d SpyBD) AgregarEntrada(nombre, telefono string) error {
	return nil
}

func TestFindByTelephone(t *testing.T) {
	//arrange

	mySpy := SpyBD{BuscarPorTelefonoWasCalled: false}
	motor := NewEngine(&mySpy)
	telefono := "12345678"

	//act
	motor.FindByTelephone(telefono)

	//assert
	assert.True(t, mySpy.BuscarPorTelefonoWasCalled)
}

func TestFindByTelephoneNotCalled(t *testing.T) {
	//arrange

	mySpy := SpyBD{BuscarPorTelefonoWasCalled: false}
	motor := NewEngine(&mySpy)
	telefono := "1234"

	//act
	motor.FindByTelephone(telefono)

	//assert
	assert.False(t, mySpy.BuscarPorTelefonoWasCalled)
}