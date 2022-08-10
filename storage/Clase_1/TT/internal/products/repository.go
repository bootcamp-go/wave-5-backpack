package products

import (
	"context"
	"database/sql"
	"errors"
	"goweb/internal/domain"
	"log"
)

const (
	Store           string = "INSERT INTO products (name,type,count,price,id_warehouse) VALUES(?,?,?,?,?)"
	GetByName       string = "SELECT id,name,type,count,price FROM products WHERE name=?"
	GetAll          string = "SELECT id,name,type,count,price,id_warehouse FROM products"
	GetFullDataById string = "SELECT p.id,p.name,p.type,p.count,p.price,p.id_warehouse, w.id,w.name,w.adress FROM products AS p INNER JOIN warehouses AS w ON p.id_warehouse=w.id WHERE p.id=? "
	queryUpdate     string = "UPDATE products SET name=?, type=?, count=?,price=?,id_warehouse=? WHERE id=?"
	queryGetOnE     string = "SELECT id,name,type,count,price,id_warehouse FROM products WHERE id=?"
)

type Repository interface {
	GetByName(name string) (domain.Product, error)
	Store(product domain.Product) (domain.Product, error)
	GetAll(ctx context.Context) ([]domain.Product, error)
	GetOne(ctx context.Context, id int) (domain.Product, error)
	GetOneFullData(ctx context.Context, id int) (domain.ProductAndWarehouse, error)
	Update(context.Context, domain.Product) (domain.Product, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Store(product domain.Product) (domain.Product, error) {
	stmt, err := r.db.Prepare(Store)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	var result sql.Result
	result, err = stmt.Exec(product.Name, product.Type, product.Count, product.Price, product.Id_warehouse)
	if err != nil {
		return domain.Product{}, err
	}

	insertedId, _ := result.LastInsertId()
	product.Id = int(insertedId)

	return product, nil
}

func (r *repository) GetByName(name string) (domain.Product, error) {
	var product domain.Product

	rows, err := r.db.Query(GetByName, name)
	if err != nil {
		log.Println(err)
		return product, err
	}

	for rows.Next() {
		if err := rows.Scan(&product.Name, &product.Type, &product.Count, &product.Price, &product.Id); err != nil {
			log.Println(err)
			return product, err
		}
	}

	if product.Id == 0 {
		return product, errors.New("Product not found")
	}
	return product, nil
}

func (r *repository) GetAll(ctx context.Context) ([]domain.Product, error) {
	var products []domain.Product

	rows, err := r.db.QueryContext(ctx, GetAll)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	for rows.Next() {
		var product domain.Product
		if err := rows.Scan(&product.Id, &product.Name, &product.Type, &product.Count, &product.Price, &product.Id_warehouse); err != nil {
			log.Fatal(err)
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func (r *repository) GetOneFullData(ctx context.Context, id int) (domain.ProductAndWarehouse, error) {
	var product domain.ProductAndWarehouse

	rows, err := r.db.QueryContext(ctx, GetFullDataById, id)
	if err != nil {
		log.Println(err)
		return domain.ProductAndWarehouse{}, err
	}

	for rows.Next() {
		if err := rows.Scan(&product.Id, &product.Name, &product.Type, &product.Count, &product.Price, &product.Id_warehouse, &product.Warehouse.Id, &product.Warehouse.Name, &product.Warehouse.Adress); err != nil {
			log.Fatal(err)
			return domain.ProductAndWarehouse{}, err
		}
	}

	return product, nil
}

func (r *repository) GetOne(ctx context.Context, id int) (domain.Product, error) {
	var product domain.Product

	rows, err := r.db.QueryContext(ctx, queryGetOnE, id)
	if err != nil {
		log.Println(err)
		return domain.Product{}, err
	}

	for rows.Next() {
		if err := rows.Scan(&product.Id, &product.Name, &product.Type, &product.Count, &product.Price, &product.Id_warehouse); err != nil {
			log.Fatal(err)
			return domain.Product{}, err
		}
	}

	return product, nil
}

func (r *repository) Update(ctx context.Context, product domain.Product) (domain.Product, error) {

	stmt, err := r.db.Prepare(queryUpdate)
	if err != nil {
		log.Println(err)
		return domain.Product{}, err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, product.Name, product.Type, product.Count, product.Price, product.Id_warehouse, product.Id)
	if err != nil {
		return domain.Product{}, err
	}

	return product, nil
}
