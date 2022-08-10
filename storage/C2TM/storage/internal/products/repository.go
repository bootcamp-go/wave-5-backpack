package products

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"storage/internal/domain"
)

type Repository interface {
	Store(domain.Product) (int, error)
	GetByName(name string) (domain.Product, error)
	GetAll() ([]domain.Product, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

const (
	queryStore     = "INSERT INTO products(name, type, count, price) VALUES( ?, ?, ?, ? )"
	queryGetByName = "SELECT id, name, type, count, price FROM products WHERE name = ?;"
	queryGetAll    = "SELECT * FROM products"
	UPDATE_QUERY   = "UPDATE products SET name=?, type=?, count=?, price=? WHERE id=?;"
)

func (r *repository) Store(product domain.Product) (int, error) { // se inicializa la base

	stmt, err := r.db.Prepare(queryStore) // se prepara la sentencia SQL a ejecutar
	if err != nil {
		return 0, err
	}
	defer stmt.Close() // se cierra la sentencia al terminar. Si quedan abiertas se genera consumos de memoria
	var result sql.Result
	result, err = stmt.Exec(product.Name, product.Type, product.Count, product.Price) // retorna un sql.Result y un error
	if err != nil {
		return 0, err
	}
	insertedId, _ := result.LastInsertId() // del sql.Resul devuelto en la ejecucion obtenemos el Id insertado
	product.ID = int(insertedId)
	return product.ID, nil
}

func (r *repository) GetByName(name string) (domain.Product, error) {
	stmt, err := r.db.Prepare(queryGetByName)
	if err != nil {
		return domain.Product{}, fmt.Errorf("error al preparar la consulta - error %v", err)
	}
	defer stmt.Close()

	var product domain.Product
	err = stmt.QueryRow(name).Scan(&product.ID, &product.Name, &product.Type, &product.Count, &product.Price)
	if err != nil {
		return domain.Product{}, fmt.Errorf("no registros para %s - error %v", name, err)
	}

	return product, nil
}

func (r *repository) GetAll() ([]domain.Product, error) {
	var products []domain.Product
	rows, err := r.db.Query(queryGetAll)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	// se recorren todas las filas
	for rows.Next() {
		// por cada fila se obtiene un objeto del tipo Product
		var product domain.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Type, &product.Count, &product.Price); err != nil {
			log.Fatal(err)
			return nil, err
		}
		//se añade el objeto obtenido al slice products
		products = append(products, product)
	}
	return products, nil
}

func (r *repository) Update(ctx context.Context, p domain.Product) error {
	stmt, err := r.db.Prepare(UPDATE_QUERY)
	if err != nil {
		return fmt.Errorf("error al preparar la consulta - error %v", err)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, p.Name, p.Type, p.Count, p.Price, p.ID)
	if err != nil {
		return fmt.Errorf("error al ejecutar la consulta - error %v", err)
	}

	return nil
}
