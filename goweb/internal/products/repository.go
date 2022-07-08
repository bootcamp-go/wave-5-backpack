package products

import (
	"fmt"
	"web-server/internal/domain"
)

type Repository interface {
	GetAll() ([]domain.Products, error)
	LastID() (int, error)
	Store(id int, nombre string, color string, precio float64, stock int, codigo string, publicado bool, fecha string) (domain.Products, error)
	Update(id int, nombre string, color string, precio float64, stock int, codigo string, publicado bool, fecha string) (domain.Products, error)
	UpdateName(id int, nombre string) (domain.Products, error)
	UpdatePrice(id int, precio float64) (domain.Products, error)
	Delete(id int) error
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

var productsSlide []domain.Products
var lastId int

func (r *repository) GetAll() ([]domain.Products, error) {
	return productsSlide, nil
}

func (r *repository) LastID() (int, error) {
	return lastId, nil
}

func (r *repository) Store(id int, nombre string, color string, precio float64, stock int, codigo string, publicado bool, fecha string) (domain.Products, error) {
	p := domain.Products{Id: id, Nombre: nombre, Color: color, Precio: precio, Stock: stock, Codigo: codigo, Publicado: publicado, Fecha: fecha}
	productsSlide = append(productsSlide, p)
	lastId = p.Id
	return p, nil
}

// func readFile() ([]domain.Products,error){

// 	products,err := os.ReadFile("./products.json")
// 	if err != nil {
// 		return nil, fmt.Errorf(err.Error())
// 	}

// }

func (r *repository) Update(id int, nombre string, color string, precio float64, stock int, codigo string, publicado bool, fecha string) (domain.Products, error) {
	p := domain.Products{Nombre: nombre, Color: color, Precio: precio, Stock: stock, Codigo: codigo, Publicado: publicado, Fecha: fecha}
	updated := false
	for i := range productsSlide {
		if productsSlide[i].Id == id {
			p.Id = id
			productsSlide[i] = p
			updated = true
		}
	}

	if !updated {
		return domain.Products{}, fmt.Errorf("producto %d no encontrado", id)
	}

	return p, nil
}

func (r *repository) UpdateName(id int, nombre string) (domain.Products, error) {
	updated := false
	var p domain.Products
	for i := range productsSlide {
		if productsSlide[i].Id == id {
			productsSlide[i].Nombre = nombre
			p = productsSlide[i]
			updated = true
		}
	}

	if !updated {
		return domain.Products{}, fmt.Errorf("producto %d no encontrado", id)
	}

	return p, nil
}

func (r *repository) UpdatePrice(id int, precio float64) (domain.Products, error) {
	updated := false
	var p domain.Products
	for i := range productsSlide {
		if productsSlide[i].Id == id {
			productsSlide[i].Precio = precio
			p = productsSlide[i]
			updated = true
		}
	}

	if !updated {
		return domain.Products{}, fmt.Errorf("producto %d no encontrado", id)
	}

	return p, nil
}

func (r *repository) Delete(id int) error {
	deleted := false
	var index int
	for i := range productsSlide {
		if productsSlide[i].Id == id {
			index = i
			deleted = true
		}
	}

	if !deleted {
		return fmt.Errorf("producto %d no encontrado", id)
	}

	productsSlide = append(productsSlide[:index], productsSlide[index+1:]...)
	return nil
}
