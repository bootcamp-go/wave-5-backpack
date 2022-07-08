package users

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Usuarios struct {
	Id       int       `json:"id"`
	Nombre   string    `json:"nombre"`
	Apellido string    `json:"apellido"`
	Email    string    `json:"email"`
	Edad     int       `json:"edad"`
	Altura   float64   `json:"altura"`
	Activo   bool      `json:"activo"`
	Fecha    time.Time `json:"fecha"`
}

type Repository interface {
	GetAll() ([]Usuarios, error)
	Store(id int, nombre string, apellido string, email string, edad int, altura float64, activo bool, fecha time.Time) (Usuarios, error)
	LastID() (int, error)
	//Ejercicio 1
	//PUT de todos los campos
	Update(id int, nombre, apellido, email string, edad int, altura float64, activo bool, fecha time.Time) (Usuarios, error)
	//Ejercicio 2
	//DELETE de un usuario de acuerdo a su id
	Delete(id int) error
	//Ejercicio 3
	//PATCH de los campos apellido y edad
	UpdateLastAge(id int, apellido string, edad int) (Usuarios, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

//Leo json de usuarios
func usersJson() []Usuarios {
	//Leo el json y lo envío como retorno
	jsonUsers, err := os.ReadFile("../../users.json")
	if err != nil {
		fmt.Print(err)
	}
	var users []Usuarios
	err = json.Unmarshal(jsonUsers, &users)
	if err != nil {
		fmt.Print(err)
	}
	return users
}

var us = usersJson()
var lastID = us[len(us)-1].Id

func (r *repository) LastID() (int, error) {
	return lastID, nil
}

func (r *repository) GetAll() ([]Usuarios, error) {
	return us, nil
}

func (r *repository) Store(id int, nombre, apellido, email string, edad int, altura float64, activo bool, fecha time.Time) (Usuarios, error) {
	//Id Nombre Apellido Email Edad Altura Activo Fecha
	u := Usuarios{Id: id, Nombre: nombre, Apellido: apellido, Email: email, Edad: edad, Altura: altura, Activo: activo, Fecha: fecha}
	us = append(us, u)
	lastID = u.Id
	return u, nil
}

func (r *repository) Update(id int, nombre, apellido, email string, edad int, altura float64, activo bool, fecha time.Time) (Usuarios, error) {
	u := Usuarios{Id: id, Nombre: nombre, Apellido: apellido, Email: email, Edad: edad, Altura: altura, Activo: activo, Fecha: fecha}
	updated := false
	for i := range us {
		if us[i].Id == id {
			u.Id = id
			us[i] = u
			updated = true
		}
	}

	if !updated {
		return Usuarios{}, fmt.Errorf("usuario %d no encontrado", id)
	}

	return u, nil
}

func (r *repository) UpdateLastAge(id int, apellido string, edad int) (Usuarios, error) {
	updated := false
	var u Usuarios
	for i := range us {
		if us[i].Id == id {
			us[i].Apellido = apellido
			us[i].Edad = edad
			u = us[i]
			updated = true
		}
	}

	if !updated {
		return Usuarios{}, fmt.Errorf("usuario %d no encontrado", id)
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
		return fmt.Errorf("usuario %d no encontrado", id)
	}

	us = append(us[:index], us[index+1:]...)
	return nil
}
