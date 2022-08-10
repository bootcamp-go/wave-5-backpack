package products

import (
	"context"
	"database/sql"
	"practica1-clase1/internal/domain"
)

type Repository interface {
	GetByName(name string) (domain.Product, error)
	Store(product domain.Product) (domain.Product, error)
	GetAll(cxt context.Context) ([]domain.ProductAndWarehouse, error)
	Update(ctx context.Context, product domain.Product) (domain.Product, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

const (
	createProduct           = `INSERT INTO products (name, type, count, price) VALUES (?, ?, ?, ?)`
	getProductByName        = `SELECT name, type, count, price FROM products WHERE name = ?`
	getAllProductsWarehouse = `SELECT p.id, p.name, p.type, p.count, p.price, w.id, w.name, w.adress 
								FROM products p INNER JOIN warehouses w ON p.id_warehouse = w.id`
	updateProduct = `UPDATE products SET name = ?, type = ?, count = ?, price = ? WHERE id = ?`
)

func (r *repository) Store(product domain.Product) (domain.Product, error) {

	stmt, err := r.db.Prepare(createProduct)
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
	rows, err := r.db.Query(getProductByName, name)
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

func (r *repository) GetAll(cxt context.Context) ([]domain.ProductAndWarehouse, error) {
	var products []domain.ProductAndWarehouse
	rows, err := r.db.QueryContext(cxt, getAllProductsWarehouse)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var product domain.ProductAndWarehouse
		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Type,
			&product.Count,
			&product.Price,
			&product.Warehouse.ID,
			&product.Warehouse.Name,
			&product.Warehouse.Adress,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (r *repository) Update(ctx context.Context, product domain.Product) (domain.Product, error) {
	stmt, err := r.db.Prepare(updateProduct)
	if err != nil {
		return domain.Product{}, err
	}
	defer stmt.Close() // se cierra la sentencia al terminar. Si quedan abiertas se genera consumos de memoria
	_, err = stmt.ExecContext(ctx, product.Name, product.Type, product.Count, product.Price, product.ID)
	if err != nil {
		return domain.Product{}, err
	}
	return product, nil
}
