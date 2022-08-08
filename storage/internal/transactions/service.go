package transactions

import (
	"fmt"

	"github.com/bootcamp-go/wave-5-backpack/tree/lopez_cristian/goweb/internal/models"
)

type Service interface {
	GetAll() ([]models.Transaction, error)
	GetByID(id int) (models.Transaction, error)
	Store(monto float64, cod, moneda, emisor, receptor string) (models.Transaction, error)
	Update(id int, monto float64, cod, moneda, emisor, receptor string) (models.Transaction, error)
	Patch(id int, monto float64, cod, moneda, emisor, receptor string) (models.Transaction, error)
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

func (s service) GetAll() ([]models.Transaction, error) {
	return s.repository.GetAll()
}

func (s service) GetByID(id int) (models.Transaction, error) {
	return s.repository.GetByID(id)
}

func (s service) Update(id int, monto float64, cod, moneda, emisor, receptor string) (models.Transaction, error) {
	_, err := s.repository.GetByID(id)
	if err != nil {
		return models.Transaction{}, fmt.Errorf("error en repository: %v", err)
	}
	return s.repository.Update(id, monto, cod, moneda, emisor, receptor)
}

func (s service) Patch(id int, monto float64, cod, moneda, emisor, receptor string) (models.Transaction, error) {
	return s.repository.Patch(id, monto, cod, moneda, emisor, receptor)
}

func (s service) Delete(id int) (int, error) {
	return s.repository.Delete(id)
}
