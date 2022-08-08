package transactions

import (
	"github.com/bootcamp-go/wave-5-backpack/tree/lopez_cristian/storage/internal/models"
)

type Repository interface {
	GetByCod(cod string) models.Transaction
}

func NewRepository() Repository {
	return &repository{}
}

type repository struct{}

func (r *repository) GetByCod(cod string) models.Transaction {
	return models.Transaction{}
}
