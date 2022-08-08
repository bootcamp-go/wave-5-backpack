package products

import (
	"database/sql"
	"log"
	"storage/internal/domain"
)

type Repository interface {
	GetById(id int) (domain.Products, error)
	CreateProduct(product domain.Products) (domain.Products, error)
	Update(id int, nombre, color string, precio float64, stock int, código string, publicado bool, fecha_de_creación string) (domain.Products, error)
	Delete(id int) error
	UpdateOne(id int, nombre string, precio float64) (domain.Products, error)
}

const (
	ProductNotFound = "producto %d no encontrado"
	FailReading     = "no se pudo leer el archivo"
	FailWriting     = "no se pudo escribir el archivo, error: %w"
)

type repository struct {
	db *sql.DB
}

func InitRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetById(id int) (domain.Products, error) {
	var product domain.Products
	rows, err := r.db.Query("SELECT id, name, color, price, stock, code, publish, creation_date FROM products WHERE id = ?", id)
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
	stmt, err := r.db.Prepare("INSERT INTO products(name, color, price, stock, code, publish, creation_date) VALUES(?, ?, ?, ?, ?, ?, ?)")
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

func (r *repository) Update(id int, nombre, color string, precio float64, stock int, código string, publicado bool, fecha_de_creación string) (domain.Products, error) {

	return domain.Products{}, nil
}

func (r *repository) Delete(id int) error {

	return nil

}

func (r *repository) UpdateOne(id int, nombre string, precio float64) (domain.Products, error) {

	return domain.Products{}, nil
}
