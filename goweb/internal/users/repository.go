package users

import (
	"fmt"
	"time"

	"github.com/bootcamp-go/wave-5-backpack/goweb/internal/domain"
	"github.com/bootcamp-go/wave-5-backpack/goweb/pkg/store"
)

//var usersLocal []domain.ModelUser
//var lastId int
const (
	errorLectura   = "no se puede leer la db, error: %s"
	errorEscritura = "no se puede escribir en la db, error: %s"
)

type Repository interface {
	GetAll() ([]domain.ModelUser, error)
	GetById(id int) (domain.ModelUser, error)
	Store(nombre string, apellido string, email string, edad int, altura float64, activo bool) (domain.ModelUser, error)
	Update(id int, nombre string, apellido string, email string, edad int, altura float64, activo bool) (domain.ModelUser, error)
	UpdateApellidoEdad(id int, nombre string, edad int) (domain.ModelUser, error)
	Delete(id int) error
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll() ([]domain.ModelUser, error) {
	// Leemos los usuarios del JSON
	var users []domain.ModelUser
	if err := r.db.Read(&users); err != nil {
		return nil, fmt.Errorf(errorLectura, err)
	}

	// Devolvemos todos los usuarios del JSON
	return users, nil
}

func (r *repository) GetById(id int) (domain.ModelUser, error) {
	var user domain.ModelUser

	// Leemos los usuarios del JSON
	var users []domain.ModelUser
	if err := r.db.Read(&users); err != nil {
		return domain.ModelUser{}, fmt.Errorf(errorLectura, err)
	}

	// Buscamos el usuario por id
	found := false
	for i := range users {
		if users[i].Id == id && !found {
			user = users[i]
			found = true
		}
	}

	if !found {
		return domain.ModelUser{}, fmt.Errorf("usuario %d no encontrado", id)
	}

	// Devolvemos el usuario por id
	return user, nil
}

func (r *repository) Store(nombre string, apellido string, email string, edad int, altura float64, activo bool) (domain.ModelUser, error) {
	user := domain.ModelUser{Nombre: nombre, Apellido: apellido, Email: email, Edad: edad, Altura: altura, Activo: activo}
	lastId := 0

	// Leemos los usuarios del JSON
	var users []domain.ModelUser
	if err := r.db.Read(&users); err != nil {
		return domain.ModelUser{}, fmt.Errorf(errorLectura, err)
	}

	// Se Calcula el id siguiente
	for i := range users {
		if users[i].Id > lastId {
			lastId = users[i].Id
		}
	}

	// Se genera el id consecutivo y la fecha de creación
	user.Id = lastId + 1
	user.FechaCreacion = time.Now()

	// Se adiciona el usuario al slice de usuarios
	users = append(users, user)

	// Se guarda la información y se verifica que no haya ocurrido un error
	if err := r.db.Write(&users); err != nil {
		return domain.ModelUser{}, fmt.Errorf(errorEscritura, err)
	}

	// Devolvemos el nuevo usuario
	return user, nil
}

func (r *repository) Update(id int, nombre string, apellido string, email string, edad int, altura float64, activo bool) (domain.ModelUser, error) {
	// Leemos los usuarios del JSON
	var users []domain.ModelUser
	if err := r.db.Read(&users); err != nil {
		return domain.ModelUser{}, fmt.Errorf(errorLectura, err)
	}

	user := domain.ModelUser{Nombre: nombre, Apellido: apellido, Email: email, Edad: edad, Altura: altura, Activo: activo}
	found := false
	for i := range users {
		if users[i].Id == id && !found {
			user.Id = id
			user.FechaCreacion = users[i].FechaCreacion
			users[i] = user
			found = true
		}
	}

	// Verificamos que haya existido el usuario a actualizar
	if !found {
		return domain.ModelUser{}, fmt.Errorf("usuario %d no encontrado", id)
	}

	// Se guarda la información y se verifica que no haya ocurrido un error
	if err := r.db.Write(&users); err != nil {
		return domain.ModelUser{}, fmt.Errorf(errorEscritura, err)
	}

	// Se devuelve el usuario actualizado
	return user, nil
}

func (r *repository) UpdateApellidoEdad(id int, apellido string, edad int) (domain.ModelUser, error) {
	// Leemos los usuarios del JSON
	var users []domain.ModelUser
	if err := r.db.Read(&users); err != nil {
		return domain.ModelUser{}, fmt.Errorf(errorLectura, err)
	}

	var user domain.ModelUser
	found := false
	for i := range users {
		if users[i].Id == id && !found {
			users[i].Apellido = apellido
			users[i].Edad = edad
			user = users[i]
			found = true
		}
	}

	// Verificamos que haya existido el usuario a actualizar su Apellido y Edad
	if !found {
		return domain.ModelUser{}, fmt.Errorf("usuario %d no encontrado", id)
	}

	// Se guarda la información y se verifica que no haya ocurrido un error
	if err := r.db.Write(&users); err != nil {
		return domain.ModelUser{}, fmt.Errorf(errorEscritura, err)
	}

	// Se devuelve el usuario actualizado del Apellido y Edad
	return user, nil
}

func (r *repository) Delete(id int) error {
	// Leemos los usuarios del JSON
	var users []domain.ModelUser
	if err := r.db.Read(&users); err != nil {
		return fmt.Errorf(errorLectura, err)
	}

	found := false
	var index int
	for i := range users {
		if users[i].Id == id && !found {
			index = i
			found = true
		}
	}

	// Verificamos que haya existido el usuario a borrar
	if !found {
		return fmt.Errorf("usuario %d no econtrado", id)
	}

	// Se quita del slice el usuario
	users = append(users[:index], users[index+1:]...)

	// Se guarda la información y se verifica que no haya ocurrido un error
	if err := r.db.Write(&users); err != nil {
		return fmt.Errorf(errorEscritura, err)
	}

	return nil
}
