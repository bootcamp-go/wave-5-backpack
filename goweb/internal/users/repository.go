package users

import (
	"errors"
	"goweb/internal/domain"
)

var lastId int
var users []domain.User

type Repository interface {
	GetAll() ([]domain.User, error)
	Store(id int, nombre, apellido, email string, edad int, altura float64, activo bool, fechaCreacion string) (domain.User, error)
	GetById(id int) (domain.User, error)
	LastId() (int, error)
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
	return domain.User{}, errors.New("no se encontró el producto")
}

func (r *repository) LastId() (int, error) {
	return lastId, nil
}
