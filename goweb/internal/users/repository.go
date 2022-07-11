package users

import (
	"errors"
	"fmt"
	"github.com/bootcamp-go/wave-5-backpack/internal/domain"
)

type Repository interface {
	GetAll() ([]domain.User, error)
	LastId() (int, error)
	GetById(id int) (domain.User, error)
	StoreUser(id int, name, lastname, email string, age int, height float32, active bool, doCreation string) (domain.User, error)
	UpdateUser(id int, name, lastname, email string, age int, height float32, active bool, doCreation string) (domain.User, error)
}

type repository struct{}

func NewRepositoy() Repository {
	return &repository{}
}

var allUsers []domain.User
var lastId int

func (r *repository) GetAll() ([]domain.User, error) {
	return allUsers, nil
}
func (r *repository) GetById(id int) (domain.User, error) {
	for _, u := range allUsers {
		if u.ID == id {
			return u, nil
		}
	}
	err := errors.New("error: User not found")
	return domain.User{}, err
}

func (r *repository) LastId() (int, error) {
	return lastId, nil

}

func (r *repository) StoreUser(id int, name, lastname, email string, age int, height float32, active bool, doCreation string) (domain.User, error) {
	newUser := domain.User{
		ID: id, 
		Name: name, 
		Lastname: lastname, 
		Email: email, 
		Age: age, 
		Height: height, 
		Active: active, 
		DoCreation: doCreation,
	}
	allUsers = append(allUsers, newUser)
	lastId = newUser.ID
	return newUser, nil
}

func (r *repository) UpdateUser(id int, name, lastname, email string, age int, height float32, active bool, doCreation string) (domain.User, error) {
	newUser := domain.User{
		ID: id, 
		Name: name, 
		Lastname: lastname, 
		Email: email, 
		Age: age, 
		Height: height, 
		Active: active, 
		DoCreation: doCreation,
	}
	for index := range allUsers {
		if allUsers[index].ID == id {
			allUsers[index] = newUser
			return newUser, nil
		}
	}
	return domain.User{}, fmt.Errorf("user id %d not found", id)

}