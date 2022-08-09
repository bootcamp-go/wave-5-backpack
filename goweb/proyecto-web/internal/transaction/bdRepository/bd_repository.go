package bdRepository

import (
	"database/sql"
	"log"
	"proyecto-web/internal/domain"
	"proyecto-web/internal/transaction/interfaces"
)

type bdRepository struct {
	db *sql.DB
}

func NewBdRepository(db *sql.DB) interfaces.IRepository {
	return &bdRepository{
		db: db,
	}
}

func (r *bdRepository) Create(codigoTransaccion, moneda string, monto float64, emisor, receptor, fecha string) (domain.Transaction, error) {
	stmt, err := r.db.Prepare(`INSERT INTO Transactions(codigo_transaccion, moneda, monto, emisor, receptor, fecha_transaccion) 
							   VALUES(?, ?, ?, ?, ?, ?)`)
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

	rows := r.db.QueryRow("SELECT * FROM Transactions WHERE codigo_transaccion = ?", codigo)

	if err := rows.Scan(&transaction.Id, &transaction.CodigoTransaccion, &transaction.Moneda, &transaction.Monto,
		&transaction.Emisor, &transaction.Receptor, &transaction.FechaTransaccion); err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (r *bdRepository) GetById(id int) (domain.Transaction, error) {
	return domain.Transaction{}, nil
}

func (r *bdRepository) Update(id int, codigoTransaccion, moneda string, monto float64, emisor, receptor, fecha string) (domain.Transaction, error) {
	return domain.Transaction{}, nil
}

func (r *bdRepository) UpdateParcial(id int, codigoTransaccion string, monto float64) (domain.Transaction, error) {
	return domain.Transaction{}, nil
}

func (r *bdRepository) Delete(id int) error {
	return nil
}

func (r *bdRepository) GetAll() ([]domain.Transaction, error) {
	return nil, nil
}
