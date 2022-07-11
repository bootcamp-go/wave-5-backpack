package transactions

import "fmt"

type Repository interface {
	LastId() (int64, error)
	GetAll() ([]Transaction, error)
	Store(id int64, codigo, moneda, emisor, receptor string, monto float64) (Transaction, error)
	Update(id int64, monto float64, codigo, emisor, receptor, moneda string) (Transaction, error)
	UpdateReceptorYMonto(id int64, receptor string, monto float64) (Transaction, error)
	Delete(id int64) error
}

type Transaction struct {
	Id       int64
	Codigo   string
	Monto    float64
	Moneda   string
	Emisor   string
	Receptor string
}

var transactions []Transaction
var lastId int64

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAll() ([]Transaction, error) {
	return transactions, nil
}

func (r *repository) Store(id int64, codigo, moneda, emisor, receptor string, monto float64) (Transaction, error) {
	transaction := Transaction{
		Id:       id,
		Codigo:   codigo,
		Moneda:   moneda,
		Monto:    monto,
		Emisor:   emisor,
		Receptor: receptor,
	}

	transactions = append(transactions, transaction)
	lastId = transaction.Id
	return transaction, nil
}

func (r *repository) LastId() (int64, error) {
	return lastId, nil
}

func (r *repository) Update(id int64, monto float64, codigo, emisor, receptor, moneda string) (Transaction, error) {
	t := Transaction{
		Monto:    monto,
		Codigo:   codigo,
		Emisor:   emisor,
		Receptor: receptor,
		Moneda:   moneda,
	}
	updated := false
	for i := range transactions {
		if transactions[i].Id == id {
			t.Id = id
			transactions[i] = t
			updated = true
		}
	}
	if !updated {
		return Transaction{}, fmt.Errorf("transaccion %d no encontrada", id)
	}
	return t, nil
}

func (r *repository) Delete(id int64) error {
	deleted := false
	var index int
	for value := range transactions {
		index = value
		deleted = true
	}
	if !deleted {
		return fmt.Errorf("la transacción id=%d no existe", id)
	}
	transactions = append(transactions[:index], transactions[index+1:]...)
	return nil
}

func (r *repository) UpdateReceptorYMonto(id int64, receptor string, monto float64) (Transaction, error) {
	update := false
	var transaction Transaction
	for value := range transactions {
		if transactions[value].Id == id {
			transactions[value].Receptor = receptor
			transactions[value].Monto = monto
			transaction = transactions[value]
			update = true
		}
	}
	if !update {
		return Transaction{}, fmt.Errorf("transaccion id=%d no encontrada", id)
	}
	return transaction, nil
}
