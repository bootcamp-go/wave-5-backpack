package users

import (
	"errors"
	"fmt"
	"goweb/internal/domain"
	"strconv"
)

var lastId int
var users []domain.User

type Repository interface {
	GetAll() ([]domain.User, error)
	Store(id int, nombre, apellido, email string, edad int, altura float64, activo bool, fechaCreacion string) (domain.User, error)
	GetById(id int) (domain.User, error)
	LastId() (int, error)
	Update(id int, nombre, apellido, email string, edad int, altura float64, activo bool, fechaCreacion string) (domain.User, error)
	Delete(id int) error
	Patch(id int, apellido string, edad int) (domain.User, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAll() ([]domain.User, error) {
	return users, nil
}

func (r *repository) Store(id int, nombre, apellido, email string, edad int, altura float64, activo bool, fechaCreacion string) (domain.User, error) {
	user := domain.User{Id: id, Nombre: nombre, Apellido: apellido, Email: email, Edad: edad, Altura: altura, Activo: activo, FechaCreacion: fechaCreacion}

	users = append(users, user)
	lastId = id

	return user, nil
}

func (r *repository) GetById(id int) (domain.User, error) {
	for _, user := range users {
		if user.Id == id {
			return user, nil
		}
	}
	return domain.User{}, fmt.Errorf("No se encontró el usuario con el ID %d", id)
}

func (r *repository) LastId() (int, error) {
	return lastId, nil
}

func (r *repository) Update(id int, nombre, apellido, email string, edad int, altura float64, activo bool, fechaCreacion string) (domain.User, error) {
	user := domain.User{Id: id, Nombre: nombre, Apellido: apellido, Email: email, Edad: edad, Altura: altura, Activo: activo, FechaCreacion: fechaCreacion}
	update := false
	for index, v := range users {
		if v.Id == id {
			user.Id = id
			users[index] = user
			update = true
		}
	}
	if !update {
		return domain.User{}, fmt.Errorf("No se encontró el usuario con el ID %d", id)
	}

	return user, nil
}

func (r *repository) Delete(id int) error {
	userV := -1
	for i, u := range users {
		if u.Id == id {
			userV = i
			break
		}
	}

	if userV == -1 {
		return errors.New("el usuario con el ID " + strconv.Itoa(id) + " no fue encontrado")
	}

	users = append(users[:userV], users[userV+1:]...)
	return nil
}

func (r *repository) Patch(id int, apellido string, edad int) (domain.User, error) {
	var user domain.User
	updated := false
	for i, v := range users {
		if v.Id == id {
			users[i].Apellido = apellido
			users[i].Edad = edad
			updated = true
			user = users[i]
		}
	}
	if !updated {
		return domain.User{}, fmt.Errorf("No se encontró el usuario con el ID %d", id)
	}

	return user, nil
}
