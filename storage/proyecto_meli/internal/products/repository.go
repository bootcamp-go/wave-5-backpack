package products

import (
	"context"
	"database/sql"
	"fmt"
	"proyecto_meli/internal/domain"
)

type Repository interface {
	GetAll(ctx context.Context) ([]domain.Product, error)
	GetById(ctx context.Context, id int) (domain.Product, error)
	FilterList(ctx context.Context, id int, name, color string, price float64, stock int, codigo string, publicado bool, fecha string) ([]domain.Product, error)
	Store(ctx context.Context, p domain.Product) (domain.Product, error)
	Update(ctx context.Context, p domain.Product) (domain.Product, error)
	Delete(ctx context.Context, id int) error
	Update_Name_Price(ctx context.Context, id int, name string, price float64) (domain.Product, error)
	GetProductByName(ctx context.Context, name string) ([]domain.Product, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

const (
	InsertProduct          = "INSERT INTO products(name, color, price, stock, code, publish, create_date) VALUES (?,?,?,?,?,?,?)"
	SelectByName           = "SELECT id, name, color, price, stock, code, publish, create_date FROM products WHERE name=?"
	SelectAllProduct       = "SELECT id, name, color, price, stock, code, publish, create_date FROM products"
	SelectByID             = "SELECT id, name, color, price, stock, code, publish, create_date FROM products WHERE id=?"
	UpdateProduct          = "UPDATE products SET name=?, color=?, price=?, stock=?, code=?, publish=?, create_date=? WHERE id=?"
	UpdateProductNamePrice = "UPDATE products SET name=?, price=? WHERE id=?"
	DeleteProduct          = "DELETE FROM products WHERE id=?"
)

func (r *repository) Store(ctx context.Context, p domain.Product) (domain.Product, error) {
	stmt, err := r.db.Prepare(InsertProduct)

	if err != nil {
		return domain.Product{}, err
	}

	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, p.Nombre, p.Color, p.Precio, p.Stock, p.Codigo, p.Publicado, p.FechaCreacion)
	if err != nil {
		return domain.Product{}, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return domain.Product{}, err
	}

	p.Id = int(id)

	return p, nil
}

func (r *repository) Update(ctx context.Context, p domain.Product) (domain.Product, error) {
	stmt, err := r.db.Prepare(UpdateProduct)
	if err != nil {
		return domain.Product{}, err
	}
	defer stmt.Close()
	fmt.Println(p)
	_, err = stmt.ExecContext(ctx, p.Nombre, p.Color, p.Precio, p.Stock, p.Codigo, p.Publicado, p.FechaCreacion, p.Id)
	if err != nil {
		return domain.Product{}, nil
	}
	return p, nil
}

func (r *repository) Update_Name_Price(ctx context.Context, id int, name string, price float64) (domain.Product, error) {
	stmt, err := r.db.Prepare(UpdateProductNamePrice)
	if err != nil {
		return domain.Product{}, err
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, name, price, id)
	if err != nil {
		return domain.Product{}, nil
	}
	p, err := r.GetById(ctx, id)
	if err != nil {
		return domain.Product{}, err
	}
	return p, nil
}

func (r *repository) GetProductByName(ctx context.Context, name string) ([]domain.Product, error) {
	rows, err := r.db.QueryContext(ctx, SelectByName, name)
	if err != nil {
		return nil, err
	}
	products := []domain.Product{}
	for rows.Next() {
		p := domain.Product{}
		publish := 0
		if err := rows.Scan(&p.Id, &p.Nombre, &p.Color, &p.Precio, &p.Stock, &p.Codigo, &publish, &p.FechaCreacion); err != nil {
			return nil, err
		}
		p.Publicado = publish != 0
		products = append(products, p)
	}
	return products, nil
}

func (r *repository) GetAll(ctx context.Context) ([]domain.Product, error) {
	rows, err := r.db.QueryContext(ctx, SelectAllProduct)
	if err != nil {
		return nil, err
	}
	products := []domain.Product{}
	for rows.Next() {
		p := domain.Product{}
		publish := 0
		if err := rows.Scan(&p.Id, &p.Nombre, &p.Color, &p.Precio, &p.Stock, &p.Codigo, &publish, &p.FechaCreacion); err != nil {
			return nil, err
		}
		p.Publicado = publish != 0
		products = append(products, p)
	}
	return products, nil
}

func (r *repository) GetById(ctx context.Context, id int) (domain.Product, error) {
	p := domain.Product{}
	row := r.db.QueryRowContext(ctx, SelectByID, id)
	publish := 0
	if err := row.Scan(&p.Id, &p.Nombre, &p.Color, &p.Precio, &p.Stock, &p.Codigo, &publish, &p.FechaCreacion); err != nil {
		return domain.Product{}, err
	}
	p.Publicado = publish != 0
	return p, nil
}
func (r *repository) FilterList(ctx context.Context, id int, name, color string, price float64, stock int, codigo string, publicado bool, fecha string) ([]domain.Product, error) {
	return nil, nil
}

func (r *repository) Delete(ctx context.Context, id int) error {
	stmt, err := r.db.Prepare(DeleteProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
