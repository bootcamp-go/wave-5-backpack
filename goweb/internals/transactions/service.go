package transactions

import (
	"fmt"
	"goweb/internals/domain"
)

type Service interface {
	GetAll() ([]domain.Transaction, error)
	Store(codigo string, moneda string, monto int, emisor string, receptor string) (domain.Transaction, error)
	Update(id int, codigo string, moneda string, monto int, emisor string, receptor string) (domain.Transaction, error)
	Delete(id int) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll() ([]domain.Transaction, error) {
	trans, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return trans, nil
}

func (s *service) Store(codigo string, moneda string, monto int, emisor string, receptor string) (domain.Transaction, error) {
	lastID, err := s.repository.LastID()
	if err != nil {
		return domain.Transaction{}, fmt.Errorf("Hubo un error al obtener el último ID: %w", err)
	}
	lastID++

	transaction, err := s.repository.Store(lastID, codigo, moneda, monto, emisor, receptor)
	if err != nil {
		return domain.Transaction{}, fmt.Errorf("Error al crear el producto: %w", err)
	}
	return transaction, nil
}

func (s *service) Update(id int, codigo string, moneda string, monto int, emisor string, receptor string) (domain.Transaction, error) {
	return s.repository.Update(id, codigo, moneda, monto, emisor, receptor)
}

func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}
