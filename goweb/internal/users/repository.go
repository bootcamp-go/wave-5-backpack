package users

import (
	"errors"
	"fmt"
	"goweb/internal/domain"
	"goweb/pkg/store"
	"strconv"
)

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

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{db: db}
}

func (r *repository) GetAll() ([]domain.User, error) {
	if err := r.db.Read(&users); err != nil {
		return nil, fmt.Errorf("error al leer el archivo")
	}
	return users, nil
}

//ANTERIOR
/* func (r *repository) Store(id int, nombre, apellido, email string, edad int, altura float64, activo bool, fechaCreacion string) (domain.User, error) {
	user := domain.User{Id: id, Nombre: nombre, Apellido: apellido, Email: email, Edad: edad, Altura: altura, Activo: activo, FechaCreacion: fechaCreacion}

	users = append(users, user)
	lastId = id

	return user, nil
} */

func (r *repository) Store(id int, nombre, apellido, email string, edad int, altura float64, activo bool, fechaCreacion string) (domain.User, error) {
	var user []domain.User
	if err := r.db.Read(&user); err != nil {
		return domain.User{}, fmt.Errorf("error al leer el archivo")
	}

	newUser := domain.User{Id: id, Nombre: nombre, Apellido: apellido, Email: email, Edad: edad, Altura: altura, Activo: activo, FechaCreacion: fechaCreacion}

	user = append(user, newUser)

	if err := r.db.Write(user); err != nil {
		return domain.User{}, fmt.Errorf("error al escribir en el archivo, error: %w", err)
	}

	return newUser, nil
}

func (r *repository) GetById(id int) (domain.User, error) {
	for _, user := range users {
		if user.Id == id {
			return user, nil
		}
	}
	return domain.User{}, fmt.Errorf("no se encontró el usuario con el ID %d", id)
}

func (r *repository) LastId() (int, error) {
	var user []domain.User
	if err := r.db.Read(&user); err != nil {
		return 0, fmt.Errorf("error al leer el archivo")
	}
	if len(user) == 0 {
		return 0, nil
	}

	return user[len(user)-1].Id, nil

}

func (r *repository) Update(id int, nombre, apellido, email string, edad int, altura float64, activo bool, fechaCreacion string) (domain.User, error) {
	if err := r.db.Read(&users); err != nil {
		return domain.User{}, fmt.Errorf("error al leer el archivo")
	}

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
		return domain.User{}, fmt.Errorf("no se encontró el usuario con el ID %d", id)
	}

	if err := r.db.Write(users); err != nil {
		return domain.User{}, fmt.Errorf("error al escribir en el archivo, error: %w", err)
	}

	return user, nil
}

func (r *repository) Delete(id int) error {
	if err := r.db.Read(&users); err != nil {
		return fmt.Errorf("error al leer el archivo")
	}

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

	if err := r.db.Write(users); err != nil {
		return fmt.Errorf("error al escribir en el archivo, error: %w", err)
	}

	return nil
}

func (r *repository) Patch(id int, apellido string, edad int) (domain.User, error) {
	var user domain.User

	if err := r.db.Read(&users); err != nil {
		return domain.User{}, fmt.Errorf("error al leer el archivo")
	}

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
		return domain.User{}, fmt.Errorf("no se encontró el usuario con el ID %d", id)
	}

	if err := r.db.Write(users); err != nil {
		return domain.User{}, fmt.Errorf("error al escribir en el archivo, error: %w", err)
	}

	return user, nil
}
