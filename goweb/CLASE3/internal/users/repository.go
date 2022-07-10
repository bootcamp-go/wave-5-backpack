package users

import (
	"clase2_parte2/internal/domain"
	"clase2_parte2/pkg/store"
	"fmt"
)

type Repository interface {
	GetAll() ([]domain.Users, error)
	Store(id int, nombre, apellido string, edad int, altura float64) (domain.Users, error)
	LastID() (int, error)
	Update(id int, nombre, apellido string, edad int, altura float64) (domain.Users, error)
	Delete(id int) error
	UpdateApellidoAndEdad(id int, apellido string, edad int) (domain.Users, error)
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll() ([]domain.Users, error) {
	var us []domain.Users
	if err := r.db.Read(&us); err != nil {
		return nil, fmt.Errorf("Error leyendo el archivo")
	}
	return us, nil
}

func (r *repository) LastID() (int, error) {
	var us []domain.Users
	if err := r.db.Read(&us); err != nil {
		return 0, fmt.Errorf("Error leyendo el archivo")
	}
	if len(us) == 0 {
		return 0, nil
	}

	return us[len(us)-1].Id, nil
}

func (r *repository) Store(id int, nombre, apellido string, edad int, altura float64) (domain.Users, error) {
	var us []domain.Users

	if err := r.db.Read(&us); err != nil {
		return domain.Users{}, fmt.Errorf("Error leyendo el archivo")
	}

	p := domain.Users{Id: id, Nombre: nombre, Apellido: apellido, Edad: edad, Altura: altura}
	us = append(us, p)

	if err := r.db.Write(us); err != nil {
		return domain.Users{}, fmt.Errorf("Error leyendo el archivo", err)
	}

	return p, nil
}

func (r *repository) Update(id int, nombre, apellido string, edad int, altura float64) (domain.Users, error) {
	var us []domain.Users

	if err := r.db.Read(&us); err != nil {
		return domain.Users{}, fmt.Errorf("Error leyendo el archivo")
	}

	p := domain.Users{Id: id, Nombre: nombre, Apellido: apellido, Edad: edad, Altura: altura}
	us = append(us, p)

	if err := r.db.Write(us); err != nil {
		return domain.Users{}, fmt.Errorf("Error leyendo el archivo", err)
	}

	return p, nil
}

func (r *repository) Delete(id int) error {
	var us []domain.Users

	if err := r.db.Read(&us); err != nil {
		return fmt.Errorf("Error leyendo el archivo")
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
		return fmt.Errorf("Usuario no encontrado", id)
	}

	us = append(us[:index], us[index+1:]...)

	if err := r.db.Write(us); err != nil {
		return fmt.Errorf("Error agregando usuario", err)
	}
	return nil
}

func (r *repository) UpdateApellidoAndEdad(id int, apellido string, edad int) (domain.Users, error) {
	var us []domain.Users

	if err := r.db.Read(&us); err != nil {
		return domain.Users{}, fmt.Errorf("Error de lectura")
	}
	update := false
	var u domain.Users

	for i := range us {
		if us[i].Id == id {
			us[i].Apellido = apellido
			us[i].Edad = edad
			u = us[i]
			update = true
		}
	}

	if !update {
		return domain.Users{}, fmt.Errorf("Usuario %d no fue encontrado", id)
	}

	if err := r.db.Write(us); err != nil {
		return domain.Users{}, fmt.Errorf("Error en la escritura a la base de datos", err)
	}

	return u, nil
}
