package products

import (
	"ejercicioTT/internal/domain"
	"fmt"
	"time"
)

type Service interface {
	GetAll() ([]domain.Usuarios, error)
	Store(nombre string, apellido string, email string, edad int, altura float64, activo bool, fecha time.Time) (domain.Usuarios, error)
	//PUT de todos los campos
	Update(id int, nombre string, apellido string, email string, edad int, altura float64, activo bool, fecha time.Time) (domain.Usuarios, error)
	//DELETE de un usuario de acuerdo a su id
	Delete(id int) error
	//PATCH de los campos apellido y edad
	UpdateLastAge(id int, apellido string, edad int) (domain.Usuarios, error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) GetAll() ([]domain.Usuarios, error) {
	us, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return us, nil
}

func (s *service) Store(nombre string, apellido string, email string, edad int, altura float64, activo bool, fecha time.Time) (domain.Usuarios, error) {
	lastID, err := s.repo.LastID()
	if err != nil {
		return domain.Usuarios{}, fmt.Errorf("error obteniendo el id del último usuario: %w", err)
	}

	lastID++

	usuario, err := s.repo.Store(lastID, nombre, apellido, email, edad, altura, true, time.Now())
	if err != nil {
		return domain.Usuarios{}, fmt.Errorf("error creando usuario: %w", err)
	}
	return usuario, nil
}

func (s *service) Update(id int, nombre string, apellido string, email string, edad int, altura float64, activo bool, fecha time.Time) (domain.Usuarios, error) {
	u, err := s.repo.Update(id, nombre, apellido, email, edad, altura, activo, fecha)
	if err != nil {
		return domain.Usuarios{}, fmt.Errorf("error actualizando usuario %w", err)
	}
	return u, nil
}

func (s *service) UpdateLastAge(id int, apellido string, edad int) (domain.Usuarios, error) {
	usuario, err := s.repo.UpdateLastAge(id, apellido, edad)
	if err != nil {
		return domain.Usuarios{}, fmt.Errorf("error actualizando usuario %w", err)
	}
	return usuario, nil
}

func (s *service) Delete(id int) error {
	if err := s.repo.Delete(id); err != nil {
		return fmt.Errorf("error al intentar eliminar usuario %w", err)
	}
	return nil
}
