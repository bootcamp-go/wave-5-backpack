package transactions

import (
	"GoWeb/internals/domain"
	"GoWeb/pkg/store"
	"fmt"
)

const (
	NotFound    = "id %d not found"
	FailReading = "cant read database"
	FailWriting = "cant write database, error %w"
)

type Repository interface {
	GetAll() ([]domain.Transanction, error)
	Store(id int, code, coin string, amount float64, emisor, receptor, date string) (domain.Transanction, error)
	lastID() (int, error)
	Update(id int, code, coin string, amount float64, emisor, receptor, date string) (domain.Transanction, error)
	Delete(id int) error
	UpdateCode(id int, code string, amount float64) (domain.Transanction, error)
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll() ([]domain.Transanction, error) {
	var transactions []domain.Transanction
	if err := r.db.Read(&transactions); err != nil {
		return nil, fmt.Errorf(FailReading)
	}
	return transactions, nil
}

func (r *repository) lastID() (int, error) {
	var transactions []domain.Transanction
	if err := r.db.Read(&transactions); err != nil {
		return 0, fmt.Errorf(FailReading)
	}
	if len(transactions) == 0 {
		return 0, nil
	}
	return transactions[len(transactions)-1].Id, nil
}

func (r *repository) Store(id int, code, coin string, amount float64, emisor, receptor, date string) (domain.Transanction, error) {

	var transactions []domain.Transanction
	if err := r.db.Read(&transactions); err != nil {
		return domain.Transanction{}, fmt.Errorf(FailReading)
	}

	tran := domain.Transanction{Id: id, Code: code, Coin: coin, Amount: amount, Emisor: emisor, Receptor: receptor, Date: date}

	transactions = append(transactions, tran)
	if err := r.db.Write(transactions); err != nil {
		return domain.Transanction{}, fmt.Errorf(FailWriting, err)
	}
	return tran, nil
}

func (r *repository) Update(id int, code, coin string, amount float64, emisor, receptor, date string) (domain.Transanction, error) {
	var transactions []domain.Transanction
	if err := r.db.Read(&transactions); err != nil {
		return domain.Transanction{}, fmt.Errorf(FailReading)
	}

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
		return domain.Transanction{}, fmt.Errorf(NotFound, id)
	}
	if err := r.db.Write(transactions); err != nil {
		return domain.Transanction{}, fmt.Errorf(FailWriting, err)
	}
	return tran, nil
}

func (r *repository) Delete(id int) error {
	var transactions []domain.Transanction
	if err := r.db.Read(&transactions); err != nil {
		return fmt.Errorf(FailReading)
	}

	deleted := false
	var index int
	for i := range transactions {
		if transactions[i].Id == id {
			index = i
			deleted = true
		}
	}
	if !deleted {
		return fmt.Errorf("el id %d no fue encontrato", id)
	}
	transactions = append(transactions[:index], transactions[index+1:]...)

	if err := r.db.Write(transactions); err != nil {
		return fmt.Errorf(FailWriting, err)
	}

	return nil
}

func (r *repository) UpdateCode(id int, code string, amount float64) (domain.Transanction, error) {
	var transactions []domain.Transanction
	if err := r.db.Read(&transactions); err != nil {
		return domain.Transanction{}, fmt.Errorf(FailReading)
	}

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

	if err := r.db.Write(transactions); err != nil {
		return domain.Transanction{}, fmt.Errorf(FailWriting, err)
	}

	return t1, nil
}
