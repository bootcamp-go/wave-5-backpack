package transactions

import (
	"fmt"

	"github.com/bootcamp-go/wave-5-backpack/tree/olivera_sebastian/goweb/clase2_parte2+siguientes/pkg/store"
)

type Repository interface {
	LastId() (int64, error)
	GetAll() ([]Transaction, error)
	Store(id int64, codigo, moneda, emisor, receptor string, monto float64) (Transaction, error)
	Update(id int64, monto float64, codigo, emisor, receptor, moneda string) (Transaction, error)
	UpdateReceptorYMonto(id int64, receptor string, monto float64) (Transaction, error)
	Delete(id int64) error
}

type Transaction struct {
	Id       int64
	Codigo   string
	Monto    float64
	Moneda   string
	Emisor   string
	Receptor string
}

//var transactions []Transaction
// var lastId int64

type repository struct {
	db store.Store
}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAll() ([]Transaction, error) {
	var ts []Transaction
	if err := r.db.Read(&ts); err != nil {
		return nil, err
	}

	return ts, nil
}

func (r *repository) Store(id int64, codigo, moneda, emisor, receptor string, monto float64) (Transaction, error) {
	var ts []Transaction
	r.db.Read(&ts)

	transaction := Transaction{
		Id:       id,
		Codigo:   codigo,
		Moneda:   moneda,
		Monto:    monto,
		Emisor:   emisor,
		Receptor: receptor,
	}

	ts = append(ts, transaction)
	if err := r.db.Write(&ts); err != nil {
		return Transaction{}, err
	}
	return transaction, nil
}

func (r *repository) LastId() (int64, error) {
	var ts []Transaction

	if err := r.db.Read(&ts); err != nil {
		return 0, err
	}

	if len(ts) == 0 {
		return 0, nil
	}

	return ts[len(ts)-1].Id, nil
}

func (r *repository) Update(id int64, monto float64, codigo, emisor, receptor, moneda string) (Transaction, error) {
	t := Transaction{
		Monto:    monto,
		Codigo:   codigo,
		Emisor:   emisor,
		Receptor: receptor,
		Moneda:   moneda,
	}
	updated := false
	var ts []Transaction
	r.db.Read(&ts)

	for value := range ts {
		if ts[value].Id == id {
			ts[value] = t
			updated = true
		}
	}

	if !updated {
		return Transaction{}, fmt.Errorf("transaccion %d no encontrada", id)
	}
	if err := r.db.Write(&ts); err != nil {
		return Transaction{}, err
	}

	return t, nil
}

func (r *repository) Delete(id int64) error {
	deleted := false
	var index int

	var ts []Transaction
	r.db.Read(&ts)
	for value := range ts {
		if ts[value].Id == id {
			index = value
			deleted = true
		}
	}
	if !deleted {
		return fmt.Errorf("la transacción id=%d no existe", id)
	}

	ts = append(ts[:index], ts[index+1:]...)
	if err := r.db.Write(&ts); err != nil {
		return err
	}
	return nil
}

func (r *repository) UpdateReceptorYMonto(id int64, receptor string, monto float64) (Transaction, error) {
	update := false
	var transaction Transaction
	var ts []Transaction
	r.db.Read(&ts)
	for value := range ts {
		if ts[value].Id == id {
			ts[value].Receptor = receptor
			ts[value].Monto = monto
			transaction = ts[value]
			update = true
		}
	}

	if !update {
		return Transaction{}, fmt.Errorf("transaccion id %d no encontrada", id)
	}

	if err := r.db.Write(&ts); err != nil {
		return Transaction{}, err
	}

	return transaction, nil

}
