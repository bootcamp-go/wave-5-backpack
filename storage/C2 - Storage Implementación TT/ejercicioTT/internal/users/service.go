package users

import (
	"context"
	"ejercicioTT/internal/domain"
	"fmt"
)

type Service interface {
	GetFullData(id int) ([]domain.UserAndWarehouse, error)
	GetByName(nombre string) (domain.Usuarios, error)
	Store(domain.Usuarios) (domain.Usuarios, error)
	Update(ctx context.Context, usuario domain.Usuarios) (domain.Usuarios, error)
	GetOne(id int) (domain.Usuarios, error)
	GetAll() ([]domain.UserAndWarehouse, error)
	Delete(id int) error
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

func (s *service) GetOne(id int) (domain.Usuarios, error) {
	usuario, err := s.repo.GetOne(id)
	if err != nil {
		return domain.Usuarios{}, err
	}
	return usuario, nil
}

func (s *service) Update(ctx context.Context, usuario domain.Usuarios) (domain.Usuarios, error) {
	usuario, err := s.repo.Update(ctx, usuario)
	if err != nil {
		return domain.Usuarios{}, fmt.Errorf("error actualizando usuario %w", err)
	}
	return usuario, nil
}

func (s *service) GetAll() ([]domain.UserAndWarehouse, error) {
	usuarioWarehouse, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return usuarioWarehouse, nil
}

func (s *service) Delete(id int) error {
	if err := s.repo.Delete(id); err != nil {
		return fmt.Errorf("error al intentar eliminar usuario %w", err)
	}
	return nil
}

func (s *service) GetFullData(id int) ([]domain.UserAndWarehouse, error) {
	usuarioWarehouse, err := s.repo.GetFullData(id)
	if err != nil {
		return nil, err
	}
	return usuarioWarehouse, nil
}
