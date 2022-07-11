package users

import "github.com/bootcamp-go/wave-5-backpack/internal/domain"

type Service interface {
	GetAll() ([]domain.User, error)
	StoreUser(name, lastname, email string, age int, height float32, active bool, doCreation string) (domain.User, error)
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
	us, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return us, nil
}

func (s *service) StoreUser(name, lastname, email string, age int, height float32, active bool, doCreation string) (domain.User, error) {
	lastID, err := s.repository.LastId()
	if err != nil {
		return domain.User{}, err
	}

	lastID++

	newUser, err := s.repository.StoreUser(lastID, name, lastname, email, age, height, active, doCreation)

	if err != nil {
		return domain.User{}, err
	}

	return newUser, nil
}
