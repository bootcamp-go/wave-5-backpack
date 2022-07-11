package users

import (
	"fmt"
	"goweb/internal/domain"
	"goweb/pkg/dataStore"
)

type Repository interface {
	GetAll() ([]domain.User, error)
	NewUser(id int, name, lastname, email string, age int, height float64, active bool, creationDate string) (domain.User, error)
	LastID() (int, error)
	Update(id int, name, lastname, email string, age int, height float64, active bool, creationDate string) (domain.User, error)
	UpdateName(id int, name string) (domain.User, error)
	Delete(id int) error
}

type repository struct {
	db dataStore.DataStore
}

func NewRepository(db dataStore.DataStore) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll() ([]domain.User, error) {
	var us []domain.User
	if err := r.db.Read(&us); err != nil {
		return nil, fmt.Errorf("error al cargar la data storage")
	}
	return us, nil
}

func (r *repository) LastID() (int, error) {
	var us []domain.User
	if err := r.db.Read(&us); err != nil {
		return 0, fmt.Errorf("error al cargar la data storage")
	}
	if len(us) == 0 {
		return 0, nil
	}

	return us[len(us)-1].ID, nil
}

func (r *repository) NewUser(id int, name, lastname, email string, age int, height float64, active bool, creationDate string) (domain.User, error) {
	var us []domain.User
	if err := r.db.Read(&us); err != nil {
		return domain.User{}, fmt.Errorf("error al cargar la data storage")
	}

	u := domain.User{ID: id, Name: name, LastName: lastname, Email: email, Age: age, Height: height, Active: active, CreationDate: creationDate}
	us = append(us, u)
	if err := r.db.Write(us); err != nil {
		return domain.User{}, fmt.Errorf("error al escribir %v", err)
	}
	return u, nil
}

func (r *repository) Update(id int, name, lastname, email string, age int, height float64, active bool, creationDate string) (domain.User, error) {
	var us []domain.User
	if err := r.db.Read(&us); err != nil {
		return domain.User{}, fmt.Errorf("error al cargar la data storage")
	}
	u := domain.User{ID: id, Name: name, LastName: lastname, Email: email, Age: age, Height: height, Active: active, CreationDate: creationDate}
	updated := false
	for i := range us {
		if us[i].ID == id {
			u.ID = id
			us[i] = u
			updated = true
		}
	}

	if !updated {
		return domain.User{}, fmt.Errorf("usuario %d no encontrado", id)
	}

	if err := r.db.Write(us); err != nil {
		return domain.User{}, fmt.Errorf("error al escribir %v", err)
	}
	return u, nil
}

func (r *repository) UpdateName(id int, name string) (domain.User, error) {
	var us []domain.User
	if err := r.db.Read(&us); err != nil {
		return domain.User{}, fmt.Errorf("error al cargar la data storage")
	}
	updated := false
	var u domain.User
	for i := range us {
		if us[i].ID == id {
			us[i].Name = name
			u = us[i]
			updated = true
		}
	}

	if !updated {
		return domain.User{}, fmt.Errorf("usuario %d no encontrado", id)
	}

	if err := r.db.Write(us); err != nil {
		return domain.User{}, fmt.Errorf("error al escribir %v", err)
	}
	return u, nil
}

func (r *repository) Delete(id int) error {
	var us []domain.User
	if err := r.db.Read(&us); err != nil {
		return fmt.Errorf("error al cargar la data storage")
	}
	deleted := false
	var index int
	for i := range us {
		if us[i].ID == id {
			index = i
			deleted = true
		}
	}

	if !deleted {
		return fmt.Errorf("usuario %d no encontrado", id)
	}

	us = append(us[:index], us[index+1:]...)
	if err := r.db.Write(us); err != nil {
		return fmt.Errorf("error al escribir %v", err)
	}
	return nil
}
