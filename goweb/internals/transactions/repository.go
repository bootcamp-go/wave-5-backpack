package transactions

import (
	"fmt"
	"goweb/internals/domain"
	"goweb/pkg/store"
	"time"
)

const (
	TransNotFound = "La transacción %d no fue encontrada"
	FailReading   = "No se puede leer la base de datos"
	FailWriting   = "No se pudo escribir en la base de datos, error: %w"
)

type Repository interface {
	GetAll() ([]domain.Transaction, error)
	Store(id int, codigo string, moneda string, monto int, emisor string, receptor string) (domain.Transaction, error)
	LastID() (int, error)
	Update(id int, codigo string, moneda string, monto int, emisor string, receptor string) (domain.Transaction, error)
	Delete(id int) error
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll() ([]domain.Transaction, error) {
	var transactions []domain.Transaction
	if err := r.db.Read(&transactions); err != nil {
		return nil, fmt.Errorf(FailReading)
	}
	return transactions, nil
}

func (r *repository) LastID() (int, error) {
	var transactions []domain.Transaction
	if err := r.db.Read(&transactions); err != nil {
		return 0, fmt.Errorf(FailReading)
	}
	if len(transactions) == 0 {
		return 0, nil
	}
	return transactions[len(transactions)-1].Id, nil
}

func (r *repository) Store(id int, codigo string, moneda string, monto int, emisor string, receptor string) (domain.Transaction, error) {
	var transactions []domain.Transaction

	if err := r.db.Read(&transactions); err != nil {
		return domain.Transaction{}, fmt.Errorf(FailReading)
	}
	date := time.Now().Local().String()
	trans := domain.Transaction{
		Id:       id,
		Codigo:   codigo,
		Moneda:   moneda,
		Monto:    monto,
		Emisor:   emisor,
		Receptor: receptor,
		Fecha:    date,
	}
	transactions = append(transactions, trans)
	if err := r.db.Write(transactions); err != nil {
		return domain.Transaction{}, fmt.Errorf(FailWriting, err)
	}
	return trans, nil
}

func (r *repository) Update(id int, codigo string, moneda string, monto int, emisor string, receptor string) (domain.Transaction, error) {
	var transactions []domain.Transaction

	if err := r.db.Read(&transactions); err != nil {
		return domain.Transaction{}, fmt.Errorf(FailReading)
	}
	trans := domain.Transaction{
		Codigo:   codigo,
		Moneda:   moneda,
		Monto:    monto,
		Emisor:   emisor,
		Receptor: receptor,
	}
	updated := false
	for i := range transactions {
		if transactions[i].Id == id {
			trans.Id = id
			transactions[i] = trans
			updated = true
		}
	}
	if !updated {
		return domain.Transaction{}, fmt.Errorf("Producto %d no encontrado", id)
	}
	return trans, nil
}

func (r *repository) Delete(id int) error {
	var transactions []domain.Transaction

	if err := r.db.Read(&transactions); err != nil {
		return fmt.Errorf(FailReading)
	}

	deleted := false
	var index int
	for i := range transactions {
		if transactions[i].Id == id {
			index = i
			deleted = true
		}
	}
	if !deleted {
		return fmt.Errorf("Producto %d no encontrado", id)
	}
	transactions = append(transactions[:index], transactions[index+1:]...)
	return nil
}
