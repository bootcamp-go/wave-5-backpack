package transactions

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"goweb/clase2-go-web-tt/internal/domain"
)

type Repository interface {
	GetAll() ([]domain.Transaction, error)
	Ecommerce(id int, codeTra string, coin string, monto float64, emisor string,
		receptor string, fecha string) (domain.Transaction, error)
	LastID() (int, error)
	GetOne(id int) (domain.Transaction, error)
}

type repository struct{}

func getData() ([]domain.Transaction, error) {
	var dataTransacciones []domain.Transaction
	file, err := os.ReadFile("../../transacciones.json")
	if err != nil {
		return dataTransacciones, err
	}
	if err := json.Unmarshal(file, &dataTransacciones); err != nil {
		return dataTransacciones, err
	}
	return dataTransacciones, nil
}

func NewRepository() Repository {
	dataFile, err := getData()
	if err != nil {
		fmt.Println(err)
	}

	maxId := 0
	for _, i := range dataFile {
		ps = append(ps, i) // Append Data
		if i.Id > maxId {  //Get the last ID integer
			maxId = i.Id
		}
	}
	lastID = maxId

	return &repository{}
}

var ps []domain.Transaction
var lastID int

func (r *repository) GetAll() ([]domain.Transaction, error) {
	return ps, nil
}

func (r *repository) LastID() (int, error) {
	return lastID, nil
}

func (r *repository) Ecommerce(id int, codeTra string, coin string, monto float64,
	emisor string, receptor string, fecha string) (domain.Transaction, error) {
	t := domain.Transaction{Id: id, CodigoTransaccion: codeTra, Moneda: coin, Monto: monto,
		Emisor: emisor, Receptor: receptor, Fecha: fecha}
	ps = append(ps, t)
	lastID = t.Id
	return t, nil
}

func (r *repository) GetOne(id int) (domain.Transaction, error) {
	for _, transaction := range ps {
		if id == transaction.Id {
			return transaction, nil
		}
	}
	return domain.Transaction{}, errors.New("> error. No hay ninguna transaccion con el Id ingresado")
}
