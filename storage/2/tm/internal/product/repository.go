package product

import (
	"context"
	"database/sql"
	"log"
	"storage/2/tm/internal/domain"
)

type Repository interface {
	Store(product domain.Product) (domain.Product, error)
	GetAll() ([]domain.Product, error)
	GetOne(id int) (domain.Product, error)
	GetByName(name string) ([]domain.Product, error)
	Update(ctx context.Context, product domain.Product) (domain.Product, error)
	Delete(id int) error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

const (
	StoreProduct     = "INSERT INTO products(name, type, count, price) VALUES (?, ?, ?, ?)"
	GetAllProducts   = "SELECT id, name, type, count, price FROM products"
	GetProduct       = "SELECT id, name, type, count, price FROM products WHERE id = ?"
	GetProductByName = "SELECT id, name, type, count, price FROM products WHERE name = ?"
	UpdateProduct    = "UPDATE products SET name = ?, type = ?, count = ?, price = ? WHERE id = ?"
	DeleteProduct    = "DELETE FROM products WHERE id = ?"
)

func (r *repository) Store(product domain.Product) (domain.Product, error) {
	stmt, err := r.db.Prepare(StoreProduct)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	var result sql.Result
	result, err = stmt.Exec(product.Name, product.Type, product.Count, product.Price)
	if err != nil {
		return domain.Product{}, err
	}
	insertedID, _ := result.LastInsertId()
	product.ID = int(insertedID)

	return product, nil
}

func (r *repository) GetOne(id int) (domain.Product, error) {
	var product domain.Product
	rows, err := r.db.Query(GetProduct, id)
	if err != nil {
		log.Println(err)
		return domain.Product{}, err
	}

	for rows.Next() {
		if err := rows.Scan(&product.ID, &product.Name, &product.Type, &product.Count, &product.Price); err != nil {
			log.Println(err)
			return domain.Product{}, err
		}
	}

	return product, nil
}

func (r *repository) GetByName(name string) ([]domain.Product, error) {
	var product domain.Product
	var products []domain.Product
	rows, err := r.db.Query(GetProductByName, name)
	if err != nil {
		log.Println(err)
		return []domain.Product{}, err
	}

	for rows.Next() {
		if err := rows.Scan(&product.ID, &product.Name, &product.Type, &product.Count, &product.Price); err != nil {
			log.Println(err)
			return []domain.Product{}, err
		}
		products = append(products, product)
	}

	return products, nil
}

func (r *repository) GetAll() ([]domain.Product, error) {
	var product domain.Product
	var products []domain.Product
	rows, err := r.db.Query(GetAllProducts)
	if err != nil {
		log.Println(err)
		return []domain.Product{}, err
	}

	for rows.Next() {
		if err := rows.Scan(&product.ID, &product.Name, &product.Type, &product.Count, &product.Price); err != nil {
			log.Println(err)
			return []domain.Product{}, err
		}
		products = append(products, product)
	}

	return products, nil
}

func (r *repository) Update(ctx context.Context, product domain.Product) (domain.Product, error) {
	stmt, err := r.db.Prepare(UpdateProduct)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, product.Name, product.Type, product.Count, product.Price, product.ID)
	if err != nil {
		return domain.Product{}, err
	}

	return product, nil
}

func (r *repository) Delete(id int) error {
	stmt, err := r.db.Prepare(DeleteProduct)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
