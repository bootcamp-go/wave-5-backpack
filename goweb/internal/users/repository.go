package users

import (
	"fmt"

	"github.com/bootcamp-go/wave-5-backpack/goweb/internal/domain"
)

type Repository interface {
	GetAll() ([]domain.Users, error)
	Store(id, age int, name, lastName, email, creationDate string, height float64, active bool) (domain.Users, error)
	LastID() (int, error)
	Update(id, age int, name, lastName, email, creationDate string, height float64, active bool) (domain.Users, error)
	UpdateLastNameAndAge(id, age int, lastName string) (domain.Users, error)
	Delete(id int) error
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

var users []domain.Users = []domain.Users{
	{},
	{Id: 1, Name: "ABC", LastName: "ABC", Email: "a@a.com", Age: 21, Height: 1.82, Active: true, CreationDate: "2022-02-20"},
	{Id: 2, Name: "BCD", LastName: "BCD", Email: "b@b.com", Age: 30, Height: 1.76, Active: false, CreationDate: "2022-02-25"},
}
var lastID int

func (r *repository) GetAll() ([]domain.Users, error) {
	return users, nil
}

func (r *repository) LastID() (int, error) {
	return lastID, nil
}

func (r *repository) Store(id, age int, name, lastName, email, creationDate string, height float64, active bool) (domain.Users, error) {
	user := domain.Users{
		Id: id, Name: name, LastName: lastName, Email: email, Age: age, Height: height, Active: active, CreationDate: creationDate}
	users = append(users, user)
	lastID = user.Id
	return user, nil
}

func (r *repository) Update(id, age int, name, lastName, email, creationDate string, height float64, active bool) (domain.Users, error) {
	for i := range users {
		user := &users[i]
		if user.Id == id {
			user.Name = name
			user.LastName = lastName
			user.Email = email
			user.CreationDate = creationDate
			user.Active = active
			user.Age = age
			user.Height = height
			return *user, nil
		}
	}
	return domain.Users{}, fmt.Errorf("producto %d no encontrado", id)
}

func (r *repository) UpdateLastNameAndAge(id, age int, lastName string) (domain.Users, error) {
	for i := range users {
		user := &users[i]
		if user.Id == id {
			user.LastName = lastName
			user.Age = age
			return *user, nil
		}
	}
	return domain.Users{}, fmt.Errorf("producto %d no encontrado", id)
}

func (r *repository) Delete(id int) error {
	for i := range users {
		user := users[i]
		if user.Id == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("producto %d no encontrado", id)
}
