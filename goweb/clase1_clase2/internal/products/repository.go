package products

import (
	"encoding/json"
	"errors"
	"fmt"
	"goweb/clase1_clase2/internal/domain"
	"io/ioutil"
	"os"
	"reflect"
)

type Repository interface {
	GetAll(nombre, color string, precio, stock int, codigo string, publicado bool, fecha string) ([]domain.Product, error)
	Store(nombre, color string, precio, stock int, codigo string, publicado bool, fecha string) (domain.Product, error)
	Update(id int, nombre, color string, precio, stock int, codigo string, publicado bool, fecha string) (domain.Product, error)
	Delete(id int) (domain.Product, error)
	UpdateFields(id int, nombre string, precio int) (domain.Product, error)
	ReadJson()
	GetById(id int) (domain.Product, error)
}

type repository struct {
}

var productos []domain.Product
var lastID int

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAll(nombre, color string, precio, stock int, codigo string, publicado bool, fecha string) ([]domain.Product, error) {
	ps := filter(nombre, color, precio, stock, codigo, publicado, fecha, productos)
	return ps, nil
}

func (r *repository) Store(nombre, color string, precio, stock int, codigo string, publicado bool, fecha string) (domain.Product, error) {
	lastID++
	p := domain.Product{Id: lastID, Nombre: nombre, Color: color, Precio: precio, Stock: stock, Codigo: codigo, Publicado: publicado, Fecha: fecha}
	productos = append(productos, p)
	writeFile()
	return p, nil
}

func (r *repository) Update(id int, nombre, color string, precio, stock int, codigo string, publicado bool, fecha string) (domain.Product, error) {
	p := domain.Product{Id: id, Nombre: nombre, Color: color, Precio: precio, Stock: stock, Codigo: codigo, Publicado: publicado, Fecha: fecha}
	update := false
	for idx, producto := range productos {
		if id == producto.Id {
			productos[idx] = p
			writeFile()
			update = true
			break
		}
	}
	if !update {
		return domain.Product{}, errors.New("error: id no encontrado")
	}
	return p, nil
}

func (r *repository) Delete(id int) (domain.Product, error) {
	deleted := false
	var p_deleted domain.Product
	for idx, producto := range productos {
		if id == producto.Id {
			deleted = true
			p_deleted = producto
			productos = append(productos[:idx], productos[idx+1:]...)
			writeFile()
		}
	}
	if !deleted {
		return domain.Product{}, errors.New("error: id no encontrado")
	}
	return p_deleted, nil
}

func (r *repository) UpdateFields(id int, nombre string, precio int) (domain.Product, error) {
	update := false
	var p domain.Product
	for idx, producto := range productos {
		if id == producto.Id {
			productos[idx].Nombre = nombre
			productos[idx].Precio = precio
			p = productos[idx]
			update = true
			writeFile()
			break
		}
	}
	if !update {
		return domain.Product{}, errors.New("error: el id es invalido")
	}
	return p, nil
}

func (r *repository) GetById(id int) (domain.Product, error) {
	for _, p := range productos {
		if id == p.Id {
			return p, nil
		}
	}
	return domain.Product{}, errors.New("error: el id no es valido")
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

func filter(nombre, color string, precio, stock int, codigo string, publicado bool, fecha string, ps []domain.Product) []domain.Product {
	mapKeys := make(map[string]interface{})
	mapKeys["string"] = ""
	mapKeys["int"] = 0
	mapKeys["bool"] = false

	mapProduct := paramsToMap(nombre, color, precio, stock, codigo, publicado, fecha)
	var keysList = []string{}
	for key, value := range mapProduct {
		if value != mapKeys[reflect.TypeOf(value).String()] {
			keysList = append(keysList, key)
		}
	}
	productos := jsonToMap(ps...)
	var filtered_products []map[string]interface{}
	var filtered_products_empty []map[string]interface{}

	for _, key := range keysList {
		for _, p := range productos {
			if p[key] == mapProduct[key] {
				filtered_products = append(filtered_products, p)
			}
		}
		productos = filtered_products
		filtered_products = filtered_products_empty
	}
	productosJson := mapToJson(productos...)
	return productosJson
}

func paramsToMap(nombre, color string, precio, stock int, codigo string, publicado bool, fecha string) map[string]interface{} {
	productoMap := make(map[string]interface{})

	productoMap["nombre"] = nombre
	productoMap["color"] = color
	productoMap["precio"] = precio
	productoMap["stock"] = stock
	productoMap["codigo"] = codigo
	productoMap["publicado"] = publicado
	productoMap["fecha"] = fecha

	return productoMap
}

func jsonToMap(productos ...domain.Product) []map[string]interface{} {
	productoMap := make(map[string]interface{})
	var productoMaps []map[string]interface{}

	for _, p := range productos {
		productoMap["id"] = p.Id
		productoMap["nombre"] = p.Nombre
		productoMap["color"] = p.Color
		productoMap["precio"] = p.Precio
		productoMap["stock"] = p.Stock
		productoMap["codigo"] = p.Codigo
		productoMap["publicado"] = p.Publicado
		productoMap["fecha"] = p.Fecha

		productoMaps = append(productoMaps, productoMap)
		productoMap = make(map[string]interface{})
	}

	return productoMaps
}

func mapToJson(productos ...map[string]interface{}) []domain.Product {
	var producto domain.Product
	var products []domain.Product

	for _, p := range productos {
		producto.Id = p["id"].(int)
		producto.Nombre = p["nombre"].(string)
		producto.Color = p["color"].(string)
		producto.Precio = p["precio"].(int)
		producto.Stock = p["stock"].(int)
		producto.Codigo = p["codigo"].(string)
		producto.Publicado = p["publicado"].(bool)
		producto.Fecha = p["fecha"].(string)

		products = append(products, producto)
	}

	return products
}
