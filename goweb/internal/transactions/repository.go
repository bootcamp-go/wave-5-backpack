package transactions

import (
	"fmt"

	"goweb/internal/domain"
	"goweb/pkg/store"
)

const (
	failedReading = "hubo un error de lectura. "
	failedWriting = "hubo un error de escritura. "
)

type IRepository interface {
	Delete(int) error
	GetAll() ([]domain.Transaction, error)
	Store(int, string, string, int, string, string, string) (domain.Transaction, error)
	Update(int, string, string, int, string, string, string) (domain.Transaction, error)
	UpdateAmount(int, int) (domain.Transaction, error)
	LastID() (int, error)
}

type Repository struct {
	db store.Store
}

func NewRepository(db store.Store) IRepository {
	return &Repository{
		db: db,
	}
}

func(repository *Repository) Delete(id int) error {
	var transactions []domain.Transaction

	if err := repository.db.Read(&transactions); err != nil {
		return fmt.Errorf(failedReading)
	}

	deleted := false
	var index int
	//Recorro el arreglo de tipo transaction uno x uno para validar si existe el ID que paso por referencia
	for i := range transactions {
		if transactions[i].Id == id {
			index = i
			deleted = true
		}
	}
	//Si la variable aun sigue falso "instruccion deleted == false" devuelvo un error
	if !deleted {
		return fmt.Errorf("Producto con ID: %d no encontrado", id)
	}
	/*
	Re-estructuro mi slice de la siguiente manera,
	CORTO justo 1 posicion ANTES en la posicion en donde encontre el elemento que coincide con el ID
	y lo SUMO junto con el elemento SIGUIENTE al de la coicidencia en mi ID,
	OMITO POR COMPLETO la posicion del elemento que coincide en el nuevo slice
	*/
	transactions = append(transactions[:index], transactions[index+1:]...)

	if err := repository.db.Write(transactions); err != nil {
		return fmt.Errorf(failedWriting)
	}

	return nil
}
func(repository *Repository) GetAll() ([]domain.Transaction, error) {
	var transactions []domain.Transaction

	if err := repository.db.Read(&transactions); err != nil {
		return []domain.Transaction{}, fmt.Errorf(failedReading, err)
	}

	repository.db.Read(&transactions)
	if len(transactions) == 0 {
		return []domain.Transaction{}, fmt.Errorf("No hay productos agregados, por favor agregar uno")
	}
	return transactions, nil
}

func (repository *Repository) Store(id int, codTransaction string, currency string, amount int, sender string, receiver string, dateOrder string) (domain.Transaction, error) {
	var transactions []domain.Transaction

	if err := repository.db.Read(&transactions); err != nil {
		return domain.Transaction{}, fmt.Errorf(failedReading, err)
	}

	transaction := domain.Transaction{id, codTransaction, currency, amount, sender, receiver, dateOrder}
	transactions = append(transactions, transaction)

	if err := repository.db.Write(transactions); err != nil {
		return domain.Transaction{}, fmt.Errorf(failedWriting, err)
	}
	return transaction, nil
}

func (repository *Repository) Update(id int, codTransaction string, currency string, amount int, sender string, receiver string, dateOrder string) (domain.Transaction, error) {
	var transactions []domain.Transaction
	if err := repository.db.Read(&transactions); err != nil {
		return domain.Transaction{}, fmt.Errorf(failedReading)
	}
	transaction := domain.Transaction{id, codTransaction, currency, amount, sender, receiver, dateOrder}
	updated := false
	for i := range transactions {
		if transactions[i].Id == id {
			transaction.Id = id
			transactions[i] = transaction
			updated = true
		}
	}
	if !updated {
		return domain.Transaction{}, fmt.Errorf("Transaccion con id %d no se ha sido encontrado", id)
	}

	if err := repository.db.Write(transactions); err != nil {
		return domain.Transaction{}, fmt.Errorf(failedWriting, err)
	}

	return transaction, nil
}

func (repository *Repository) UpdateAmount(id, amount int) (domain.Transaction, error) {
	var transaction domain.Transaction
	var transactions []domain.Transaction
	if err := repository.db.Read(&transactions); err != nil {
		return domain.Transaction{}, fmt.Errorf(failedReading)
	}

	updated := false
	for i := range transactions {
		if transactions[i].Id == id {
			transactions[i].Amount = amount
			updated = true
			transaction = transactions[i]
		}
	}
	if !updated {
		return domain.Transaction{}, fmt.Errorf("Transaccion con id %d no se ha sido encontrado", id)
	}

	if err := repository.db.Write(transactions); err != nil {
		return domain.Transaction{}, fmt.Errorf(failedWriting, err)
	}

	return transaction, nil
}

func (repository *Repository) LastID() (int, error) {
	var transactions []domain.Transaction
	if err := repository.db.Read(&transactions); err != nil {
		return 0, err
	}
	if len(transactions) == 0 {
		return 0, nil
	}
	return transactions[len(transactions) - 1].Id, nil
}

