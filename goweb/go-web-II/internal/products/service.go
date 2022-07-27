package products

import (
	"goweb/go-web-II/internal/domain"
)

/*
Acá va a ir todo lo que es EXTERNO, consultas a API y lógica.
*/

type Service interface {
	GetAll() (*[]domain.User, error)
	Store(age int, name, surname, email, created string, active bool) (domain.User, error)
	Update(id, age int, name, surname, email, created string, active bool) (domain.User, error)
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

func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}

func (s *service) Update(id, age int, name, surname, email, created string, active bool) (domain.User, error) {
	return s.repository.Update(id, age, name, surname, email, created, active)
}

func (s *service) GetAll() (*[]domain.User, error) {
	return s.repository.GetAll()
}

func (s *service) Store(age int, name, surname, email, created string, active bool) (domain.User, error) {
	lastId, err := s.repository.LastId()
	if err != nil {
		return domain.User{}, err
	}

	lastId++
	user, err := s.repository.Store(age, name, surname, email, created, active)

	if err != nil {
		return domain.User{}, err
	}

	return user, nil

}
