package productos

import (
	"fmt"

	"github.com/bootcamp-go/wave-5-backpack/tree/ospina_christian/goweb/clase3_parte2/internal/domain"
	"github.com/bootcamp-go/wave-5-backpack/tree/ospina_christian/goweb/clase3_parte2/pkg/store"
)

const (
	ProductNotFound = "product %d not found"
	FailReading     = "cant read database"
	FailWriting     = "cant write database, error: %w"
)

type Repository interface {
	LastID() (int, error)
	Store(id int, nombre string, color string, precio float64, stock int, codigo string, publicado bool, fechaCreacion string) (domain.Productos, error)
	GetAll() ([]domain.Productos, error)
	Update(id int, nombre string, color string, precio float64, stock int, codigo string, publicado bool, fechaCreacion string) (domain.Productos, error)
	UpdatePrecio(id int, precion float64) (domain.Productos, error)
	Delete(id int) error
	GetForId(id int) (domain.Productos, error)
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Store(id int, nombre string, color string, precio float64, stock int, codigo string, publicado bool, fechaCreacion string) (domain.Productos, error) {
	var ps []domain.Productos

	if err := r.db.Read(&ps); err != nil {
		return domain.Productos{}, fmt.Errorf(FailReading)
	}
	p := domain.Productos{Id: id, Nombre: nombre, Color: color, Precio: precio, Stock: stock, Codigo: codigo, Publicado: publicado, FechaCreación: fechaCreacion}

	ps = append(ps, p)

	if err := r.db.Write(ps); err != nil {
		return domain.Productos{}, fmt.Errorf(FailWriting, err)
	}
	return p, nil
}

func (r *repository) Update(id int, nombre string, color string, precio float64, stock int, codigo string, publicado bool, fechaCreacion string) (domain.Productos, error) {
	var ps []domain.Productos
	if err := r.db.Read(&ps); err != nil {
		return domain.Productos{}, fmt.Errorf(FailReading)
	}

	p := domain.Productos{Nombre: nombre, Color: color, Precio: precio, Stock: stock, Codigo: codigo, Publicado: publicado, FechaCreación: fechaCreacion}
	updated := false
	for i := range ps {
		if ps[i].Id == id {
			p.Id = id
			ps[i] = p
			updated = true
		}
	}
	if !updated {
		return domain.Productos{}, fmt.Errorf(ProductNotFound, id)
	}
	if err := r.db.Write(ps); err != nil {
		return domain.Productos{}, fmt.Errorf(FailWriting, err)
	}
	return p, nil
}

func (r *repository) UpdatePrecio(id int, precio float64) (domain.Productos, error) {
	var ps []domain.Productos
	if err := r.db.Read(&ps); err != nil {
		return domain.Productos{}, fmt.Errorf(FailReading)
	}
	updated := false
	var p domain.Productos
	for i := range ps {
		if ps[i].Id == id {
			ps[i].Precio = precio
			p = ps[i]
			updated = true
		}
	}
	if !updated {
		return domain.Productos{}, fmt.Errorf(ProductNotFound, id)
	}
	if err := r.db.Write(ps); err != nil {
		return domain.Productos{}, fmt.Errorf(FailWriting, err)
	}
	return p, nil
}

func (r *repository) Delete(id int) error {
	var ps []domain.Productos
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

func (r *repository) GetForId(id int) (domain.Productos, error) {
	var ps []domain.Productos
	if err := r.db.Read(&ps); err != nil {
		return domain.Productos{}, fmt.Errorf(FailReading)
	}
	founded := false
	var p domain.Productos
	for i := range ps {
		if ps[i].Id == id {
			founded = true
			p = ps[i]
		}
	}
	if !founded {
		return domain.Productos{}, fmt.Errorf(ProductNotFound, id)
	}
	return p, nil
}

func (r *repository) GetAll() ([]domain.Productos, error) {
	var ps []domain.Productos
	if err := r.db.Read(&ps); err != nil {
		return nil, fmt.Errorf(FailReading)
	}
	return ps, nil
}

func (r *repository) LastID() (int, error) {
	var ps []domain.Productos
	if err := r.db.Read(&ps); err != nil {
		return 0, fmt.Errorf(FailReading)
	}
	if len(ps) == 0 {
		return 0, nil
	}
	return ps[len(ps)-1].Id, nil
}
