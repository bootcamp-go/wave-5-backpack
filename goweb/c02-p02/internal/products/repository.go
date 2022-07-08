package products

import "github.com/bootcamp-go/wave-5-backpack/tree/lugo_abelardo/internal/domain"

type Repository interface {
	GetAll() ([]domain.Product, error)
	Store(id int, nombre string, cantidad int, precio float64) (domain.Product, error)
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

func (r *repository) Store(id int, nombre string, cantidad int, precio float64) (domain.Product, error) {
	p := domain.Product{ID: id, Name: nombre, Count: cantidad, Price: precio}
	ps = append(ps, p)
	lastID = p.ID
	return p, nil
}
