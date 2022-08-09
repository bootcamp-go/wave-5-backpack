package transactions

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/bootcamp-go/wave-5-backpack/tree/lopez_cristian/storage/internal/models"
	"github.com/gin-gonic/gin"
)

type Repository interface {
	Store(ctx *gin.Context, monto float64, cod, moneda, emisor, receptor string) (models.Transaction, error)
	GetByCod(ctx *gin.Context, cod string) (models.Transaction, error)
	GetByID(ctx *gin.Context, id int) (models.Transaction, error)
	GetAll(ctx *gin.Context) ([]models.Transaction, error)
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
)

func (r *repository) Store(ctx *gin.Context, monto float64, cod, moneda, emisor, receptor string) (models.Transaction, error) {
	stmt, err := r.db.Prepare(queryStore)
	if err != nil {
		return models.Transaction{}, err
	}

	transaction := models.Transaction{
		Monto:    monto,
		Cod:      cod,
		Emisor:   emisor,
		Receptor: receptor,
		Fecha:    time.Now().Format("2006-01-01"),
	}

	res, err := stmt.Exec(monto, cod, moneda, emisor, receptor, transaction.Fecha)
	if err != nil {
		return models.Transaction{}, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return models.Transaction{}, err
	}

	transaction.ID = int(id)

	return transaction, nil
}

func (r *repository) GetByCod(ctx *gin.Context, cod string) (models.Transaction, error) {
	rows, err := r.db.Query(queryGetByCod, cod)
	if err != nil {
		return models.Transaction{}, err
	}

	var transaction models.Transaction
	for rows.Next() {
		if err := rows.Scan(&transaction); err != nil {
			return models.Transaction{}, err
		}
	}

	return transaction, nil
}

func (r *repository) GetByID(ctx *gin.Context, id int) (models.Transaction, error) {
	rows, err := r.db.Query(queryGetByID, id)
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

func (r *repository) GetAll(ctx *gin.Context) ([]models.Transaction, error) {
	rows, err := r.db.Query(queryGetAll)
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
