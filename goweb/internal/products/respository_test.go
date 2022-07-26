package products

import (
	"fmt"
	"goweb/internal/domain"
	"testing"

	"github.com/go-playground/assert"
)

// Se importa testify

type StubSearchEngine struct{}

func (d *StubSearchEngine) GetAll() ([]domain.Products, error) {
	sliceProducts := []domain.Products{}
	respuestaEsperada := domain.Products{
		Id:            1,
		Nombre:        "aguacate",
		Color:         "verde",
		Precio:        30000,
		Stock:         5,
		Codigo:        "23fe2",
		Publicado:     true,
		FechaCreacion: "23/10/2022",
	}

	t1 := domain.Products{
		Id:            2,
		Nombre:        "Banana",
		Color:         "Amarillo",
		Precio:        60000,
		Stock:         13,
		Codigo:        "d7fe2",
		Publicado:     true,
		FechaCreacion: "30/11/2022",
	}

	sliceProducts = append(sliceProducts, respuestaEsperada, t1)
	return sliceProducts, nil
}

func (d *StubSearchEngine) CreateProduct(id int, nombre, color string, precio float64, stock int, código string, publicado bool, fecha_de_creación string) (domain.Products, error) {
	return domain.Products{}, nil
}

func (d *StubSearchEngine) Update(id int, nombre, color string, precio float64, stock int, código string, publicado bool, fecha_de_creación string) (domain.Products, error) {
	return domain.Products{}, nil
}

func (d *StubSearchEngine) Delete(id int) error {
	return nil
}

func (d *StubSearchEngine) UpdateOne(id int, nombre string, precio float64) (domain.Products, error) {
	return domain.Products{}, nil
}

func (d *StubSearchEngine) LastID() (int, error) {
	return 0, nil
}
func TestGetAll(t *testing.T) {
	sliceProducts := []domain.Products{}
	myStubSearchEngine := StubSearchEngine{}
	motor := InitService(&myStubSearchEngine)
	respuestaEsperada := domain.Products{
		Id:            1,
		Nombre:        "aguacate",
		Color:         "verde",
		Precio:        30000,
		Stock:         5,
		Codigo:        "23fe2",
		Publicado:     true,
		FechaCreacion: "23/10/2022",
	}

	t1 := domain.Products{
		Id:            2,
		Nombre:        "Banana",
		Color:         "Amarillo",
		Precio:        60000,
		Stock:         13,
		Codigo:        "d7fe2",
		Publicado:     true,
		FechaCreacion: "30/11/2022",
	}

	sliceProducts = append(sliceProducts, respuestaEsperada, t1)
	fmt.Println(sliceProducts)

	res, _ := motor.GetAll()

	assert.Equal(t, sliceProducts, res)

}
