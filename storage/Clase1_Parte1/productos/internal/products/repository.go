package products

import (
	"Clase1_Parte1/productos/internal/domain"
	"database/sql"
)

const (
	ProductNotFound = "product %d not found"
	FailReading     = "cant read database"
	FailWriting     = "cant write database, error: %w"
)

type Repository interface {
	GetAll() ([]domain.Product, error)
	Store(nombre, color string, precio, stock int, codigo string, publicado bool, fechaCreacion string) (domain.Product, error)
	GetByID(id int) (domain.Product, error)
	GetByName(nombre string) ([]domain.Product, error)
	Update(id int, nombre, color string, precio, stock int, codigo string, publicado bool, fechaCreacion string) (domain.Product, error)
	UpdateNamePrice(id int, nombre string, precio int) (domain.Product, error)
	Delete(id int) (domain.Product, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll() ([]domain.Product, error) {
	var products []domain.Product
	db := r.db
	rows, err := db.Query("select * from products")
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

func (r *repository) Store(nombre, color string, precio, stock int, codigo string, publicado bool, fechaCreacion string) (domain.Product, error) {
	db := r.db
	stmt, err := db.Prepare("INSERT INTO products(nombre, color, precio, stock, codigo, publicado, fecha_creacion) VALUES( ?, ?, ?, ?, ?, ?, ?)")
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

func (r *repository) GetByID(id int) (domain.Product, error) {
	var product domain.Product
	db := r.db
	rows, err := db.Query("select * from products where id = ?", id)
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

func (r *repository) GetByName(nombre string) ([]domain.Product, error) {
	var products []domain.Product
	db := r.db
	rows, err := db.Query("select * from products where nombre = ?", nombre)
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

func (r *repository) Update(id int, nombre, color string, precio, stock int, codigo string, publicado bool, fechaCreacion string) (domain.Product, error) {
	db := r.db
	stmt, err := db.Prepare("UPDATE products SET nombre = ?, color = ?, precio = ?, stock = ?, codigo = ?, publicado = ?, fecha_creacion = ? WHERE id = ?")
	if err != nil {
		return domain.Product{}, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(nombre, color, precio, stock, codigo, publicado, fechaCreacion, id)
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

func (r *repository) UpdateNamePrice(id int, nombre string, precio int) (domain.Product, error) {
	db := r.db
	stmt, err := db.Prepare("UPDATE products SET nombre = ?, precio = ? WHERE id = ?")
	if err != nil {
		return domain.Product{}, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(nombre, precio, id)
	if err != nil {
		return domain.Product{}, err
	}
	product, err := r.GetByID(id)
	if err != nil {
		return domain.Product{}, err
	}
	return product, nil
}

func (r *repository) Delete(id int) (domain.Product, error) {
	db := r.db
	stmt, err := db.Prepare("DELETE FROM products WHERE id = ?")
	if err != nil {
		return domain.Product{}, err
	}
	defer stmt.Close()
	product, err := r.GetByID(id)
	if err != nil {
		return domain.Product{}, err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return domain.Product{}, err
	}
	return product, nil
}
