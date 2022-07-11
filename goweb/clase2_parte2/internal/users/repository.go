package users

import "goweb/clase2_parte2/internal/domain"

// Se deben crear las variables globales donde guardar las entidades
var us []domain.User
var lastID int

// Se debe generar la interface Repository con todos sus métodos
type Repository interface {
	GetAll() ([]domain.User, error)
	Store(id int, nombre, apellido, email string, edad int, altura float64, activo bool, fechaCreacion string) (domain.User, error)
	LastID() (int, error)
}

// Se debe generar la estructura repository
type repository struct {}

// Se debe generar una función que devuelva el Repositorio
func NewRepository() Repository {
	return &repository{}
}

// Se deben implementar todos los métodos correspondientes a las operaciones a realizar (GetAll, Store, etc..)
func (r *repository) GetAll() ([]domain.User, error) {
	return us, nil
}

func (r *repository) LastID() (int, error) {
	return lastID, nil
}

func (r *repository) Store(id int, nombre, apellido, email string, edad int, altura float64, activo bool, fechaCreacion string) (domain.User, error) {
	u := domain.User{id, nombre, apellido, email, edad, altura, activo, fechaCreacion}
	us = append(us, u)
	lastID = u.ID
	return u, nil
}



