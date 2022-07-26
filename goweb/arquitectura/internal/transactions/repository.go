package transactions

import (
	"arquitectura/internal/domain"
	"arquitectura/pkg/store"
	"fmt"
)

const (
	FailReading         = "can`t read database"
	FailWriting         = "can`t write database, error: %w"
	TransactionNotFound = "transaction with id %d doesn`t exists en database"
)

type Repository interface {
	GetAll() ([]domain.Transaction, error)
	Store(id int, tranCode, currency string, amount float64, transmitter, receiver, tranDate string) (domain.Transaction, error)
	Update(id int, tranCode, currency string, amount float64, transmitter, receiver, tranDate string) (domain.Transaction, error)
	UpdateTranCode(id int, tranCode string) (domain.Transaction, error)
	UpdateAmount(id int, amount float64) (domain.Transaction, error)
	Delete(id int) error
	LastID() (int, error)
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
	return lista[len(lista)-1].Id, nil
}

func (r *repository) Store(id int, tranCode, currency string, amount float64, transmitter, receiver, tranDate string) (domain.Transaction, error) {

	var lista []domain.Transaction

	if err := r.db.Read(&lista); err != nil {
		return domain.Transaction{}, fmt.Errorf(FailReading)
	}

	t := domain.Transaction{
		Id:          id,
		TranCode:    tranCode,
		Currency:    currency,
		Amount:      amount,
		Transmitter: transmitter,
		Reciever:    receiver,
		TranDate:    tranDate,
	}
	lista = append(lista, t)

	if err := r.db.Write(lista); err != nil {
		return domain.Transaction{}, fmt.Errorf(FailWriting, err)
	}
	return t, nil
}

func (r *repository) Update(id int, tranCode, currency string, amount float64, transmitter, receiver, tranDate string) (domain.Transaction, error) {
	var lista []domain.Transaction

	if err := r.db.Read(&lista); err != nil {
		return domain.Transaction{}, fmt.Errorf(FailReading)
	}

	t := domain.Transaction{
		Id:          id,
		TranCode:    tranCode,
		Currency:    currency,
		Amount:      amount,
		Transmitter: transmitter,
		Reciever:    receiver,
		TranDate:    tranDate,
	}
	updated := false
	for i := range lista {
		if lista[i].Id == id {
			lista[i] = t
			updated = true
			break
		}
	}

	if !updated {
		return domain.Transaction{}, fmt.Errorf(TransactionNotFound, id)
	}

	if err := r.db.Write(lista); err != nil {
		return domain.Transaction{}, fmt.Errorf(FailWriting, err)
	}

	return t, nil
}

func (r *repository) UpdateTranCode(id int, tranCode string) (domain.Transaction, error) {
	var lista []domain.Transaction

	if err := r.db.Read(&lista); err != nil {
		return domain.Transaction{}, fmt.Errorf(FailReading)
	}

	updated := false
	var t domain.Transaction
	for i := range lista {
		if lista[i].Id == id {
			lista[i].TranCode = tranCode
			t = lista[i]
			updated = true
		}
	}

	if !updated {
		return domain.Transaction{}, fmt.Errorf(TransactionNotFound, id)
	}

	if err := r.db.Write(lista); err != nil {
		return domain.Transaction{}, fmt.Errorf(FailWriting, err)
	}

	return t, nil
}

func (r *repository) UpdateAmount(id int, amount float64) (domain.Transaction, error) {
	var lista []domain.Transaction

	if err := r.db.Read(&lista); err != nil {
		return domain.Transaction{}, fmt.Errorf(FailReading)
	}

	updated := false
	var t domain.Transaction
	for i := range lista {
		if lista[i].Id == id {
			lista[i].Amount = amount
			t = lista[i]
			updated = true
		}
	}

	if !updated {
		return domain.Transaction{}, fmt.Errorf("la transaccion con id %d no existe", id)
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
		return fmt.Errorf(TransactionNotFound, id)
	}

	lista = append(lista[:index], lista[index+1:]...)

	if err := r.db.Write(lista); err != nil {
		return fmt.Errorf(FailWriting, err)
	}

	return nil
}
