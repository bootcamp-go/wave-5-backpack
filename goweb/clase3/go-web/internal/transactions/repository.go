package transactions

import (
	"fmt"

	"github.com/bootcamp-go/go-web/pkg/store"
)

type Repository interface {
	LastId() (int64, error)
	GetAll() ([]*Transaction, error)
	Store(id int64, codigo, moneda, emisor, receptor string, monto float64) (*Transaction, error)
	Update(id int64, codigo, moneda, emisor, receptor string, monto float64) (*Transaction, error)
	UpdateReceptorYMonto(id int64, receptor string, monto float64) (*Transaction, error)
	Delete(id int64) error
}

type Transaction struct {
	Id       int64
	Codigo   string
	Moneda   string
	Emisor   string
	Receptor string
	Monto    float64
}

// No lo usamos más var transactions []*Transaction

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{db: db}
}

func (r *repository) GetAll() ([]*Transaction, error) {
	var ts []*Transaction
	if err := r.db.Read(&ts); err != nil {
		return nil, err
	}

	return ts, nil
}

func (r *repository) Store(id int64, codigo, moneda, emisor, receptor string, monto float64) (*Transaction, error) {
	var ts []Transaction
	r.db.Read(&ts) // manejamos el error

	transaction := Transaction{
		Id:       id,
		Codigo:   codigo,
		Moneda:   moneda,
		Emisor:   emisor,
		Receptor: receptor,
		Monto:    monto,
	}
	ts = append(ts, transaction)
	if err := r.db.Write(&ts); err != nil {
		return nil, err
	}

	return &transaction, nil
}

func (r *repository) Update(id int64, codigo, moneda, emisor, receptor string, monto float64) (*Transaction, error) {
	update := false
	transactionNew := &Transaction{
		Id:       id,
		Codigo:   codigo,
		Moneda:   moneda,
		Receptor: receptor,
		Emisor:   emisor,
		Monto:    monto,
	}

	var ts []*Transaction
	r.db.Read(&ts)

	for value := range ts {
		if ts[value].Id == id {
			ts[value] = transactionNew
			update = true
		}
	}

	if !update {
		return nil, fmt.Errorf("transacción id %d no encontrada", id)
	}

	if err := r.db.Write(&ts); err != nil {
		return nil, err
	}

	return transactionNew, nil
}

func (r *repository) UpdateReceptorYMonto(id int64, receptor string, monto float64) (*Transaction, error) {
	update := false
	var transaction *Transaction

	var ts []*Transaction
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
		return nil, fmt.Errorf("transacción id %d no encontrada", id)
	}

	if err := r.db.Write(&ts); err != nil {
		return nil, err
	}

	return transaction, nil
}

func (r *repository) Delete(id int64) error {
	deleted := false
	var indice int

	var ts []*Transaction
	r.db.Read(&ts)
	for value := range ts {
		if ts[value].Id == id {
			indice = value
			deleted = true
		}
	}

	if !deleted {
		return fmt.Errorf("la transacción id %d no existe", id)
	}

	ts = append(ts[:indice], ts[indice+1:]...)
	if err := r.db.Write(&ts); err != nil {
		return err
	}

	return nil
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
