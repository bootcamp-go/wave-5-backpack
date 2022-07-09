package transactions

import (
	"github.com/bootcamp-go/wave-5-backpack/tree/lopez_cristian/goweb/internal/models"
)

type Service interface {
  GetAll() ([]models.Transaction, error)
  GetByID(id int) (models.Transaction, error)
  Store(monto float64, cod, moneda, emisor, receptor string) (models.Transaction, error)
  Update(id int, monto float64, cod, moneda, emisor, receptor string) (models.Transaction, error)
  Delete(id int) (int, error)
}

func NewService(r Repository) Service {
  return &service{
    repository: r,
  }
}

type service struct {
  repository Repository
}

func (s service) Store(monto float64, cod, moneda, emisor, receptor string) (models.Transaction, error) {
  return s.repository.Store(monto, cod, moneda, emisor, receptor)
}

func (s service) Update(id int, monto float64, cod, moneda, emisor, receptor string) (models.Transaction, error) {
	return s.repository.Update(id, monto, cod, moneda, emisor, receptor)
}

func (s service) GetByID(id int) (models.Transaction, error) {
  return s.repository.GetByID(id)
}

func (s service) GetAll() ([]models.Transaction, error) {
  transactions, err := s.repository.GetAll()

  if err != nil {
    return nil, err
  }

  return transactions, nil
}

func (s service) Delete(id int) (int, error) {
	return s.repository.Delete(id)
}
