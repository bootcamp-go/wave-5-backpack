package products

import (
	"context"
	"database/sql"
	"log"

	"storage/internal/domain"
	//"storage/pkg/db"
)

type Repository interface {
	Store(product domain.Product) (domain.Product, error)
	GetByName(name string) (domain.Product, error)
	GetOne(id int) (domain.Product, error)
	Update(ctx context.Context, product domain.Product) (domain.Product, error)
	GetAll() ([]domain.Product, error)
	Delete(id int) error
}
type repository struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

const (
	InsertQuery    = "INSERT INTO products(name, type, count, price) VALUES( ?, ?, ?, ? )"
	GetByNameQuery = "SELECT id, name, type, count, price FROM products WHERE name = ?"
	GetOneQuery    = "SELECT id, name, type, count, price FROM products WHERE id = ?"
	GetAllQuery    = "SELECT id, name, type, count, price FROM products"
	UpdateQuery    = "UPDATE products SET name = ?, type = ?, count = ?, price = ? WHERE id = ?"
	DeleteQuery    = "DELETE FROM products WHERE id = ?"
)

func (r *repository) Store(product domain.Product) (domain.Product, error) {
	stmt, err := r.db.Prepare(InsertQuery)
	if err != nil {
		log.Println(err)
		return product, err
	}
	defer stmt.Close()
	var result sql.Result
	result, err = stmt.Exec(product.Name, product.Type, product.Count, product.Price)
	if err != nil {
		log.Fatal(err)
		return product, err
	}
	insertedId, _ := result.LastInsertId()
	product.ID = int(insertedId)

	return product, nil
}
func (r *repository) GetOne(id int) (domain.Product, error) {
	var product domain.Product
	rows, err := r.db.Query(GetOneQuery, id)
	if err != nil {
		log.Println(err)
		return product, err
	}
	for rows.Next() {
		if err := rows.Scan(&product.ID, &product.Name, &product.Type, &product.Count, &product.Price); err != nil {
			log.Fatal(err)
			return product, err
		}
	}
	return product, nil
}
func (r *repository) GetByName(name string) (domain.Product, error) {
	var product domain.Product
	rows, err := r.db.Query(GetByNameQuery, name)
	if err != nil {
		log.Println(err)
		return product, err
	}
	for rows.Next() {
		if err := rows.Scan(&product.ID, &product.Name, &product.Type, &product.Count, &product.Price); err != nil {
			log.Fatal(err)
			return product, err
		}
	}
	return product, nil
}
func (r *repository) GetAll() ([]domain.Product, error) {
	var products []domain.Product
	rows, err := r.db.Query(GetAllQuery)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	for rows.Next() {
		var product domain.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Type, &product.Count, &product.Price); err != nil {
			log.Fatal(err)
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}
func (r *repository) Update(ctx context.Context, product domain.Product) (domain.Product, error) {
	stmt, err := r.db.PrepareContext(ctx, UpdateQuery)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(product.Name, product.Type, product.Count, product.Price, product.ID)
	if err != nil {
		return domain.Product{}, err
	}
	return product, nil
}
func (r *repository) Delete(id int) error {
	stmt, err := r.db.Prepare("DELETE FROM products WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}
