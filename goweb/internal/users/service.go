package users

import "github.com/bootcamp-go/wave-5-backpack/goweb/internal/domain"

type Service interface {
	GetAll() ([]domain.Users, error)
	Store(age int, name, lastName, email, creationDate string, height float64, active bool) (domain.Users, error)
}
type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll() ([]domain.Users, error) {
	ps, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return ps, nil
}

func (s *service) Store(age int, name, lastName, email, creationDate string, height float64, active bool) (domain.Users, error) {
	lastID, err := s.repository.LastID()
	if err != nil {
		return domain.Users{}, err
	}

	lastID++

	producto, err := s.repository.Store(lastID, age, name, lastName, email, creationDate, height, active)
	if err != nil {
		return domain.Users{}, err
	}

	return producto, nil
}
