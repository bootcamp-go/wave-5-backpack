package transactions

import (
	"encoding/json"
	"fmt"
	"goweb/internal/domain"
	"os"
)

type Repository interface {
	GetAll() ([]domain.Transaction, error)
	LastID() (int, error)
	Create(id int, code, currency string, amount float64, issuer, recipient, date string) (domain.Transaction, error)
}

type repository struct{}

func loadFile() {
	file, _ := os.ReadFile("transactions.json")
	json.Unmarshal([]byte(file), &transactions)
}

func NewRepository() Repository {

	r := &repository{}
	loadFile()
	fmt.Println(transactions)
	return r
}

var transactions []domain.Transaction

func (r *repository) GetAll() ([]domain.Transaction, error) {
	return transactions, nil
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
