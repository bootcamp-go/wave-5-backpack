package users

import (
	"context"
	"time"

	"github.com/google/uuid"
)

//Implementación de métodos de repository
type Service interface {
	GetOne(id string) (User, error)
	Store(nombre string, apellido string, email string, edad int, altura float64) (User, error)
	Delete(id string) error
	Update(id string, nombre string, apellido string, email string, edad int, altura float64, activo bool) (User, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetOne(id string) (User, error) {
	user, err := s.repository.GetOne(context.Background(), id)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (s *service) Store(nombre string, apellido string, email string, edad int, altura float64) (User, error) {
	//Asignación de identificador
	id := uuid.New().String()
	user := User{
		Id:       id,
		Nombre:   nombre,
		Apellido: apellido,
		Email:    email,
		Edad:     edad,
		Altura:   altura,
		Activo:   true,
		Fecha:    time.Now(),
	}
	err := s.repository.Store(context.Background(), &user)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (s *service) Delete(id string) error {
	return s.repository.Delete(context.Background(), id)
}

func (s *service) Update(id string, nombre string, apellido string, email string, edad int, altura float64, activo bool) (User, error) {
	return s.repository.Update(id, nombre, apellido, email, edad, altura, activo)
}
