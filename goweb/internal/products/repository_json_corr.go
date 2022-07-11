package products

import (
	"fmt"
	"goweb/internal/domain"
	"goweb/pkg/store"
)

const (
	ProductNotFound = "producto %d no encontrado"
	FailReading     = "no se pudo leer la base de datos"
	FailWriting     = "no se pudo escribir en la base de datos, error: %w"
)

// --------------------------------------------
// --------------- Estructuras ----------------
// --------------------------------------------

type repositoryJsonCorrDB struct {
	db store.Store
}

func NewRepositoryJsonCorrDB(db store.Store) Repository {
	return &repositoryJsonCorrDB{db: db}
}

// --------------------------------------------
// ------------------- CRUD -------------------
// --------------------------------------------

func (r *repositoryJsonCorrDB) GetAll() ([]domain.Product, error) {
	var products []domain.Product
	if err := r.db.Read(&products); err != nil {
		return nil, fmt.Errorf(FailReading)
	}
	return products, nil
}

func (r *repositoryJsonCorrDB) Store(id int, nombre string, color string, precio float64, stock int, codigo string, publicado bool, fechaCreacion string) (domain.Product, error) {

	var products []domain.Product

	if err := r.db.Read(&products); err != nil {
		return domain.Product{}, fmt.Errorf(FailReading)
	}

	producto := domain.Product{
		Id:            id,
		Nombre:        nombre,
		Color:         color,
		Precio:        precio,
		Stock:         stock,
		Codigo:        codigo,
		Publicado:     publicado,
		FechaCreacion: fechaCreacion,
	}

	products = append(products, producto)

	if err := r.db.Write(&products); err != nil {
		return domain.Product{}, fmt.Errorf(FailWriting, err)
	}

	return producto, nil
}

func (r *repositoryJsonCorrDB) LastId() (int, error) {
	var products []domain.Product
	if err := r.db.Read(&products); err != nil {
		return 0, fmt.Errorf(FailReading)
	}

	if len(products) == 0 {
		return 0, nil
	}

	return products[len(products)-1].Id, nil
}

func (r *repositoryJsonCorrDB) Delete(id int) error {
	var products []domain.Product
	if err := r.db.Read(&products); err != nil {
		return fmt.Errorf(FailReading)
	}

	deleted := false
	var index int
	for i, product := range products {
		if product.Id == id {
			index = i
			deleted = true
			break
		}
	}

	if !deleted {
		return fmt.Errorf(ProductNotFound, id)
	}

	products = append(products[:index], products[index+1:]...)

	if err := r.db.Write(&products); err != nil {
		return fmt.Errorf(FailWriting, err)
	}

	return nil
}

func (r *repositoryJsonCorrDB) GetById(id int) (domain.Product, error) {
	var products []domain.Product
	if err := r.db.Read(&products); err != nil {
		return domain.Product{}, fmt.Errorf(FailReading)
	}

	for _, product := range products {
		if product.Id == id {
			return product, nil
		}
	}

	return domain.Product{}, fmt.Errorf(ProductNotFound, id)
}

func (r *repositoryJsonCorrDB) Update(id int, nombre string, color string, precio float64, stock int, codigo string, publicado bool, fechaCreacion string) (domain.Product, error) {
	var products []domain.Product
	if err := r.db.Read(&products); err != nil {
		return domain.Product{}, fmt.Errorf(FailReading)
	}

	product := domain.Product{
		Nombre:        nombre,
		Color:         color,
		Precio:        precio,
		Stock:         stock,
		Codigo:        codigo,
		Publicado:     publicado,
		FechaCreacion: fechaCreacion,
	}
	updated := false
	for i := range products {
		if products[i].Id == id {
			product.Id = id
			products[i] = product
			updated = true
			break
		}
	}

	if !updated {
		return domain.Product{}, fmt.Errorf(ProductNotFound, id)
	}

	if err := r.db.Write(&products); err != nil {
		return domain.Product{}, fmt.Errorf(FailWriting, err)
	}

	return product, nil
}

func (r *repositoryJsonCorrDB) UpdateNombre(id int, nombre string) (domain.Product, error) {
	var products []domain.Product
	if err := r.db.Read(&products); err != nil {
		return domain.Product{}, fmt.Errorf(FailReading)
	}

	updated := false
	for i := range products {
		if products[i].Id == id {
			products[i].Nombre = nombre
			updated = true
			break
		}
	}

	if !updated {
		return domain.Product{}, fmt.Errorf(ProductNotFound, id)
	}

	if err := r.db.Write(&products); err != nil {
		return domain.Product{}, fmt.Errorf(FailWriting, err)
	}

	return products[len(products)-1], nil
}

func (r *repositoryJsonCorrDB) UpdatePrecio(id int, precio float64) (domain.Product, error) {
	var products []domain.Product
	if err := r.db.Read(&products); err != nil {
		return domain.Product{}, fmt.Errorf(FailReading)
	}

	updated := false
	for i := range products {
		if products[i].Id == id {
			products[i].Precio = precio
			updated = true
			break
		}
	}

	if !updated {
		return domain.Product{}, fmt.Errorf(ProductNotFound, id)
	}

	if err := r.db.Write(&products); err != nil {
		return domain.Product{}, fmt.Errorf(FailWriting, err)
	}

	return products[len(products)-1], nil
}
