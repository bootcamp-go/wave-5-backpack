package transactions

import (
	"arquitectura/internal/domain"
	"fmt"
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

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAll() ([]domain.Transaction, error) {
	return lista, nil
}

func (r *repository) LastID() (int, error) {
	return lastID, nil
}

func (r *repository) Store(id int, tranCode, currency string, amount float64, transmitter, receiver, tranDate string) (domain.Transaction, error) {
	t := domain.Transaction{
		Id:          id,
		TranCode:    tranCode,
		Currency:    currency,
		Amount:      amount,
		Transmitter: transmitter,
		Reciever:    receiver,
		TranDate:    tranCode,
	}

	lista = append(lista, t)
	lastID = t.Id
	return t, nil
}

func (r *repository) Update(id int, tranCode, currency string, amount float64, transmitter, receiver, tranDate string) (domain.Transaction, error) {
	t := domain.Transaction{
		Id:          id,
		TranCode:    tranCode,
		Currency:    currency,
		Amount:      amount,
		Transmitter: transmitter,
		Reciever:    receiver,
		TranDate:    tranCode,
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
		return domain.Transaction{}, fmt.Errorf("la transaccion con id %d no existe", id)
	}

	return t, nil
}

func (r *repository) UpdateTranCode(id int, tranCode string) (domain.Transaction, error) {
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
		return domain.Transaction{}, fmt.Errorf("la transaccion con id %d no existe", id)
	}

	return t, nil
}

func (r *repository) UpdateAmount(id int, amount float64) (domain.Transaction, error) {
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

	return t, nil
}

func (r *repository) Delete(id int) error {
	deleted := false
	var index int
	for i := range lista {
		if lista[i].Id == id {
			index = i
			deleted = true
		}
	}

	if !deleted {
		return fmt.Errorf("la transaccion con id %d no existe", id)
	}

	lista = append(lista[:index], lista[index+1:]...)
	return nil
}

// variables globales
var lista []domain.Transaction
var lastID int
