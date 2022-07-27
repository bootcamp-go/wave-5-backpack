package transactions

import (
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
	Store(id int, tranCode, currency string, amount float64, transmitter, receiver, tranDate string) (domain.Transaction, error)
	LastID() (int, error)
	Update(id int, tranCode, currency string, amount float64, transmitter, receiver, tranDate string) (domain.Transaction, error)
	Delete(id int) error
	UpdateCodeAmount(id int, tranCode string, amount float64) (domain.Transaction, error)
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {

	return &repository{
		db: db,
	}
}

func (r *repository) GetAll() ([]domain.Transaction, error) {
	var lista []domain.Transaction

	if err := r.db.Read(&lista); err != nil {
		return nil, fmt.Errorf(FailReading)
	}

	return lista, nil
}

func (r *repository) LastID() (int, error) {

	var lista []domain.Transaction
	if err := r.db.Read(&lista); err != nil {
		return 0, fmt.Errorf(FailReading)
	}

	if len(lista) == 0 {
		return 0, nil
	}

	//var lastID int
	lastID := lista[len(lista)-1].Id

	return lastID, nil
}

func (r *repository) Store(id int, tranCode, currency string, amount float64, transmitter, receiver, tranDate string) (domain.Transaction, error) {

	var lista []domain.Transaction

	t := domain.Transaction{
		Id:          id,
		TranCode:    tranCode,
		Currency:    currency,
		Amount:      amount,
		Transmitter: transmitter,
		Reciever:    receiver,
		TranDate:    tranCode,
	}

	if err := r.db.Read(&lista); err != nil {
		return domain.Transaction{}, fmt.Errorf(FailReading)
	}
	lista = append(lista, t)

	if err := r.db.Write(lista); err != nil {
		return domain.Transaction{}, fmt.Errorf(FailWriting, err)
	}

	return t, nil

	// Solución guardando en memorio local

	// lista = append(lista, t)
	// lastID = t.Id
	// return t, nil

}

// Función Update(PUT)

func (r *repository) Update(id int, tranCode, currency string, amount float64, transmitter, receiver, tranDate string) (domain.Transaction, error) {
	var lista []domain.Transaction
	if err := r.db.Read(&lista); err != nil {
		return domain.Transaction{}, fmt.Errorf(FailReading)
	}

	t := domain.Transaction{
		TranCode:    tranCode,
		Currency:    currency,
		Amount:      amount,
		Transmitter: transmitter,
		Reciever:    receiver,
		TranDate:    tranDate,
	}

	update := false
	for i := range lista {
		if lista[i].Id == id {
			t.Id = id
			lista[i] = t
			update = true
		}
	}

	if !update {
		return domain.Transaction{}, fmt.Errorf(ProductNotFound, id)
	}

	if err := r.db.Write(lista); err != nil {
		return domain.Transaction{}, fmt.Errorf(FailWriting, err)
	}

	return t, nil
}

func (r *repository) Delete(id int) error {
	var lista []domain.Transaction
	if err := r.db.Read(&lista); err != nil {
		return fmt.Errorf(FailReading)
	}

	deleted := false
	var index int
	for i := range lista {
		if lista[i].Id == id {
			index = i
			deleted = true
		}
	}

	if !deleted {
		return fmt.Errorf("Producto %d no encontrado", id)
	}

	lista = append(lista[:index], lista[index+1:]...)

	if err := r.db.Write(lista); err != nil {
		return fmt.Errorf(FailWriting, err)
	}

	return nil
}

func (r *repository) UpdateCodeAmount(id int, tranCode string, amount float64) (domain.Transaction, error) {

	var lista []domain.Transaction
	if err := r.db.Read(&lista); err != nil {
		return domain.Transaction{}, fmt.Errorf(FailReading)
	}

	var t domain.Transaction
	update := false
	for i := range lista {
		if lista[i].Id == id {
			lista[i].TranCode = tranCode
			lista[i].Amount = amount
			t = lista[i]
			update = true
		}
	}

	if !update {
		return domain.Transaction{}, fmt.Errorf("Producto %d no encontrado", id)
	}

	if err := r.db.Write(lista); err != nil {
		return domain.Transaction{}, fmt.Errorf(FailWriting, err)
	}

	return t, nil

}
