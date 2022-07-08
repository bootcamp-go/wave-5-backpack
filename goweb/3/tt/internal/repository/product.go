package repository

import (
	"fmt"
	"goweb/3/tt/internal/domain"
)

var products []domain.Product

func (r *repository) ReadAll() ([]domain.Product, error) {
	if err := r.db.Read(&products); err != nil {
		return []domain.Product{}, nil
	}

	return products, nil
}

func (r *repository) Read(id int) (domain.Product, error) {
	if err := r.db.Read(&products); err != nil {
		return domain.Product{}, nil
	}

	for _, product := range products {
		if product.ID == id {
			return product, nil
		}
	}
	return domain.Product{}, fmt.Errorf("no se encontro el producto de id %d", id)
}

func (r *repository) Create(name string, price float64, quantity int) (domain.Product, error) {
	if err := r.db.Read(&products); err != nil {
		return domain.Product{}, err
	}

	lastId, err := r.LastId()
	if err != nil {
		return domain.Product{}, err
	}
	newId := lastId + 1

	newProduct := domain.NewProduct(newId, name, price, quantity)

	products = append(products, newProduct)

	if err := r.db.Write(products); err != nil {
		return domain.Product{}, err
	}

	return newProduct, nil
}

func (r *repository) Update(id int, name string, price float64, quantity int) (domain.Product, error) {
	updatedProduct := domain.NewProduct(id, name, price, quantity)

	if err := r.db.Read(&products); err != nil {
		return domain.Product{}, err
	}

	for i, product := range products {
		if product.ID == id {
			products[i] = updatedProduct

			if err := r.db.Write(products); err != nil {
				return domain.Product{}, err
			}

			return updatedProduct, nil
		}
	}

	return domain.Product{}, fmt.Errorf("no se encontro el producto de id %d", id)
}

func (r *repository) UpdateNamePrice(id int, name string, price float64) (domain.Product, error) {
	if err := r.db.Read(&products); err != nil {
		return domain.Product{}, err
	}

	for i, product := range products {
		if product.ID == id {
			product.Name = name
			product.Price = price

			products[i] = product

			if err := r.db.Write(products); err != nil {
				return domain.Product{}, err
			}

			return product, nil
		}
	}

	return domain.Product{}, fmt.Errorf("no se encontro el producto de id %d", id)
}

func (r *repository) Delete(id int) error {
	if err := r.db.Read(&products); err != nil {
		return err
	}

	for i, product := range products {
		if product.ID == id {
			if i != len(products)-1 {
				products = append(products[:i], products[i+1:]...)
			} else {
				products = products[:i]
			}

			if err := r.db.Write(products); err != nil {
				return err
			}
			return nil
		}
	}

	return fmt.Errorf("no se encontro el producto de id %d", id)
}

func (r *repository) LastId() (int, error) {
	if err := r.db.Read(&products); err != nil {
		return 0, err
	}

	if len(products) == 0 {
		return 0, nil
	}

	return products[len(products)-1].ID, nil
}
