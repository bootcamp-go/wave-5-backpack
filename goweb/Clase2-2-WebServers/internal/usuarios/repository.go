package usuarios

import (
	"goweb/Clase2-2-WebServers/internal/domain"
)

type Repository interface {
	GetAll() ([]domain.Usuario, error)
	Store(id int, nombre, apellido, email string, edad, altura int, activo bool, fecha string) (domain.Usuario, error)
	LastID() (int, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

var us []domain.Usuario
var lastID int

func (r *repository) GetAll() ([]domain.Usuario, error) {
	return us, nil
}

func (r *repository) LastID() (int, error) {
	return lastID, nil
}

func (r *repository) Store(id int, nombre, apellido, email string, edad, altura int, activo bool, fecha string) (domain.Usuario, error) {
	u := domain.Usuario{
		Id:              id,
		Nombre:          nombre,
		Apellido:        apellido,
		Email:           email,
		Edad:            edad,
		Altura:          altura,
		Activo:          activo,
		FechaDeCreacion: fecha,
	}
	us = append(us, u)
	lastID = u.Id
	return u, nil
}
