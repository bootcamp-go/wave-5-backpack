package products

import (
	"encoding/json"
	"fmt"
	"goweb/clase1_clase2/internal/domain"
	"io/ioutil"
	"os"
)

type Repository interface {
	GetAll() ([]domain.Product, error)
	Store(id int, nombre, color string, precio, stock int, codigo string, publicado bool, fecha string) (domain.Product, error)
	LastID() (int, error)
	ReadJson()
}

type repository struct {
}

var productos []domain.Product
var lastID int

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAll() ([]domain.Product, error) {
	return productos, nil
}

func (r *repository) LastID() (int, error) {
	return lastID, nil
}

func (r *repository) Store(id int, nombre, color string, precio, stock int, codigo string, publicado bool, fecha string) (domain.Product, error) {
	p := domain.Product{Id: id, Nombre: nombre, Color: color, Precio: precio, Stock: stock, Codigo: codigo, Publicado: publicado, Fecha: fecha}
	productos = append(productos, p)
	lastID = p.Id
	writeFile()
	return p, nil
}

func (r *repository) ReadJson() {
	jsonFile, err := os.Open("./internal/domain/products.json")
	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal([]byte(byteValue), &productos)
	lastID = productos[len(productos)-1].Id
	defer jsonFile.Close()
}

func writeFile() {
	jsonFile, err := json.Marshal(productos)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("guarda el archivo")
	err = ioutil.WriteFile("./internal/domain/products.json", jsonFile, 0644)
	if err != nil {
		fmt.Println(err)
	}
}
