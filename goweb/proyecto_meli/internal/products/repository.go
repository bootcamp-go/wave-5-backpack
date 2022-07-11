package products

import (
	"errors"
	"fmt"
	"proyecto_meli/internal/domain"
	"proyecto_meli/pkg/store"
)

type Repository interface {
	GetAll() ([]domain.Product, error)
	GetById(id int) (domain.Product, error)
	FilterList(id int, name, color string, price float64, stock int, codigo string, publicado bool, fecha string) ([]domain.Product, error)
	Store(id int, name, color string, price float64, stock int, codigo string, publicado bool, fecha string) (domain.Product, error)
	LastID() (int, error)
	Update(id int, name, color string, price float64, stock int, codigo string, publicado bool, fecha string) (domain.Product, error)
	Delete(id int) error
	Update_Name_Price(id int, name string, price float64) (domain.Product, error)
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
		return nil, errors.New("Fallo la lectura de la BD")
	}
	return ps, nil
}

func (r *repository) GetById(id int) (domain.Product, error) {
	var ps []domain.Product
	if err := r.db.Read(&ps); err != nil {
		return domain.Product{}, errors.New("Fallo la lectura de la BD")
	}
	for _, p := range ps {
		if p.Id == id {
			return p, nil
		}
	}
	return domain.Product{}, errors.New("No se encontro el producto.")
}
func (r *repository) FilterList(id int, name, color string, price float64, stock int, codigo string, publicado bool, fecha string) ([]domain.Product, error) {
	var ps []domain.Product
	if err := r.db.Read(&ps); err != nil {
		return nil, errors.New("Fallo la lectura de la BD")
	}
	if id > 0 {
		ps = filterListById(ps, id)
		fmt.Println(ps)
	}
	if name != "" {
		ps = filterListByName(ps, name)
	}
	if color != "" {
		ps = filterListByColor(ps, color)
	}
	if price > 0 {
		ps = filterListByPrice(ps, price)
	}
	if stock >= 0 {
		ps = filterListByStock(ps, stock)
	}
	if codigo != "" {
		ps = filterListByCode(ps, codigo)
	}
	ps = filterListByPublish(ps, publicado)
	if fecha != "" {
		ps = filterListByDate(ps, fecha)
	}
	return ps, nil
}

func filterListById(productos []domain.Product, id int) (filtroProductos []domain.Product) {
	fmt.Println(id)
	for _, p := range productos {
		fmt.Println(p)
		if p.Id == id {
			filtroProductos = append(filtroProductos, p)
		}
	}
	fmt.Println(filtroProductos)
	return
}

func filterListByName(productos []domain.Product, name string) (filtroProductos []domain.Product) {
	for _, p := range productos {
		if p.Nombre == name {
			filtroProductos = append(filtroProductos, p)
		}
	}
	return
}

func filterListByColor(productos []domain.Product, color string) (filtroProductos []domain.Product) {
	for _, p := range productos {
		if p.Color == color {
			filtroProductos = append(filtroProductos, p)
		}
	}
	return
}

func filterListByPrice(productos []domain.Product, price float64) (filtroProductos []domain.Product) {
	for _, p := range productos {
		if p.Precio == price {
			filtroProductos = append(filtroProductos, p)
		}
	}
	return
}

func filterListByStock(productos []domain.Product, stock int) (filtroProductos []domain.Product) {
	for _, p := range productos {
		if p.Stock == stock {
			filtroProductos = append(filtroProductos, p)
		}
	}
	return
}

func filterListByCode(productos []domain.Product, code string) (filtroProductos []domain.Product) {
	for _, p := range productos {
		if p.Codigo == code {
			filtroProductos = append(filtroProductos, p)
		}
	}
	return
}

func filterListByPublish(productos []domain.Product, publish bool) (filtroProductos []domain.Product) {
	for _, p := range productos {
		if p.Publicado == publish {
			filtroProductos = append(filtroProductos, p)
		}
	}
	return
}

func filterListByDate(productos []domain.Product, date string) (filtroProductos []domain.Product) {
	for _, p := range productos {
		if p.FechaCreacion == date {
			filtroProductos = append(filtroProductos, p)
		}
	}
	return
}

func (r *repository) Store(id int, name, color string, price float64, stock int, codigo string, publicado bool, fecha string) (domain.Product, error) {
	var ps []domain.Product

	if err := r.db.Read(&ps); err != nil {
		return domain.Product{}, errors.New("Fallo la lectura de la BD")
	}

	p := domain.Product{Id: id, Nombre: name, Color: color, Precio: price, Stock: stock, Codigo: codigo, Publicado: publicado, FechaCreacion: fecha}
	ps = append(ps, p)

	if err := r.db.Write(ps); err != nil {
		return domain.Product{}, errors.New("Fallo la escritura de la BD")
	}

	return p, nil
}

func (r *repository) LastID() (int, error) {
	var ps []domain.Product
	if err := r.db.Read(&ps); err != nil {
		return 0, errors.New("Error al leer la BD")
	}
	if len(ps) == 0 {
		return 0, nil
	}

	return ps[len(ps)-1].Id, nil
}

func (r *repository) Update(id int, name, color string, price float64, stock int, codigo string, publicado bool, fecha string) (domain.Product, error) {
	var ps []domain.Product

	if err := r.db.Read(&ps); err != nil {
		return domain.Product{}, errors.New("Error al leer la BD")
	}

	p := domain.Product{Id: id, Nombre: name, Color: color, Precio: price, Stock: stock, Codigo: codigo, Publicado: publicado, FechaCreacion: fecha}
	updated := false
	for i := range ps {
		if ps[i].Id == id {
			p.Id = id
			ps[i] = p
			updated = true
		}
	}

	if !updated {
		return domain.Product{}, errors.New("Producto no encontrado")
	}

	if err := r.db.Write(ps); err != nil {
		return domain.Product{}, errors.New("Error al guardar en la BD")
	}

	return p, nil
}

func (r *repository) Delete(id int) error {
	var ps []domain.Product

	if err := r.db.Read(&ps); err != nil {
		return errors.New("Error al leer la BD")
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
		return errors.New("Producto no encontrado")
	}

	ps = append(ps[:index], ps[index+1:]...)

	if err := r.db.Write(ps); err != nil {
		return errors.New("Error al guardar en la BD")
	}
	return nil
}

func (r *repository) Update_Name_Price(id int, name string, price float64) (domain.Product, error) {
	var ps []domain.Product

	if err := r.db.Read(&ps); err != nil {
		return domain.Product{}, errors.New("Error al leer la BD")
	}

	updated := false
	var p domain.Product
	for i := range ps {
		if ps[i].Id == id {
			ps[i].Nombre = name
			ps[i].Precio = price
			p = ps[i]
			updated = true
		}
	}

	if !updated {
		return domain.Product{}, errors.New("Producto no encontrado")
	}

	if err := r.db.Write(ps); err != nil {
		return domain.Product{}, errors.New("Error al guardar en la BD")
	}

	return p, nil
}
