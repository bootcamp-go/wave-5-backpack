package products

import (
	"context"
	"database/sql"
	"errors"
	"goweb/internal/domain"
	"log"
)

const (
	Store     = "INSERT INTO products (name,type,count,price) VALUES(?,?,?,?)"
	GetByName = "SELECT id,name,type,count,price FROM products WHERE name=?"
	GetAll    = "SELECT id,name,type,count,price,id_warehouse FROM products"
)

type Repository interface {
	GetByName(name string) (domain.Product, error)
	Store(product domain.Product) (domain.Product, error)
	GetAll(context.Context) ([]domain.Product, error)
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
	stmt, err := r.db.Prepare(Store)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	var result sql.Result
	result, err = stmt.Exec(product.Name, product.Type, product.Count, product.Price)
	if err != nil {
		return domain.Product{}, err
	}

	insertedId, _ := result.LastInsertId()
	product.Id = int(insertedId)

	return product, nil
}

func (r *repository) GetByName(name string) (domain.Product, error) {
	var product domain.Product

	rows, err := r.db.Query(GetByName, name)
	if err != nil {
		log.Println(err)
		return product, err
	}

	for rows.Next() {
		if err := rows.Scan(&product.Id, &product.Name, &product.Type, &product.Count, &product.Price); err != nil {
			log.Println(err)
			return product, err
		}
	}

	if product.Id == 0 {
		return product, errors.New("Product not found")
	}
	return product, nil
}

func (r *repository) GetAll(ctx context.Context) ([]domain.Product, error) {
	var products []domain.Product

	rows, err := r.db.QueryContext(ctx, GetAll)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	for rows.Next() {
		var product domain.Product
		if err := rows.Scan(&product.Id, &product.Name, &product.Type, &product.Count, &product.Price, &product.Id_warehouse); err != nil {
			log.Fatal(err)
			log.Println("estoy aqui")
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}
