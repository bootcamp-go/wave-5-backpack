package transactions

import (
	"errors"
	"goweb/internal/domain"
	"time"
)

type Repository interface {
	GetAll() ([]domain.Transaction, error)
	Store(Id int, TransactionCode string, Currency string, Amount float64, Sender string, Reciever string, TransactionDate time.Time) (domain.Transaction, error)
	GetById(Id int) (domain.Transaction, error)
	lastId() (int, error)
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
func (r *repository) GetById(searchId int) (domain.Transaction, error) {
	for _, transaction := range transactions {
		if transaction.Id == searchId {
			return transaction, nil
		}
	}
	return domain.Transaction{}, errors.New("error: id not found in database")
}
func (r *repository) lastId() (int, error) {
	return lastId, nil
}
