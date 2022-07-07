package products

import (
	"errors"
	"goweb/productos_capas/internal/domain"
	"reflect"
	"strconv"
)

type Service interface {
	GetAll(nombre, color string, precio, stock int, codigo string, publicado bool, fechaCreacion string) ([]domain.Product, error)
	Store(nombre, color string, precio, stock int, codigo string, publicado bool, fechaCreacion string) (domain.Product, error)
	GetByID(id string) (domain.Product, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll(nombre, color string, precio, stock int, codigo string, publicado bool, fechaCreacion string) ([]domain.Product, error) {
	ps, err := s.repository.GetAll()
	ps = filterProducts(nombre, color, precio, stock, codigo, publicado, fechaCreacion, ps)
	if err != nil {
		return nil, err
	}

	return ps, nil
}

func (s *service) Store(nombre, color string, precio, stock int, codigo string, publicado bool, fechaCreacion string) (domain.Product, error) {

	if err := checkFields(nombre, color, precio, stock, codigo, fechaCreacion); err != nil {
		return domain.Product{}, err
	}

	lastID, err := s.repository.LastID()
	if err != nil {
		return domain.Product{}, err
	}

	lastID++

	producto, err := s.repository.Store(lastID, nombre, color, precio, stock, codigo, publicado, fechaCreacion)
	if err != nil {
		return domain.Product{}, err
	}

	return producto, nil
}

func (s *service) GetByID(id string) (domain.Product, error) {
	int_id, _ := strconv.Atoi(id)
	ps, err := s.repository.GetAll()
	if err != nil {
		return domain.Product{}, err
	}

	var product domain.Product
	for _, p := range ps {
		if p.Id == int_id {
			product = p
			break
		}
	}

	if product.Id == 0 {
		return domain.Product{}, errors.New("id no encontrado")
	}

	return product, nil
}

// Revisa que se hayan suministrado todos los campos requeridos
func checkFields(nombre, color string, precio, stock int, codigo string, fechaCreacion string) error {
	if nombre == "" {
		return errors.New("Nombre")
	}
	if color == "" {
		return errors.New("Color")
	}
	if precio == 0 {
		return errors.New("Precio")
	}
	if stock == 0 {
		return errors.New("Stock")
	}
	if codigo == "" {
		return errors.New("Codigo")
	}
	if fechaCreacion == "" {
		return errors.New("FechaCreacion")
	}
	return nil
}

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
