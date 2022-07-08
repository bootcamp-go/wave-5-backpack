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

// Método llamado por PUT para actualizar toda la entidad
func (s *service) Update(id int, nombre string, apellido string, email string, edad int, altura float64, activo bool) (domain.ModelUser, error) {
	return s.repository.Update(id, nombre, apellido, email, edad, altura, activo)
}

// Método llamado por PATCH para actualizar 2 campos de la entidad
func (s *service) UpdateApellidoEdad(id int, apellido string, edad int) (domain.ModelUser, error) {
	return s.repository.UpdateApellidoEdad(id, apellido, edad)
}

// Método llamado por DELETE para borrar la entidad
func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}
