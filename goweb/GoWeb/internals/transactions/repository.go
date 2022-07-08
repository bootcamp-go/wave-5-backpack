package transactions

import "GoWeb/internals/domain"

type Repository interface {
	GetAll() ([]domain.Transanction, error)
	Store(id int, code, coin string, amount float64, emisor, receptor, date string) (domain.Transanction, error)
	lastID() (int, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

var transactions []domain.Transanction
var lastID int

func (r *repository) GetAll() ([]domain.Transanction, error) {
	return transactions, nil
}

func (r *repository) lastID() (int, error) {
	return lastID, nil
}

func (r *repository) Store(id int, code, coin string, amount float64, emisor, receptor, date string) (domain.Transanction, error) {
	tran := domain.Transanction{Id: id, Code: code, Coin: coin, Amount: amount, Emisor: emisor, Receptor: receptor, Date: date}

	transactions = append(transactions, tran)
	lastID = tran.Id
	return tran, nil
}
