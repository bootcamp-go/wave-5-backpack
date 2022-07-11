package users

import (
	"CLASE3/internal/domain"
	"fmt"
)

type Service interface {
	GetAll() ([]domain.Users, error)
	Store(nombre, apellido string, edad int, altura float64) (domain.Users, error)
	Update(id int, nombre, apeliido string, edad int, altura float64) (domain.Users, error)
	Delete(id int) error
	UpdateApellidoAndEdad(id int, apellido string, edad int) (domain.Users, error)
}
type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{repository: r}
}

func (s *service) GetAll() ([]domain.Users, error) {
	us, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return us, nil
}

func (s *service) Store(nombre, apellido string, edad int, altura float64) (domain.Users, error) {
	lastID, err := s.repository.LastID()
	if err != nil {
		return domain.Users{}, fmt.Errorf("Error extrayendo el usuario con el ultimo id %w", err)
	}
	lastID++
	users, err := s.repository.Store(lastID, nombre, apellido, edad, altura)
	if err != nil {
		return domain.Users{}, fmt.Errorf("Error creando usuario: %w", err)
	}
	return users, nil
}

func (s *service) Update(id int, nombre, apellido string, edad int, altura float64) (domain.Users, error) {
	p, err := s.repository.Update(id, nombre, apellido, edad, altura)
	if err != nil {
		return domain.Users{}, fmt.Errorf("Error actualizando usuario %w", err)
	}
	return p, nil
}

func (s *service) Delete(id int) error {
	if err := s.repository.Delete(id); err != nil {
		return fmt.Errorf("Error eliminando el usuario %w", err)
	}
	return nil

}

func (s *service) UpdateApellidoAndEdad(id int, apellido string, edad int) (domain.Users, error) {
	product, err := s.repository.UpdateApellidoAndEdad(id, apellido, edad)
	if err != nil {
		return domain.Users{}, fmt.Errorf("Error actualizando el usuario %w", err)
	}
	return product, nil

}
