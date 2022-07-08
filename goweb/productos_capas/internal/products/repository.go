package products

import (
	"encoding/json"
	"errors"
	"goweb/productos_capas/internal/domain"
	"io/ioutil"
	"os"
	"reflect"
	"strconv"
)

type Repository interface {
	GetAll(nombre, color string, precio, stock int, codigo string, publicado bool, fechaCreacion string) ([]domain.Product, error)
	Store(nombre, color string, precio, stock int, codigo string, publicado bool, fechaCreacion string) (domain.Product, error)
	GetByID(id int) (domain.Product, error)
	LastID() (int, error)
	ReadJSON()
	Update(id int, nombre, color string, precio, stock int, codigo string, publicado bool, fechaCreacion string) (domain.Product, error)
	UpdateNamePrice(id int, nombre string, precio int) (domain.Product, error)
	Delete(id int) (domain.Product, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

var products []domain.Product
var lastID int

func (r *repository) GetAll(nombre, color string, precio, stock int, codigo string, publicado bool, fechaCreacion string) ([]domain.Product, error) {
	ps := filterProducts(nombre, color, precio, stock, codigo, publicado, fechaCreacion, products)
	return ps, nil
}

func (r *repository) Store(nombre, color string, precio, stock int, codigo string, publicado bool, fechaCreacion string) (domain.Product, error) {
	lastID++
	p := domain.Product{Id: lastID, Nombre: nombre, Color: color, Precio: precio, Stock: stock, Codigo: codigo, Publicado: publicado, FechaCreacion: fechaCreacion}
	products = append(products, p)
	writeJSON()
	return p, nil
}

func (r *repository) GetByID(id int) (domain.Product, error) {
	var product domain.Product
	for _, p := range products {
		if p.Id == id {
			product = p
			break
		}
	}

	if product.Id == 0 {
		return domain.Product{}, errors.New("id no encontrado")
	}
	return product, nil
}

func (r *repository) Update(id int, nombre, color string, precio, stock int, codigo string, publicado bool, fechaCreacion string) (domain.Product, error) {
	product := domain.Product{Id: id, Nombre: nombre, Color: color, Precio: precio, Stock: stock, Codigo: codigo, Publicado: publicado, FechaCreacion: fechaCreacion}
	var updated bool
	for idx, p := range products {
		if p.Id == id {
			products[idx] = product
			writeJSON()
			updated = true
		}
	}

	if !updated {
		return domain.Product{}, errors.New("Producto con id " + strconv.Itoa(id) + " no encontrado")
	}
	return product, nil
}

func (r *repository) UpdateNamePrice(id int, nombre string, precio int) (domain.Product, error) {
	var product domain.Product
	var updated bool
	for idx, p := range products {
		if p.Id == id {
			products[idx].Nombre = nombre
			products[idx].Precio = precio
			product = products[idx]
			writeJSON()
			updated = true
		}
	}

	if !updated {
		return domain.Product{}, errors.New("Producto con id " + strconv.Itoa(id) + " no encontrado")
	}
	return product, nil
}

func (r *repository) Delete(id int) (domain.Product, error) {
	var product domain.Product
	var deleted bool
	for idx, p := range products {
		if p.Id == id {
			product = p
			products = append(products[:idx], products[idx+1:]...)
			writeJSON()
			deleted = true
		}
	}
	if !deleted {
		return domain.Product{}, errors.New("Producto con id" + strconv.Itoa(id) + "no encontrado")
	}
	return product, nil
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

// Filtra la lista de productos por los parámetros especificados
func filterProducts(nombre, color string, precio, stock int, codigo string, publicado bool, fechaCreacion string, ps []domain.Product) []domain.Product {
	mapKeys := make(map[string]interface{})
	mapKeys["string"] = ""
	mapKeys["int"] = 0
	mapKeys["bool"] = false
	mapProduct := paramsToMap(nombre, color, precio, stock, codigo, publicado, fechaCreacion)

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

// Convierte una lista de parametros a map[string]interface
func paramsToMap(nombre, color string, precio, stock int, codigo string, publicado bool, fechaCreacion string) map[string]interface{} {
	productoMap := make(map[string]interface{})
	productoMap["nombre"] = nombre
	productoMap["color"] = color
	productoMap["precio"] = precio
	productoMap["stock"] = stock
	productoMap["codigo"] = codigo
	productoMap["publicado"] = publicado
	productoMap["fecha_creacion"] = fechaCreacion
	return productoMap
}

// Convierte una lista de struct productos a []map[string]interface
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
		productoMap["fecha_creacion"] = p.FechaCreacion

		productoMaps = append(productoMaps, productoMap)
		productoMap = make(map[string]interface{})
	}

	return productoMaps
}

// Convierte una lista de map[string]interface a struct domain.product
func mapToJson(productos ...map[string]interface{}) []domain.Product {
	var producto domain.Product
	var productosJson []domain.Product

	for _, p := range productos {
		producto.Id = p["id"].(int)
		producto.Nombre = p["nombre"].(string)
		producto.Color = p["color"].(string)
		producto.Precio = p["precio"].(int)
		producto.Stock = p["stock"].(int)
		producto.Codigo = p["codigo"].(string)
		producto.Publicado = p["publicado"].(bool)
		producto.FechaCreacion = p["fecha_creacion"].(string)

		productosJson = append(productosJson, producto)
	}

	return productosJson
}
