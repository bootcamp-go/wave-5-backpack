package products

import (
	"encoding/json"
	"goweb/productos_capas/internal/domain"
	"io/ioutil"
	"os"
)

type Repository interface {
	GetAll() ([]domain.Product, error)
	Store(id int, nombre, color string, precio, stock int, codigo string, publicado bool, fechaCreacion string) (domain.Product, error)
	LastID() (int, error)
	ReadJSON()
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

var products []domain.Product
var lastID int

func (r *repository) GetAll() ([]domain.Product, error) {
	return products, nil
}

func (r *repository) Store(id int, nombre, color string, precio, stock int, codigo string, publicado bool, fechaCreacion string) (domain.Product, error) {
	p := domain.Product{Id: id, Nombre: nombre, Color: color, Precio: precio, Stock: stock, Codigo: codigo, Publicado: publicado, FechaCreacion: fechaCreacion}
	products = append(products, p)
	writeJSON()
	lastID = p.Id
	return p, nil
}

func (r *repository) LastID() (int, error) {
	return lastID, nil
}

// Obtiene la lista de productos del archivo products.json
func (r *repository) ReadJSON() {
	jsonFile, err := os.Open("internal/domain/products.json")
	if err != nil {
		panic(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &products)
	lastID = products[len(products)-1].Id
	jsonFile.Close()
}

// Guarda una lista de productos en products.json
func writeJSON() error {
	file, err := json.Marshal(products)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile("internal/domain/products.json", file, 0644)
	if err != nil {
		return err
	}
	return nil
}
