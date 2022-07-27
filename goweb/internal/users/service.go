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
	user, err := s.repository.Update(id, nombre, apellido, email, edad, altura, activo, fechaCreacion)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (s *service) Delete(id int) error {
	if err := s.repository.Delete(id); err != nil {
		return fmt.Errorf("error al eliminar el usuario %d", id)
	}
	return nil
}

func (s *service) Patch(id int, apellido string, edad int) (domain.User, error) {
	user, err := s.repository.Patch(id, apellido, edad)

	if err != nil {
		return domain.User{}, fmt.Errorf("error al actualizar datos del usuario %d", id)
	}

	return user, nil
}
