/*Repositorio, debe tener el acceso a la variable guardada en memoria.
OK Se debe crear el archivo repository.go
OK Se deben crear las variables globales donde guardar las entidades
OK Se debe generar la interface Repository con todos sus métodos
OK Se debe generar la estructura repository
OK Se debe generar una función que devuelva el Repositorio
OK Se deben implementar todos los métodos correspondientes a las operaciones a realizar (GetAll, Store, etc..)
*/
package usuarios

import (
	"fmt"

	"github.com/bootcamp-go/wave-5-backpack/goweb/internal/domain"
)

var us []domain.Usuarios
var lastID int

type Repository interface {
	GetAll() ([]domain.Usuarios, error)
	Guardar(id int, nombre string, apellido string, email string, edad int, altura float64, actico bool, fecha string) (domain.Usuarios, error)
	lastId() (int, error)
	Update(id int, nombre, apellido, email string, edad int, altura float64, activo bool, fecha string) (domain.Usuarios, error)
	Delete(id int) error
	UpdateNameAndLastName(id int, name string, apellido string) (domain.Usuarios, error)
}

type repository struct {
}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) UpdateNameAndLastName(id int, name string, last string) (domain.Usuarios, error) {
	var u domain.Usuarios
	updated := false
	for i := range us {
		if us[i].Id == id {
			us[i].Nombre = name
			us[i].Apellido = last
			updated = true
			u = us[i]
		}
	}
	if !updated {
		return domain.Usuarios{}, fmt.Errorf("producto con id %d no encontrado", id)
	}
	return u, nil
}

func (r *repository) Update(id int, nombre, apellido, email string, edad int, altura float64, activo bool, fecha string) (domain.Usuarios, error) {
	u := domain.Usuarios{Nombre: nombre, Apellido: apellido, Email: email, Edad: edad, Altura: altura, Activo: activo, FechaCreacion: fecha}
	actualizado := false
	for i := range us {
		if us[i].Id == id {
			u.Id = id
			us[i] = u
			actualizado = true
		}
	}
	if !actualizado {
		return domain.Usuarios{}, fmt.Errorf("el producto con id %d no fue encontrado", id)
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
		return fmt.Errorf("producto %d no encontrado", id)
	}
	us = append(us[:index], us[index+1:]...)
	return nil
}

func (r *repository) GetAll() ([]domain.Usuarios, error) {
	return us, nil
}
func (r *repository) Guardar(id int, nombre string, apellido string, email string, edad int, altura float64, actico bool, fecha string) (domain.Usuarios, error) {
	u := domain.Usuarios{Id: id, Nombre: nombre, Apellido: apellido, Email: email, Edad: edad, Altura: altura, Activo: actico, FechaCreacion: fecha}
	us = append(us, u)
	lastID = u.Id
	return u, nil
}
func (r *repository) lastId() (int, error) {
	return lastID, nil
}
