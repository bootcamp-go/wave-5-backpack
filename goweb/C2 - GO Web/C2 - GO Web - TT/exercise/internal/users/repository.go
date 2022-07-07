package users

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"exercise/internal/domain"
)

type Repository interface {
	GetAll() ([]domain.Usuarios, error)
	Store(id int, nombre string, apellido string, email string, edad int, altura float64, activo bool, fecha time.Time) (domain.Usuarios, error)
	LastID() (int, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

//Leo json de usuarios
func usersJson() []domain.Usuarios {
	//Leo el json y lo envío como retorno
	jsonUsers, err := os.ReadFile("../users.json")
	if err != nil {
		fmt.Print(err)
	}
	var users []domain.Usuarios
	err = json.Unmarshal(jsonUsers, &users)
	if err != nil {
		fmt.Print(err)
	}
	return users
}

var us = usersJson()
var lastID = us[len(us)-1].Id

func (r *repository) GetAll() ([]domain.Usuarios, error) {
	return us, nil
}

func (r *repository) LastID() (int, error) {
	return lastID, nil
}

func (r *repository) Store(id int, nombre string, apellido string, email string, edad int, altura float64, activo bool, fecha time.Time) (domain.Usuarios, error) {
	//Id Nombre Apellido Email Edad Altura Activo Fecha
	u := domain.Usuarios{Id: id, Nombre: nombre, Apellido: apellido, Email: email, Edad: edad, Altura: altura, Activo: activo, Fecha: fecha}
	us = append(us, u)
	lastID = u.Id
	return u, nil
}
