package transactions

import (
	"ejer02-TT/internal/domain"
	"ejer02-TT/pkg/store"
	"fmt"
)

type Repository interface {
	GetAll() ([]domain.Transaction, error)
	Store(id int, tranCode, currency string, amount float64, transmitter, receiver, tranDate string) (domain.Transaction, error)
	LastID() (int, error)
	Update(id int, tranCode, currency string, amount float64, transmitter, receiver, tranDate string) (domain.Transaction, error)
	UpdateCodeAndAmount(id int, tranCode string, amount float64) (domain.Transaction, error)
	Delete(id int) error
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) LastID() (int, error) {
	var ps []domain.Transaction
	if err := r.db.Read(&ps); err != nil {
		return 0, err
	}
	if len(ps) == 0 {
		return 0, nil
	}

	return ps[len(ps)-1].Id, nil
}

func (r *repository) GetAll() ([]domain.Transaction, error) {
	var tl []domain.Transaction
	if err := r.db.Read(&tl); err != nil {
		return tl, err
	}

	return tl, nil
}

func (r *repository) Store(id int, tranCode, currency string, amount float64, transmitter, receiver, tranDate string) (domain.Transaction, error) {

	var tl []domain.Transaction
	if err := r.db.Read(&tl); err != nil {
		return domain.Transaction{}, err
	}
	t := domain.Transaction{
		Id:          id,
		TranCode:    tranCode,
		Currency:    currency,
		Amount:      amount,
		Transmitter: transmitter,
		Reciever:    receiver,
		TranDate:    tranCode,
	}

	tl = append(tl, t)

	if err := r.db.Write(tl); err != nil {
		return domain.Transaction{}, err
	}

	return t, nil

}

func (r *repository) Update(id int, tranCode, currency string, amount float64, transmitter, receiver, tranDate string) (domain.Transaction, error) {
	var tl []domain.Transaction

	if err := r.db.Read(&tl); err != nil {
		return domain.Transaction{}, err
	}
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
	for i := range tl {
		if tl[i].Id == id {
			t.Id = id
			tl[i] = t
			updated = true
		}
	}

	if !updated {
		return domain.Transaction{}, fmt.Errorf("transaccion %d no encontrada", id)
	}

	if err := r.db.Write(tl); err != nil {
		return domain.Transaction{}, err
	}

	return t, nil
}

func (r *repository) UpdateCodeAndAmount(id int, tranCode string, amount float64) (domain.Transaction, error) {
	var tl []domain.Transaction

	if err := r.db.Read(&tl); err != nil {
		return domain.Transaction{}, err
	}
	updated := false
	var t domain.Transaction
	for i := range tl {
		if tl[i].Id == id {
			tl[i].TranCode = tranCode
			tl[i].Amount = amount
			t = tl[i]
			updated = true
		}
	}

	if !updated {
		return domain.Transaction{}, fmt.Errorf("transaccion %d no encontrada", id)
	}

	if err := r.db.Write(tl); err != nil {
		return domain.Transaction{}, err
	}

	return t, nil
}

func (r *repository) Delete(id int) error {
	var ps []domain.Transaction

	if err := r.db.Read(&ps); err != nil {
		return err
	}

	deleted := false
	var index int
	for i := range ps {
		if ps[i].Id == id {
			index = i
			deleted = true
		}
	}

	if !deleted {
		return fmt.Errorf("err")
	}

	ps = append(ps[:index], ps[index+1:]...)

	if err := r.db.Write(ps); err != nil {
		return fmt.Errorf("error")
	}
	return nil
}
