package products

import (
	"errors"
	"fmt"
	"github.com/bootcamp-go/wave-5-backpack/goweb/internal/models"
	"github.com/bootcamp-go/wave-5-backpack/goweb/pkg/storage"
	"time"
)

type Repository interface {
	Store(nombre, color, precio, stock, codigo, publicado, fechaCreacion) (models.Products, error)
	Update(id, nombre, color, precio, stock, codigo, publicado, fechaCreacion) (models.Products, error)
	UpdatePrecioStock(id, precio, stock) (models.Products, error)
	GetAll() ([]models.Products, error)
	GetByID(id int) (models.Products, error)
	GetLastID() (int, error)
	Delete(id int) (int, error)
}

func NewRepository(storage storage.Storage) Repository {
	return &repository{storage}
}

type repository struct {
	storage storage.Storage
}

func (r repository) Store(nombre, color, precio, stock, codigo, publicado, fechaCreacion) (models.Products, error) {
	var pr []models.Products
	if err := r.storage.Read(&pr); err != nil {
		return models.Products{}, fmt.Errorf("error: al leer el archivo %v", err)
	}
	newID := (pr[len(pr)-1].ID) + 1
	p := models.Products{
		ID:            newID,
		Nombre:        nombre,
		Color:         color,
		Precio:        precio,
		Stock:         stock,
		Codigo:        codigo,
		Publicado:     publicado,
		FechaCreacion: time.Now().Local().String(),
	}
	pr = append(pr, p)

	err := r.storage.Write(pr)
	if err != nil {
		return models.Products{}, fmt.Errorf("error: al escribir el archivo %v", err)
	}
	return p, nil
}

func (r repository) Update(id, nombre, color, precio, stock, codigo, publicado, fechaCreacion) (models.Products, error) {
	var pr []models.Products
	if err := r.storage.Read(&pr); err != nil {
		return models.Products{}, fmt.Errorf("error: al leer el archivo %v", err)
	}
	for i, pp := range pr {
		if pp.ID == id {
			p := models.Products{
				ID:            id,
				Nombre:        nombre,
				Color:         color,
				Precio:        precio,
				Stock:         stock,
				Codigo:        codigo,
				Publicado:     publicado,
				FechaCreacion: pp.FechaCreacion,
			}
			pr[i] = p

			err := r.storage.Write(pr)
			if err != nil {
				return models.Products{}, fmt.Errorf("error: al escribir el archivo %v", err)
			}
			return p, nil
		}
	}
	return models.Products{}, fmt.Errorf("error: no existe el ID: %v", id)
}

func (r repository) UpdatePrecioStock(id, precio, stock) (models.Products, error) {
	var pr []models.Products
	if err := r.storage.Read(&pr); err != nil {
		return models.Products{}, fmt.Errorf("error: al leer el archivo %v", err)
	}
	for i, pp := range pr {
		if pp.ID == id {
			p := models.Products{
				ID:     pp.ID,
				Precio: pp.Precio,
				Stock:  pp.Stock,
			}
			if precio != 0 {
				p.Precio = precio
			} else {
				p.Precio = pp.Precio
			}
			if stock != 0 {
				p.Stock = stock
			} else {
				p.Stock = pp.Stock
			}

			pr[i] = p

			if err := r.storage.Write(pr); err != nil {
				return models.Products{}, fmt.Errorf("error: al escribit el archivo %v\n", err)
			}
			return p, nil
		}
	}
	return models.Products{}, fmt.Errorf("error: ID: %v no encontrado", id)
}

func (r repository) GetAll() ([]models.Products, error) {
	var pr []models.Products
	if err := r.storage.Read(&pr); err != nil {
		return nil, err
	}
	if len(pr) == 0 {
		return nil, errors.New("No existen Registros")
	}
	return pr, nil
}

func (r repository) GetByID(id int) (models.Products, error) {
	var pr []models.Products
	if err := r.storage.Read(&pr); err != nil {
		return models.Products{}, err
	}
	for _, p := range pr {
		if p.ID == id {
			return p, nil
		}
	}
	return models.Products{}, fmt.Errorf("producto con ID: %v no encontrado", id)
}

func (r repository) GetLastID() (int, error) {
	var pr []models.Products
	if err := r.storage.Read(&pr); err != nil {
		return 0, err
	}
	if len(pr) == 0 {
		return 0, errors.New("error: no existen productos")
	}

	id := pr[len(pr)-1].ID

	return id, nil
}

func (r repository) Delete(id int) (int, error) {
	var pr []models.Products
	if err := r.storage.Read(&pr); err != nil {
		return 0, err
	}
	for i, p := range pr {
		if p.ID == id {
			if i == len(pr)-1 {
				pr = pr[:i]
				r.storage.Write(pr)
				return id, nil
			}
			pr = append(pr[:i], pr[i+1:]...)
			r.storage.Write(pr)
			return id, nil
		}
	}
	return 0, fmt.Errorf("error: ID %v no existe", id)
}
