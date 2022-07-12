package products

import (
	"fmt"
	"goweb/go-web-II/internal/domain"
	"goweb/go-web-II/pkg/store"
)

/*
En repository va a ir todo lo que sea manejo de BD (Consultas)
y variables globales.
*/

type Repository interface {
	GetAll() ([]*domain.User, error)
	LastId() (int, error)
	Store(id, age int, name, surname, email, created string, active bool) (*domain.User, error)
	Update(id, age int, name, surname, email, created string, active bool) (*domain.User, error)
	Delete(id int) error
}

type repository struct {
	db store.Store
}

var users []*domain.User
var lastId int

func NewRepository(db store.Store) Repository {
	return &repository{db: db}
}

func (r *repository) Delete(id int) error {
	deleted := false
	var index int
	for i := range users {
		if users[i].Id == id {
			index = i
			deleted = true
		}
	}
	if !deleted {
		return fmt.Errorf("producto %d no encotrado", id)
	}
	users = append(users[:index], users[index+1:]...) /*agarra users desde el principio hasta el indice del user encontrado y lo "recorta" y lo sigue desde el indice + 1*/
	return nil
}

func (r *repository) Update(id, age int, name, surname, email, created string, active bool) (*domain.User, error) {
	u := &domain.User{Name: name, Surname: surname, Email: email, Created: created, Age: age, Active: active}
	updated := false
	var us []*domain.User
	r.db.Read(&us)
	for i := range users {
		if users[i].Id == id {
			u.Id = id
			users[i] = u
			updated = true
		}
	}

	if !updated {
		return &domain.User{}, fmt.Errorf("producto %d no encontrado", id)
	}

	if err := r.db.Write(&us); err != nil {
		return nil, err
	}

	return u, nil
}

func (r *repository) Store(id, age int, name, surname, email, created string, active bool) (*domain.User, error) {
	var us []domain.User
	r.db.Read(&us)

	user := domain.User{
		Id:      id,
		Name:    name,
		Age:     age,
		Surname: surname,
		Email:   email,
		Created: created,
		Active:  active,
	}
	us = append(us, user)
	if err := r.db.Write(&us); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *repository) LastId() (int, error) {
	return lastId, nil
}

func (r *repository) GetAll() ([]*domain.User, error) {
	var us []*domain.User
	if err := r.db.Read(&us); err != nil {
		return nil, err
	}
	return us, nil
}
