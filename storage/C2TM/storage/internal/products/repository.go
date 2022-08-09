package products

import (
	"database/sql"
	"fmt"
	"storage/internal/domain"
)

type Repository interface {
	Store(domain.Product) (int, error)
	GetByName(name string) (domain.Product, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Store(product domain.Product) (int, error) { // se inicializa la base

	stmt, err := r.db.Prepare("INSERT INTO products(name, type, count, price) VALUES( ?, ?, ?, ? )") // se prepara la sentencia SQL a ejecutar
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
	stmt, err := r.db.Prepare("SELECT id, name, type, count, price FROM products WHERE name = ?;")
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
