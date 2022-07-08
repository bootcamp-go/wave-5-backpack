package products

import (
	"goweb/internal/domain"
)

type Repository interface {
	GetAll() ([]domain.Product, error)
	Store(id int, name, color string, price, stock int, code string, published bool, date string) (domain.Product, error)
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

func (r *repository) Store(id int, name, color string, price, stock int, code string, published bool, date string) (domain.Product, error) {
	p := domain.Product{Id: id, Name: name, Color: color, Price: price, Stock: stock, Code: code, Published: published, Date: date}
	ps = append(ps, p)
	lastID = p.Id
	return p, nil
}
