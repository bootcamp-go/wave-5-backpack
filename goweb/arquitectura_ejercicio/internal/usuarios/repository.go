package usuarios

import (
	"errors"

	"github.com/anesquivel/wave-5-backpack/goweb/arquitectura_ejercicio/internal/domain"
)

type Repository interface {
	GetAll() ([]domain.Usuario, error)
	Store(id, age int, names, lastname, email, dateCreated string, estatura float64) (domain.Usuario, error)
	LastID() (int, error)
	Update(id, age int, names, lastname, email, dateCreated string, estatura float64, activo bool) (domain.Usuario, error)
	UpdateLastNameAndAge(id, age int, lastname string) (domain.Usuario, error)
	Delete(id int) error
}

type repository struct{}

var usersList []domain.Usuario
var lastID int

func (r *repository) GetAll() ([]domain.Usuario, error) {
	if len(usersList) == 0 {
		return []domain.Usuario{}, errors.New("No hay usuarios registrados.")
	}
	return usersList, nil
}

func (r *repository) Store(id, age int, names, lastname, email, dateCreated string, estatura float64) (domain.Usuario, error) {
	nwUsuario := domain.Usuario{
		Id:          id,
		Names:       names,
		LastName:    lastname,
		Age:         age,
		DateCreated: dateCreated,
		Estatura:    estatura,
		Email:       email,
		IsActivo:    true,
	}
	usersList = append(usersList, nwUsuario)
	lastID++
	return nwUsuario, nil
}

func (r *repository) Update(id, age int, names, lastname, email, dateCreated string, estatura float64, activo bool) (domain.Usuario, error) {

	upUsuario := domain.Usuario{
		Id:          id,
		Names:       names,
		LastName:    lastname,
		Age:         age,
		DateCreated: dateCreated,
		Estatura:    estatura,
		Email:       email,
		IsActivo:    activo,
	}

	update := false

	for i := range usersList {
		if usersList[i].Id == id {
			update = true
			usersList[i] = upUsuario
		}
	}

	if !update {
		return domain.Usuario{}, errors.New("No se encontró el usuario a actualizar.")
	}
	return upUsuario, nil
}

func (r *repository) UpdateLastNameAndAge(id, age int, lastname string) (domain.Usuario, error) {

	upUsuario := domain.Usuario{}

	update := false

	for i := range usersList {
		if usersList[i].Id == id {
			update = true

			usersList[i].Age = age
			usersList[i].LastName = lastname
			upUsuario = usersList[i]
		}
	}

	if !update {
		return domain.Usuario{}, errors.New("No se encontró el usuario a actualizar.")
	}
	return upUsuario, nil
}

func (r *repository) LastID() (int, error) {
	return lastID, nil
}

func (r *repository) Delete(id int) error {
	deleted := false
	var indexAux int

	for i := range usersList {
		if usersList[i].Id == id {
			deleted = true
			indexAux = i
		}
	}

	if !deleted {
		return errors.New("No se encontró el usuario a actualizar.")
	}

	usersList = append(usersList[:indexAux], usersList[indexAux+1:]...)
	return nil
}
func NewRepository() Repository {
	return &repository{}
}
