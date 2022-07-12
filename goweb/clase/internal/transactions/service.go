package transactions

import (
	"goweb/internal/domain"
)

type Service interface {
	GetAll() ([]domain.Transaction, error)
	GetOne(id int) (domain.Transaction, error)
	Create(code, currency string, amount float64, issuer, recipient, date string) (domain.Transaction, error)
	Update(id int, code, currency string, amount float64, issuer, recipient, date string) (domain.Transaction, error)
	Delete(id int) error
	Update2(id int, code string, amount float64) (domain.Transaction, error)
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
	transactions, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (s *service) GetOne(id int) (domain.Transaction, error) {
	return s.repository.GetOne(id)
}

func (s *service) Create(code, currency string, amount float64, issuer, recipient, date string) (domain.Transaction, error) {
	lastID, err := s.repository.LastID()
	if err != nil {
		return domain.Transaction{}, err
	}
	lastID++

	transaction, err := s.repository.Create(lastID, code, currency, amount, issuer, recipient, date)
	if err != nil {
		return domain.Transaction{}, err
	}
	return transaction, nil
}

func (s *service) Update(id int, code, currency string, amount float64, issuer, recipient, date string) (domain.Transaction, error) {
	transaction, err := s.repository.Update(id, code, currency, amount, issuer, recipient, date)
	if err != nil {
		return domain.Transaction{}, err
	}
	return transaction, nil
}

func (s *service) Delete(id int) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) Update2(id int, code string, amount float64) (domain.Transaction, error) {
	return s.repository.Update2(id, code, amount)
}
