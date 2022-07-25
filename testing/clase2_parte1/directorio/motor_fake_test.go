package directorio

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type FakeSearchEngine struct {
	DB map[string]string
}

func (m *FakeSearchEngine) BuscarPorNombre(nombre string) string {
	return m.DB[nombre]
}

func (m *FakeSearchEngine) BuscarPorTelefono(telefono string) string {
	for key, value := range m.DB {
		if value == telefono {
			return key
		}
	}
	return ""
}

func (m *FakeSearchEngine) AgregarEntrada(nombre, telefono string) error {
	if m.DB == nil {
		m.DB = map[string]string{}
	}
	m.DB[nombre] = telefono
	return nil
}

func TestFindByNameFaked(t *testing.T) {
	//arrange
	testValues := map[string]string{"Nacho": "123456", "Nico": "234567"}

	myFakeSearchEngine := FakeSearchEngine{DB:testValues}
	motor := NewEngine(&myFakeSearchEngine)

	//act
	resultadoNacho := motor.FindByName("Nacho")
	resultadoNico := motor.FindByName("Nico")
	resultadoTelefono := motor.FindByTelephone("123456")

	//assert
	assert.Equal(t, testValues["Nacho"], resultadoNacho)
	assert.Equal(t, testValues["Nico"], resultadoNico)
	assert.Equal(t, "Nacho", resultadoTelefono)

}
