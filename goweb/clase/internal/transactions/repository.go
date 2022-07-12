package transactions

import (
	"errors"
	"fmt"
	"goweb/internal/domain"
	"goweb/pkg/store"
)

const (
	ProductNotFound = "product %d not found"
	FailReading     = "cant read database"
	FailWriting     = "cant write database, error: %w"
)

type Repository interface {
	GetAll() ([]domain.Transaction, error)
	GetOne(id int) (domain.Transaction, error)
	LastID() (int, error)
	Create(id int, code, currency string, amount float64, issuer, recipient, date string) (domain.Transaction, error)
	Update(id int, code, currency string, amount float64, issuer, recipient, date string) (domain.Transaction, error)
	Update2(id int, code string, amount float64) (domain.Transaction, error)
	Delete(id int) error
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {

	r := &repository{
		db: db,
	}
	return r
}

func (r *repository) GetAll() ([]domain.Transaction, error) {
	var transactions []domain.Transaction
	if err := r.db.Read(&transactions); err != nil {
		return nil, fmt.Errorf(FailReading)
	}
	return transactions, nil
}

func (r *repository) GetOne(id int) (domain.Transaction, error) {
	var transactions []domain.Transaction
	if err := r.db.Read(&transactions); err != nil {
		return domain.Transaction{}, fmt.Errorf(FailReading)
	}
	for i := range transactions {
		if id == transactions[i].Id {
			return transactions[i], nil
		}
	}
	return domain.Transaction{}, errors.New("Transaccion no existente")
}

func (r *repository) LastID() (int, error) {
	var transactions []domain.Transaction
	if err := r.db.Read(&transactions); err != nil {
		return 0, fmt.Errorf(FailReading)
	}
	if len(transactions) == 0 {
		return 0, nil
	} else {
		return transactions[len(transactions)-1].Id, nil
	}
}

func (r *repository) Create(id int, code, currency string, amount float64, issuer, recipient, date string) (domain.Transaction, error) {
	var transactions []domain.Transaction
	if err := r.db.Read(&transactions); err != nil {
		return domain.Transaction{}, fmt.Errorf(FailReading)
	}
	t := domain.Transaction{
		Id:        id,
		Code:      code,
		Currency:  currency,
		Amount:    amount,
		Issuer:    issuer,
		Recipient: recipient,
		Date:      date,
	}
	transactions = append(transactions, t)
	if err := r.db.Write(transactions); err != nil {
		return domain.Transaction{}, fmt.Errorf(FailWriting, err)
	}
	return t, nil
}

func (r *repository) Update(id int, code, currency string, amount float64, issuer, recipient, date string) (domain.Transaction, error) {
	var transactions []domain.Transaction
	if err := r.db.Read(&transactions); err != nil {
		return domain.Transaction{}, fmt.Errorf(FailReading)
	}
	newTransaction := domain.Transaction{
		Id:        id,
		Code:      code,
		Currency:  currency,
		Amount:    amount,
		Issuer:    issuer,
		Recipient: recipient,
		Date:      date,
	}
	for i := range transactions {
		if id == transactions[i].Id {
			transactions[i] = newTransaction
			if err := r.db.Write(transactions); err != nil {
				return domain.Transaction{}, fmt.Errorf(FailWriting, err)
			}
			return newTransaction, nil
		}
	}
	return domain.Transaction{}, fmt.Errorf(ProductNotFound, id)
}

func (r *repository) Delete(id int) error {
	var transactions []domain.Transaction
	if err := r.db.Read(&transactions); err != nil {
		return fmt.Errorf(FailReading)
	}
	for i := range transactions {
		if transactions[i].Id == id {
			transactions = append(transactions[:i], transactions[i+1:]...)
			if err := r.db.Write(transactions); err != nil {
				return fmt.Errorf(FailWriting, err)
			}
			return nil
		}
	}
	return fmt.Errorf(ProductNotFound, id)
}

func (r *repository) Update2(id int, code string, amount float64) (domain.Transaction, error) {
	var transactions []domain.Transaction
	if err := r.db.Read(&transactions); err != nil {
		return domain.Transaction{}, fmt.Errorf(FailReading)
	}
	for i := range transactions {
		if transactions[i].Id == id {
			transactions[i].Code = code
			transactions[i].Amount = amount
			if err := r.db.Write(transactions); err != nil {
				return domain.Transaction{}, fmt.Errorf(FailWriting, err)
			}
			return transactions[i], nil
		}
	}
	return domain.Transaction{}, fmt.Errorf(ProductNotFound, id)
}
