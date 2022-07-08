package users

import (
	"errors"
	"fmt"
	"goweb/internal/domain"
)

type Service interface {
	GetAll() ([]domain.User, error)
	Store(Nombre string, Apellido string, Email string, Edad int, Altura float64, Activo bool, FechaCreacion string) (domain.User, error)
	GetById(id int) (domain.User, error)
	Update(id int, nombre, apellido, email string, edad int, altura float64, activo bool, fechaCreacion string) (domain.User, error)
	Delete(id int) error
	Patch(id int, apellido string, edad int) (domain.User, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll() ([]domain.User, error) {
	users, err := s.repository.GetAll()
	if err != nil {
		return nil, errors.New("no se pudo obtener los usuarios")
	}

	return users, nil
}
//ANTERIOR
/* func (s *service) Store(nombre string, apellido string, email string, edad int, altura float64, activo bool, fechaCreacion string) (domain.User, error) {
	id, err := s.repository.LastId()
	if err != nil {
		return domain.User{}, errors.New("no se pudo cargar el último id de los usuarios")
	}

	id++

	user, err := s.repository.Store(id, nombre, apellido, email, edad, altura, activo, fechaCreacion)
	if err != nil {
		return domain.User{}, errors.New("no se pudo guardar el usuario")
	}

	return user, nil
} */

func (s *service) Store(nombre string, apellido string, email string, edad int, altura float64, activo bool, fechaCreacion string) (domain.User, error) {
	lastID, err := s.repository.LastId()
	if err != nil {
		return domain.User{}, fmt.Errorf("no se pudo cargar el último id de los usuarios")
	}

	lastID++

	user, err := s.repository.Store(lastID, nombre, apellido, email, edad, altura, activo, fechaCreacion)
	if err != nil {
		return domain.User{}, errors.New("no se pudo guardar el usuario")
	}

	return user, nil
}

func (s *service) GetById(id int) (domain.User, error) {
	user, err := s.repository.GetById(id)
	if err != nil {
		return domain.User{}, fmt.Errorf("no se pudo encontrar el usuario con el id: %d", id)
	}
	return user, nil
}

func (s *service) Update(id int, nombre, apellido, email string, edad int, altura float64, activo bool, fechaCreacion string) (domain.User, error) {
	return s.repository.Update(id, nombre, apellido, email, edad, altura, activo, fechaCreacion)
}

func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}

func (s *service) Patch(id int, apellido string, edad int) (domain.User, error) {
	return s.repository.Patch(id, apellido, edad)
}
