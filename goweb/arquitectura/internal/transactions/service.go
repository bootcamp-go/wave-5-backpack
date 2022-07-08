package transactions

import (
	"arquitectura/internal/domain"
)

type Service interface {
	GetAll() ([]domain.Transaction, error)
	Store(tranCode, currency string, amount float64, transmitter, receiver, tranDate string) (domain.Transaction, error)
	Update(id int, tranCode, currency string, amount float64, transmitter, receiver, tranDate string) (domain.Transaction, error)
	UpdateFields(id int, tranCode string, amount float64) (domain.Transaction, error)
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

func (s *service) UpdateFields(id int, tranCode string, amount float64) (domain.Transaction, error) {
	var tran domain.Transaction
	if tranCode != "" {
		t, err := s.repository.UpdateTranCode(id, tranCode)
		if err != nil {
			return t, err
		}
		tran = t
	}
	if amount != 0 {
		t, err := s.repository.UpdateAmount(id, amount)
		if err != nil {
			return t, nil
		}
		tran = t
	}

	return tran, nil
}

func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}
