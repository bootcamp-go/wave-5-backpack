package products

import (
	"context"
	"database/sql"
	"products_project/internal/domain"
)

const (
	GetById        string = "SELECT p.id, p.nombre, p.color, p.precio, p.stock, p.codigo, p.publicado, p.fecha FROM products p WHERE p.id = ?"
	GetByName      string = "SELECT p.id, p.nombre, p.color, p.precio, p.stock, p.codigo, p.publicado, p.fecha FROM products p WHERE p.nombre = ?"
	GetAllProducts string = "SELECT p.id, p.nombre, p.color, p.precio, p.stock, p.codigo, p.publicado, p.fecha FROM products p"
	Store          string = "INSERT INTO products(nombre, color, precio, stock, codigo, publicado, fecha) VALUES( ?, ?, ?, ?, ?, ?, ?)"
	Update         string = "UPDATE products SET nombre = ?, color = ?, precio = ?, stock = ?, codigo = ?, publicado = ?, fecha = ? WHERE id = ?"
	UpdateFields   string = "UPDATE products SET nombre = ?, precio = ? WHERE id = ?"
	Delete         string = "DELETE FROM products WHERE id = ?"
)

type Repository interface {
	GetAll(ctx context.Context) ([]domain.Product, error)
	Store(ctx context.Context, nombre, color string, precio, stock int, codigo string, publicado bool, fecha string) (domain.Product, error)
	Update(ctx context.Context, id int, nombre, color string, precio, stock int, codigo string, publicado bool, fecha string) (domain.Product, error)
	UpdateFields(ctx context.Context, id int, nombre string, precio int) (domain.Product, error)
	Delete(ctx context.Context, id int) (domain.Product, error)
	GetById(ctx context.Context, id int) (domain.Product, error)
	GetByName(ctx context.Context, nombre string) ([]domain.Product, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetById(ctx context.Context, id int) (domain.Product, error) {
	var product domain.Product
	db := r.db
	rows, err := db.QueryContext(ctx, GetById, id)
	if err != nil {
		return domain.Product{}, err
	}
	for rows.Next() {
		if err := rows.Scan(&product.Id, &product.Nombre, &product.Color, &product.Precio, &product.Stock, &product.Codigo, &product.Publicado, &product.Fecha); err != nil {
			return domain.Product{}, err
		}
	}
	return product, nil
}

func (r *repository) GetByName(ctx context.Context, nombre string) ([]domain.Product, error) {
	var products []domain.Product
	db := r.db
	rows, err := db.QueryContext(ctx, GetByName, nombre)
	if err != nil {
		return []domain.Product{}, err
	}
	for rows.Next() {
		var product domain.Product
		if err := rows.Scan(&product.Id, &product.Nombre, &product.Color, &product.Precio, &product.Stock, &product.Codigo, &product.Publicado, &product.Fecha); err != nil {
			return []domain.Product{}, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (r *repository) GetAll(ctx context.Context) ([]domain.Product, error) {
	var products []domain.Product
	db := r.db
	rows, err := db.QueryContext(ctx, GetAllProducts)
	if err != nil {
		return []domain.Product{}, err
	}
	for rows.Next() {
		var product domain.Product
		if err := rows.Scan(&product.Id, &product.Nombre, &product.Color, &product.Precio, &product.Stock, &product.Codigo, &product.Publicado, &product.Fecha); err != nil {
			return []domain.Product{}, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (r *repository) Store(ctx context.Context, nombre, color string, precio, stock int, codigo string, publicado bool, fecha string) (domain.Product, error) {
	db := r.db
	stmt, err := db.Prepare(Store)
	if err != nil {
		return domain.Product{}, err
	}
	defer stmt.Close()
	var result sql.Result
	result, err = stmt.ExecContext(ctx, nombre, color, precio, stock, codigo, publicado, fecha)
	if err != nil {
		return domain.Product{}, err
	}
	insertedId, _ := result.LastInsertId()
	id := int(insertedId)
	product := domain.Product{Id: id, Nombre: nombre, Color: color, Precio: precio, Stock: stock, Codigo: codigo, Publicado: publicado, Fecha: fecha}
	return product, nil
}

func (r *repository) Update(ctx context.Context, id int, nombre, color string, precio, stock int, codigo string, publicado bool, fecha string) (domain.Product, error) {
	db := r.db
	stmt, err := db.Prepare(Update)
	if err != nil {
		return domain.Product{}, err
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, nombre, color, precio, stock, codigo, publicado, fecha, id)
	if err != nil {
		return domain.Product{}, err
	}
	product := domain.Product{Id: id, Nombre: nombre, Color: color, Precio: precio, Stock: stock, Codigo: codigo, Publicado: publicado, Fecha: fecha}
	return product, nil
}

func (r *repository) UpdateFields(ctx context.Context, id int, nombre string, precio int) (domain.Product, error) {
	db := r.db
	stmt, err := db.Prepare(UpdateFields)
	if err != nil {
		return domain.Product{}, err
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, nombre, precio, id)
	if err != nil {
		return domain.Product{}, err
	}
	product, err := r.GetById(ctx, id)
	if err != nil {
		return domain.Product{}, err
	}
	return product, nil
}

func (r *repository) Delete(ctx context.Context, id int) (domain.Product, error) {
	db := r.db
	stmt, err := db.Prepare(Delete)
	if err != nil {
		return domain.Product{}, err
	}
	defer stmt.Close()
	product, err := r.GetById(ctx, id)
	if err != nil {
		return domain.Product{}, err
	}
	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return domain.Product{}, err
	}
	return product, nil
}
