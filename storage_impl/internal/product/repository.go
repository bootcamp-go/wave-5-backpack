package product

import (
	"context"
	"database/sql"
	"storage/impl/internal/domain"
)

type Repository interface {
	GetAll(ctx context.Context) ([]domain.Product, error)
	Get(ctx context.Context, id int) (domain.Product, error)
	Save(ctx context.Context, p domain.Product) (int, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

func (r *repository) GetAll(ctx context.Context) ([]domain.Product, error) {
	query := "SELECT * FROM products;"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	var products []domain.Product
	for rows.Next() {
		p := domain.Product{}
		_ = rows.Scan(&p.ID, &p.Nombre, &p.Color, &p.Precio, &p.Stock, &p.Codigo, &p.Publicado, &p.FechaCreacion)
		products = append(products, p)
	}
	return products, nil
}

func (r *repository) Get(ctx context.Context, id int) (domain.Product, error) {
	query := "SELECT * FROM products WHERE id=?;"
	row := r.db.QueryRow(query, id)
	p := domain.Product{}
	err := row.Scan(&p.ID, &p.Nombre, &p.Color, &p.Precio, &p.Stock, &p.Codigo, &p.Publicado, &p.FechaCreacion)
	if err != nil {
		return domain.Product{}, err
	}
	return p, nil
}

func (r repository) Save(ctx context.Context, p domain.Product) (int, error) {
	query := "insert into products(id, nombre, color, precio, stock, codigo, publicado, fecha_creacion) VALUES (?,?,?,?,?,?,?,?)"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(p.Nombre, p.Color, p.Precio, p.Stock, p.Codigo, p.Publicado, p.FechaCreacion)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}
