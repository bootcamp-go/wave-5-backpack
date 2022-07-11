package users

import (
	"fmt"

	"github.com/bootcamp-go/wave-5-backpack/tree/gonzalez_jose/goweb/internal/domain"
	"github.com/bootcamp-go/wave-5-backpack/tree/gonzalez_jose/goweb/pkg/store"
)

type Repository interface {
	GetAll() ([]domain.User, error)
	Store(edad int, nombre, apellido, email, fecha_creacion string, altura float64, activo bool) (domain.User, error)
	Update(id, edad int, nombre, apellido, email, fecha_creacion string, altura float64, activo bool) (domain.User, error)
	Delete(id int) error
	UpdateLastNameAndAge(id, edad int, apellido string) (domain.User, error)
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll() ([]domain.User, error) {
	var users []domain.User
	r.db.Read(&users)
	return users, nil
}

func (r *repository) Store(edad int, nombre, apellido, email, fecha_creacion string, altura float64, activo bool) (domain.User, error) {
	var users []domain.User
	r.db.Read(&users)
	u := domain.User{Id: users[len(users)-1].Id + 1, Edad: edad, Nombre: nombre, Apellido: apellido, Email: email, Fecha_creacion: fecha_creacion, Altura: altura, Activo: activo}
	users = append(users, u)
	if err := r.db.Write(users); err != nil {
		return domain.User{}, nil
	}
	return u, nil
}

func (r *repository) Update(id, edad int, nombre, apellido, email, fecha_creacion string, altura float64, activo bool) (domain.User, error) {
	u := domain.User{Id: id, Edad: edad, Nombre: nombre, Apellido: apellido, Email: email, Fecha_creacion: fecha_creacion, Altura: altura, Activo: activo}

	var users []domain.User
	r.db.Read(&users)

	update := false

	for i := range users {
		if users[i].Id == id {
			u.Id = id
			users[i] = u
			update = true
		}
	}
	if err := r.db.Write(users); err != nil {
		return domain.User{}, nil
	}

	if !update {
		return domain.User{}, fmt.Errorf("Usuario %v no encontrado", id)
	}
	return u, nil
}

func (r *repository) Delete(id int) error {
	delete := false
	var index int
	var users []domain.User
	r.db.Read(&users)
	for i := range users {
		if users[i].Id == id {
			index = i
			delete = true
		}
	}
	if !delete {
		return fmt.Errorf("Usuario %v no encontrado", id)
	}
	users = append(users[:index], users[index+1:]...)
	if err := r.db.Write(users); err != nil {
		return nil
	}
	return nil
}

func (r *repository) UpdateLastNameAndAge(id, edad int, apellido string) (domain.User, error) {
	update := false
	var u domain.User
	var users []domain.User
	r.db.Read(&users)
	for i := range users {
		if users[i].Id == id {
			users[i].Apellido = apellido
			users[i].Edad = edad
			update = true
			u = users[i]
		}
	}
	if !update {
		return domain.User{}, fmt.Errorf("Usuario %v no encontrado", id)
	}
	if err := r.db.Write(users); err != nil {
		return domain.User{}, nil
	}
	return u, nil
}
