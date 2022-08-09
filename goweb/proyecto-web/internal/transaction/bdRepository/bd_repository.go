package bdRepository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"proyecto-web/internal/domain"
	"proyecto-web/internal/transaction/interfaces"
)

type bdRepository struct {
	db *sql.DB
}

const (
	TransactionNotFound = "transaction %d not found"
	FailReading         = "cant read database"
	FailWriting         = "cant write database"
)

var createQuery = `INSERT INTO Transactions(codigo_transaccion, moneda, monto, emisor, receptor, fecha_transaccion) 
				   VALUES(?, ?, ?, ?, ?, ?)`

var getByTransactionCodeQuery = `SELECT id, codigo_transaccion, moneda, monto, emisor, receptor, fecha_transaccion 
								 FROM Transactions WHERE codigo_transaccion = ?`

var getAllQuery = `SELECT id, codigo_transaccion, moneda, monto, emisor, receptor, fecha_transaccion FROM Transactions`

var updateQuery = `UPDATE Transactions set codigo_transaccion = ?, moneda = ?, monto = ?, emisor = ?, receptor = ?, fecha_transaccion = ? WHERE id = ?`

func NewBdRepository(db *sql.DB) interfaces.IRepository {
	return &bdRepository{
		db: db,
	}
}

func (r *bdRepository) Create(codigoTransaccion, moneda string, monto float64, emisor, receptor, fecha string) (domain.Transaction, error) {
	stmt, err := r.db.Prepare(createQuery)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(codigoTransaccion, moneda, monto, emisor, receptor, fecha)
	if err != nil {
		return domain.Transaction{}, err
	}

	lastId, err := result.LastInsertId()

	return domain.Transaction{Id: int(lastId), CodigoTransaccion: codigoTransaccion, Moneda: moneda, Monto: monto, Emisor: emisor, Receptor: receptor, FechaTransaccion: fecha}, nil
}

func (r *bdRepository) GetByCodigoTransaccion(codigo string) (domain.Transaction, error) {
	var transaction domain.Transaction

	rows := r.db.QueryRow(getByTransactionCodeQuery, codigo)

	if err := rows.Scan(&transaction.Id, &transaction.CodigoTransaccion, &transaction.Moneda, &transaction.Monto,
		&transaction.Emisor, &transaction.Receptor, &transaction.FechaTransaccion); err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (r *bdRepository) GetById(id int) (domain.Transaction, error) {
	return domain.Transaction{}, nil
}

func (r *bdRepository) Update(ctx context.Context, id int, codigoTransaccion, moneda string, monto float64, emisor, receptor, fecha string) (domain.Transaction, error) {
	transacciones, err := r.GetAll()
	if err != nil {
		return domain.Transaction{}, err
	}
	_, encontrada := findById(id, transacciones)

	if !encontrada {
		return domain.Transaction{}, fmt.Errorf(TransactionNotFound, id)
	}

	stmt, err := r.db.Prepare(updateQuery)
	defer stmt.Close()
	if err != nil {
		return domain.Transaction{}, err
	}

	_, err = stmt.ExecContext(ctx, codigoTransaccion, moneda, monto, emisor, receptor, fecha, id)
	if err != nil {
		return domain.Transaction{}, err
	}

	return domain.Transaction{Id: id, CodigoTransaccion: codigoTransaccion, Moneda: moneda, Monto: monto,
		Emisor: emisor, Receptor: receptor, FechaTransaccion: fecha}, nil
}

func (r *bdRepository) UpdateParcial(id int, codigoTransaccion string, monto float64) (domain.Transaction, error) {
	return domain.Transaction{}, nil
}

func (r *bdRepository) Delete(id int) error {
	return nil
}

func (r *bdRepository) GetAll() ([]domain.Transaction, error) {
	var transactions []domain.Transaction

	rows, err := r.db.Query(getAllQuery)

	if err != nil {
		return []domain.Transaction{}, err
	}

	for rows.Next() {
		var transaction domain.Transaction
		if err := rows.Scan(&transaction.Id, &transaction.CodigoTransaccion, &transaction.Moneda, &transaction.Monto,
			&transaction.Emisor, &transaction.Receptor, &transaction.FechaTransaccion); err != nil {
			return []domain.Transaction{}, err
		}
		transactions = append(transactions, transaction)
	}
	return transactions, nil
}

func findById(id int, transacciones []domain.Transaction) (*domain.Transaction, bool) {
	var transaccionBuscada *domain.Transaction
	var encontrada bool
	for i, transaccion := range transacciones {
		if transaccion.Id == id {
			transaccionBuscada = &transacciones[i]
			encontrada = true
			break
		}
	}
	return transaccionBuscada, encontrada
}
