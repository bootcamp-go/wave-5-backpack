package transactions

import (
	"github.com/bootcamp-go/wave-5-backpack/tree/lopez_cristian/storage/internal/models"
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
	return models.Transaction{}, nil
}

func (s service) GetAll() ([]models.Transaction, error) {
	return []models.Transaction{}, nil
}

func (s service) GetByID(id int) (models.Transaction, error) {
	return models.Transaction{}, nil
}

func (s service) Update(id int, monto float64, cod, moneda, emisor, receptor string) (models.Transaction, error) {
	return models.Transaction{}, nil
}

func (s service) Patch(id int, monto float64, cod, moneda, emisor, receptor string) (models.Transaction, error) {
	return models.Transaction{}, nil
}

func (s service) Delete(id int) (int, error) {
	return 0, nil
}
