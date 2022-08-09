package transactions

import (
	"context"

	"github.com/bootcamp-go/wave-5-backpack/tree/lopez_cristian/storage/internal/models"
)

type Service interface {
	GetAll(ctx context.Context) ([]models.Transaction, error)
	GetByID(ctx context.Context, id int) (models.Transaction, error)
	Store(ctx context.Context, monto float64, cod, moneda, emisor, receptor string) (models.Transaction, error)
	Update(ctx context.Context, id int, monto float64, cod, moneda, emisor, receptor string) (models.Transaction, error)
	Patch(ctx context.Context, id int, monto float64, cod, moneda, emisor, receptor string) (models.Transaction, error)
	Delete(ctx context.Context, id int) (int, error)
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

type service struct {
	repository Repository
}

func (s service) Store(ctx context.Context, monto float64, cod, moneda, emisor, receptor string) (models.Transaction, error) {
	return s.repository.Store(ctx, monto, cod, moneda, emisor, receptor)
}

func (s service) GetAll(ctx context.Context) ([]models.Transaction, error) {
	return s.repository.GetAll(ctx)
}

func (s service) GetByID(ctx context.Context, id int) (models.Transaction, error) {
	return s.repository.GetByID(ctx, id)
}

func (s service) Update(ctx context.Context, id int, monto float64, cod, moneda, emisor, receptor string) (models.Transaction, error) {
	return s.repository.Update(ctx, id, monto, cod, moneda, emisor, receptor)
}

func (s service) Patch(ctx context.Context, id int, monto float64, cod, moneda, emisor, receptor string) (models.Transaction, error) {
	return models.Transaction{}, nil
}

func (s service) Delete(ctx context.Context, id int) (int, error) {
	return 0, nil
}
