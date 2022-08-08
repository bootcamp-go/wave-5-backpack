package products

import (
	"ejercicioTT/internal/domain"
	"fmt"
)

type Service interface {
	GetByName(nombre string) (domain.Usuarios, error)
	Store(domain.Usuarios) (domain.Usuarios, error)
	Update(domain.Usuarios) (domain.Usuarios, error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) GetByName(nombre string) (domain.Usuarios, error) {
	usuario, err := s.repo.GetByName(nombre)
	if err != nil {
		return domain.Usuarios{}, err
	}
	return usuario, nil
}

func (s *service) Store(usuario domain.Usuarios) (domain.Usuarios, error) {
	usuario, err := s.repo.Store(usuario)
	if err != nil {
		return domain.Usuarios{}, fmt.Errorf("error creando usuario: %w", err)
	}
	return usuario, nil
}

func (s *service) Update(usuario domain.Usuarios) (domain.Usuarios, error) {
	usuario, err := s.repo.Update(usuario)
	if err != nil {
		return domain.Usuarios{}, fmt.Errorf("error actualizando usuario %w", err)
	}
	return usuario, nil
}
