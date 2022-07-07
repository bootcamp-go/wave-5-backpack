package transactions

import "goweb/internal/domain"

type Service interface {
	GetAll() ([]domain.Transaction, error)
	Create(code, currency string, amount float64, issuer, recipient, date string) (domain.Transaction, error)
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
