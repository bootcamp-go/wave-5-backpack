package users

import (
	"goweb/internal/domain"
)

type Service interface{
	GetAllUsers() ([]domain.User, error)
	GetUserById(id int) (domain.User, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAllUsers() ([]domain.User, error) {
	users, err := s.repository.GetAllUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *service) GetUserById(id int) (domain.User, error) {
	user, err := s.repository.GetUserById(id)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}


