package repository

import (
	"encoding/json"
	"goweb/4/tt/internal/domain"
)

type StubStore struct{}

func (s *StubStore) Read(data interface{}) error {
	var products []domain.Product

	product1 := domain.NewProduct(1, "Banana", 1.5, 2)
	product2 := domain.NewProduct(2, "Manzana", 1.25, 5)

	products = append(products, product1, product2)
	jsonProducts, _ := json.Marshal(products)

	return json.Unmarshal(jsonProducts, &data)
}

func (s *StubStore) Write(data interface{}) error {
	return nil
}
