package transactions

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"goweb/clase3-go-web-tm/internal/domain"
)

//	GLOBAL : Variables
var ps []domain.Transaction
var lastID int

//	INTERFACE : Repository
type Repository interface {
	GetAll() ([]domain.Transaction, error)
	Ecommerce(id int, codeTra string, coin string, monto float64, emisor string,
		receptor string, fecha string) (domain.Transaction, error)
	LastID() (int, error)
	GetOne(id int) (domain.Transaction, error)
	Update(id int, codeTra string, coin string, monto float64, emisor string,
		receptor string, fecha string) (domain.Transaction, error)
	UpdateOne(id int, codeTra string, monto float64) (domain.Transaction, error)
	Delete(id int) error
}

type repository struct{}

//	FUNCTIONS
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

func (r *repository) Update(id int, codeTra string, coin string, monto float64, emisor string,
	receptor string, fecha string) (domain.Transaction, error) {
	update := false
	tmpT := domain.Transaction{
		CodigoTransaccion: codeTra,
		Moneda:            coin,
		Monto:             monto,
		Emisor:            emisor,
		Receptor:          receptor,
		Fecha:             fecha,
	}

	for i := range ps {
		if ps[i].Id == id {
			tmpT.Id = id
			ps[i] = tmpT
			update = true
			break
		}
	}

	if !update {
		return domain.Transaction{}, fmt.Errorf("Transaccion con id <%d> no fue encontrado 😵‍💫", id)
	}

	return tmpT, nil
}

func (r *repository) UpdateOne(id int, codeTra string, monto float64) (domain.Transaction, error) {
	var t domain.Transaction
	update := false

	for i := range ps {
		if ps[i].Id == id {
			ps[i].CodigoTransaccion = codeTra
			ps[i].Monto = monto
			update = true
			t = ps[i]
			break
		}
	}

	if !update {
		return domain.Transaction{}, fmt.Errorf("Actualizacion con id <%d> no fue encontrado 😵‍💫", id)
	}

	return t, nil
}

func (r *repository) Delete(id int) error {
	deleted := false
	var index int
	for i := range ps {
		if ps[i].Id == id {
			index = i
			deleted = true
			break
		}
	}
	if !deleted {
		return fmt.Errorf("Eliminacion con id <%d> no fue encontrado 😵‍💫", id)
	}

	ps = append(ps[:index], ps[index+1:]...)

	return nil
}
