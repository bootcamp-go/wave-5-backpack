package products

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/bootcamp-go/storage/internal/domains"
)

type Repository interface {
	Store(context.Context, domains.Product) (int, error)
	GetAll(context.Context, int) ([]domains.Products, error)
	Get(context.Context, int) (domains.Product, error)
	GetByName(context.Context, string) (domains.Product, error)
	Update(context.Context, domains.Product) error
	Exists(ctx context.Context, id int) bool
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

type repository struct {
	db *sql.DB
}

func (r *repository) Get(ctx context.Context, id int) (domains.Product, error) {
	stmt, err := r.db.Prepare(GET_BY_ID_QUERY)
	if err != nil {
		return domains.Product{}, fmt.Errorf("error al preparar la consulta - error %v", err)
	}
	defer stmt.Close()

	var p domains.Product // id, name, type, count, price, id_warehouse
	err = stmt.QueryRowContext(ctx, id).Scan(&p.ID, &p.Name, &p.Type, &p.Count, &p.Price, &p.WarehouseId)
	if err != nil {
		return domains.Product{}, err
	}

	return p, nil
}

func (r *repository) GetAll(ctx context.Context, id int) ([]domains.Products, error) {
	stmt, err := r.db.Prepare(GET_ALL_QUERY)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, id)
	if err != nil {
		return nil, err
	}

	var products []domains.Products
	for rows.Next() {
		var p domains.Products
		err := rows.Scan(&p.ID, &p.Name, &p.Type, &p.Count, &p.Price, &p.WarehouseName, &p.WarehouseAddress)
		if err != nil {
			return nil, err
		}

		products = append(products, p)
	}

	if err := rows.Err(); err != nil {
		return nil, rows.Err()
	}

	return products, nil
}

func (r *repository) Store(ctx context.Context, p domains.Product) (int, error) {
	stmt, err := r.db.Prepare(INSERT_QUERY)
	if err != nil {
		return 0, fmt.Errorf("error al preparar la consulta - error %v", err)
	}
	defer stmt.Close()

	result, err := stmt.ExecContext(ctx, p.Name, p.Type, p.Count, p.Price, p.WarehouseId)
	if err != nil {
		return 0, fmt.Errorf("error al ejecutar la consulta - error %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("error al obtener último id - error %v", err)
	}

	return int(id), nil
}

func (r *repository) GetByName(ctx context.Context, name string) (domains.Product, error) {
	stmt, err := r.db.Prepare(GET_BY_NAME_QUERY)
	if err != nil {
		return domains.Product{}, fmt.Errorf("error al preparar la consulta - error %v", err)
	}
	defer stmt.Close()

	var product domains.Product
	err = stmt.QueryRowContext(ctx, name).Scan(&product.ID, &product.Name, &product.Type, &product.Count, &product.Price)
	if err != nil {
		return domains.Product{}, fmt.Errorf("no registros para %s - error %v", name, err)
	}

	return product, nil
}

/* Ejercicio 3 - Implementar Context en Update
Hacer los cambios necesarios para implementar el context en el método Update del repository. Para esto será necesario:
    1. Cambiar la interfaz de Repository.
    2. Aplicar el conext en la ejecución del Update. Para lograrlo debe usar el método db.ExecContext en vez de db.Exec.
*/
func (r *repository) Update(ctx context.Context, p domains.Product) error {
	stmt, err := r.db.Prepare(UPDATE_QUERY)
	if err != nil {
		return fmt.Errorf("error al preparar la consulta - error %v", err)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, p.Name, p.Type, p.Count, p.Price, p.WarehouseId, p.ID)
	if err != nil {
		return fmt.Errorf("error al ejecutar la consulta - error %v", err)
	}

	return nil
}

func (r *repository) Exists(ctx context.Context, id int) bool {
	row := r.db.QueryRow(EXISTS_QUERY, id)
	err := row.Scan(&id)
	return err == nil
}
