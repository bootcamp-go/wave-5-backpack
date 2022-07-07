package repository

import (
	"goweb/2/tt/internal/domain"
)

var id int
var products []domain.Product

func (r *repository) SelectAll() ([]domain.Product, error) {
	return products, nil
}

func (r *repository) Insert(name string, price float64) (domain.Product, error) {
	newId := r.LastId()
	newProduct := domain.Product{ID: newId, Name: name, Price: price}

	products = append(products, newProduct)

	return newProduct, nil
}

func (r *repository) LastId() int {
	id++
	return id
}
