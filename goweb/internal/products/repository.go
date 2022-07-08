package products

import (
	"fmt"
	"goweb/internal/domain"
	"goweb/pkg/store"
	"time"
)

const(
	ERROR_WRITING = "cannot write file"
	ERROR_READING = "cannot read the file"
)

type Repository interface {
	GetAll() ([]domain.Product, error)
	Create(name string, color string, price float64, stock int, code string, publisher bool) (domain.Product, error)
	Update(id int, name string, color string, price float64, stock int, code string, publisher bool) (domain.Product, error)
	Delete(id int) error
	ParcialUpdate(id int, name string, price float64) (domain.Product, error)
}

var ps []domain.Product
var lastId int

type repository struct{
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{db: db}
}

// ParcialUpdate implements Repository
func (r *repository) ParcialUpdate(id int, name string, price float64) (domain.Product, error) {

	updated := false
	ps, err := r.GetAll()
	var p domain.Product
	if err != nil {
		fmt.Println("Error:", err)
	}

	for i := range ps {
		if ps[i].Id == id {
		
			if name != "" && price > 0 {
				ps[i].Name = name
				ps[i].Price = price
				p = ps[i]
				updated = true
			}
			if name == "" && price > 0{
				ps[i].Price = price
				p = ps[i]
				updated = true
			}
			if name != "" && price == 0 {
				ps[i].Name = name
				p = ps[i]
				updated = true
			}
		}
	}
	if !updated {
		return domain.Product{}, fmt.Errorf("producto %d no encontrado", id)
	}


	if err:= r.db.Write(ps); err != nil {
		fmt.Println("Error:", err.Error())

	}

	return p, nil
}

// Delete implements Repository
func (r *repository) Delete(id int) error {
	deleted := false
	var index int
	ps, err := r.GetAll()
	if err != nil {
		fmt.Println("Error:", err)
	}

	for i := range ps {
		if ps[i].Id == id {
			index = i
			deleted = true
		}
	}

	if !deleted {
		return fmt.Errorf("producto %d no encontrado", id)
	}

	ps = append(ps[:index], ps[index+1:]...)


	if err := r.db.Write(ps); err != nil {
		fmt.Println("Error:", err.Error())

	}

	return nil
}

// Update implements Repository
func (r *repository) Update(id int, name string, color string, price float64, stock int, code string, publisher bool) (domain.Product, error) {
	p := domain.Product{Name: name, Color: color, Price: price, Stock: stock, Code: code, Publisher: publisher}
	update := false
	ps, err := r.GetAll()
	if err != nil {
		fmt.Println("Error:", err)
	}

	for i := range ps {
		if ps[i].Id == id {
			p.Id = id
			p.CreatedAt = ps[i].CreatedAt
			ps[i] = p
			update = true
		}
	}
	if !update {
		return domain.Product{}, fmt.Errorf("producto %d no encontrado", id)
	}

	if err:= r.db.Write(ps); err != nil {
		fmt.Println("Error:", err.Error())

	}

	return p, nil
}

//Created Product
func (r *repository) Create(name string, color string, price float64, stock int, code string, publisher bool) (domain.Product, error) {

	var t time.Time = time.Now()
	created := t.Format("2006-01-02 15:04:05")

	err := r.db.Read(&ps)
	if err != nil {
		fmt.Println("Error:", err)
	}
	lastId = ps[len(ps)-1].Id + 1
	p := domain.Product{Id: lastId, Name: name, Color: color, Price: price, Stock: stock, Code: code, Publisher: publisher, CreatedAt: created}

	ps = append(ps, p)

	if err1 := r.db.Write(ps); err1 != nil {
		return domain.Product{}, fmt.Errorf(ERROR_WRITING)
	}

	return p, nil
}

//Get All Products
func (r *repository) GetAll() ([]domain.Product, error) {
	err := r.db.Read(&ps)
	if err != nil {
		return []domain.Product{}, fmt.Errorf(ERROR_READING)
	}
	return ps, nil
}
