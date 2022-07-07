package users

import (
	"github.com/bootcamp-go/wave-5-backpack/goweb/internal/domain"
)

type Repository interface {
	GetAll() ([]domain.Users, error)
	Store(id, age int, name, lastName, email, creationDate string, height float64, active bool) (domain.Users, error)
	LastID() (int, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

var ps []domain.Users
var lastID int

func (r *repository) GetAll() ([]domain.Users, error) {
	return ps, nil
}

func (r *repository) LastID() (int, error) {
	return lastID, nil
}

func (r *repository) Store(id, age int, name, lastName, email, creationDate string, height float64, active bool) (domain.Users, error) {
	p := domain.Users{
		Id: id, Name: name, LastName: lastName, Email: email, Age: age, Height: height, Active: active, CreationDate: creationDate}
	ps = append(ps, p)
	lastID = p.Id
	return p, nil
}
