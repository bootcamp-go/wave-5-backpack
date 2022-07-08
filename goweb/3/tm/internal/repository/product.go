package repository

import (
	"fmt"
	"goweb/3/tm/internal/domain"
)

var id int
var products []domain.Product

func (r *repository) ReadAll() ([]domain.Product, error) {
	return products, nil
}

func (r *repository) Read(id int) (domain.Product, error) {
	for _, product := range products {
		if product.ID == id {
			return product, nil
		}
	}
	return domain.Product{}, fmt.Errorf("no se encontro el producto de id %d", id)
}

func (r *repository) Create(name string, price float64, quantity int) (domain.Product, error) {
	newId := r.LastId()
	newProduct := domain.NewProduct(newId, name, price, quantity)

	products = append(products, newProduct)

	return newProduct, nil
}

func (r *repository) Update(id int, name string, price float64, quantity int) (domain.Product, error) {
	updatedProduct := domain.NewProduct(id, name, price, quantity)

	for i, product := range products {
		if product.ID == id {
			products[i] = updatedProduct

			return updatedProduct, nil
		}
	}
	return domain.Product{}, fmt.Errorf("no se encontro el producto de id %d", id)
}

func (r *repository) UpdateNamePrice(id int, name string, price float64) (domain.Product, error) {
	for i, product := range products {
		if product.ID == id {
			product.Name = name
			product.Price = price

			products[i] = product
			return product, nil
		}
	}
	return domain.Product{}, fmt.Errorf("no se encontro el producto de id %d", id)
}

func (r *repository) Delete(id int) error {
	for i, product := range products {
		if product.ID == id {
			if i != len(products)-1 {
				products = append(products[:i], products[i+1:]...)
			} else {
				products = products[:i]
			}
			return nil
		}
	}

	return fmt.Errorf("no se encontro el producto de id %d", id)
}

func (r *repository) LastId() int {
	id++
	return id
}
