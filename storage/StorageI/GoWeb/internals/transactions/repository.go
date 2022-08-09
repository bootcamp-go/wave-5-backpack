package transactions

import (
	"GoWeb/internals/domain"
	"database/sql"
)

const (
	NotFound    = "id %d not found"
	FailReading = "cant read database"
	FailWriting = "cant write database, error %w"
)

type Repository interface {
	GetAll() ([]domain.Transanction, error)
	Store(code, coin string, amount float64, emisor, receptor, date string) (domain.Transanction, error)
	Update(id int, code, coin string, amount float64, emisor, receptor, date string) (domain.Transanction, error)
	UpdateCode(id int, code string, amount float64) (domain.Transanction, error)
	Delete(id int) (domain.Transanction, error)
	GetById(id int) (domain.Transanction, error)
	GetByName(name string) ([]domain.Transanction, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll() ([]domain.Transanction, error) {
	var tran []domain.Transanction
	db := r.db
	rows, err := db.Query("SELECT * FROM transactions")
	if err != nil {
		return []domain.Transanction{}, nil
	}
	for rows.Next() {
		var tra domain.Transanction
		if err := rows.Scan(&tra.Id, &tra.Code, &tra.Coin, &tra.Amount, &tra.Emisor, &tra.Receptor, &tra.Date); err != nil {
			return []domain.Transanction{}, nil
		}
		tran = append(tran, tra)
	}

	return tran, nil
}

func (r *repository) Store(code, coin string, amount float64, emisor, receptor, date string) (domain.Transanction, error) {

	db := r.db
	stmt, err := db.Prepare("INSERT INTO transactions(code,coin,amount,emisor,receptor,date) VALUES(?,?,?,?,?,?)")

	if err != nil {
		return domain.Transanction{}, nil
	}
	defer stmt.Close()
	var result sql.Result
	result, err = stmt.Exec(code, coin, amount, emisor, receptor, date)
	if err != nil {
		return domain.Transanction{}, nil
	}
	insertId, _ := result.LastInsertId()
	id := int(insertId)
	transaction := domain.Transanction{Id: id, Code: code, Coin: coin, Amount: amount, Emisor: emisor, Receptor: receptor, Date: date}
	return transaction, nil
}

func (r *repository) Update(id int, code, coin string, amount float64, emisor, receptor, date string) (domain.Transanction, error) {
	db := r.db
	stmt, err := db.Prepare("UPDATE transactions SET code=?, coin =?, amount =?, emisor=?, receptor=?, date=? WHERE id=?")

	if err != nil {
		return domain.Transanction{}, nil
	}
	defer stmt.Close()
	_, err = stmt.Exec(code, coin, amount, emisor, receptor, date, id)
	if err != nil {
		return domain.Transanction{}, nil
	}
	tran := domain.Transanction{Id: id, Code: code, Coin: coin, Amount: amount, Emisor: emisor, Receptor: receptor, Date: date}
	return tran, nil
}

func (r *repository) Delete(id int) (domain.Transanction, error) {
	db := r.db
	stmt, err := db.Prepare("DELETE FROM transactions WHERE id=?")
	if err != nil {
		return domain.Transanction{}, nil
	}
	defer db.Close()
	tran, err := r.GetById(id)
	if err != nil {
		return domain.Transanction{}, nil
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return domain.Transanction{}, nil
	}

	return tran, nil
}

func (r *repository) UpdateCode(id int, code string, amount float64) (domain.Transanction, error) {

	db := r.db
	stmt, err := db.Prepare("UPDATE transactions SET code=?,amount=? WHERE id=?")

	if err != nil {
		return domain.Transanction{}, nil
	}
	defer stmt.Close()
	_, err = stmt.Exec(code, amount, id)
	if err != nil {
		return domain.Transanction{}, nil
	}
	tran, err := r.GetById(id)
	if err != nil {
		return domain.Transanction{}, nil
	}

	return tran, nil
}

func (r *repository) GetById(id int) (domain.Transanction, error) {
	var tran domain.Transanction
	db := r.db

	rows, err := db.Query("SELECT * FROM transactions WHERE id =?", id)
	if err != nil {
		return domain.Transanction{}, nil
	}
	for rows.Next() {
		if err := rows.Scan(&tran.Id, &tran.Code, &tran.Coin, &tran.Amount, &tran.Emisor, &tran.Receptor, &tran.Date); err != nil {
			return domain.Transanction{}, nil
		}
	}
	return tran, nil
}

func (r *repository) GetByName(name string) ([]domain.Transanction, error) {
	var tran []domain.Transanction
	db := r.db

	rows, err := db.Query("SELECT * FROM transactions WHERE name = ?", name)
	if err != nil {
		return []domain.Transanction{}, nil
	}
	for rows.Next() {
		var tra domain.Transanction
		if err := rows.Scan(&tra.Id, &tra.Code, &tra.Coin, &tra.Amount, &tra.Emisor, &tra.Receptor, &tra.Date); err != nil {
			return []domain.Transanction{}, nil
		}
		tran = append(tran, tra)
	}
	return tran, nil
}
