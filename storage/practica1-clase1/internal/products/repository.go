package products

import (
	"database/sql"
	"practica1-clase1/internal/domain"
)

type Repository interface {
	GetByName(name string) (domain.Product, error)
	Store(product domain.Product) (domain.Product, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Store(product domain.Product) (domain.Product, error) {

	stmt, err := r.db.Prepare("INSERT INTO products (name, type, count, price) VALUES (?, ?, ?, ?)")
	if err != nil {
		return domain.Product{}, err
	}
	defer stmt.Close() // se cierra la sentencia al terminar. Si quedan abiertas se genera consumos de memoria

	var result sql.Result
	result, err = stmt.Exec(product.Name, product.Type, product.Count, product.Price)
	if err != nil {
		return domain.Product{}, err
	}
	insertedId, _ := result.LastInsertId()
	product.ID = int(insertedId)

	return product, nil
}

func (r *repository) GetByName(name string) (domain.Product, error) {
	rows, err := r.db.Query("SELECT * FROM products WHERE name = ?", name)
	if err != nil {
		return domain.Product{}, err
	}

	var product domain.Product
	for rows.Next() {
		if err := rows.Scan(&product.ID, &product.Name, &product.Type, &product.Count, &product.Price); err != nil {
			return product, nil
		}
	}
	return product, nil
}
