package users

import (
	"goweb/internal/domain"
)

// variables globales----------------------------------------------------------------------------------------------------------------------
var users []domain.User
var user domain.User
//var lastID int
//-----------------------------------------------------------------------------------------------------------------------------------------


type Repository interface {
	GetAllUsers() ([]domain.User, error)
	GetUserById(id int) (domain.User, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAllUsers() ([]domain.User, error) {
	return users, nil
}

func (r *repository) GetUserById(id int) (domain.User, error) {
	return user, nil
}