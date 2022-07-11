package users

import (
	"github.com/bootcamp-go/wave-5-backpack/tree/gonzalez_jose/goweb/internal/domain"
)

type Service interface {
	GetAll() ([]domain.User, error)
	Store(edad int, nombre, apellido, email, fecha_creacion string, altura float64, activo bool) (domain.User, error)
	Update(id, edad int, nombre, apellido, email, fecha_creacion string, altura float64, activo bool) (domain.User, error)
	Delete(id int) error
	UpdateLastNameAndAge(id, edad int, apellido string) (domain.User, error)
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
	u, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (s *service) Store(edad int, nombre, apellido, email, fecha_creacion string, altura float64, activo bool) (domain.User, error) {

	user, err := s.repository.Store(edad, nombre, apellido, email, fecha_creacion, altura, activo)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (s *service) Update(id, edad int, nombre, apellido, email, fecha_creacion string, altura float64, activo bool) (domain.User, error) {
	return s.repository.Update(id, edad, nombre, apellido, email, fecha_creacion, altura, activo)
}

func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}

func (s *service) UpdateLastNameAndAge(id, edad int, apellido string) (domain.User, error) {
	return s.repository.UpdateLastNameAndAge(id, edad, apellido)
}
