package transactions

import (
	"goweb/internal/domain"
	"goweb/pkg/store"
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

var lastId int = 0

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {

	return &repository{
		db,
	}
}

func (r *repository) GetAll() ([]domain.Transaction, error) {
	var transactions []domain.Transaction
	if err := r.db.Read(&transactions); err != nil {
		return nil, err
	}
	return transactions, nil
}
func (r *repository) Store(Id int, TransactionCode string, Currency string, Amount float64, Sender string, Reciever string, TransactionDate time.Time) (domain.Transaction, error) {
	transactions, err := r.GetAll()
	if err != nil {
		return domain.Transaction{}, err
	}
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
	if err := r.db.Write(transactions); err != nil {
		return domain.Transaction{}, err
	}
	return transaction, nil
}

func (r *repository) Update(id int, Currency string, Amount float64, Sender string, Reciever string) (domain.Transaction, error) {
	transactions, err := r.GetAll()
	if err != nil {
		return domain.Transaction{}, err
	}
	transactionUpdate := domain.Transaction{}
	updated := false
	for i, transaction := range transactions {
		if transaction.Id == id {
			transaction.Currency = Currency
			transaction.Amount = Amount
			transaction.Sender = Sender
			transaction.Reciever = Reciever
			transactions[i] = transaction

			transactionUpdate = transaction
			updated = true
		}
	}
	if !updated {
		return domain.Transaction{}, &NotFound{searchValue: strconv.Itoa(id), fieldName: "Id"}
	}

	if err := r.db.Write(transactions); err != nil {
		return domain.Transaction{}, err
	}

	return transactionUpdate, nil
}

func (r *repository) GetById(searchId int) (domain.Transaction, error) {
	transactions, err := r.GetAll()
	if err != nil {
		return domain.Transaction{}, err
	}
	for _, transaction := range transactions {
		if transaction.Id == searchId {
			return transaction, nil
		}
	}
	return domain.Transaction{}, &NotFound{searchValue: strconv.Itoa(searchId), fieldName: "Id"}
}
func (r *repository) lastId() (int, error) {
	transactions, err := r.GetAll()
	if err != nil {
		return 0, err
	}

	return transactions[len(transactions)-1].Id, nil
}

func (r *repository) Delete(id int) error {
	transactions, err := r.GetAll()
	if err != nil {
		return err
	}
	deleted := false
	for i, trans := range transactions {
		if trans.Id == id {
			transactions = append(transactions[:i], transactions[i+1:]...)
			break
		}
	}
	if !deleted {
		return &NotFound{searchValue: strconv.Itoa(id), fieldName: "Id"}
	}
	if err := r.db.Write(transactions); err != nil {
		return err
	}
	return nil
}

func (r *repository) UpdateCurrencyAndAmount(id int, Currency string, Amount float64) (domain.Transaction, error) {
	transactions, err := r.GetAll()
	if err != nil {
		return domain.Transaction{}, err
	}
	updated := false
	transactionUpdate := domain.Transaction{}
	for i, transaction := range transactions {
		if transaction.Id == id {
			transaction.Currency = Currency
			transaction.Amount = Amount
			transactions[i] = transaction

			transactionUpdate = transaction
			updated = true
			break
		}
	}

	if !updated {
		return domain.Transaction{}, &NotFound{searchValue: strconv.Itoa(id), fieldName: "Id"}
	}
	if err := r.db.Write(transactions); err != nil {
		return domain.Transaction{}, err
	}
	return transactionUpdate, nil
}
