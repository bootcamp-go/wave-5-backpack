package transactions

import (
	"arquitectura/internal/domain"
	"fmt"
)

type Repository interface {
	GetAll() ([]domain.Transaction, error)
	Store(id int, tranCode, currency string, amount float64, transmitter, receiver, tranDate string) (domain.Transaction, error)
	LastID() (int, error)
	Update(id int, tranCode, currency string, amount float64, transmitter, receiver, tranDate string) (domain.Transaction, error)
	Delete(id int) error
	UpdateCodeAmount(id int, tranCode string, amount float64) (domain.Transaction, error)
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

// Función Update(PUT)

func (r *repository) Update(id int, tranCode, currency string, amount float64, transmitter, receiver, tranDate string) (domain.Transaction, error) {

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
		return domain.Transaction{}, fmt.Errorf("Producto %d no encontrado", id)
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
		return fmt.Errorf("Producto %d no encontrado", id)
	}

	lista = append(lista[:index], lista[index+1:]...)

	return nil
}

func (r *repository) UpdateCodeAmount(id int, tranCode string, amount float64) (domain.Transaction, error) {

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

	return t, nil

}

// variables globales
var lista []domain.Transaction
var lastID int
