package transactions

import (
	"arquitectura/internal/domain"
)

type Service interface {
	GetAll() ([]domain.Transaction, error)
	Store(tranCode, currency string, amount float64, transmitter, receiver, tranDate string) (domain.Transaction, error)
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
	lista, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return lista, nil
}

func (s *service) Store(tranCode, currency string, amount float64, transmitter, receiver, tranDate string) (domain.Transaction, error) {
	lastId, err := s.repository.LastID()
	if err != nil {
		return domain.Transaction{}, nil
	}

	lastId++

	transaction, err := s.repository.Store(lastId, tranCode, currency, amount, transmitter, receiver, tranDate)
	if err != nil {
		return domain.Transaction{}, nil
	}

	return transaction, nil
}
