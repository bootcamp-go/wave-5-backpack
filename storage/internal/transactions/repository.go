package transactions

import (
	"github.com/bootcamp-go/wave-5-backpack/tree/lopez_cristian/storage/cmd/db"
	"github.com/bootcamp-go/wave-5-backpack/tree/lopez_cristian/storage/internal/models"
)

type Repository interface {
	GetByCod(cod string) (models.Transaction, error)
}

func NewRepository() Repository {
	return &repository{}
}

type repository struct{}

var getByCod = "SELECT id, monto, cod, moneda, emisor, receptor, fecha FROM transactions WHERE cod = ?;"

func (r *repository) GetByCod(cod string) (models.Transaction, error) {
	db := db.StorageDB

	rows, err := db.Query(getByCod, cod)
	if err != nil {
		return models.Transaction{}, err
	}

	var transaction models.Transaction
	for rows.Next() {
		if err := rows.Scan(&transaction); err != nil {
			return models.Transaction{}, err
		}
	}

	return transaction, nil
}
