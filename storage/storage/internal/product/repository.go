package product

import (
	"context"
	"database/sql"
	"storage/internal/domain"
)

const (
	GetAll    = "SELECT id, name, type, count, price FROM products"
	GetByName = "SELECT id, name, type, count, price FROM products WHERE name = ?"
	Store     = "INSERT INTO products(name, type, count, price) VALUES(?, ?, ?, ?)"
	Update    = "UPDATE products SET name = ?, type = ?, count = ?, price = ? WHERE id = ?"
	Delete    = "DELETE FROM products WHERE id = ?"
)

type Repository interface {
	GetAll(ctx context.Context) ([]domain.Product, error)
	GetByName(ctx context.Context, name string) (domain.Product, error)
	Store(ctx context.Context, product domain.Product) (domain.Product, error)
	Update(ctx context.Context, product domain.Product) (domain.Product, error)
	Delete(ctx context.Context, id int) error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll(ctx context.Context) ([]domain.Product, error) {
	var products []domain.Product
	rows, err := r.db.QueryContext(ctx, GetAll)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var product domain.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Type, &product.Count, &product.Price); err != nil {
			return nil, err
		}
		if product.ID > 0 {
			products = append(products, product)
		}
	}
	return products, nil
}

func (r *repository) GetByName(ctx context.Context, name string) (domain.Product, error) {
	var product domain.Product
	rows, err := r.db.QueryContext(ctx, GetByName, name)
	if err != nil {
		return product, err
	}
	for rows.Next() {
		if err := rows.Scan(&product.ID, &product.Name, &product.Type, &product.Count, &product.Price); err != nil {
			return domain.Product{}, err
		}
	}
	if product.ID == 0 {
		return domain.Product{}, nil
	}
	return product, nil
}

func (r *repository) Store(ctx context.Context, product domain.Product) (domain.Product, error) {
	stmt, err := r.db.PrepareContext(ctx, Store)
	if err != nil {
		return domain.Product{}, err
	}
	defer stmt.Close()
	result, err := stmt.ExecContext(ctx, product.Name, product.Type, product.Count, product.Price)
	if err != nil {
		return domain.Product{}, err
	}
	id, _ := result.LastInsertId()
	product.ID = int(id)
	return product, nil
}

func (r *repository) Update(ctx context.Context, product domain.Product) (domain.Product, error) {
	stmt, err := r.db.PrepareContext(ctx, Update)
	if err != nil {
		return domain.Product{}, err
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, product.Name, product.Type, product.Count, product.Price, product.ID)
	if err != nil {
		return domain.Product{}, err
	}
	return product, nil
}

func (r *repository) Delete(ctx context.Context, id int) error {
	stmt, err := r.db.PrepareContext(ctx, Delete)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
