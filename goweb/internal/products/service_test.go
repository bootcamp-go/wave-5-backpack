package products

import (
	"github.com/bootcamp-go/wave-5-backpack/goweb/internal/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

type repositoryStub struct {
	GetAllWasCalled bool
	BeforUpdate     domain.Products
}

func (r *repositoryStub) GetAll() ([]domain.Products, error) {
	r.GetAllWasCalled = true
	p1 := domain.Products{
		ID:            1,
		Nombre:        "Razer",
		Color:         "Negro",
		Precio:        99990,
		Stock:         25,
		Codigo:        "SKU000001",
		Publicado:     true,
		FechaCreacion: "01-01-2022",
	}
	p2 := domain.Products{
		ID:            2,
		Nombre:        "Corsair",
		Color:         "Amarillo",
		Precio:        88880,
		Stock:         15,
		Codigo:        "SKU000002",
		Publicado:     true,
		FechaCreacion: "02-01-2022",
	}
	list := []domain.Products{p1, p2}
	return list, nil
}

func (r *repositoryStub) Store(nombre, color string, precio float64, stock int, codigo string,
	publicado bool, fechaCreacion string) (domain.Products, error) {
	return r.BeforUpdate, nil
}
func (r *repositoryStub) Update(id int, nombre, color string, precio float64, stock int, codigo string,
	publicado bool, fechaCreacion string) (domain.Products, error) {
	return r.BeforUpdate, nil
}
func (r repositoryStub) UpdatePrecioStock(id int, precio float64, stock int) (domain.Products, error) {
	r.GetAll()
	r.BeforUpdate.Precio = precio
	r.BeforUpdate.Stock = stock
	return r.BeforUpdate, nil
}

func (r *repositoryStub) GetByID(id int) (domain.Products, error) {
	return r.BeforUpdate, nil
}

func (r *repositoryStub) Delete(id int) (int, error) {
	return 0, nil
}

func (r repositoryStub) GetLastID() (int, error) {
	return 0, nil
}

func TestGetAll(t *testing.T) {
	p1 := domain.Products{
		ID:            1,
		Nombre:        "Razer",
		Color:         "Negro",
		Precio:        99990,
		Stock:         25,
		Codigo:        "SKU000001",
		Publicado:     true,
		FechaCreacion: "01-01-2022",
	}
	p2 := domain.Products{
		ID:            2,
		Nombre:        "Corsair",
		Color:         "Amarillo",
		Precio:        88880,
		Stock:         15,
		Codigo:        "SKU000002",
		Publicado:     true,
		FechaCreacion: "02-01-2022",
	}

	stubRepo := repositoryStub{}
	service := NewService(&stubRepo)
	expectedResult := []domain.Products{p1, p2}

	result, err := service.GetAll()

	assert.NoError(t, err)
	assert.Equal(t, expectedResult, result)
}
func TestUpdatePrecioStock(t *testing.T) {
	before := domain.Products{
		ID:            1,
		Nombre:        "Razer",
		Color:         "Negro",
		Precio:        99990,
		Stock:         25,
		Codigo:        "SKU000001",
		Publicado:     true,
		FechaCreacion: "01-01-2022",
	}
	stubRepo := repositoryStub{false, before}
	service := NewService(&stubRepo)
	expectedResult := domain.Products{
		ID:            1,
		Nombre:        "Razer",
		Color:         "Negro",
		Precio:        88880,
		Stock:         45,
		Codigo:        "SKU000001",
		Publicado:     true,
		FechaCreacion: "01-01-2022",
	}

	result, err := service.UpdatePrecioStock(1, 88880, 45)

	assert.NoError(t, err)
	assert.Equal(t, expectedResult, result)
	assert.True(t, stubRepo.GetAllWasCalled)
}
