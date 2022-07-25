package productos

import (
	"testing"

	"github.com/bootcamp-go/wave-5-backpack/tree/ospina_christian/testing/c2_testingTM/internal/domain"
	"github.com/stretchr/testify/assert"
)

type MockSearchEngine struct {
	ReadWasCalled bool
}

func (m MockSearchEngine) GetAll() ([]domain.Productos, error) {
	return []domain.Productos{}, nil
}

func (m MockSearchEngine) LastID() (int, error) {
	return 0, nil
}
func (m MockSearchEngine) Store(id int, nombre string, color string, precio float64, stock int, codigo string, publicado bool, fechaCreacion string) (domain.Productos, error) {
	return domain.Productos{}, nil
}

func (m MockSearchEngine) Update(id int, nombre string, color string, precio float64, stock int, codigo string, publicado bool, fechaCreacion string) (domain.Productos, error) {
	return domain.Productos{}, nil
}

func (m MockSearchEngine) UpdatePrecio(id int, precion float64) (domain.Productos, error) {

	return domain.Productos{}, nil
}

func (m MockSearchEngine) Delete(id int) error {
	return nil
}

func (m *MockSearchEngine) GetForId(id int) (domain.Productos, error) {
	m.ReadWasCalled = true
	return domain.Productos{
		Id:            1,
		Nombre:        "Esparragos",
		Color:         "Verde",
		Precio:        12300,
		Stock:         12,
		Codigo:        "@123",
		Publicado:     true,
		FechaCreación: "12/08/2022",
	}, nil
}

func TestUpdatePrice(t *testing.T) {
	//arrange
	myMockSearchEngine := MockSearchEngine{}
	motor := Repository(&myMockSearchEngine)
	resultadoEsperado := domain.Productos{
		Id:            1,
		Nombre:        "Esparragos",
		Color:         "Verde",
		Precio:        12300,
		Stock:         12,
		Codigo:        "@123",
		Publicado:     true,
		FechaCreación: "12/08/2022",
	}

	//act
	resultado, _ := motor.GetForId(1)

	//assert
	assert.Equal(t, resultadoEsperado, resultado)
	assert.True(t, myMockSearchEngine.ReadWasCalled)
}
