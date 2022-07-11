package users

import "goweb/internal/domain"

type Repository interface {
	GetAll() ([]domain.User, error)
}

var users []domain.User = make([]domain.User, 0)

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAll() ([]domain.User, error) {
	return users, nil
}
