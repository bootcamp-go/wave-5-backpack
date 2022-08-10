package transactions

import (
	"context"
	"database/sql"
	"goweb/internal/domain"
	"time"
)

var (
	insert           = "INSERT INTO transactions(transaction_code, currency, amount, sender, reciever, transaction_date) VALUES (?,?,?,?,?,?)"
	selectByCurrency = "SELECT id, transaction_code, currency, amount, sender, reciever, transaction_date FROM transactions WHERE currency = ?"
	selectAll        = "SELECT id, transaction_code, currency, amount, sender, reciever, transaction_date FROM transactions"
	update           = "UPDATE transactions SET currency = ?, amount = ?, sender = ?, reciever = ? WHERE id = ?"
	selectById       = "SELECT id,transaction_code,currency, amount,sender,reciever,transaction_date FROM transactions WHERE id = ?"
	deleteById       = "DELETE FROM transactions WHERE id = ?"
)

type Repository interface {
	GetAll() ([]domain.Transaction, error)
	Store(TransactionCode string, Currency string, Amount float64, Sender string, Reciever string, TransactionDate time.Time) (domain.Transaction, error)
	GetById(Id int) (domain.Transaction, error)
	GetByCurrency(Currency string) ([]domain.Transaction, error)
	lastId() (int, error)
	Update(id int, Currency string, Amount float64, Sender string, Reciever string) (domain.Transaction, error)
	Delete(id int) error
	UpdateCurrencyAndAmount(id int, Currency string, Amount float64) (domain.Transaction, error)
	UpdateWithContext(ctx context.Context, id int, Currency string, Amount float64, Sender string, Reciever string) (domain.Transaction, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {

	return &repository{
		db,
	}
}

func (r *repository) GetAll() ([]domain.Transaction, error) {
	var transactions = []domain.Transaction{}
	rows, err := r.db.Query(selectAll)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var t domain.Transaction
		if err := rows.Scan(&t.Id, &t.TransactionCode, &t.Currency, &t.Amount, &t.Sender, &t.Reciever, &t.TransactionDate); err != nil {
			return nil, err
		}
		transactions = append(transactions, t)
	}
	return transactions, nil
}
func (r *repository) Store(TransactionCode string, Currency string, Amount float64, Sender string, Reciever string, TransactionDate time.Time) (domain.Transaction, error) {
	t := domain.Transaction{
		TransactionCode: TransactionCode,
		Currency:        Currency,
		Amount:          Amount,
		Sender:          Sender,
		Reciever:        Reciever,
		TransactionDate: TransactionDate,
	}
	stmt, err := r.db.Prepare(insert)

	if err != nil {
		return domain.Transaction{}, err
	}
	defer stmt.Close()
	result, err := stmt.Exec(t.TransactionCode, t.Currency, t.Amount, t.Sender, t.Reciever, t.TransactionDate)
	if err != nil {
		return domain.Transaction{}, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return domain.Transaction{}, err
	}
	t.Id = int(id)
	return t, nil
}

func (r *repository) GetByCurrency(Currency string) ([]domain.Transaction, error) {
	var transactions = []domain.Transaction{}
	rows, err := r.db.Query(selectByCurrency, Currency)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var t domain.Transaction
		if err := rows.Scan(&t.Id, &t.TransactionCode, &t.Currency, &t.Amount, &t.Sender, &t.Reciever, &t.TransactionDate); err != nil {
			return nil, err
		}
		transactions = append(transactions, t)
	}
	return transactions, nil
}

func (r *repository) UpdateWithContext(ctx context.Context, id int, Currency string, Amount float64, Sender string, Reciever string) (domain.Transaction, error) {
	t := domain.Transaction{
		Currency: Currency,
		Amount:   Amount,
		Sender:   Sender,
		Reciever: Reciever,
	}

	stmt, err := r.db.PrepareContext(ctx, update)
	if err != nil {
		return domain.Transaction{}, err
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, Currency, Amount, Sender, Reciever, id)
	if err != nil {
		return domain.Transaction{}, err
	}
	return t, nil
}

func (r *repository) Update(id int, Currency string, Amount float64, Sender string, Reciever string) (domain.Transaction, error) {
	t := domain.Transaction{
		Currency: Currency,
		Amount:   Amount,
		Sender:   Sender,
		Reciever: Reciever,
	}
	stmt, err := r.db.Prepare(update)
	if err != nil {
		return domain.Transaction{}, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(Currency, Amount, Sender, Reciever, id)
	if err != nil {
		return domain.Transaction{}, err
	}
	return t, nil
}

func (r *repository) GetById(searchId int) (domain.Transaction, error) {
	var t domain.Transaction
	row := r.db.QueryRow(selectById, searchId)
	if err := row.Scan(&t.Id, &t.TransactionCode, &t.Currency, &t.Amount, &t.Sender, &t.Reciever, &t.TransactionDate); err != nil {
		return domain.Transaction{}, err
	}
	return t, nil
}
func (r *repository) lastId() (int, error) {
	return 1, nil
}

func (r *repository) Delete(id int) error {
	stmt, err := r.db.Prepare(deleteById)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) UpdateCurrencyAndAmount(id int, Currency string, Amount float64) (domain.Transaction, error) {
	return domain.Transaction{}, nil
}
