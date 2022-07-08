package products

import "fmt"

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"nombre"`
	Type  string  `json:"tipo"`
	Count int     `json:"cantidad"`
	Price float64 `json:"precio"`
}

var ps []Product
var lastID int

type Repository interface {
	LastID() (int, error)
	GetAll() ([]Product, error)
	Store(id int, name, producType string, count int, price float64) (Product, error)
	Update(id int, name, productType string, count int, price float64) (Product, error)
	UpdateName(id int, name string, price float64) (Product, error)
	Delete(id int) error
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAll() ([]Product, error) {
	if len(ps) == 0 {
		return nil, fmt.Errorf("no hay productos registrados")
	}

	return ps, nil
}

func (r *repository) Store(id int, name, producType string, count int, price float64) (Product, error) {
	p := Product{id, name, producType, count, price}
	ps = append(ps, p)
	lastID = p.ID

	return p, nil
}

func (r *repository) Update(id int, name, productType string, count int, price float64) (Product, error) {
	p := Product{Name: name, Type: productType, Count: count, Price: price}
	updated := false
	for i := range ps {
		if ps[i].ID == id {
			p.ID = id
			ps[i] = p
			updated = true
		}
	}

	if !updated {
		return Product{}, fmt.Errorf("producto %d no encontrado", id)
	}

	return p, nil
}

func (r *repository) UpdateName(id int, name string, price float64) (Product, error) {
	updated := false
	var p Product
	for i := range ps {
		if ps[i].ID == id {
			ps[i].Name = name
			ps[i].Price = price
			p = ps[i]
			updated = true
		}
	}

	if !updated {
		return Product{}, fmt.Errorf("producto %d no encontrado", id)
	}

	return p, nil
}

func (r *repository) Delete(id int) error {
	deleted := false
	var index int
	for i := range ps {
		if ps[i].ID == id {
			index = i
			deleted = true
		}
	}

	if !deleted {
		return fmt.Errorf("producto %d no encontrado", id)
	}

	ps = append(ps[:index], ps[index+1:]...)
	return nil
}

func (r *repository) LastID() (int, error) {
	return lastID, nil
}
