package products

import (
	"C3ejercicioTT/internal/domain"
	"C3ejercicioTT/pkg/store"
	"fmt"
	"time"
)

const (
	UserNotFound = "usuario %d no encontrado"
	FailReading  = "No se puede leer la base de datos"
	FailWriting  = "No puede escribir la base de datos, error: %w"
)

type Repository interface {
	GetAll() ([]domain.Usuarios, error)
	Store(id int, nombre string, apellido string, email string, edad int, altura float64, activo bool, fecha time.Time) (domain.Usuarios, error)
	LastID() (int, error)
	//Ejercicio 1
	//PUT de todos los campos
	Update(id int, nombre, apellido, email string, edad int, altura float64, activo bool, fecha time.Time) (domain.Usuarios, error)
	//Ejercicio 2
	//DELETE de un usuario de acuerdo a su id
	Delete(id int) error
	//Ejercicio 3
	//PATCH de los campos apellido y edad
	UpdateLastAge(id int, apellido string, edad int) (domain.Usuarios, error)
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll() ([]domain.Usuarios, error) {
	var us []domain.Usuarios
	if err := r.db.Read(&us); err != nil {
		return nil, fmt.Errorf(FailReading)
	}
	return us, nil
}

func (r *repository) Store(id int, nombre, apellido, email string, edad int, altura float64, activo bool, fecha time.Time) (domain.Usuarios, error) {
	var us []domain.Usuarios

	if err := r.db.Read(&us); err != nil {
		return domain.Usuarios{}, fmt.Errorf(FailReading)
	}

	u := domain.Usuarios{Id: id, Nombre: nombre, Apellido: apellido, Email: email, Edad: edad, Altura: altura, Activo: activo, Fecha: fecha}
	us = append(us, u)

	if err := r.db.Write(us); err != nil {
		return domain.Usuarios{}, fmt.Errorf(FailWriting, err)
	}

	return u, nil
}

func (r *repository) Update(id int, nombre, apellido, email string, edad int, altura float64, activo bool, fecha time.Time) (domain.Usuarios, error) {
	var us []domain.Usuarios

	if err := r.db.Read(&us); err != nil {
		return domain.Usuarios{}, fmt.Errorf(FailReading)
	}

	u := domain.Usuarios{Id: id, Nombre: nombre, Apellido: apellido, Email: email, Edad: edad, Altura: altura, Activo: true, Fecha: time.Now()}
	updated := false
	for i := range us {
		if us[i].Id == id {
			u.Id = id
			us[i] = u
			updated = true
		}
	}

	if !updated {
		return domain.Usuarios{}, fmt.Errorf(UserNotFound, id)
	}

	if err := r.db.Write(us); err != nil {
		return domain.Usuarios{}, fmt.Errorf(FailWriting, err)
	}

	return u, nil
}

func (r *repository) UpdateLastAge(id int, apellido string, edad int) (domain.Usuarios, error) {
	var us []domain.Usuarios

	if err := r.db.Read(&us); err != nil {
		return domain.Usuarios{}, fmt.Errorf(FailReading)
	}

	updated := false
	var u domain.Usuarios
	for i := range us {
		if us[i].Id == id {
			us[i].Apellido = apellido
			us[i].Edad = edad
			u = us[i]
			updated = true
		}
	}

	if !updated {
		return domain.Usuarios{}, fmt.Errorf(UserNotFound, id)
	}

	if err := r.db.Write(us); err != nil {
		return domain.Usuarios{}, fmt.Errorf(FailWriting, err)
	}

	return u, nil
}

func (r *repository) Delete(id int) error {
	var us []domain.Usuarios

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

	/* if err := r.db.Write(us); err != nil {
		return fmt.Errorf(FailWriting, err)
	} */

	if err := r.db.Write(&us); err != nil {
		return fmt.Errorf(FailWriting, err)
	}
	return nil
}

func (r *repository) LastID() (int, error) {
	var us []domain.Usuarios
	if err := r.db.Read(&us); err != nil {
		return 0, fmt.Errorf(FailReading)
	}
	if len(us) == 0 {
		return 0, nil
	}

	return us[len(us)-1].Id, nil
}
