package products

import "fmt"

type Producto struct {
	ID       int     `json:"id"`
	Nombre   string  `json:"nombre"`
	Cantidad int     `json:"cantidad"`
	Precio   float64 `json:"precio"`
}

var ps []Producto
var lastID int

type Repository interface {
	GetAll() ([]Producto, error)
	Store(id int, nombre string, cantidad int, precio float64) (Producto, error)
	Update(id int, nombre string, cantidad int, precio float64) (Producto, error)
	UpdateName(id int, nombre string) (Producto, error)
	UpdateNamePrice(id int, nombre string, price float64) (Producto, error)
	Delete(id int) error
	LastID() (int, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAll() ([]Producto, error) {
	if len(ps) == 0 {
		return nil, fmt.Errorf("no hay productoos registrados")
	}
	return ps, nil
}

func (r *repository) Store(id int, nombre string, cantidad int, precio float64) (Producto, error) {
	p := Producto{id, nombre, cantidad, precio}
	ps = append(ps, p)
	lastID = p.ID
	return p, nil
}

func (r *repository) Update(id int, nombre string, cantidad int, precio float64) (Producto, error) {
	p := Producto{Nombre: nombre, Cantidad: cantidad, Precio: precio}
	updated := false
	for i := range ps {
		if ps[i].ID == id {
			p.ID = id
			ps[i] = p
			updated = true
		}
	}

	if !updated {
		return Producto{}, fmt.Errorf("producto %d no encontrado", id)
	}

	return p, nil

}

func (r *repository) UpdateName(id int, nombre string) (Producto, error) {
	updated := false
	var p Producto
	for i := range ps {
		if ps[i].ID == id {
			ps[i].Nombre = nombre
			p = ps[i]
			updated = true
		}
	}

	if !updated {
		return Producto{}, fmt.Errorf("producto %d no encontrado", id)
	}

	return p, nil
}

func (r *repository) UpdateNamePrice(id int, nombre string, precio float64) (Producto, error) {
	updated := false
	var p Producto
	for i := range ps {
		if ps[i].ID == id {
			ps[i].Nombre = nombre
			ps[i].Precio = precio
			p = ps[i]
			updated = true
		}
	}

	if !updated {
		return Producto{}, fmt.Errorf("producto %d no encontrado", id)
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
		return fmt.Errorf("productoo %d no encontrado", id)
	}

	ps = append(ps[:index], ps[index+1:]...)
	return nil
}

func (r *repository) LastID() (int, error) {
	return lastID, nil
}
