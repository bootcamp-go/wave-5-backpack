package users

import (
	"github.com/bootcamp-go/wave-5-backpack/goweb/internal/domain"
)

type Service interface {
	GetAll() ([]domain.ModelUser, error)
	GetById(id int) (domain.ModelUser, error)
	Store(nombre string, apellido string, email string, edad int, altura float64, activo bool) (domain.ModelUser, error)
	Update(id int, nombre string, apellido string, email string, edad int, altura float64, activo bool) (domain.ModelUser, error)
	UpdateApellidoEdad(id int, nombre string, edad int) (domain.ModelUser, error)
	Delete(id int) error
	SearchUser(nombreQuery string, apellidoQuery string, emailQuery string, edadQuery string, alturaQuery string, activoQuery string, fechaCreacionQuery string) ([]domain.ModelUser, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll() ([]domain.ModelUser, error) {
	users, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return users, nil
}

// Función para devolver una entidad por id
func (s *service) GetById(id int) (domain.ModelUser, error) {
	return s.repository.GetById(id)
}

// Función para guardar una entidad
func (s *service) Store(nombre string, apellido string, email string, edad int, altura float64, activo bool) (domain.ModelUser, error) {
	user, err := s.repository.Store(nombre, apellido, email, edad, altura, activo)
	if err != nil {
		return domain.ModelUser{}, err
	}

	return user, nil
}

// Función para actualizar una entidad completa
func (s *service) Update(id int, nombre string, apellido string, email string, edad int, altura float64, activo bool) (domain.ModelUser, error) {
	return s.repository.Update(id, nombre, apellido, email, edad, altura, activo)
}

// Función para actualizar 2 campos de una entidad
func (s *service) UpdateApellidoEdad(id int, apellido string, edad int) (domain.ModelUser, error) {
	return s.repository.UpdateApellidoEdad(id, apellido, edad)
}

// Función para borrar una entidad
func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}

// Función para buscar una entidad
func (s *service) SearchUser(nombreQuery string, apellidoQuery string, emailQuery string, edadQuery string, alturaQuery string, activoQuery string, fechaCreacionQuery string) ([]domain.ModelUser, error) {
	return s.repository.SearchUser(nombreQuery, apellidoQuery, emailQuery, edadQuery, alturaQuery, activoQuery, fechaCreacionQuery)
}
