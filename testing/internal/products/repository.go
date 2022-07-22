package products

import (
	"errors"
	"fmt"
	"testing/internal/domain"
	"testing/pkg/file"
)

var lastId uint64

type Repository interface {
	Store(id uint64, nombre string, color string, precio float64, stock uint64, codigo string, publicado bool, fechaCreacion string) (domain.Product, error)
	GetAll() ([]domain.Product, error)
	GetById(id uint64) (domain.Product, error)
	UpdateTotal(id uint64, nombre string, color string, precio float64, stock uint64, codigo string, publicado bool, fechaCreacion string) (domain.Product, error)
	UpdatePartial(id uint64, nombre string, color string, precio float64, stock uint64, codigo string, publicado bool, fechaCreacion string) (domain.Product, error)
	Delete(id uint64) (domain.Product, error)
	LastId() (uint64, error)
}

type repository struct {
	db file.File
}

func NewRepository(db file.File) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Store(id uint64, name string, color string, price float64, stock uint64, code string, published bool, createdAt string) (domain.Product, error) {
	product := domain.Product{
		Id:         id,
		Name:       name,
		Color:      color,
		Price:      price,
		Stock:      stock,
		Code:       code,
		Published:  published,
		Created_at: createdAt,
	}

	var products []domain.Product
	if err := r.db.Read(&products); err != nil {
		return domain.Product{}, err
	}

	products = append(products, product)
	lastId = id
	if err := r.db.Write(&products); err != nil {
		return domain.Product{}, err
	}
	return product, nil
}

func (r *repository) GetAll() ([]domain.Product, error) {
	var products []domain.Product
	if err := r.db.Read(&products); err != nil {
		return []domain.Product{}, err
	}
	return products, nil
}

func (r *repository) GetById(id uint64) (domain.Product, error) {
	var products []domain.Product
	if err := r.db.Read(&products); err != nil {
		return domain.Product{}, err
	}
	for _, product := range products {
		if product.Id == id {
			return product, nil
		}
	}
	return domain.Product{}, errors.New("no se encontró el producto")
}

func (r *repository) UpdateTotal(id uint64, name string, color string, price float64, stock uint64, code string, published bool, createdAt string) (domain.Product, error) {
	newProduct := domain.Product{
		Id:         id,
		Name:       name,
		Color:      color,
		Price:      price,
		Stock:      stock,
		Code:       code,
		Published:  published,
		Created_at: createdAt,
	}
	var products []domain.Product
	if err := r.db.Read(&products); err != nil {
		return domain.Product{}, err
	}
	for i, product := range products {
		if product.Id == id {
			products[i] = newProduct
			if err := r.db.Write(&products); err != nil {
				return domain.Product{}, err
			}
			return newProduct, nil
		}
	}
	return domain.Product{}, fmt.Errorf("elemento %d no existe", id)
}

func (r *repository) UpdatePartial(id uint64, name string, color string, price float64, stock uint64, code string, published bool, createdAt string) (domain.Product, error) {
	newProduct := domain.Product{
		Id:         id,
		Name:       name,
		Color:      color,
		Price:      price,
		Stock:      stock,
		Code:       code,
		Published:  published,
		Created_at: createdAt,
	}
	var products []domain.Product
	if err := r.db.Read(&products); err != nil {
		return domain.Product{}, err
	}
	for i, product := range products {
		if product.Id == id {
			updated := partialUpdate(product, newProduct)
			products[i] = updated
			if err := r.db.Write(&products); err != nil {
				return domain.Product{}, err
			}
			return updated, nil
		}
	}
	return domain.Product{}, fmt.Errorf("elemento %d no existe", id)
}

func (r *repository) Delete(id uint64) (domain.Product, error) {
	var products []domain.Product
	if err := r.db.Read(&products); err != nil {
		return domain.Product{}, err
	}
	for i, product := range products {
		if product.Id == id {
			products = append(products[:i], products[i+1:]...)
			if err := r.db.Write(&products); err != nil {
				return domain.Product{}, err
			}
			return product, nil
		}
	}
	return domain.Product{}, errors.New("no se encontró el producto")
}

func (r *repository) LastId() (uint64, error) {
	prList, err := r.GetAll()
	if err != nil {
		return 0, err
	}
	return prList[len(prList)-1].Id, nil
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
