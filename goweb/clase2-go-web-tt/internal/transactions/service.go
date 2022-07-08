package transactions

import (
	"goweb/clase2-go-web-tt/internal/domain"
)

type Service interface {
	GetAll() ([]domain.Transaction, error)
	Ecommerce(codeTra string, coin string, monto float64,
		emisor string, receptor string, fecha string) (domain.Transaction, error)
	GetOne(id int) (domain.Transaction, error)
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
	ts, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return ts, nil
}

func (s *service) Ecommerce(codeTra string, coin string, monto float64,
	emisor string, receptor string, fecha string) (domain.Transaction, error) {
	lastID, err := s.repository.LastID()
	if err != nil {
		return domain.Transaction{}, err
	}

	lastID++

	transaccion, err := s.repository.Ecommerce(lastID, codeTra, coin, monto, emisor, receptor, fecha)
	if err != nil {
		return domain.Transaction{}, err
	}

	return transaccion, nil
}

func (s *service) GetOne(id int) (domain.Transaction, error) {
	ts, err := s.repository.GetOne(id)
	if err != nil {
		return domain.Transaction{}, err
	}

	return ts, nil
}
