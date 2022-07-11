package users

import (
	"errors"
	"goweb/internal/domain"
)

type Service interface {
	GetAll() ([]domain.User, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) GetAll() ([]domain.User, error) {
	users, err := s.repo.GetAll()
	if err != nil {
		return nil, errors.New("error loading users")
	}

	return users, nil
}
