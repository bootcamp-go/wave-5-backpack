package usuarios

import (
	"github.com/anesquivel/wave-5-backpack/goweb/arquitectura_ejercicio/internal/domain"
)

type Service interface {
	GetAll() ([]domain.Usuario, error)
	Store(age int, names, lastname, email string, estatura float64) (domain.Usuario, error)
	Update(id, age int, names, lastname, email, dateCreated string, estatura float64, activo bool) (domain.Usuario, error)
	UpdateLastNameAndAge(id, age int, lastname string) (domain.Usuario, error)
}

type service struct {
	repository Repository
}

func (s *service) GetAll() ([]domain.Usuario, error) {
	usuarios, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return usuarios, nil
}

func (s *service) Store(age int, names, lastname, email string, estatura float64) (domain.Usuario, error) {
	lastID, err := s.repository.LastID()
	if err != nil {
		return domain.Usuario{}, err
	}

	lastID++

	usuario, err := s.repository.Store(lastID, age, names, lastname, email, "7 JUL 2022", estatura)
	if err != nil {
		return domain.Usuario{}, err
	}

	return usuario, nil
}

func (s *service) Update(id, age int, names, lastname, email, dateCreated string, estatura float64, activo bool) (domain.Usuario, error) {
	return s.repository.Update(id, age, names, lastname, email, dateCreated, estatura, activo)
}

func (s *service) UpdateLastNameAndAge(id, age int, lastname string) (domain.Usuario, error) {
	return s.repository.UpdateLastNameAndAge(id, age, lastname)
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}
