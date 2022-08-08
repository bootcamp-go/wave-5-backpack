package product

import (
	"database/sql"
	"errors"
	"storage/internal/domain"
)

type Repository interface {
	GetByName(name string) (domain.Product, error)
	Store(product domain.Product) (domain.Product, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetByName(name string) (domain.Product, error) {
	var product domain.Product
	rows, err := r.db.Query("SELECT id, name, type, count, price FROM products WHERE name = ?", name)
	if err != nil {
		return product, err
	}
	for rows.Next() {
		if err := rows.Scan(&product.ID, &product.Name, &product.Type, &product.Count, &product.Price); err != nil {
			return domain.Product{}, err
		}
	}
	if product.ID == 0 {
		return domain.Product{}, errors.New("not found")
	}
	return product, nil
}

func (r *repository) Store(product domain.Product) (domain.Product, error) {
	stmt, err := r.db.Prepare("INSERT INTO products(name, type, count, price) VALUES(?, ?, ?, ?)")
	if err != nil {
		return domain.Product{}, err
	}
	defer stmt.Close()
	result, err := stmt.Exec(product.Name, product.Type, product.Count, product.Price)
	if err != nil {
		return domain.Product{}, err
	}
	id, _ := result.LastInsertId()
	product.ID = int(id)
	return product, nil
}
