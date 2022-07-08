package transactions

import (
	"GoWeb/internals/domain"
	"fmt"
)

type Repository interface {
	GetAll() ([]domain.Transanction, error)
	Store(id int, code, coin string, amount float64, emisor, receptor, date string) (domain.Transanction, error)
	lastID() (int, error)
	Update(id int, code, coin string, amount float64, emisor, receptor, date string) (domain.Transanction, error)
	Delete(id int) error
	UpdateCode(id int, code string, amount float64) (domain.Transanction, error)
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

func (r *repository) Update(id int, code, coin string, amount float64, emisor, receptor, date string) (domain.Transanction, error) {
	tran := domain.Transanction{Code: code, Coin: coin, Amount: amount, Emisor: emisor, Receptor: receptor, Date: date}
	updated := false
	for i := range transactions {
		if transactions[i].Id == id {
			tran.Id = id
			transactions[i] = tran
			updated = true
		}
	}

	if !updated {
		return domain.Transanction{}, fmt.Errorf("emisor %d no encontrado", id)
	}
	return tran, nil
}

func (r *repository) Delete(id int) error {
	deleted := false
	var index int
	for i := range transactions {
		if transactions[i].Id == id {
			index = i
			deleted = true
		}
	}
	if !deleted {
		return fmt.Errorf("El id %d no fue encontrato", id)
	}
	transactions = append(transactions[:index], transactions[index+1:]...)

	return nil
}

func (r *repository) UpdateCode(id int, code string, amount float64) (domain.Transanction, error) {
	updated := false
	var t1 domain.Transanction
	for i := range transactions {
		if transactions[i].Id == id {
			transactions[i].Code = code
			transactions[i].Amount = amount
			t1 = transactions[i]
			updated = true
		}
	}
	if !updated {
		return domain.Transanction{}, fmt.Errorf("el codigo %s o la cantidad %f no son validos", code, amount)
	}

	return t1, nil
}
