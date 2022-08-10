package transactions

import (
	"context"
	"database/sql"
	"fmt"

	"goweb/internal/domain"
)

const (
	GetAllTransactions string = "SELECT id, cod_transaction, currency, amount, sender, receiver, date_order FROM TRANSACTIONS"
	GetTransactionBySender string = "SELECT id, cod_transaction, currency, amount, sender, receiver, date_order FROM TRANSACTIONS WHERE sender = ?;"
    InsertTransaction string = "INSERT INTO TRANSACTIONS (cod_transaction, currency, amount, sender, receiver, date_order) VALUES (?, ?, ?, ?, ?, ?)"
)

type IRepository interface {
	Delete(context.Context, int) error
	GetAll(context.Context) ([]domain.Transaction, error)
	GetBySender(context.Context, string) (domain.Transaction, error)
	Store(context.Context, int, string, string, int, string, string, string) (domain.Transaction, error)
	Update(context.Context, int, string, string, int, string, string, string) (domain.Transaction, error)
	UpdateAmount(context.Context, int, int) (domain.Transaction, error)
	LastID(context.Context) (int, error)
}

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) IRepository {
	return &Repository{
		db: db,
	}
}

func(repository *Repository) Delete(ctx context.Context, id int) error {
	return nil
}
func(repository *Repository) GetAll(ctx context.Context) ([]domain.Transaction, error) {

	rows, err := repository.db.Query(GetAllTransactions)
	if err != nil {
		return []domain.Transaction{}, fmt.Errorf(err.Error())
	}
	var transactions []domain.Transaction

	for rows.Next() {
		transaction := domain.Transaction{}
		if err := rows.Scan(&transaction.Id, &transaction.CodTransaction, &transaction.Currency, &transaction.Amount, &transaction.Sender, &transaction.Receiver, &transaction.DateOrder); err != nil {
			return nil, err
		}
		transactions =append(transactions, transaction)
	}

	return transactions, nil
}

func(repository *Repository) GetBySender(ctx context.Context, sender string) (domain.Transaction, error) {

	stmt, err := repository.db.Prepare(GetTransactionBySender)
	if err != nil {
		return domain.Transaction{}, fmt.Errorf(err.Error())
	}
	//Cierro la instancia. Si se quedan abiertos se generan consumos de memoria innecesarios.
	defer stmt.Close()

	transaction := domain.Transaction{}

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

func (repository *Repository) Store(ctx context.Context, id int, codTransaction string, currency string, amount int, sender string, receiver string, dateOrder string) (domain.Transaction, error) {

	stmt, err := repository.db.Prepare(InsertTransaction)
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

func (repository *Repository) Update(ctx context.Context, id int, codTransaction string, currency string, amount int, sender string, receiver string, dateOrder string) (domain.Transaction, error) {
	return domain.Transaction{}, nil
}

func (repository *Repository) UpdateAmount(ctx context.Context, id, amount int) (domain.Transaction, error) {
	return domain.Transaction{}, nil
}

func (repository *Repository) LastID(ctx context.Context) (int, error) {
	return 0, nil
}

