package users

import (
	"fmt"
	"goweb/internal/domain"
)

// variables globales----------------------------------------------------------------------------------------------------------------------
var users []domain.User
var user domain.User
var lastID int
//-----------------------------------------------------------------------------------------------------------------------------------------



type Repository interface {
	GetAllUsers() ([]domain.User, error)
	GetUserById(id int) (domain.User, error)
	StoreUser(id int, name, lastname, email string, age int, height float32, active bool, createdat string) (domain.User, error)
	LastID() (int,error)
	UpdateTotal(id int, name, lastname, email string, age int, height float32, active bool, createdat string) (domain.User, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAllUsers() ([]domain.User, error) {
	return users, nil
}

func (r *repository) GetUserById(id int) (domain.User, error) {
	var userFounded domain.User
	for _,u :=range users{
		if u.Id == id{
			userFounded = u
		}
	}
	return userFounded, nil
}

func (r *repository) LastID() (int, error) {
	return lastID, nil
}

func (r *repository) StoreUser(id int, name, lastname, email string, age int, height float32, active bool, createdat string) (domain.User, error){
	user = domain.User{Id: id, Name:name, LastName: lastname, Email: email, Age: age, Height: height, Active: active, CreatedAt: createdat}

	users = append(users, user)
	lastID= user.Id

	return user, nil
}

func(r *repository) UpdateTotal(id int, name, lastname, email string, age int, height float32, active bool, createdat string) (domain.User, error) {
	userToUpdate := domain.User{Name: name, LastName: lastname, Email: email, Age: age, Height: height, Active: active, CreatedAt: createdat}
	updated := false
	for i:= range users {
		if users[i].Id == id{
			userToUpdate.Id = id
			users[i] = userToUpdate
			updated = true
			break
		}
	}
	if !updated{
		return domain.User{}, fmt.Errorf("Usuario %d no encontrado", id)
	}
	return userToUpdate, nil
}