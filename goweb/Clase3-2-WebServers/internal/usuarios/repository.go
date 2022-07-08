package usuarios

import (
	"fmt"
	"goweb/Clase3-2-WebServers/internal/domain"
	"goweb/Clase3-2-WebServers/pkg/store"
	"time"
)

const (
	UserNotFound = "user %d not found"
	FailReading  = "cant read database"
	FailWriting  = "cant write database, error: %w"
)

type Repository interface {
	GetAll() ([]domain.Usuario, error)
	Store(id int, nombre, apellido, email string, edad, altura int, activo bool) (domain.Usuario, error)
	LastID() (int, error)
	Update(id int, nombre, apellido, email string, edad, altura int, activo bool) (domain.Usuario, error)
	UpdateSurnameAndAge(id int, surname string, age int) (domain.Usuario, error)
	Delete(id int) error
}

type repository struct {
	db store.Store
}

func NewRepository(database store.Store) Repository {
	return &repository{
		db: database,
	}
}

func (r *repository) GetAll() ([]domain.Usuario, error) {
	var us []domain.Usuario
	if err := r.db.Read(&us); err != nil {
		return nil, fmt.Errorf(FailReading)
	}
	return us, nil
}

func (r *repository) LastID() (int, error) {
	var us []domain.Usuario
	if err := r.db.Read(&us); err != nil {
		return 0, fmt.Errorf(FailReading)
	}
	if len(us) == 0 {
		return 0, nil
	}

	return us[len(us)-1].Id, nil
}

func (r *repository) Store(id int, nombre, apellido, email string, edad, altura int, activo bool) (domain.Usuario, error) {
	var us []domain.Usuario
	if err := r.db.Read(&us); err != nil {
		return domain.Usuario{}, fmt.Errorf(FailReading)
	}
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
	if err := r.db.Write(us); err != nil {
		return domain.Usuario{}, fmt.Errorf(FailWriting, err)
	}
	return u, nil
}
func (r *repository) Update(id int, nombre, apellido, email string, edad, altura int, activo bool) (domain.Usuario, error) {
	var us []domain.Usuario
	if err := r.db.Read(&us); err != nil {
		return domain.Usuario{}, fmt.Errorf(FailReading)
	}
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
		return domain.Usuario{}, fmt.Errorf(UserNotFound, id)
	}
	if err := r.db.Write(us); err != nil {
		return domain.Usuario{}, fmt.Errorf(FailWriting, err)
	}
	return u, nil
}

func (r *repository) UpdateSurnameAndAge(id int, surname string, age int) (domain.Usuario, error) {
	var us []domain.Usuario
	if err := r.db.Read(&us); err != nil {
		return domain.Usuario{}, fmt.Errorf(FailReading)
	}
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
		return domain.Usuario{}, fmt.Errorf(UserNotFound, id)
	}

	if err := r.db.Write(us); err != nil {
		return domain.Usuario{}, fmt.Errorf(FailWriting, err)
	}

	return u, nil
}

func (r *repository) Delete(id int) error {
	var us []domain.Usuario
	if err := r.db.Read(&us); err != nil {
		return fmt.Errorf(FailReading)
	}
	deleted := false
	var index int
	for i := range us {
		if us[i].Id == id {
			index = i
			deleted = true
		}
	}

	if !deleted {
		return fmt.Errorf(UserNotFound, id)
	}

	us = append(us[:index], us[index+1:]...)

	if err := r.db.Write(us); err != nil {
		return fmt.Errorf(FailWriting, err)
	}

	return nil
}
