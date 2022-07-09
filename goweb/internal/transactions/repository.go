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
  Store(monto float64, cod, moneda, emisor, receptor string) (models.Transaction, error)
  Update(id int, monto float64, cod, moneda, emisor, receptor string) (models.Transaction, error)
  UpdateMontoCod(id int, monto float64, cod string) (models.Transaction, error)
  GetAll() ([]models.Transaction, error)
  GetByID(id int) (models.Transaction, error)
  GetLastID() (int, error)
  Delete(id int) (int, error)
}

func NewRepository() Repository {
  return &repository{}
}

type repository struct {}

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

func (r repository) Update(id int, monto float64, cod, moneda, emisor, receptor string) (models.Transaction, error) {
	for i , tt := range transactions {
		if tt.ID == id {
			t := models.Transaction{
				ID: id,
				Monto: monto,
				Cod: cod,
				Moneda: moneda,
				Emisor: emisor,
				Receptor: receptor,
				Fecha: tt.Fecha,
			}

			// Actualiza la memoria
			transactions[i] = t

			return t, nil
		}
	}

	return models.Transaction{}, fmt.Errorf("error: no existe el ID: %v", id)
}

func (r repository) UpdateMontoCod(id int, monto float64, cod string) (models.Transaction, error) {
	for i, tt := range transactions {
		if tt.ID == id {
			t := models.Transaction{
				ID: tt.ID,
				Moneda: tt.Moneda,
				Emisor: tt.Emisor,
				Receptor: tt.Receptor,
				Fecha: tt.Fecha,
			}

			if monto != 0 {
				t.Monto = monto
			} else {
				t.Monto = tt.Monto
			}

			if cod != "" {
				t.Cod = cod
			} else {
				t.Cod = tt.Cod
			}

			// Actualiza la memoria
			transactions[i] = t

			return t, nil
		}
	}

	return models.Transaction{}, fmt.Errorf("error: recurso con ID: %v no encontrado", id)
}

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

func (r repository) GetLastID() (int, error) {
  if len(transactions) == 0 {
  	return 0, errors.New("no hay registros")
  }

  return lastID, nil
}

func (r repository) Delete(id int) (int, error) {
	for i , t := range transactions {
		if t.ID == id {
			if i == len(transactions) - 1 {
				transactions = transactions[:i]
				return id, nil
			}

			transactions = append(transactions[:i], transactions[i+1:]...)
			return id, nil
		}
	}

	return 0, fmt.Errorf("error: ID %v no existe", id)
}
