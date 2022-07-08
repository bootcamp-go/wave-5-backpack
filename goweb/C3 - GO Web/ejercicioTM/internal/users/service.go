package users

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

type Service interface {
	GetAll() ([]Usuarios, error)
	Store(nombre string, apellido string, email string, edad int, altura float64, activo bool, fecha time.Time) (Usuarios, error)
	//Ejercicio 1
	//PUT de todos los campos
	Update(id int, nombre string, apellido string, email string, edad int, altura float64, activo bool, fecha time.Time) (Usuarios, error)
	//Ejercicio 2
	//DELETE de un usuario de acuerdo a su id
	Delete(id int) error
	//Ejercicio 3
	//PATCH de los campos apellido y edad
	UpdateLastAge(id int, apellido string, edad int) (Usuarios, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll() ([]Usuarios, error) {
	us, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return us, nil
}

func (s *service) Store(nombre string, apellido string, email string, edad int, altura float64, activo bool, fecha time.Time) (Usuarios, error) {
	lastID, err := s.repository.LastID()
	if err != nil {
		return Usuarios{}, err
	}

	lastID++

	usuario, err := s.repository.Store(lastID, nombre, apellido, email, edad, altura, true, time.Now())
	if err != nil {
		return Usuarios{}, err
	}

	data, err := s.repository.GetAll()
	if err != nil {
		fmt.Println(err)
	}
	jsonData, errMarshal := json.Marshal(data)
	if errMarshal != nil {
		log.Fatal(errMarshal)
	}

	os.WriteFile("../users.json", jsonData, 0644)

	return usuario, nil
}

func (s *service) Update(id int, nombre string, apellido string, email string, edad int, altura float64, activo bool, fecha time.Time) (Usuarios, error) {
	return s.repository.Update(id, nombre, apellido, email, edad, altura, activo, fecha)
}

func (s *service) UpdateLastAge(id int, apellido string, edad int) (Usuarios, error) {
	return s.repository.UpdateLastAge(id, apellido, edad)
}

func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}
