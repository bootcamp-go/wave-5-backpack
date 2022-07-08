package users

import "github.com/bootcamp-go/wave-5-backpack/goweb/internal/domain"

type Service interface {
	GetAll() ([]domain.ModelUser, error)
	Store(nombre string, apellido string, email string, edad int, altura float64) (domain.ModelUser, error)
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

func (s *service) Store(nombre string, apellido string, email string, edad int, altura float64) (domain.ModelUser, error) {
	lastId, err := s.repository.LastId()
	if err != nil {
		return domain.ModelUser{}, err
	}

	lastId++

	user, err := s.repository.Store(lastId, nombre, apellido, email, edad, altura)
	if err != nil {
		return domain.ModelUser{}, err
	}

	return user, nil
}
