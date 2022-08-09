package products

import (
	"bootcamp/wave-5-backpack/storage/internal/domain"
	"context"
	"database/sql"
)

type Repository interface {
	Store(domain.Product) (domain.Product, error)
	GetOne(id int) (domain.Product, error)
	Update(product domain.Product) (domain.Product, error)

	GetAll() ([]domain.Product, error)
	Delete(id int) error
	GetFullData(id int) ([]domain.ProductAndWarehouse, error)
	GetOneWithcontext(ctx context.Context, id int) (domain.Product, error)
}

const (
	GetProduct    string = "SELECT p.id, p.name, p.type, p.count, p.price, p.warehouse_id FROM products p WHERE id = ?"
	InsertProduct string = "INSERT INTO products(name, type, count, price, warehouse_id) VALUES( ?, ?, ?, ?, ? )"
	UpdateProduct string = "UPDATE products SET name = ?, type = ?, count = ?, price = ?, p.warehouse_id = ? WHERE id = ?"
	DeleteProduct string = "DELETE FROM products WHERE id = ?"

	GetFullData string = `SELECT p.id, p.name, p.type, p.count, p.price, w.id, w.name, w.address
	FROM products AS p INNER JOIN warehouses AS w ON p.warehouse_id = w.id 
	WHERE p.id = ?`
	GetAllProducts  string = "SELECT p.id, p.name, p.type, p.count, p.price FROM products p"
	GetProductSleep string = "SELECT SLEEP(10) FROM DUAL where 0 < ?"
)

type repository struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Store(product domain.Product) (domain.Product, error) {
	stmt, err := r.db.Prepare(InsertProduct)
	if err != nil {
		return domain.Product{}, err
	}
	defer stmt.Close()
	var result sql.Result
	result, err = stmt.Exec(product.Name, product.Type, product.Count, product.Price, product.Warehouse)
	if err != nil {
		return domain.Product{}, err
	}
	insertedId, _ := result.LastInsertId()
	product.ID = int(insertedId)
	return product, nil
}

func (r *repository) GetOne(id int) (domain.Product, error) {
	var product domain.Product
	row := r.db.QueryRow(GetProduct, id)
	if err := row.Scan(&product.ID, &product.Name, &product.Type, &product.Count, &product.Price, &product.Warehouse); err != nil {
		return domain.Product{}, err
	}
	return product, nil
}

func (r *repository) Update(product domain.Product) (domain.Product, error) {
	stmt, err := r.db.Prepare(UpdateProduct)
	if err != nil {
		return domain.Product{}, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(product.Name, product.Type, product.Count, product.Price, product.Warehouse, product.ID)
	if err != nil {
		return domain.Product{}, err
	}
	return product, nil

}

func (r *repository) GetAll() ([]domain.Product, error) {
	var products []domain.Product
	rows, err := r.db.Query(GetAllProducts)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var product domain.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Type, &product.Count, &product.Price, &product.Warehouse); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (r *repository) Delete(id int) error {
	stmt, err := r.db.Prepare(DeleteProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) GetFullData(id int) ([]domain.ProductAndWarehouse, error) {
	var productAndWarehouses []domain.ProductAndWarehouse
	rows, err := r.db.Query(GetFullData, id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var productAndWarehouse domain.ProductAndWarehouse
		err := rows.Scan(
			&productAndWarehouse.ID,
			&productAndWarehouse.Name,
			&productAndWarehouse.Type,
			&productAndWarehouse.Count,
			&productAndWarehouse.Price,
			&productAndWarehouse.Warehouse,
			&productAndWarehouse.Warehouse.ID,
			&productAndWarehouse.Warehouse.Name,
			&productAndWarehouse.Warehouse.Address)
		if err != nil {
			return nil, err
		}
		productAndWarehouses = append(productAndWarehouses, productAndWarehouse)
	}
	return productAndWarehouses, nil
}

func (r *repository) GetOneWithcontext(ctx context.Context, id int) (domain.Product, error) {
	var product domain.Product
	rows, err := r.db.QueryContext(ctx, GetProduct, id)
	if err != nil {
		return product, err
	}
	for rows.Next() {
		if err := rows.Scan(&product.ID, &product.Name, &product.Type, &product.Count, &product.Price, &product.Warehouse); err != nil {
			return product, err
		}
	}
	return product, nil
}
