package transactions

import (
	"errors"
	"fmt"
	"time"

	"github.com/bootcamp-go/wave-5-backpack/tree/lopez_cristian/goweb/internal/models"
)

// Simula DB
var transactions []models.Transaction
var lastID int

type Repository interface {
  GetAll() ([]models.Transaction, error)
  GetByID(id int) (models.Transaction, error)
  Store(monto float64, cod, moneda, emisor, receptor string) (models.Transaction, error)
  GetLastID() (int, error)
}

func NewRepository() Repository {
  return &repository{}
}

type repository struct {}

func (r repository) GetAll() ([]models.Transaction, error) {
	if len(transactions) == 0 {
		return nil, errors.New("no hay registros")
	}

  return transactions, nil
}

func (r repository) GetByID(id int) (models.Transaction, error) {
	for _ , t := range transactions {
		if t.ID == id {
			return t, nil
		}
	}

	return models.Transaction{}, fmt.Errorf("trasaction con ID: %v no encontrado", id)
}

func (r repository) Store(monto float64, cod, moneda, emisor, receptor string) (models.Transaction, error) {
	lastID += 1
  t := models.Transaction{
    ID: lastID,
    Monto: monto,
    Cod: cod,
    Moneda: moneda,
    Emisor: emisor,
    Receptor: receptor,
    Fecha: time.Now().Local().String(),
  }

  transactions = append(transactions, t)

  return t, nil
}

func (r repository) GetLastID() (int, error) {
  if len(transactions) == 0 {
  	return 0, errors.New("no hay registros")
  }

  return lastID, nil
}
