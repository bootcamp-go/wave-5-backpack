package products

import (
	"database/sql"
	"fmt"

	"github.com/bootcamp-go/storage/internal/domains"
)

type Repository interface {
	Store(p domains.Product) (int, error)
	GetByName(name string) (domains.Product, error)
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

type repository struct {
	db *sql.DB
}

/* Ejercicio 2 - Replicar Store()
Tomar el ejemplo visto en la clase y diseñar el método Store():
Puede tomar de ejemplo la definición del método Store visto en clase para incorporarlo en la interfaz.
Implementar el método Store.
*/
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

/* Ejercicio 1 - Implementar GetByName()
Desarrollar un método en el repositorio que permita hacer búsquedas de un producto por nombre. Para lograrlo se deberá:
Diseñar interfaz “Repository” en la que exista un método GetByName() que reciba por parámetro un string y retorna una estructura del tipo Product.
Implementar el método de forma que con el string recibido lo use para buscar en la DB por el campo “name”.
*/
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
