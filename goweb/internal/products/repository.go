package products

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/bootcamp-go/wave-5-backpack/tree/Ramos_Andres/goweb/practica/internal/domain"
)

var lastId uint64
var products []domain.Product

type Repository interface {
	GetAll() ([]domain.Product, error)
	Store(id uint64, nombre string, color string, precio float64, stock uint64, codigo string, publicado bool, fechaCreacion string) (domain.Product, error)
	GetById(id uint64) (domain.Product, error)
	LastId() (uint64, error)
	UpdateTotal(id uint64, nombre string, color string, precio float64, stock uint64, codigo string, publicado bool, fechaCreacion string) (domain.Product, error)
	UpdatePartial(id uint64, nombre string, color string, precio float64, stock uint64, codigo string, publicado bool, fechaCreacion string) (domain.Product, error)
	Delete(id uint64) (domain.Product, error)
}

type repository struct{}

func getProductList() {
	data, err := os.ReadFile("./resources/products.json")
	if err != nil {
		panic(err)
	}
	var res []domain.Product
	err = json.Unmarshal(data, &res)
	if err != nil {
		panic(err)
	}
	lastId = uint64(len(res))
	products = res
}

func NewRepository() Repository {
	getProductList()
	return &repository{}
}

func (r *repository) GetAll() ([]domain.Product, error) {
	return products, nil
}

func (r *repository) GetById(id uint64) (domain.Product, error) {
	for _, product := range products {
		if product.Id == id {
			return product, nil
		}
	}
	return domain.Product{}, errors.New("no se encontró el producto")
}

func (r *repository) LastId() (uint64, error) {
	return lastId, nil
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

	products = append(products, product)
	lastId = id

	return product, nil
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
	for i, product := range products {
		if product.Id == id {
			products[i] = newProduct
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
	for i, product := range products {
		if product.Id == id {
			updated := partialUpdate(product, newProduct)
			products[i] = updated
			return updated, nil
		}
	}
	return domain.Product{}, fmt.Errorf("elemento %d no existe", id)
}

func (r *repository) Delete(id uint64) (domain.Product, error) {
	for i, product := range products {
		if product.Id == id {
			products = append(products[:i], products[i+1:]...)
			return product, nil
		}
	}
	return domain.Product{}, errors.New("no se encontró el producto")
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
