package users

import (
	"testing"
	"time"

	"github.com/bootcamp-go/wave-5-backpack/goweb/internal/domain"
	"github.com/go-playground/assert"
)

type StubRepository struct{}

func (r *StubRepository) GetAll() ([]domain.ModelUser, error) {
	fecha, _ := time.Parse("2006-01-02", "2022-07-25")
	users := []domain.ModelUser{
		{
			Id:            1,
			Nombre:        "Juan",
			Apellido:      "Perez",
			Email:         "juan.perez@gmail.com",
			Edad:          22,
			Altura:        1.60,
			Activo:        true,
			FechaCreacion: fecha,
			Borrado:       false,
		},
		{
			Id:            2,
			Nombre:        "Norma",
			Apellido:      "Carrasco",
			Email:         "norma.carrasco@gmail.com",
			Edad:          28,
			Altura:        1.56,
			Activo:        false,
			FechaCreacion: fecha,
			Borrado:       false,
		},
	}
	return users, nil
}

func (r *StubRepository) GetById(id int) (domain.ModelUser, error) {
	return domain.ModelUser{}, nil
}

func (r *StubRepository) Store(nombre string, apellido string, email string, edad int, altura float64) (domain.ModelUser, error) {
	return domain.ModelUser{}, nil
}

func (r *StubRepository) Update(id int, nombre string, apellido string, email string, edad int, altura float64) (domain.ModelUser, error) {
	return domain.ModelUser{}, nil
}

func (r *StubRepository) UpdateApellidoEdad(id int, apellido string, edad int) (*domain.ModelUser, error) {
	return nil, nil
}

func (r *StubRepository) Delete(id int) error {
	return nil
}

func (r *StubRepository) SearchUser(nombreQuery string, apellidoQuery string, emailQuery string, edadQuery string, alturaQuery string, activoQuery string, fechaCreacionQuery string) ([]domain.ModelUser, error) {
	return nil, nil
}

func TestGetAll(t *testing.T) {
	// Arrange
	myStubRepository := StubRepository{}
	service := NewService(&myStubRepository) // Se testea el Servicio
	fecha, _ := time.Parse("2006-01-02", "2022-07-25")
	esperado := []domain.ModelUser{
		{
			Id:            1,
			Nombre:        "Juan",
			Apellido:      "Perez",
			Email:         "juan.perez@gmail.com",
			Edad:          22,
			Altura:        1.60,
			Activo:        true,
			FechaCreacion: fecha,
			Borrado:       false,
		},
		{
			Id:            2,
			Nombre:        "Norma",
			Apellido:      "Carrasco",
			Email:         "norma.carrasco@gmail.com",
			Edad:          28,
			Altura:        1.56,
			Activo:        false,
			FechaCreacion: fecha,
			Borrado:       false,
		},
	}

	// Act
	users, _ := service.GetAll()

	// Assert
	assert.Equal(t, esperado, users)
}
