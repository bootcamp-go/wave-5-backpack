package products

import (
	"context"
	"database/sql"
	"log"
	"storage/internal/domain"
)

type Repository interface {
	GetAll() ([]domain.Products, error)
	GetById(id int) (domain.Products, error)
	CreateProduct(product domain.Products) (domain.Products, error)
	Update(ctx context.Context, p domain.Products) error
	Delete(id int) error
	UpdateOne(id int, nombre string, precio float64) (domain.Products, error)
}

const (
	ProductNotFound = "producto %d no encontrado"
	FailReading     = "no se pudo leer el archivo"
	FailWriting     = "no se pudo escribir el archivo, error: %w"
)

//Querys
const (
	GetAllQuery = "SELECT * FROM products"
	GetOneQuery = "SELECT id, name, color, price, stock, code, publish, creation_date FROM products WHERE id = ?"
	CreateQuery = "INSERT INTO products(name, color, price, stock, code, publish, creation_date) VALUES(?, ?, ?, ?, ?, ?, ?)"
	UpdateQuery = "UPDATE products SET name=?, color=?, price=?, stock=?, publish=?  WHERE id=?"
)

type repository struct {
	db *sql.DB
}

func InitRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll() ([]domain.Products, error) {
	var products []domain.Products
	rows, err := r.db.Query(GetAllQuery)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		p := domain.Products{}
		_ = rows.Scan(&p.Id, &p.Nombre, &p.Color, &p.Precio, &p.Stock, &p.Codigo, &p.Publicado, &p.FechaCreacion)
		products = append(products, p)
	}
	return products, nil
}

func (r *repository) GetById(id int) (domain.Products, error) {
	var product domain.Products
	rows, err := r.db.Query(GetOneQuery, id)
	if err != nil {
		return domain.Products{}, err
	}
	for rows.Next() {
		if err := rows.Scan(&product.Id, &product.Nombre, &product.Color, &product.Precio, &product.Stock, &product.Codigo, &product.Publicado, &product.FechaCreacion); err != nil {
			return domain.Products{}, err
		}
	}
	return product, nil
}

func (r *repository) CreateProduct(product domain.Products) (domain.Products, error) {
	stmt, err := r.db.Prepare(CreateQuery)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	var result sql.Result
	result, err = stmt.Exec(product.Nombre, product.Color, product.Precio, product.Stock, product.Codigo, product.Publicado, product.FechaCreacion)
	if err != nil {
		return domain.Products{}, err
	}
	insertedId, _ := result.LastInsertId()
	product.Id = int(insertedId)

	return product, nil
}

func (r *repository) Update(ctx context.Context, p domain.Products) error {
	stmt, err := r.db.PrepareContext(ctx, UpdateQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	res, err := stmt.ExecContext(ctx, &p.Nombre, &p.Color, &p.Precio, &p.Stock, &p.Publicado, &p.Id)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) Delete(id int) error {

	return nil

}

func (r *repository) UpdateOne(id int, nombre string, precio float64) (domain.Products, error) {

	return domain.Products{}, nil
}
