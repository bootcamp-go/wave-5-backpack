package usuarios

import (
	"fmt"
	"goweb/Clase3-1-WebServers/internal/domain"
	"time"
)

type Repository interface {
	GetAll() ([]domain.Usuario, error)
	Store(id int, nombre, apellido, email string, edad, altura int, activo bool) (domain.Usuario, error)
	LastID() (int, error)
	Update(id int, nombre, apellido, email string, edad, altura int, activo bool) (domain.Usuario, error)
	UpdateSurnameAndAge(id int, surname string, age int) (domain.Usuario, error)
	Delete(id int) error
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

func (r *repository) Store(id int, nombre, apellido, email string, edad, altura int, activo bool) (domain.Usuario, error) {
	u := domain.Usuario{
		Id:              id,
		Nombre:          nombre,
		Apellido:        apellido,
		Email:           email,
		Edad:            edad,
		Altura:          altura,
		Activo:          activo,
		FechaDeCreacion: time.Now().String(),
	}
	us = append(us, u)
	lastID = u.Id
	return u, nil
}
func (r *repository) Update(id int, nombre, apellido, email string, edad, altura int, activo bool) (domain.Usuario, error) {
	u := domain.Usuario{
		Id:              id,
		Nombre:          nombre,
		Apellido:        apellido,
		Email:           email,
		Edad:            edad,
		Altura:          altura,
		Activo:          activo,
		FechaDeCreacion: time.Now().String(),
	}
	updated := false
	for i := range us {
		if us[i].Id == id {
			u.Id = id
			us[i] = u
			updated = true
		}
	}

	if !updated {
		return domain.Usuario{}, fmt.Errorf("usuario %d no encontrado", id)
	}

	return u, nil
}

func (r *repository) UpdateSurnameAndAge(id int, surname string, age int) (domain.Usuario, error) {
	updated := false
	var u domain.Usuario
	for i := range us {
		if us[i].Id == id {
			us[i].Apellido = surname
			us[i].Edad = age
			u = us[i]
			updated = true
		}
	}

	if !updated {
		return domain.Usuario{}, fmt.Errorf("usuario %d no encontrado", id)
	}

	return u, nil
}

func (r *repository) Delete(id int) error {
	deleted := false
	var index int
	for i := range us {
		if us[i].Id == id {
			index = i
			deleted = true
		}
	}

	if !deleted {
		return fmt.Errorf("usuario %d no encontrado", id)
	}

	us = append(us[:index], us[index+1:]...)
	return nil
}
