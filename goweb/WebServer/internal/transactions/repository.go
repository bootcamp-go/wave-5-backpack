package transactions

import (
	"WebServer/pkg/store"
	"errors"
	"fmt"
)

// en la capa repositorio se contienen todas los datos de la estructura
type Transaction struct {
	Id                int     `json:"id"`
	CodigoTransaccion int     `json:"codigo_de_transaccion"`
	Moneda            string  `json:"moneda"`
	Monto             float64 `json:"monto"`
	Emisor            string  `json:"emisor"`
	Receptor          string  `json:"receptor"`
	FechaTransaccion  string  `json:"fecha_de_transaccion"`
}

var Transactions []Transaction
var lastID int

type Repository interface {
	GetAll() ([]Transaction, error)
	Create(id, codigoTransaccion int, moneda string, monto float64, emisor, receptor, fechaTransaccion string) (Transaction, error)
	LastId() (int, error)
	Update(id, codigoTransaccion int, moneda string, monto float64, emisor, receptor, fechaTransaccion string) (Transaction, error)
	UpdatePartial(id, codigoTransaccion int, monto float64) (Transaction, error)
	Delete(id int) (Transaction, error)
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}
func (r *repository) GetAll() ([]Transaction, error) {
	var tr []Transaction
	if err := r.db.Read(&tr); err != nil {
		return []Transaction{}, fmt.Errorf("FailReading")
	}
	return tr, nil
}

func (r *repository) LastId() (int, error) {
	var tr []Transaction
	if err := r.db.Read(&tr); err != nil {
		return 0, fmt.Errorf("FailReading")
	}
	if len(tr) == 0 {
		return 0, nil
	}

	return tr[len(tr)-1].Id, nil
}

func (r *repository) Create(id, codigoTransaccion int, moneda string, monto float64, emisor, receptor, fechaTransaccion string) (Transaction, error) {
	var tr []Transaction

	if err := r.db.Read(&tr); err != nil {
		return Transaction{}, err
	}
	t := Transaction{Id: id, CodigoTransaccion: codigoTransaccion, Moneda: moneda, Monto: monto, Emisor: emisor, Receptor: receptor, FechaTransaccion: fechaTransaccion}
	tr = append(tr, t)

	if err := r.db.Write(tr); err != nil {
		return Transaction{}, fmt.Errorf("FailWriting %w", err)
	}

	return t, nil
}

func (r *repository) Update(id, codigoTransaccion int, moneda string, monto float64, emisor, receptor, fechaTransaccion string) (Transaction, error) {
	var tr []Transaction

	if err := r.db.Read(&tr); err != nil {
		return Transaction{}, err
	}
	t := Transaction{Id: id, CodigoTransaccion: codigoTransaccion, Moneda: moneda, Monto: monto, Emisor: emisor, Receptor: receptor, FechaTransaccion: fechaTransaccion}
	var updateReg bool
	for i := range tr {
		if tr[i].Id == id {
			tr[i] = t
			updateReg = true
		}
	}
	if !updateReg {
		return Transaction{}, errors.New("id not found")
	}
	if err := r.db.Write(tr); err != nil {
		return Transaction{}, fmt.Errorf("FailWriting %w", err)
	}
	return t, nil
}
func (r *repository) UpdatePartial(id, codigoTransaccion int, monto float64) (Transaction, error) {
	var tr []Transaction
	var tr_i Transaction
	if err := r.db.Read(&tr); err != nil {
		return Transaction{}, err
	}
	var updateReg bool
	for i := range tr {
		if tr[i].Id == id {
			tr[i].CodigoTransaccion = codigoTransaccion
			tr[i].Monto = monto
			tr_i = tr[i]
			updateReg = true
		}
	}
	if !updateReg {
		return Transaction{}, errors.New("id not found")
	}
	if err := r.db.Write(tr); err != nil {
		return Transaction{}, fmt.Errorf("FailWriting %w", err)
	}

	return tr_i, nil
}

func (r *repository) Delete(id int) (Transaction, error) {
	var deleteReg bool
	deleteTransaction := Transaction{}
	var tr []Transaction
	if err := r.db.Read(&tr); err != nil {
		return Transaction{}, err
	}
	for i := range tr {
		if tr[i].Id == id {
			deleteTransaction = tr[i]
			tr = append(tr[:i], tr[i+1:]...)
			deleteReg = true
			break
		}
	}
	if !deleteReg {
		return Transaction{}, errors.New("id not found")
	}
	if err := r.db.Write(tr); err != nil {
		return Transaction{}, fmt.Errorf("FailWriting %w", err)
	}
	return deleteTransaction, nil
}
