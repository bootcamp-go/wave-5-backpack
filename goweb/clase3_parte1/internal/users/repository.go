package users

import (
	"goweb/clase3_parte1/internal/domain"
	"fmt"
)

// Se crean las variables globales donde se guardarán las entidades
var users []domain.User
var lastID int

// Se genera la interface Repository con todos sus métodos
type Repository interface {
	GetAll() ([]domain.User, error)
	Store(nombre, apellido, email string, edad int, altura float64, activo *bool, fechaCreacion string) (domain.User, error)
	Update(id int, nombre, apellido, email string, edad int, altura float64, activo *bool, fechaCreacion string) (domain.User, error)
	UpdateLastNameAndAge(id int, apellido string, edad int) (domain.User, error)
	Delete(id int) error
}

// Se genera la estructura repository
type repository struct {}

// Se genera una función que devuelve el Repositorio
func NewRepository() Repository {
	return &repository{}
}

/* Se implementan todos los métodos correspondientes a las operaciones a realizar (GetAll, Store, etc..) */

func (r *repository) GetAll() ([]domain.User, error) {
	if len(users) == 0 {
		return users, fmt.Errorf("No hay usuarios registrados aún")
	}
	return users, nil
}

func (r *repository) Store(nombre, apellido, email string, edad int, altura float64, activo *bool, fechaCreacion string) (domain.User, error) {
	lastID++
	u := domain.User{lastID, nombre, apellido, email, edad, altura, activo, fechaCreacion}
	users = append(users, u)
	return u, nil	
}

/* Se implementa la funcionalidad para actualizar el usuario en memoria, en caso que coincida
con el ID enviado; en caso contrario, retorna un error */
func (r *repository) Update(id int, nombre, apellido, email string, edad int, altura float64, activo *bool, fechaCreacion string) (domain.User, error) {
	u := domain.User{Nombre: nombre, Apellido: apellido, Email: email, Edad: edad, Altura: altura, Activo: activo, FechaCreacion: fechaCreacion}
	updated := false
	for i := range users {
		if users[i].ID == id {
			u.ID = id
			users[i] = u
			updated = true
		}
	}
	if !updated {
		return domain.User{}, fmt.Errorf("Usuario %d no encontrado", id)
	}
	return u, nil
}

/* Se implementa la funcionalidad para actualizar el nombre y la edad del usuario en memoria,
en caso que coincida con el ID enviado; en caso contrario, se retorna un error */
func (r *repository) UpdateLastNameAndAge(id int, apellido string, edad int) (domain.User, error) {
	var u domain.User
	updated := false
	for i := range users {
		if users[i].ID == id {
			users[i].Apellido = apellido
			users[i].Edad = edad
			updated = true
			u = users[i]
		}
	}
	if !updated {
		return domain.User{}, fmt.Errorf("Usuario %d no encontrado", id)
	}
	return u, nil
}

/* Se implementa la funcionalidad para eliminar el usuario en memoria, en caso que 
coincida con el ID enviado; en caso contrario, retorna un error */
func (r *repository) Delete(id int) error {
	if len(users) == 0 {
		return fmt.Errorf("No hay usuarios registrados aún")
	}

	deleted := false
	var index int
	for i := range users {
		if users[i].ID == id {
			index = i
			deleted = true
		} 
	}
	if !deleted {
		return fmt.Errorf("Usuario %d no encontrado", id)
	}

	users = append(users[:index], users[index+1:]...)
	return nil
}



