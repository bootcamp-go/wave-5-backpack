package products

import (
	"clase4_parte1/internal/domain"
	"clase4_parte1/pkg/store"
	"fmt"
)

type Repository interface {
	GetAll() ([]domain.Product, error)
	Store(id int, nombre, tipo string, cantidad int, precio float64) (domain.Product, error)
	LastID() (int, error)
	Update(id int, name, productType string, count int, price float64) (domain.Product, error)
	UpdateName(id int, name string) (domain.Product, error)
	Delete(id int) error
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll() ([]domain.Product, error) {
	var ps []domain.Product
	err := r.db.Read(&ps)
	if err != nil {
		return nil, err
	}
	return ps, nil
}

func (r *repository) LastID() (int, error) {
	var ps []domain.Product
	if err := r.db.Read(&ps); err != nil {
		return 0, err
	}
	if len(ps) == 0 {
		return 0, nil
	}

	return ps[len(ps)-1].ID, nil

}

func (r *repository) Store(id int, nombre, tipo string, cantidad int, precio float64) (domain.Product, error) {
	var ps []domain.Product
	r.db.Read(&ps)
	p := domain.Product{ID: id, Name: nombre, Type: tipo, Count: cantidad, Price: precio}
	ps = append(ps, p)
	if err := r.db.Write(ps); err != nil {
		return domain.Product{}, err
	}
	return p, nil

}

func (r *repository) UpdateName(id int, name string) (domain.Product, error) {
	var ps []domain.Product
	if err := r.db.Read(&ps); err != nil {
		return domain.Product{}, nil
	}
	var i int
	for i = 0; i < len(ps) && ps[i].ID != id; i++ {
	}
	if i < len(ps) {
		ps[i].Name = name
		r.db.Write(ps)
		return ps[i], nil
	}
	return domain.Product{}, fmt.Errorf("not found")

}

func (r *repository) Delete(id int) error {
	var ps []domain.Product
	if err := r.db.Read(&ps); err != nil {
		return nil
	}

	var i int
	for i = 0; i < len(ps) && ps[i].ID != id; i++ {
	}
	if i < len(ps) {
		ps = append(ps[:i], ps[i+1:]...)
		r.db.Write(ps)
		return nil
	}
	return fmt.Errorf("not found")
}

func (r *repository) Update(id int, name, productType string, count int, price float64) (domain.Product, error) {
	var ps []domain.Product

	if err := r.db.Read(&ps); err != nil {
		return domain.Product{}, err
	}

	p := domain.Product{Name: name, Type: productType, Count: count, Price: price}
	updated := false
	for i := range ps {
		if ps[i].ID == id {
			p.ID = id
			ps[i] = domain.Product(p)
			updated = true
		}
	}

	if !updated {
		return domain.Product{}, fmt.Errorf("producto %d no se pudo actualizar", id)
	}

	if err := r.db.Write(ps); err != nil {
		return domain.Product{}, err
	}

	return p, nil
}
