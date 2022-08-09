package products

import (
	"database/sql"

	"github.com/bootcamp-go/wave-5-backpack/tree/Ramos_Andres/goweb/practica/internal/domain"
)

var (
	createStmt    = "INSERT INTO PRODUCTS (Name, Color, Price, Stock, Code, Published, Created_at) VALUES (?, ?, ?, ?, ?, ?, CURDATE())"
	getByNameStmt = "SELECT ID, Name, Color, Price, Stock, Code, Published, Created_at FROM PRODUCTS WHERE name = ?"
)

type Repository interface {
	Store(domain.Product) (domain.Product, error)
	GetAll() ([]domain.Product, error)
	GetById(id uint64) (domain.Product, error)
	GetByName(name string) (domain.Product, error)
	Update(domain.Product) (domain.Product, error)
	Delete(id uint64) (domain.Product, error)
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
	db := r.db
	stmt, err := db.Prepare(createStmt)
	if err != nil {
		return domain.Product{}, err
	}
	defer stmt.Close()
	sqlRes, err := stmt.Exec(product.Name, product.Color, product.Price, product.Stock, product.Code, product.Published)
	if err != nil {
		return domain.Product{}, err
	}
	insertedId, err := sqlRes.LastInsertId()
	if err != nil {
		return domain.Product{}, err
	}
	product.Id = uint64(insertedId)
	return product, nil
}

func (r *repository) GetAll() ([]domain.Product, error) {
	return []domain.Product{}, nil
}

func (r *repository) GetById(id uint64) (domain.Product, error) {
	return domain.Product{}, nil
}

func (r *repository) GetByName(name string) (domain.Product, error) {
	db := r.db
	rows, err := db.Query(getByNameStmt, name)
	if err != nil {
		return domain.Product{}, err
	}
	defer rows.Close()

	prList := []domain.Product{}

	for rows.Next() {
		var pr domain.Product
		if err := rows.Scan(&pr.Id, &pr.Name, &pr.Color, &pr.Price, &pr.Stock, &pr.Code, &pr.Published, &pr.Created_at); err != nil {
			return prList[0], err
		}
		prList = append(prList, pr)
	}
	if err := rows.Err(); err != nil {
		return domain.Product{}, err
	}
	return prList[0], nil
}

func (r *repository) Update(domain.Product) (domain.Product, error) {
	return domain.Product{}, nil
}

func (r *repository) Delete(id uint64) (domain.Product, error) {
	return domain.Product{}, nil
}

func partialUpdate(oldProduct domain.Product, newProduct domain.Product) domain.Product {
	if newProduct.Name != "" {
		oldProduct.Name = newProduct.Name
	}

	if newProduct.Color != "" {
		oldProduct.Color = newProduct.Color
	}

	if newProduct.Price != 0 {
		oldProduct.Price = newProduct.Price
	}

	oldProduct.Stock = newProduct.Stock

	if newProduct.Code != "" {
		oldProduct.Code = newProduct.Code
	}
	return oldProduct
}
