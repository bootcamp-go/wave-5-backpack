package users

import (
	"github.com/bootcamp-go/wave-5-backpack/internal/domain"
)

type Repository interface {
	GetAll() ([]domain.User, error)
	LastId() (int, error)
	StoreUser(id int, name, lastname, email string, age int, height float32, active bool, doCreation string) (domain.User, error)
}

type repository struct{}

func NewRepositoy() Repository {
	return &repository{}
}

var us []domain.User
var lastId int

func (r *repository) GetAll() ([]domain.User, error) {
	return us, nil
}

func (r *repository) LastId() (int, error) {
	return lastId, nil

}

func (r *repository) StoreUser(id int, name, lastname, email string, age int, height float32, active bool, doCreation string) (domain.User, error) {
	u := domain.User{ID: id, Name: name, Lastname: lastname, Email: email, Age: age, Height: height, Active: active, DoCreation: doCreation}
	us = append(us, u)
	lastId = u.ID
	return u, nil
}
