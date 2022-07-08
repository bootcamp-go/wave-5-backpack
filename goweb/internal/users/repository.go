package users

import (
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