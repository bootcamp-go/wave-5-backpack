package transactions

import (
	"arquitectura/internal/domain"
)

type Service interface {
	GetAll() ([]domain.Transaction, error)
	Store(tranCode, currency string, amount float64, transmitter, receiver, tranDate string) (domain.Transaction, error)
	Update(id int, tranCode, currency string, amount float64, transmitter, receiver, tranDate string) (domain.Transaction, error)
	UpdateTranCode(id int, tranCode string) (domain.Transaction, error)
	UpdateAmount(id int, amount float64) (domain.Transaction, error)
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

func (s *service) Update(id int, tranCode, currency string, amount float64, transmitter, receiver, tranDate string) (domain.Transaction, error) {
	return s.repository.Update(id, tranCode, currency, amount, transmitter, receiver, tranDate)
}

func (s *service) UpdateTranCode(id int, tranCode string) (domain.Transaction, error) {
	return s.repository.UpdateTranCode(id, tranCode)
}

func (s *service) UpdateAmount(id int, amount float64) (domain.Transaction, error) {
	return s.repository.UpdateAmount(id, amount)
}

func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}
