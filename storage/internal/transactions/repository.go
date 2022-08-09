package transactions

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/bootcamp-go/wave-5-backpack/tree/lopez_cristian/storage/internal/models"
)

type Repository interface {
	Store(ctx context.Context, monto float64, cod, moneda, emisor, receptor string) (models.Transaction, error)
	GetByCod(ctx context.Context, cod string) (models.Transaction, error)
	GetByID(ctx context.Context, id int) (models.Transaction, error)
	GetAll(ctx context.Context) ([]models.Transaction, error)
	Update(ctx context.Context, id int, monto float64, cod, moneda, emisor, receptor string) (models.Transaction, error)
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db}
}

type repository struct {
	db *sql.DB
}

const (
	queryStore    = `INSERT INTO transactions (monto, cod, moneda, emisor, receptor, fecha) VALUES (?, ?, ?, ?, ?, ?);`
	queryGetByCod = `SELECT id, monto, cod, moneda, emisor, receptor, fecha FROM transactions WHERE cod = ?;`
	queryGetByID  = `SELECT id, monto, cod, moneda, emisor, receptor, fecha FROM transactions WHERE id = ?;`
	queryGetAll   = `SELECT id, monto, cod, moneda, emisor, receptor, fecha FROM transactions;`
	queryUpdate   = `UPDATE transactions SET monto = ?, cod = ?, moneda = ?, emisor = ?, receptor = ? WHERE id = ?;`
)

func (r *repository) Store(ctx context.Context, monto float64, cod, moneda, emisor, receptor string) (models.Transaction, error) {
	stmt, err := r.db.Prepare(queryStore)
	if err != nil {
		return models.Transaction{}, err
	}
	defer stmt.Close()

	t := models.Transaction{
		Monto:    monto,
		Cod:      cod,
		Moneda:   moneda,
		Emisor:   emisor,
		Receptor: receptor,
		Fecha:    time.Now().Format("2006-01-01"),
	}

	res, err := stmt.ExecContext(ctx, monto, cod, moneda, emisor, receptor, t.Fecha)
	if err != nil {
		return models.Transaction{}, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return models.Transaction{}, err
	}

	t.ID = int(id)

	return t, nil
}

func (r *repository) GetByCod(ctx context.Context, cod string) (models.Transaction, error) {
	rows, err := r.db.QueryContext(ctx, queryGetByCod, cod)
	if err != nil {
		return models.Transaction{}, err
	}

	var t models.Transaction
	for rows.Next() {
		if err := rows.Scan(&t); err != nil {
			return models.Transaction{}, err
		}
	}

	return t, nil
}

func (r *repository) GetByID(ctx context.Context, id int) (models.Transaction, error) {
	rows, err := r.db.QueryContext(ctx, queryGetByID, id)
	if err != nil {
		return models.Transaction{}, err
	}

	var t models.Transaction
	for rows.Next() {
		if err := rows.Scan(&t.ID, &t.Monto, &t.Cod, &t.Moneda, &t.Emisor, &t.Receptor, &t.Fecha); err != nil {
			return models.Transaction{}, err
		}
	}

	if t == (models.Transaction{}) {
		return models.Transaction{}, fmt.Errorf("transaction by ID %v not found", id)
	}

	return t, nil
}

func (r *repository) GetAll(ctx context.Context) ([]models.Transaction, error) {
	rows, err := r.db.QueryContext(ctx, queryGetAll)
	if err != nil {
		return []models.Transaction{}, err
	}

	var transactions []models.Transaction
	for rows.Next() {
		var t models.Transaction
		if err := rows.Scan(&t.ID, &t.Monto, &t.Cod, &t.Moneda, &t.Emisor, &t.Receptor, &t.Fecha); err != nil {
			return []models.Transaction{}, err
		}
		transactions = append(transactions, t)
	}

	return transactions, nil
}

func (r *repository) Update(ctx context.Context, id int, monto float64, cod, moneda, emisor, receptor string) (models.Transaction, error) {
	stmt, err := r.db.PrepareContext(ctx, queryUpdate)
	if err != nil {
		return models.Transaction{}, err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, monto, cod, moneda, emisor, receptor, id)
	if err != nil {
		return models.Transaction{}, err
	}

	t := models.Transaction{
		ID:       id,
		Monto:    monto,
		Cod:      cod,
		Moneda:   moneda,
		Emisor:   emisor,
		Receptor: receptor,
	}

	return t, nil
}
