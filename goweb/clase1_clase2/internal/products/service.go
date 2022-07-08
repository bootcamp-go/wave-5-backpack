package products

import (
	"errors"
	"goweb/clase1_clase2/internal/domain"
	"reflect"
	"strconv"
)

type Service interface {
	GetAll(nombre, color string, precio, stock int, codigo string, publicado bool, fecha string) ([]domain.Product, error)
	Store(nombre, color string, precio, stock int, codigo string, publicado bool, fecha string) (domain.Product, error)
	GetById(id string) (domain.Product, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll(nombre, color string, precio, stock int, codigo string, publicado bool, fecha string) ([]domain.Product, error) {
	ps, err := s.repository.GetAll()
	ps = filter(nombre, color, precio, stock, codigo, publicado, fecha, ps)
	if err != nil {
		return nil, err
	}

	return ps, nil
}

func (s *service) Store(nombre, color string, precio, stock int, codigo string, publicado bool, fecha string) (domain.Product, error) {
	if err := validateFields(nombre, color, precio, stock, codigo, fecha); err != nil {
		return domain.Product{}, err
	}

	lastID, err := s.repository.LastID()
	if err != nil {
		return domain.Product{}, err
	}

	lastID++

	producto, err := s.repository.Store(lastID, nombre, color, precio, stock, codigo, publicado, fecha)
	if err != nil {
		return domain.Product{}, err
	}

	return producto, nil
}

func validateFields(nombre, color string, precio, stock int, codigo string, fecha string) error {
	if nombre == "" {
		return errors.New("el campo nombre es requerido")
	}
	if color == "" {
		return errors.New("el campo color es requerido")
	}
	if precio == 0 {
		return errors.New("el campo precio es requerido")
	}
	if stock == 0 {
		return errors.New("el campo stock es requerido")
	}
	if codigo == "" {
		return errors.New("el campo codigo es requerido")
	}
	if fecha == "" {
		return errors.New("el campo fecha es requerido")
	}
	return nil
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

func (s *service) GetById(id string) (domain.Product, error) {
	id_int, err := strconv.Atoi(id)
	if err != nil {
		return domain.Product{}, err
	}
	ps, err := s.repository.GetAll()
	if err != nil {
		return domain.Product{}, err
	}
	for _, p := range ps {
		if id_int == p.Id {
			return p, nil
		}
	}
	return domain.Product{}, errors.New("el id no es válido")
}
