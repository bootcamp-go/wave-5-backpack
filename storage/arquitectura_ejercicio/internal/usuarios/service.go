package usuarios

import (
	"context"
	"time"

	"github.com/anesquivel/wave-5-backpack/storage/arquitectura_ejercicio/internal/domain"
)

type Service interface {
	GetAll() ([]domain.Usuario, error)
	Store(age int, names, lastname, email string, estatura float64) (domain.Usuario, error)
	Update(id, age int, names, lastname, email, dateCreated string, estatura float64, activo bool) (domain.Usuario, error)
	UpdateLastNameAndAge(ctx context.Context, id, age int, lastname string) (domain.Usuario, error)
	Delete(id int) error
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
	newUser := domain.Usuario{
		Names:       names,
		LastName:    lastname,
		Email:       email,
		Age:         age,
		Estatura:    estatura,
		IsActivo:    true,
		DateCreated: time.Now().GoString(),
	}
	usuario, err := s.repository.Store(newUser)
	if err != nil {
		return domain.Usuario{}, err
	}

	return usuario, nil
}

func (s *service) Update(id, age int, names, lastname, email, dateCreated string, estatura float64, activo bool) (domain.Usuario, error) {
	user := domain.Usuario{
		Names:       names,
		LastName:    lastname,
		Email:       email,
		Age:         age,
		Estatura:    estatura,
		DateCreated: dateCreated,
		IsActivo:    activo,
	}
	return s.repository.Update(id, user)
}

func (s *service) UpdateLastNameAndAge(ctx context.Context, id, age int, lastname string) (domain.Usuario, error) {
	return s.repository.UpdateLastNameAndAge(ctx, id, age, lastname)
}
func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}
func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}
