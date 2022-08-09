package products

import (
	"Clase1_Parte1/productos/internal/domain"
	"context"
	"database/sql"
)

const (
	ProductNotFound = "product %d not found"
	FailReading     = "cant read database"
	FailWriting     = "cant write database, error: %w"
)

type Repository interface {
	GetAll(ctx context.Context) ([]domain.Product, error)
	Store(ctx context.Context, nombre, color string, precio, stock int, codigo string, publicado bool, fechaCreacion string) (domain.Product, error)
	GetByID(ctx context.Context, id int) (domain.Product, error)
	GetByName(ctx context.Context, nombre string) ([]domain.Product, error)
	Update(ctx context.Context, id int, nombre, color string, precio, stock int, codigo string, publicado bool, fechaCreacion string) (domain.Product, error)
	UpdateNamePrice(ctx context.Context, id int, nombre string, precio int) (domain.Product, error)
	Delete(ctx context.Context, id int) (domain.Product, error)
}

const (
	GetAllProducts         = "SELECT p.id, p.nombre, p.color, p.precio, p.stock, p.codigo, p.publicado, p.fecha_creacion FROM products p"
	SaveProduct            = "INSERT INTO products(nombre, color, precio, stock, codigo, publicado, fecha_creacion) VALUES( ?, ?, ?, ?, ?, ?, ?)"
	GetProdyctByID         = "SELECT p.id, p.nombre, p.color, p.precio, p.stock, p.codigo, p.publicado, p.fecha_creacion FROM products p WHERE p.id = ?"
	GetProductByName       = "SELECT p.id, p.nombre, p.color, p.precio, p.stock, p.codigo, p.publicado, p.fecha_creacion FROM products p WHERE p.nombre = ?"
	UpdateProduct          = "UPDATE products SET nombre = ?, color = ?, precio = ?, stock = ?, codigo = ?, publicado = ?, fecha_creacion = ? WHERE id = ?"
	UpdateProductNamePrice = "UPDATE products SET nombre = ?, precio = ? WHERE id = ?"
	DeleteProduct          = "DELETE FROM products WHERE id = ?"
	GetFullData            = `SELECT p.id, p.name, p.type, p.count, p.price, w.id, w.name, w.address
	FROM products AS p INNER JOIN warehouses AS w ON p.warehouse_id = w.id WHERE p.id = ?`
	GetProductSleep = "SELECT sleep(10) FROM DUAL where 0 < ?"
)

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll(ctx context.Context) ([]domain.Product, error) {
	var products []domain.Product
	db := r.db
	rows, err := db.Query(GetAllProducts)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var product domain.Product
		if err := rows.Scan(&product.Id, &product.Nombre, &product.Color, &product.Precio, &product.Stock, &product.Codigo, &product.Publicado, &product.FechaCreacion); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (r *repository) Store(ctx context.Context, nombre, color string, precio, stock int, codigo string, publicado bool, fechaCreacion string) (domain.Product, error) {
	db := r.db
	stmt, err := db.Prepare(SaveProduct)
	if err != nil {
		return domain.Product{}, err
	}
	defer stmt.Close()
	var result sql.Result
	result, err = stmt.Exec(nombre, color, precio, stock, codigo, publicado, fechaCreacion)
	if err != nil {
		return domain.Product{}, err
	}
	insertedId, _ := result.LastInsertId()
	product := domain.Product{
		Id:            int(insertedId),
		Nombre:        nombre,
		Color:         color,
		Precio:        precio,
		Stock:         stock,
		Codigo:        codigo,
		Publicado:     publicado,
		FechaCreacion: fechaCreacion,
	}
	return product, nil
}

func (r *repository) GetByID(ctx context.Context, id int) (domain.Product, error) {
	var product domain.Product
	db := r.db
	rows, err := db.Query(GetProdyctByID, id)
	if err != nil {
		return product, err
	}
	for rows.Next() {
		if err := rows.Scan(&product.Id, &product.Nombre, &product.Color, &product.Precio, &product.Stock, &product.Codigo, &product.Publicado, &product.FechaCreacion); err != nil {
			return product, err
		}
	}
	return product, nil
}

func (r *repository) GetByName(ctx context.Context, nombre string) ([]domain.Product, error) {
	var products []domain.Product
	db := r.db
	rows, err := db.Query(GetProductByName, nombre)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var product domain.Product
		if err := rows.Scan(&product.Id, &product.Nombre, &product.Color, &product.Precio, &product.Stock, &product.Codigo, &product.Publicado, &product.FechaCreacion); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (r *repository) Update(ctx context.Context, id int, nombre, color string, precio, stock int, codigo string, publicado bool, fechaCreacion string) (domain.Product, error) {
	db := r.db
	stmt, err := db.Prepare(UpdateProduct)
	if err != nil {
		return domain.Product{}, err
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, nombre, color, precio, stock, codigo, publicado, fechaCreacion, id)
	if err != nil {
		return domain.Product{}, err
	}
	product := domain.Product{
		Id:            id,
		Nombre:        nombre,
		Color:         color,
		Precio:        precio,
		Stock:         stock,
		Codigo:        codigo,
		Publicado:     publicado,
		FechaCreacion: fechaCreacion,
	}
	return product, nil
}

func (r *repository) UpdateNamePrice(ctx context.Context, id int, nombre string, precio int) (domain.Product, error) {
	db := r.db
	stmt, err := db.Prepare(UpdateProductNamePrice)
	if err != nil {
		return domain.Product{}, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(nombre, precio, id)
	if err != nil {
		return domain.Product{}, err
	}
	product, err := r.GetByID(ctx, id)
	if err != nil {
		return domain.Product{}, err
	}
	return product, nil
}

func (r *repository) Delete(ctx context.Context, id int) (domain.Product, error) {
	db := r.db
	stmt, err := db.Prepare(DeleteProduct)
	if err != nil {
		return domain.Product{}, err
	}
	defer stmt.Close()
	product, err := r.GetByID(ctx, id)
	if err != nil {
		return domain.Product{}, err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return domain.Product{}, err
	}
	return product, nil
}
