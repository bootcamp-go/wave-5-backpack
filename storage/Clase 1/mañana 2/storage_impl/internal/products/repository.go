package products

import (
	"database/sql"
	"fmt"

	"github.com/bootcamp-go/storage/internal/domains"
)

type Repository interface {
	Store(p domains.Product) (int, error)
	GetByName(name string) (domains.Product, error)
	GetAll()([]domains.Product, error)
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

type repository struct {
	db *sql.DB
}
const (
	Store         = "INSERT INTO products (name, type, count, price) VALUES (?,?,?,?)"
	GetByName =  "SELECT id, name, type, count, price FROM products WHERE name = ?;"
	GetAll       = "SELECT id, name, type, count, price FROM products"
)

func (r *repository) Store(p domains.Product) (int, error) {
	stmt, err := r.db.Prepare("INSERT INTO products (name, type, count, price) VALUES (?,?,?,?)")
	if err != nil {
		return 0, fmt.Errorf("error al preparar la consulta - error %v", err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(p.Name, p.Type, p.Count, p.Price)
	if err != nil {
		return 0, fmt.Errorf("error al ejecutar la consulta - error %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("error al obtener último id - error %v", err)
	}

	return int(id), nil
}


func (r *repository) GetByName(name string) (domains.Product, error) {
	stmt, err := r.db.Prepare("SELECT id, name, type, count, price FROM products WHERE name = ?;")
	if err != nil {
		return domains.Product{}, fmt.Errorf("error al preparar la consulta - error %v", err)
	}
	defer stmt.Close()

	var product domains.Product
	err = stmt.QueryRow(name).Scan(&product.ID, &product.Name, &product.Type, &product.Count, &product.Price)
	if err != nil {
		return domains.Product{}, fmt.Errorf("no registros para %s - error %v", name, err)
	}

	return product, nil
}


func (r *repository) GetAll() ([]domains.Product, error) {
	var products []domains.Product
	rows, err := r.db.Query(GetAll)
	if err != nil {
		return nil, fmt.Errorf("error al preparar la consulta - error %v", err)
	}
	for rows.Next() {
		var product domains.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Type, &product.Count, &product.Price); err != nil {
			return nil, err
		}
		products= append(products, product)
	}
	return products, nil 
}
