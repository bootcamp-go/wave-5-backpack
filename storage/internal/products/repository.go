package products

import (
	"database/sql"

	"github.com/bootcamp-go/wave-5-backpack/storage/internal/domain"
)

type Repository interface {
	Store(p domain.Product) (int, error)
	GetProductByName(name string) (domain.Product, error)
	Update(product domain.Product) (domain.Product, error)
	GetAll() ([]domain.Product, error)
	Delete(id int) error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Store(p domain.Product) (int, error) {
	query := "INSERT INTO products(name, type, price, count, code, public) VALUES (?,?,?,?,?,?)"
	stmt, err := r.db.Prepare(query)

	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	res, err := stmt.Exec(p.Name, p.Type, p.Price, p.Count, p.Code, p.Public)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()

	return int(id), nil
}

func (r *repository) GetProductByName(name string) (domain.Product, error) {
	query := "SELECT id, name, type, price, count, code, public FROM products WHERE name = ?;"
	row := r.db.QueryRow(query, name)
	p := domain.Product{}

	if err := row.Scan(&p.ID, &p.Name, &p.Type, &p.Price, &p.Count, &p.Code, &p.Public); err != nil {
		return domain.Product{}, err
	}

	return p, nil
}

func (r *repository) Update(product domain.Product) (domain.Product, error) {
	return domain.Product{}, nil
}

func (r *repository) GetAll() ([]domain.Product, error) {
	return nil, nil
}

func (r *repository) Delete(id int) error {
	return nil
}
