package transactions

import (
	"clase1-storage-implementation-tt/internal/domain"
)

// Service ...
type Service interface {
	GetAll() ([]domain.Transaction, error)
	Ecommerce(codeTra string, coin string, monto float64,
		emisor string, receptor string, fecha string) (domain.Transaction, error)
	GetOne(id int) (domain.Transaction, error)
	GetByName(code string) ([]domain.Transaction, error)
	Update(id int, codeTra string, coin string, monto float64, emisor string,
		receptor string, fecha string) (domain.Transaction, error)
	UpdateOne(id int, codeTra string, monto float64) (domain.Transaction, error)
	Delete(id int) (domain.Transaction, error)
}
type service struct {
	repository Repository
}

// NewService ...
func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll() ([]domain.Transaction, error) {
	return s.repository.GetAll()
}

func (s *service) Ecommerce(codeTra string, currency string, monto float64,
	emisor string, receptor string, fecha string) (domain.Transaction, error) {
	req := domain.Transaction{
		CodigoTransaccion: codeTra,
		Moneda:            currency,
		Monto:             monto,
		Emisor:            emisor,
		Receptor:          receptor,
		Fecha:             fecha,
	}
	return s.repository.Ecommerce(req)
}

func (s *service) GetOne(id int) (domain.Transaction, error) {
	return s.repository.GetOne(id)
}

func (s *service) GetByName(code string) ([]domain.Transaction, error) {
	return s.repository.GetByName(code)
}

func (s *service) Update(id int, codeTra string, coin string, monto float64, emisor string,
	receptor string, fecha string) (domain.Transaction, error) {
	return s.repository.Update(id, codeTra, coin, monto, emisor, receptor, fecha)
}

func (s *service) UpdateOne(id int, codeTra string, monto float64) (domain.Transaction, error) {
	return s.repository.UpdateOne(id, codeTra, monto)
}

func (s *service) Delete(id int) (domain.Transaction, error) {
	return s.repository.Delete(id)
}
