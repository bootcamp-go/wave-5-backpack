package products

import (
	"database/sql"
	"log"
	"practica/internal/domain"
)

const (
	GetProduct   = "SELECT * FROM products WHERE id = ?"
	StoreProduct = "INSERT INTO products(name, type, count, price) VALUES( ?, ?, ?, ? )"
)

type Repository interface {
	Store(product domain.Product) (domain.Product, error)
	GetOne(id int) (domain.Product, error)
}

type repository struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Store(product domain.Product) (domain.Product, error) { // se inicializa la base
	stmt, err := r.db.Prepare(StoreProduct) // se prepara el SQL
	if err != nil {
		log.Fatal(err)
		return domain.Product{}, err
	}
	defer stmt.Close() // se cierra la sentencia al terminar. Si quedan abiertas se genera consumos de memoria
	var result sql.Result
	result, err = stmt.Exec(product.Name, product.Type, product.Count, product.Price) // retorna un sql.Result y un error
	if err != nil {
		return domain.Product{}, err
	}
	insertedId, _ := result.LastInsertId() // del sql.Resul devuelto en la ejecución obtenemos el Id insertado
	product.ID = int(insertedId)

	return product, nil
}

func (r *repository) GetOne(id int) (domain.Product, error) {
	var product domain.Product
	rows, err := r.db.Query(GetProduct, id)
	if err != nil {
		return product, err
	}
	for rows.Next() {
		err := rows.Scan(&product.ID, &product.Name, &product.Type, &product.Count, &product.Price)
		if err != nil {
			log.Fatal(err)
			return product, err
		}
	}
	return product, nil
}
