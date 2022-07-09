package transactions

import (
	"fmt"
	"goweb/pkg/store"
)

type Transactions struct {
	Id              int64   `json:"id"`
	TransactionCode string  `json:"transaction_code"`
	TypeCurrency    string  `json:"type_of_currency"`
	Amount          float64 `json:"amount"`
	Transmitter     string  `json:"transmitter"`
	Receiver        string  `json:"receiver"`
	Date            string  `json:"date"`
	Completed       bool    `json:"completed"`
}

type repository struct {
	db store.Store
}

var transactions []*Transactions
var lastId int64

type Repository interface {
	LastId() (int64, error)
	GetAll() ([]*Transactions, error)
	Store(code, currency, transmitter, receiver, date string, amount float64, completed bool) (*Transactions, error)
	Update(id int64, code, currency, transmitter, receiver, date string, amount float64, completed bool) (*Transactions, error)
	UpdateTransmitter(id int64, transmitter string) (*Transactions, error)
	Delete(id int64) error
}

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

func (r repository) GetAll() ([]*Transactions, error) {
	if err := r.db.Read(&transactions); err != nil {
		return nil, err
	}
	return transactions, nil
}

func (r *repository) Store(code, currency, transmitter, receiver, date string, amount float64, completed bool) (*Transactions, error) {
	lastId += 1

	r.db.Read(&transactions)

	transaction := &Transactions{
		Id:              lastId,
		TransactionCode: code,
		TypeCurrency:    currency,
		Amount:          amount,
		Transmitter:     transmitter,
		Receiver:        receiver,
		Date:            date,
		Completed:       completed,
	}
	transactions = append(transactions, transaction)
	if err := r.db.Write(&transactions); err != nil {
		return nil, err
	}
	return transaction, nil
}

func (r *repository) Update(id int64, code, currency, transmitter, receiver, date string, amount float64, completed bool) (*Transactions, error) {
	updated := false
	p := &Transactions{
		TransactionCode: code,
		TypeCurrency:    currency,
		Amount:          amount,
		Transmitter:     transmitter,
		Receiver:        receiver,
		Date:            date,
		Completed:       completed,
	}
	r.db.Read(&transactions)

	for i := range transactions {
		if transactions[i].Id == id {
			p.Id = id
			transactions[i] = p
			updated = true
		}
	}

	if !updated {
		return &Transactions{}, fmt.Errorf("producto %d no encontrado", id)
	}

	if err := r.db.Write(&transactions); err != nil {
		return nil, err
	}

	return p, nil
}

func (r *repository) UpdateTransmitter(id int64, transmitter string) (*Transactions, error) {
	var p Transactions
	updatedT := false

	r.db.Read(&transactions)

	for i := range transactions {
		if transactions[i].Id == id {
			transactions[i].Transmitter = transmitter
			updatedT = true
			p = *transactions[i]
		}
	}
	if !updatedT {
		return &Transactions{}, fmt.Errorf("product id-%d not found", id)
	}

	if err := r.db.Write(&transactions); err != nil {
		return nil, err
	}

	return &p, nil
}

func (r *repository) Delete(id int64) error {
	deleted := false
	var index int

	r.db.Read(&transactions)

	for i := range transactions {
		if transactions[i].Id == id {
			index = i
			deleted = true
		}
	}
	if !deleted {
		return fmt.Errorf("product id-%d not found", id)
	}
	transactions = append(transactions[:index], transactions[index+1:]...)

	if err := r.db.Write(&transactions); err != nil {
		return err
	}

	return nil
}

func (r *repository) LastId() (int64, error) {

	if err := r.db.Read(&transactions); err != nil {
		return 0, err
	}

	if len(transactions) == 0 {
		return 0, nil
	}

	return transactions[len(transactions)-1].Id, nil
}
