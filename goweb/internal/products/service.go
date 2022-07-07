package products

import (
	"goweb/internal/domain"
	"time"
)

type Service interface {
	GetAll() ([]domain.Usuarios, error)
	Store(id int, nombre string, apellido string, email string, edad int, altura float64, activo bool, fecha time.Time) (domain.Usuarios, error)
}
type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll() ([]domain.Usuarios, error) {
	us, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return us, nil
}

func (s *service) Store(id int, nombre string, apellido string, email string, edad int, altura float64, activo bool, fecha time.Time) (domain.Usuarios, error) {
	lastID, err := s.repository.LastID()
	if err != nil {
		return domain.Usuarios{}, err
	}

	lastID++

	usuario, err := s.repository.Store(id, nombre, apellido, email, edad, altura, activo, fecha)
	if err != nil {
		return domain.Usuarios{}, err
	}

	return usuario, nil
}
