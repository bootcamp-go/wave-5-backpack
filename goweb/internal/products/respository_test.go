package products

import (
	"fmt"
	"goweb/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

type StubSearchEngine struct{}

type MockStorage struct {
	dataMock []domain.Products
	errWrite string
	errRead  string
}

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

func (m *MockStorage) Read(data interface{}) error {
	if m.errRead != "" {
		return fmt.Errorf(m.errRead)
	}
	a := data.(*[]domain.Products)
	*a = m.dataMock
	return nil
}

func (m *MockStorage) Write(data interface{}) error {
	if m.errWrite != "" {
		return fmt.Errorf(m.errWrite)
	}
	a := data.([]domain.Products)
	m.dataMock = append(m.dataMock, a...)
	return nil
}
func (m *MockStorage) Ping() error {
	return nil
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

func TestGetAllWithReadAndWrite(t *testing.T) {

	database := []domain.Products{

		{
			Id:            1,
			Nombre:        "aguacate",
			Color:         "verde",
			Precio:        30000,
			Stock:         5,
			Codigo:        "23fe2",
			Publicado:     true,
			FechaCreacion: "23/10/2022",
		},

		{
			Id:            2,
			Nombre:        "Banana",
			Color:         "Amarillo",
			Precio:        60000,
			Stock:         13,
			Codigo:        "d7fe2",
			Publicado:     true,
			FechaCreacion: "30/11/2022",
		},
	}

	mockStorage := MockStorage{
		dataMock: database,
		errWrite: "",
		errRead:  "",
	}

	repo := InitRepository(&mockStorage)
	result, err := repo.GetAll()

	assert.Nil(t, err)
	assert.Equal(t, mockStorage.dataMock, result)
}

func TestUpdate(t *testing.T) {

	updateData := []domain.Products{
		{Id: 1,
			Nombre:        "Sandia",
			Color:         "verde",
			Precio:        12332,
			Stock:         13,
			Codigo:        "asd7sd",
			Publicado:     true,
			FechaCreacion: "12/02/2022"},
	}

	mockUpdate := MockStorage{
		dataMock: updateData,
		errWrite: "",
		errRead:  "",
	}

	repo := InitRepository(&mockUpdate)
	result, err := repo.Update(updateData[0].Id, updateData[0].Nombre, updateData[0].Color, updateData[0].Precio, updateData[0].Stock, updateData[0].Codigo, updateData[0].Publicado, updateData[0].FechaCreacion)
	assert.Nil(t, err)
	assert.Equal(t, mockUpdate.dataMock[0], result)

}

func TestDelete(t *testing.T) {
	deleteData := []domain.Products{
		{Id: 1,
			Nombre:        "Sandia",
			Color:         "verde",
			Precio:        12332,
			Stock:         13,
			Codigo:        "asd7sd",
			Publicado:     true,
			FechaCreacion: "12/02/2022"},
	}
	mockUpdate := MockStorage{
		dataMock: deleteData,
		errWrite: "",
		errRead:  "",
	}

	repo := InitRepository(&mockUpdate)
	result := repo.Delete(1)
	assert.Equal(t, nil, result)

}
