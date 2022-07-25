package productos

import (
	"testing"

	"github.com/bootcamp-go/wave-5-backpack/tree/ospina_christian/testing/c2_testingTM/internal/domain"
	"github.com/stretchr/testify/assert"
)

type StubSearchEngine struct{}

func (d StubSearchEngine) GetAll() ([]domain.Productos, error) {
	return []domain.Productos{{
		Id:            1,
		Nombre:        "Esparragos",
		Color:         "Verde",
		Precio:        12300,
		Stock:         12,
		Codigo:        "@123",
		Publicado:     true,
		FechaCreación: "12/08/2022",
	}}, nil
}

func (d StubSearchEngine) LastID() (int, error) {
	return 0, nil
}
func (d StubSearchEngine) Store(id int, nombre string, color string, precio float64, stock int, codigo string, publicado bool, fechaCreacion string) (domain.Productos, error) {
	return domain.Productos{}, nil
}

func (d StubSearchEngine) Update(id int, nombre string, color string, precio float64, stock int, codigo string, publicado bool, fechaCreacion string) (domain.Productos, error) {
	return domain.Productos{}, nil
}

func (d StubSearchEngine) UpdatePrecio(id int, precion float64) (domain.Productos, error) {
	return domain.Productos{}, nil
}

func (d StubSearchEngine) Delete(id int) error {
	return nil
}

func (d StubSearchEngine) GetForId(id int) (domain.Productos, error) {
	return domain.Productos{}, nil
}

func TestRead(t *testing.T) {
	//arrange
	myStubSearchEngine := StubSearchEngine{}
	motor := Repository(myStubSearchEngine)

	resultEsperado := []domain.Productos{{
		Id:            1,
		Nombre:        "Esparragos",
		Color:         "Verde",
		Precio:        12300,
		Stock:         12,
		Codigo:        "@123",
		Publicado:     true,
		FechaCreación: "12/08/2022",
	}}
	//resultadoEsperado := []domain.Productos([]domain.Productos{})

	//act
	resultado, _ := motor.GetAll()

	//assert
	assert.Equal(t, resultEsperado, resultado)

}
