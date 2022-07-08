package products

import (
	"encoding/json"
	"errors"
	"goweb/internal/domain"
	"log"
	"os"
)

var PATH string = "../../resources/products.json"

type repositoryJsonDB struct{}

//func NewRepositoryJsonDB() Repository {
//	return &repositoryJsonDB{}
//}

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

	// Suponiendo que en el json este ordenado por id
	lastId = len(productos)

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

	productos, err := r.GetAll()
	if err != nil {
		log.Println(err)
		return domain.Product{}, errors.New("error al leer el archivo products.json")
	}

	productos = append(productos, producto)

	// Crear los datos del json
	data, err := json.Marshal(productos)
	if err != nil {
		log.Println(err)
		return domain.Product{}, errors.New("error al serializar el producto")
	}

	// Guardar la nueva lista de productos
	err = os.WriteFile(PATH, data, 0644)

	if err != nil {
		log.Println(err)
		return domain.Product{}, errors.New("error al escribir el archivo products.json")
	}

	lastId = id

	return producto, nil
}

func (r *repositoryJsonDB) GetById(id int) (domain.Product, error) {
	productos, err := r.GetAll()
	if err != nil {
		return domain.Product{}, err
	}

	for _, product := range productos {
		if product.Id == id {
			return product, nil
		}
	}

	return domain.Product{}, errors.New("no se encontró el producto")
}

func (r *repositoryJsonDB) LastId() (int, error) {
	r.GetAll()
	return lastId, nil
}
