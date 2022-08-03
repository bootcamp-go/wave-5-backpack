/*
Ejercicio 1 - Test Unitario GetAll()
Generar un Stub del Store cuya función “Read” retorne dos productos con las especificaciones que deseen. Comprobar que GetAll() retorne la información exactamente igual a la esperada. Para esto:
Dentro de la carpeta /internal/(producto/usuario/transacción), crear un archivo repository_test.go con el test diseñado.
*/
package products

import (
	"clase4_parte2/internal/domain"
	"fmt"
	"testing"

	"github.com/go-playground/assert"
)

// Se importa testify

type StubSearchEngine struct{}

func (d *StubSearchEngine) GetAll() ([]domain.Product, error) {
	sliceProducts := []domain.Product{}
	respuestaEsperada := domain.Product{
		ID:    1,
		Name:  "aguacate",
		Price: 30000,
		Count: 5,
	}

	t1 := domain.Product{
		ID:    2,
		Name:  "Banana",
		Price: 60000,
		Count: 13,
	}
	sliceProducts = append(sliceProducts, respuestaEsperada, t1)
	return sliceProducts, nil
}

func (d *StubSearchEngine) CreateProduct(id int, nombre, color string, precio float64, stock int, código string, publicado bool, fecha_de_creación string) (domain.Product, error) {
	return domain.Product{}, nil
}

func (d *StubSearchEngine) Update(id int, name, productType string, count int, price float64) (domain.Product, error) {
	return domain.Product{}, nil
}

func (d *StubSearchEngine) UpdateName(id int, name string) (domain.Product, error) {
	return domain.Product{}, nil
}

func (d *StubSearchEngine) Delete(id int) error {
	return nil
}

func (d *StubSearchEngine) UpdateOne(id int, nombre string, precio float64) (domain.Product, error) {
	return domain.Product{}, nil
}

func (d *StubSearchEngine) LastID() (int, error) {
	return 0, nil
}

func (d *StubSearchEngine) Store(id int, name, producType string, count int, price float64) (domain.Product, error) {
	return domain.Product{}, nil
}

func TestGetAll(t *testing.T) {
	sliceProducts := []domain.Product{}
	myStubSearchEngine := StubSearchEngine{}
	service := NewService(&myStubSearchEngine)
	respuestaEsperada := domain.Product{
		ID:    1,
		Name:  "aguacate",
		Price: 30000,
		Count: 5,
	}

	t1 := domain.Product{
		ID:    2,
		Name:  "Banana",
		Price: 60000,
		Count: 13,
	}

	sliceProducts = append(sliceProducts, respuestaEsperada, t1)
	fmt.Println(sliceProducts)

	res, _ := service.GetAll()

	assert.Equal(t, sliceProducts, res)

}
