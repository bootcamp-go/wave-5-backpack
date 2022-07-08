package transactions

import (
	"goweb/internal/domain"
	"strconv"
	"time"
)

type Repository interface {
	GetAll() ([]domain.Transaction, error)
	Store(Id int, TransactionCode string, Currency string, Amount float64, Sender string, Reciever string, TransactionDate time.Time) (domain.Transaction, error)
	GetById(Id int) (domain.Transaction, error)
	lastId() (int, error)
	Update(id int, Currency string, Amount float64, Sender string, Reciever string) (domain.Transaction, error)
	Delete(id int) error
	UpdateCurrencyAndAmount(id int, Currency string, Amount float64) (domain.Transaction, error)
}

var transactions []domain.Transaction = make([]domain.Transaction, 0)
var lastId int = 0

type repository struct {
}

func NewRepository() Repository {

	return &repository{}
}

func (r *repository) GetAll() ([]domain.Transaction, error) {
	return transactions, nil
}
func (r *repository) Store(Id int, TransactionCode string, Currency string, Amount float64, Sender string, Reciever string, TransactionDate time.Time) (domain.Transaction, error) {
	transaction := domain.Transaction{
		Id:              Id,
		TransactionCode: TransactionCode,
		Currency:        Currency,
		Amount:          Amount,
		Sender:          Sender,
		Reciever:        Reciever,
		TransactionDate: TransactionDate,
	}
	transactions = append(transactions, transaction)
	return transaction, nil
}

func (r *repository) Update(id int, Currency string, Amount float64, Sender string, Reciever string) (domain.Transaction, error) {

	for i, transaction := range transactions {
		if transaction.Id == id {
			transaction.Currency = Currency
			transaction.Amount = Amount
			transaction.Sender = Sender
			transaction.Reciever = Reciever
			transactions[i] = transaction
			return transaction, nil
		}
	}
	return domain.Transaction{}, &NotFound{searchValue: strconv.Itoa(id), fieldName: "Id"}
}

func (r *repository) GetById(searchId int) (domain.Transaction, error) {
	for _, transaction := range transactions {
		if transaction.Id == searchId {
			return transaction, nil
		}
	}
	return domain.Transaction{}, &NotFound{searchValue: strconv.Itoa(searchId), fieldName: "Id"}
}
func (r *repository) lastId() (int, error) {
	return lastId, nil
}

func (r *repository) Delete(id int) error {
	for i, trans := range transactions {
		if trans.Id == id {

			transactions = append(transactions[:i], transactions[i+1:]...)
			return nil
		}
	}
	return &NotFound{searchValue: strconv.Itoa(id), fieldName: "Id"}
}

func (r *repository) UpdateCurrencyAndAmount(id int, Currency string, Amount float64) (domain.Transaction, error) {
	for i, transaction := range transactions {
		if transaction.Id == id {
			transaction.Currency = Currency
			transaction.Amount = Amount
			transactions[i] = transaction
			return transaction, nil
		}
	}
	return domain.Transaction{}, &NotFound{searchValue: strconv.Itoa(id), fieldName: "Id"}
}
