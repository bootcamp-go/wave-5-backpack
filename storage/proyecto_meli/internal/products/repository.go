package products

import (
	"database/sql"
	"fmt"
	"proyecto_meli/internal/domain"
)

type Repository interface {
	GetAll() ([]domain.Product, error)
	GetById(id int) (domain.Product, error)
	FilterList(id int, name, color string, price float64, stock int, codigo string, publicado bool, fecha string) ([]domain.Product, error)
	Store(p domain.Product) (domain.Product, error)
	LastID() (int, error)
	Update(id int, name, color string, price float64, stock int, codigo string, publicado bool, fecha string) (domain.Product, error)
	Delete(id int) error
	Update_Name_Price(id int, name string, price float64) (domain.Product, error)
	GetProductByName(name string) ([]domain.Product, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Store(p domain.Product) (domain.Product, error) {
	query := "INSERT INTO products(name, color, price, stock, code, publish, create_date) VALUES (?,?,?,?,?,?,?)"
	stmt, err := r.db.Prepare(query)

	if err != nil {
		return domain.Product{}, err
	}

	defer stmt.Close()

	res, err := stmt.Exec(p.Nombre, p.Color, p.Precio, p.Stock, p.Codigo, p.Publicado, p.FechaCreacion)
	fmt.Println(p)
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

func (r *repository) GetProductByName(name string) ([]domain.Product, error) {
	query := "SELECT id, name, color, price, stock, code, publish, create_date FROM products WHERE name=?"
	rows, err := r.db.Query(query, name)
	fmt.Println(name)
	if err != nil {
		return nil, err
	}
	products := []domain.Product{}
	for rows.Next() {
		p := domain.Product{}
		publish := 0
		if err := rows.Scan(&p.Id, &p.Nombre, &p.Color, &p.Precio, &p.Stock, &p.Codigo, &publish, &p.FechaCreacion); err != nil {
			fmt.Println(err)
			return nil, err
		}
		p.Publicado = publish != 0
		products = append(products, p)
	}
	return products, nil
}

func (r *repository) GetAll() ([]domain.Product, error) {
	return nil, nil
}

func (r *repository) GetById(id int) (domain.Product, error) {
	return domain.Product{}, nil
}
func (r *repository) FilterList(id int, name, color string, price float64, stock int, codigo string, publicado bool, fecha string) ([]domain.Product, error) {
	return nil, nil
}

func (r *repository) LastID() (int, error) {
	return 0, nil
}

func (r *repository) Update(id int, name, color string, price float64, stock int, codigo string, publicado bool, fecha string) (domain.Product, error) {
	return domain.Product{}, nil
}

func (r *repository) Delete(id int) error {
	return nil
}

func (r *repository) Update_Name_Price(id int, name string, price float64) (domain.Product, error) {
	return domain.Product{}, nil
}
