package transactions

import (
	"goweb/clase3-go-web-tm/internal/domain"
)

type Service interface {
	GetAll() ([]domain.Transaction, error)
	Ecommerce(codeTra string, coin string, monto float64,
		emisor string, receptor string, fecha string) (domain.Transaction, error)
	GetOne(id int) (domain.Transaction, error)
	Update(id int, codeTra string, coin string, monto float64, emisor string,
		receptor string, fecha string) (domain.Transaction, error)
	UpdateOne(id int, codeTra string, monto float64) (domain.Transaction, error)
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

func (s *service) Update(id int, codeTra string, coin string, monto float64, emisor string,
	receptor string, fecha string) (domain.Transaction, error) {

	return s.repository.Update(id, codeTra, coin, monto, emisor, receptor, fecha)
}

func (s *service) UpdateOne(id int, codeTra string, monto float64) (domain.Transaction, error) {
	return s.repository.UpdateOne(id, codeTra, monto)
}

func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}
