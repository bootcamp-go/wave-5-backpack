package transactions

import (
	"database/sql"
	"fmt"

	"goweb/internal/domain"
)

const (
	failedReading = "hubo un error de lectura. "
	failedWriting = "hubo un error de escritura. "
)

type IRepository interface {
	Delete(int) error
	GetAll() ([]domain.Transaction, error)
	GetBySender(string) (domain.Transaction, error)
	Store(int, string, string, int, string, string, string) (domain.Transaction, error)
	Update(int, string, string, int, string, string, string) (domain.Transaction, error)
	UpdateAmount(int, int) (domain.Transaction, error)
	LastID() (int, error)
}

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) IRepository {
	return &Repository{
		db: db,
	}
}

func(repository *Repository) Delete(id int) error {
	return nil
}
func(repository *Repository) GetAll() ([]domain.Transaction, error) {
	return []domain.Transaction{}, nil
}

func(repository *Repository) GetBySender(sender string) (domain.Transaction, error) {

	stmt, err := repository.db.Prepare("SELECT id, cod_transaction, currency, amount, sender, receiver, date_order FROM TRANSACTIONS WHERE sender = ?;")
	if err != nil {
		return domain.Transaction{}, fmt.Errorf(err.Error())
	}
	//Cierro la instancia. Si se quedan abiertos se generan consumos de memoria innecesarios.
	defer stmt.Close()

	transaction := domain.Transaction{}
	//Retorna un sql.Return y un error
	err = stmt.QueryRow(sender).Scan(
		&transaction.Id,
		&transaction.CodTransaction,
		&transaction.Currency,
		&transaction.Amount,
		&transaction.Sender,
		&transaction.Receiver,
		&transaction.DateOrder,
	)
	if err != nil {
		return domain.Transaction{}, fmt.Errorf(err.Error())
	}

	return transaction, nil
}

func (repository *Repository) Store(id int, codTransaction string, currency string, amount int, sender string, receiver string, dateOrder string) (domain.Transaction, error) {

	stmt, err := repository.db.Prepare("INSERT INTO TRANSACTIONS (cod_transaction, currency, amount, sender, receiver, date_order) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return domain.Transaction{}, fmt.Errorf(err.Error())
	}
	//Cierro la instancia. Si se quedan abiertos se generan consumos de memoria innecesarios.
	defer stmt.Close()

	var result sql.Result
	//Retorna un sql.Return y un error
	result, err = stmt.Exec(codTransaction, currency, amount, sender, receiver, dateOrder)
	if err != nil {
		return domain.Transaction{}, err
	}
	// Del sql.Return devuelto en la ejecucion se obtiene el id insertado
	insertedId, _ := result.LastInsertId()

	transaction := domain.Transaction{
		Id:             int(insertedId),
		CodTransaction: codTransaction,
		Currency:       currency,
		Amount:         amount,
		Sender:         sender,
		Receiver:       receiver,
		DateOrder:      dateOrder,
	}

	return transaction, nil
}

func (repository *Repository) Update(id int, codTransaction string, currency string, amount int, sender string, receiver string, dateOrder string) (domain.Transaction, error) {
	return domain.Transaction{}, nil
}

func (repository *Repository) UpdateAmount(id, amount int) (domain.Transaction, error) {
	return domain.Transaction{}, nil
}

func (repository *Repository) LastID() (int, error) {
	return 0, nil
}

