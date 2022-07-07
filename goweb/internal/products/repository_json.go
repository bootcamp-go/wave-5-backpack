package products

import (
	"encoding/json"
	"errors"
	"goweb/internal/domain"
	"log"
	"os"
)

var PATH string = "./resources/products.json"

type repositoryJsonDB struct{}

func NewRepositoryJsonDB() Repository {
	return &repositoryJsonDB{}
}

func (r *repositoryJsonDB) GetAll() ([]domain.Product, error) {
	data, err := os.ReadFile(PATH)
	if err != nil {
		log.Println(err)
		return nil, errors.New("error al leer el archivo products.json")
	}

	productos := []domain.Product{}

	if err := json.Unmarshal([]byte(data), &productos); err != nil {
		log.Println(err)
		return nil, errors.New("error al deserializar el archivo products.json")
	}

	return productos, nil
}

func (r *repositoryJsonDB) Store(id int, nombre string, color string, precio float64, stock int, codigo string, publicado bool, fechaCreacion string) (domain.Product, error) {
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
	lastId = id

	return producto, nil
}

func (r *repositoryJsonDB) GetById(id int) (domain.Product, error) {
	for _, product := range products {
		if product.Id == id {
			return product, nil
		}
	}
	return domain.Product{}, errors.New("no se encontró el producto")
}

func (r *repositoryJsonDB) LastId() (int, error) {
	return lastId, nil
}
