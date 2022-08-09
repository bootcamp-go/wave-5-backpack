package products

import (
	"database/sql"

	"github.com/bootcamp-go/wave-5-backpack/tree/Ramos_Andres/goweb/practica/internal/domain"
)

var (
	createStmt    = "INSERT INTO products (name, color, price, stock, code, published, created_at, warehouse_id) VALUES (?, ?, ?, ?, ?, ?, CURDATE(), ?)"
	getAllStmt    = "SELECT ID, name, color, price, stock, code, published, created_at, warehouse_id FROM products"
	getByNameStmt = "SELECT ID, name, color, price, stock, code, published, created_at, warehouse_id FROM products WHERE name = ?"
)

type Repository interface {
	Store(domain.Product) (domain.Product, error)
	GetAll() ([]domain.Product, error)
	GetById(id uint64) (domain.Product, error)
	GetByName(name string) (domain.Product, error)
	Update(domain.Product) (domain.Product, error)
	Delete(id uint64) (domain.Product, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Store(product domain.Product) (domain.Product, error) {
	db := r.db
	stmt, err := db.Prepare(createStmt)
	if err != nil {
		return domain.Product{}, err
	}
	defer stmt.Close()
	sqlRes, err := stmt.Exec(product.Name, product.Color, product.Price, product.Stock, product.Code, product.Published, product.Warehouse_id)
	if err != nil {
		return domain.Product{}, err
	}
	insertedId, err := sqlRes.LastInsertId()
	if err != nil {
		return domain.Product{}, err
	}
	product.Id = uint64(insertedId)
	return product, nil
}

func (r *repository) GetAll() ([]domain.Product, error) {
	db := r.db
	rows, err := db.Query(getAllStmt)
	if err != nil {
		return []domain.Product{}, err
	}
	defer rows.Close()

	productList := []domain.Product{}

	for rows.Next() {
		var product domain.Product
		if err := rows.Scan(&product.Id, &product.Name, &product.Color, &product.Price, &product.Stock, &product.Code, &product.Published, &product.Created_at, &product.Warehouse_id); err != nil {
			return productList, err
		}
		productList = append(productList, product)
	}
	if err := rows.Err(); err != nil {
		return []domain.Product{}, err
	}
	return productList, nil
}

func (r *repository) GetById(id uint64) (domain.Product, error) {
	return domain.Product{}, nil
}

func (r *repository) GetByName(name string) (domain.Product, error) {
	db := r.db
	rows, err := db.Query(getByNameStmt, name)
	if err != nil {
		return domain.Product{}, err
	}
	defer rows.Close()

	productList := []domain.Product{}

	for rows.Next() {
		var product domain.Product
		if err := rows.Scan(&product.Id, &product.Name, &product.Color, &product.Price, &product.Stock, &product.Code, &product.Published, &product.Created_at, &product.Warehouse_id); err != nil {
			return productList[0], err
		}
		productList = append(productList, product)
	}
	if err := rows.Err(); err != nil {
		return domain.Product{}, err
	}
	return productList[0], nil
}

func (r *repository) Update(domain.Product) (domain.Product, error) {
	return domain.Product{}, nil
}

func (r *repository) Delete(id uint64) (domain.Product, error) {
	return domain.Product{}, nil
}
