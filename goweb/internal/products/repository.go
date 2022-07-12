package products

import (
	"fmt"
	"goweb/internal/domain"
	"goweb/pkg/store"
)

const (
	ProductNotFound = "product %d not found"
	FailReading     = "cant read database"
	FailWriting     = "cant write database, error: %w"
)

type Repository interface {
	GetAll() ([]domain.Product, error)
	Store(id int, name, color string, price, stock int, code string, published bool, date string) (domain.Product, error)
	LastID() (int, error)
	Update(id int, name, color string, price, stock int, code string, published bool, date string) (domain.Product, error)
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
	if err := r.db.Read(&ps); err != nil {
		return nil, fmt.Errorf(FailReading)
	}
	return ps, nil
}

func (r *repository) LastID() (int, error) {
	var ps []domain.Product
	if err := r.db.Read(&ps); err != nil {
		return 0, fmt.Errorf(FailReading)
	}
	if len(ps) == 0 {
		return 0, nil
	}

	return ps[len(ps)-1].Id, nil
}

func (r *repository) Store(id int, name, color string, price, stock int, code string, published bool, date string) (domain.Product, error) {
	var ps []domain.Product

	if err := r.db.Read(&ps); err != nil {
		return domain.Product{}, fmt.Errorf(FailReading)
	}

	p := domain.Product{Id: id, Name: name, Color: color, Price: price, Stock: stock, Code: code, Published: published, Date: date}

	ps = append(ps, p)

	if err := r.db.Write(ps); err != nil {
		return domain.Product{}, fmt.Errorf(FailWriting, err)
	}

	return p, nil
}

func (r *repository) Update(id int, name, color string, price, stock int, code string, published bool, date string) (domain.Product, error) {
	var ps []domain.Product

	if err := r.db.Read(&ps); err != nil {
		return domain.Product{}, fmt.Errorf(FailReading)
	}

	p := domain.Product{Name: name, Color: color, Price: price, Stock: stock, Code: code, Published: published, Date: date}
	updated := false
	for i := range ps {
		if ps[i].Id == id {
			p.Id = id
			ps[i] = p
			updated = true
		}
	}

	if !updated {
		return domain.Product{}, fmt.Errorf(ProductNotFound, id)
	}

	if err := r.db.Write(ps); err != nil {
		return domain.Product{}, fmt.Errorf(FailWriting, err)
	}

	return p, nil
}

func (r *repository) UpdateName(id int, name string) (domain.Product, error) {
	var ps []domain.Product

	if err := r.db.Read(&ps); err != nil {
		return domain.Product{}, fmt.Errorf(FailReading)
	}

	updated := false
	var p domain.Product
	for i := range ps {
		if ps[i].Id == id {
			ps[i].Name = name
			p = ps[i]
			updated = true
		}
	}

	if !updated {
		return domain.Product{}, fmt.Errorf(ProductNotFound, id)
	}

	if err := r.db.Write(ps); err != nil {
		return domain.Product{}, fmt.Errorf(FailWriting, err)
	}

	return p, nil
}

func (r *repository) Delete(id int) error {
	var ps []domain.Product

	if err := r.db.Read(&ps); err != nil {
		return fmt.Errorf(FailReading)
	}

	deleted := false
	var index int
	for i := range ps {
		if ps[i].Id == id {
			index = i
			deleted = true
		}
	}

	if !deleted {
		return fmt.Errorf(ProductNotFound, id)
	}

	ps = append(ps[:index], ps[index+1:]...)

	if err := r.db.Write(ps); err != nil {
		return fmt.Errorf(FailWriting, err)
	}
	return nil
}
