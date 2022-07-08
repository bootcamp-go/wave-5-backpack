package transactions

import (
	"encoding/json"
	"errors"
	"goweb/internal/domain"
	"os"
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

type repository struct{}

func loadFile() {
	file, _ := os.ReadFile("transactions.json")
	json.Unmarshal([]byte(file), &transactions)
}

func NewRepository() Repository {

	r := &repository{}
	loadFile()
	return r
}

var transactions []domain.Transaction

func (r *repository) GetAll() ([]domain.Transaction, error) {
	return transactions, nil
}

func (r *repository) GetOne(id int) (domain.Transaction, error) {
	for i := range transactions {
		if id == transactions[i].Id {
			return transactions[i], nil
		}
	}
	return domain.Transaction{}, errors.New("Transaccion no existente")
}

func (r *repository) LastID() (int, error) {
	if len(transactions) == 0 {
		return 0, nil
	} else {
		return transactions[len(transactions)-1].Id, nil
	}
}

func (r *repository) Create(id int, code, currency string, amount float64, issuer, recipient, date string) (domain.Transaction, error) {
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
	return t, nil
}

func (r *repository) Update(id int, code, currency string, amount float64, issuer, recipient, date string) (domain.Transaction, error) {
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
			return newTransaction, nil
		}
	}
	return domain.Transaction{}, errors.New("transaccion no existente")
}

func (r *repository) Delete(id int) error {
	for i := range transactions {
		if transactions[i].Id == id {
			transactions = append(transactions[:i], transactions[i+1:]...)
			return nil
		}
	}
	return errors.New("transaccion no existente")
}

func (r *repository) Update2(id int, code string, amount float64) (domain.Transaction, error) {
	for i := range transactions {
		if transactions[i].Id == id {
			transactions[i].Code = code
			transactions[i].Amount = amount
			return transactions[i], nil
		}
	}
	return domain.Transaction{}, errors.New("transaccion no existente")
}
