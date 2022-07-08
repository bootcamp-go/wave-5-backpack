package products

import (
	"clase2_parte2/internal/domain"
)

type Repository interface {
	GetAll() ([]domain.Product, error)
	Store(id int, nombre, tipo string, cantidad int, precio float64) (domain.Product, error)
	LastID() (int, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

var ps []domain.Product
var lastID int

func (r *repository) GetAll() ([]domain.Product, error) {
	return ps, nil
}

func (r *repository) LastID() (int, error) {
	return lastID, nil
}

func (r *repository) Store(id int, nombre, tipo string, cantidad int, precio float64) (domain.Product, error) {
	p := domain.Product{ID: id, Name: nombre, Type: tipo, Count: cantidad, Price: precio}
	ps = append(ps, p)
	lastID = p.ID
	return p, nil
}
