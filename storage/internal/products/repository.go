package products

import (
	"context"
	"database/sql"
	"errors"

	"github.com/bootcamp-go/wave-5-backpack/storage/internal/domain"
)

type Repository interface {
	Store(p domain.Product) (int, error)
	GetProductByName(name string) (domain.Product, error)
	GetProductAndWareHouse() ([]domain.Product_Warehouse, error)
	Update(ctx context.Context, p domain.Product) error
	GetAll() ([]domain.Product, error)
	Delete(id int) error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Store(p domain.Product) (int, error) {
	query := Store
	stmt, err := r.db.Prepare(query)

	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	res, err := stmt.Exec(p.Name, p.Type, p.Price, p.Count, p.Code, p.Public, p.WarehouseID)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()

	return int(id), nil
}

func (r *repository) GetProductAndWareHouse() ([]domain.Product_Warehouse, error) {
	query := GetProductAndWareHouse
	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}

	var pW []domain.Product_Warehouse
	var p domain.Product_Warehouse
	for rows.Next() {
		if err := rows.Scan(&p.ID, &p.Name, &p.Type, &p.Price, &p.Count, &p.Code, &p.Public, &p.Warehouse.ID, &p.Warehouse.Name, &p.Warehouse.Address); err != nil {
			return nil, err
		}
		pW = append(pW, p)
	}

	return pW, nil
}

func (r *repository) GetProductByName(name string) (domain.Product, error) {
	query := GetProductByName
	row := r.db.QueryRow(query, name)
	p := domain.Product{}

	if err := row.Scan(&p.ID, &p.Name, &p.Type, &p.Price, &p.Count, &p.Code, &p.Public, &p.WarehouseID); err != nil {
		return domain.Product{}, err
	}

	return p, nil
}

func (r *repository) Update(ctx context.Context, p domain.Product) error {
	query := UpdateAll
	stmt, err := r.db.Prepare(query)

	if err != nil {
		return nil
	}

	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, p.Name, p.Type, p.Price, p.Count, p.Code, p.Public, p.ID)

	if err != nil {
		return err
	}

	if _, err = res.RowsAffected(); err != nil {
		return err
	}

	return nil
}

func (r *repository) GetAll() ([]domain.Product, error) {
	query := GetAll
	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}

	var products []domain.Product
	var p domain.Product
	for rows.Next() {
		if err := rows.Scan(&p.ID, &p.Name, &p.Type, &p.Price, &p.Count, &p.Code, &p.Public); err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}

func (r *repository) Delete(id int) error {
	query := Delete
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	res, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	affect, err := res.RowsAffected()
	if affect < 1 {
		return errors.New("no fue posible eliminar el producto")
	}

	return nil
}
