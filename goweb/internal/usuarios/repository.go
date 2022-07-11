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
	"github.com/bootcamp-go/wave-5-backpack/goweb/pkg/store"
)

const (
	UsuarioNotFound = "producto %d not found"
	FailReading     = "error al leer la bd"
	FailWriting     = "error al escribir la bd"
)

var us []domain.Usuarios

type Repository interface {
	GetAll() ([]domain.Usuarios, error)
	Guardar(id int, nombre string, apellido string, email string, edad int, altura float64, actico bool, fecha string) (domain.Usuarios, error)
	LastId() (int, error)
	Update(id int, nombre, apellido, email string, edad int, altura float64, activo bool, fecha string) (domain.Usuarios, error)
	Delete(id int) error
	UpdateNameAndLastName(id int, name string, apellido string) (domain.Usuarios, error)
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
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
	var us []domain.Usuarios

	if err := r.db.Read(&us); err != nil {
		return domain.Usuarios{}, fmt.Errorf(FailReading)
	}

	for i := range us {
		user := &us[i]
		if user.Id == id {
			user.Nombre = nombre
			user.Apellido = apellido
			user.Email = email
			user.FechaCreacion = fecha
			user.Activo = activo
			user.Edad = edad
			user.Altura = altura

			if err := r.db.Write(us); err != nil {
				return domain.Usuarios{}, fmt.Errorf(FailWriting)
			}
			return *user, nil
		}
	}
	return domain.Usuarios{}, fmt.Errorf(UsuarioNotFound, id)
}

func (r *repository) Delete(id int) error {
	var listaUs []domain.Usuarios
	if err := r.db.Read(&listaUs); err != nil {
		return fmt.Errorf(FailReading)
	}

	for i := range listaUs {
		user := listaUs[i]
		if user.Id == id {
			listaUs = append(listaUs[:i], listaUs[i+1:]...)
			if err := r.db.Write(listaUs); err != nil {
				return fmt.Errorf(FailWriting)
			}
			return nil
		}
	}
	return nil
}

func (r *repository) GetAll() ([]domain.Usuarios, error) {
	var us []domain.Usuarios
	if err := r.db.Read(&us); err != nil {
		return nil, fmt.Errorf(FailReading)
	}
	return us, nil
}
func (r *repository) Guardar(id int, nombre string, apellido string, email string, edad int, altura float64, actico bool, fecha string) (domain.Usuarios, error) {
	var us []domain.Usuarios

	if err := r.db.Read(&us); err != nil {
		return domain.Usuarios{}, fmt.Errorf(FailReading)
	}

	u := domain.Usuarios{Id: id, Nombre: nombre, Apellido: apellido, Email: email, Edad: edad, Altura: altura, Activo: actico, FechaCreacion: fecha}
	us = append(us, u)

	if err := r.db.Write(us); err != nil {
		return domain.Usuarios{}, fmt.Errorf(FailWriting)
	}
	return u, nil
}
func (r *repository) LastId() (int, error) {
	var us []domain.Usuarios
	if err := r.db.Read(&us); err != nil {
		return 0, fmt.Errorf(FailReading)
	}
	if len(us) == 0 {
		return 0, nil
	}
	return us[len(us)-1].Id, nil
}
