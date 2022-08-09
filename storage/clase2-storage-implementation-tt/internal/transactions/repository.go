package transactions

import (
	"context"
	"database/sql"

	"clase2-storage-implementation-tt/internal/domain"
)

//	Constants
const (
	//	ERRORS 			= messages
	TransactionNotFound = "transaction with id: %d, not found 😵‍💫"
	FailReading         = "cant read database 🫠"
	FailWriting         = "cant write database 😱, error: %w"

	//	Query	string = message
	GetAllTransactions  = "SELECT id, codeTransaction, currency, amount, transmitter, receiver, date FROM transactions"
	GetTransaction      = "SELECT codeTransaction, currency, amount, transmitter, receiver, date FROM transactions WHERE id = ?"
	GetTransactionByID  = "SELECT * FROM transactions WHERE id = ?"
	GetCodeTransaction  = "SELECT * FROM transactions WHERE codeTransaction = ?"
	GetTransactionSleep = "SELECT SLEEP(5) FROM DUAL WHERE 0 < ?"

	InsertTransaction = "INSERT INTO transactions(codeTransaction, currency, amount, transmitter, receiver, date) VALUES( ?, ?, ?, ?, ?, ? )"
	UpdateTransaction = "UPDATE transactions SET codeTransaction = ?, currency = ?, amount = ?, transmitter = ?, receiver = ?, date = ? WHERE id = ?"
	DeleteTransaction = "DELETE FROM transactions WHERE id = ?"
)

// Repository ...
type Repository interface {
	GetAll(ctx context.Context) ([]domain.Transaction, error)
	Ecommerce(transaction domain.Transaction) (domain.Transaction, error)
	GetOne(id int) (domain.Transaction, error)
	GetByName(code string) ([]domain.Transaction, error)
	GetOneWithContext(ctx context.Context, id int) (domain.Transaction, error)
	Update(ctx context.Context, id int, codeTra, currency string, amount float64, transmitter,
		receiver, date string) (domain.Transaction, error)
	UpdateOne(id int, codeTra string, monto float64) (domain.Transaction, error)
	Delete(id int) (domain.Transaction, error)
}

type repository struct {
	db *sql.DB
}

// NewRepository ...
func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Ecommerce(transaction domain.Transaction) (domain.Transaction, error) {

	stmt, err := r.db.Prepare(InsertTransaction) // se prepara la sentencia SQL a ejecutar
	if err != nil {
		return domain.Transaction{}, err
	}
	defer stmt.Close() // se cierra la sentencia al terminar. Si quedan abiertas se genera consumos de memoria
	var result sql.Result

	result, err = stmt.Exec(transaction.CodigoTransaccion, transaction.Moneda, transaction.Monto, transaction.Emisor,
		transaction.Receptor, transaction.Fecha) // retorna un sql.Result y un error
	if err != nil {
		return domain.Transaction{}, err
	}

	insertedID, _ := result.LastInsertId() // del sql.Resul devuelto en la ejecucion obtenemos el Id insertado
	transaction.ID = int(insertedID)
	return transaction, nil
}

func (r *repository) GetAll(ctx context.Context) ([]domain.Transaction, error) {
	var transactions []domain.Transaction
	db := r.db
	rows, err := db.QueryContext(ctx, GetAllTransactions)
	if err != nil {
		return []domain.Transaction{}, err
	}
	for rows.Next() {
		var transaction domain.Transaction
		if err := rows.Scan(&transaction.ID, &transaction.CodigoTransaccion, &transaction.Moneda, &transaction.Monto,
			&transaction.Emisor, &transaction.Receptor, &transaction.Fecha); err != nil {
			return []domain.Transaction{}, err
		}
		transactions = append(transactions, transaction)
	}
	return transactions, nil
}

func (r *repository) GetByName(code string) ([]domain.Transaction, error) {
	var transactions []domain.Transaction
	db := r.db
	rows, err := db.Query(GetCodeTransaction, code)
	if err != nil {
		return []domain.Transaction{}, err
	}
	for rows.Next() {
		var transaction domain.Transaction
		if err := rows.Scan(&transaction.ID, &transaction.CodigoTransaccion, &transaction.Moneda, &transaction.Monto, &transaction.Emisor,
			&transaction.Receptor, &transaction.Fecha); err != nil {
			return []domain.Transaction{}, err
		}
		transactions = append(transactions, transaction)
	}
	return transactions, nil
}

func (r *repository) GetOne(id int) (domain.Transaction, error) {

	var transaction domain.Transaction

	rows, err := r.db.Query(GetTransaction, id)
	if err != nil {
		return domain.Transaction{}, err
	}
	for rows.Next() {
		if err := rows.Scan(&transaction.CodigoTransaccion, &transaction.Moneda, &transaction.Monto, &transaction.Emisor,
			&transaction.Receptor, &transaction.Fecha); err != nil {
			return domain.Transaction{}, err
		}
	}
	return transaction, nil
}

func (r *repository) Update(ctx context.Context, id int, codeTra, currency string, amount float64, transmitter,
	receiver, date string) (domain.Transaction, error) {

	stmt, err := r.db.Prepare(UpdateTransaction) // se prepara la sentencia SQL a ejecutar
	if err != nil {
		return domain.Transaction{}, err
	}
	defer stmt.Close() // se cierra la sentencia al terminar. Si quedan abiertas se genera consumos de memoria
	transaction := domain.Transaction{ID: id, CodigoTransaccion: codeTra, Moneda: currency, Monto: amount,
		Emisor: transmitter, Receptor: receiver, Fecha: date}

	_, err = stmt.ExecContext(ctx, codeTra, currency, amount, transmitter, receiver, date, id)
	if err != nil {
		return domain.Transaction{}, err
	}
	return transaction, nil
}

func (r *repository) UpdateOne(id int, codeTra string, monto float64) (domain.Transaction, error) {

	stmt, err := r.db.Prepare("UPDATE transactions SET codeTransaction = ?, amount = ? WHERE id = ?")
	if err != nil {
		return domain.Transaction{}, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(codeTra, monto, id)
	if err != nil {
		return domain.Transaction{}, err
	}

	transaction, err := r.GetOne(id)
	if err != nil {
		return domain.Transaction{}, err
	}

	return transaction, nil
}

func (r *repository) Delete(id int) (domain.Transaction, error) {
	db := r.db

	stmt, err := db.Prepare(DeleteTransaction)
	if err != nil {
		return domain.Transaction{}, err
	}
	defer stmt.Close()
	product, err := r.GetOne(id)
	if err != nil {
		return domain.Transaction{}, err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return domain.Transaction{}, err
	}
	return product, nil
}

func (r *repository) GetOneWithContext(ctx context.Context, id int) (domain.Transaction, error) {
	var transaction domain.Transaction
	rows, err := r.db.QueryContext(ctx, GetTransactionSleep, id)
	if err != nil {
		return transaction, err
	}
	for rows.Next() {
		if err := rows.Scan(&transaction.ID); err != nil {
			return transaction, err
		}
	}
	return transaction, nil
}
