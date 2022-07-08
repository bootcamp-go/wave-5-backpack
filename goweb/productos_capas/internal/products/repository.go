package products

import (
	"errors"
	"fmt"
	"goweb/productos_capas/internal/domain"
	"goweb/productos_capas/pkg/store"
	"reflect"
	"strconv"
)

const (
	ProductNotFound = "product %d not found"
	FailReading     = "cant read database"
	FailWriting     = "cant write database, error: %w"
)

type Repository interface {
	GetAll(nombre, color string, precio, stock int, codigo string, publicado bool, fechaCreacion string) ([]domain.Product, error)
	Store(nombre, color string, precio, stock int, codigo string, publicado bool, fechaCreacion string) (domain.Product, error)
	GetByID(id int) (domain.Product, error)
	LastID() (int, error)
	Update(id int, nombre, color string, precio, stock int, codigo string, publicado bool, fechaCreacion string) (domain.Product, error)
	UpdateNamePrice(id int, nombre string, precio int) (domain.Product, error)
	Delete(id int) (domain.Product, error)
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll(nombre, color string, precio, stock int, codigo string, publicado bool, fechaCreacion string) ([]domain.Product, error) {
	var ps []domain.Product
	if err := r.db.Read(&ps); err != nil {
		return nil, fmt.Errorf(FailReading)
	}
	filtered_ps := filterProducts(nombre, color, precio, stock, codigo, publicado, fechaCreacion, ps)
	return filtered_ps, nil
}

func (r *repository) Store(nombre, color string, precio, stock int, codigo string, publicado bool, fechaCreacion string) (domain.Product, error) {
	var ps []domain.Product
	if err := r.db.Read(&ps); err != nil {
		return domain.Product{}, fmt.Errorf(FailReading)
	}

	lastID, err := r.LastID()
	if err != nil {
		return domain.Product{}, fmt.Errorf("error obteniendo el ultimo id: %w", err)
	}
	lastID++

	p := domain.Product{Id: lastID, Nombre: nombre, Color: color, Precio: precio, Stock: stock, Codigo: codigo, Publicado: publicado, FechaCreacion: fechaCreacion}
	ps = append(ps, p)

	if err := r.db.Write(ps); err != nil {
		return domain.Product{}, fmt.Errorf(FailWriting, err)
	}

	return p, nil
}

func (r *repository) GetByID(id int) (domain.Product, error) {
	var ps []domain.Product
	if err := r.db.Read(&ps); err != nil {
		return domain.Product{}, fmt.Errorf(FailReading)
	}

	for _, p := range ps {
		if p.Id == id {
			return p, nil
		}
	}

	return domain.Product{}, errors.New("id no encontrado")

}

func (r *repository) Update(id int, nombre, color string, precio, stock int, codigo string, publicado bool, fechaCreacion string) (domain.Product, error) {
	var ps []domain.Product
	if err := r.db.Read(&ps); err != nil {
		return domain.Product{}, fmt.Errorf(FailReading)
	}

	product := domain.Product{Id: id, Nombre: nombre, Color: color, Precio: precio, Stock: stock, Codigo: codigo, Publicado: publicado, FechaCreacion: fechaCreacion}
	var updated bool
	for idx, p := range ps {
		if p.Id == id {
			ps[idx] = product
			updated = true
		}
	}

	if !updated {
		return domain.Product{}, errors.New("Producto con id " + strconv.Itoa(id) + " no encontrado")
	}

	if err := r.db.Write(ps); err != nil {
		return domain.Product{}, fmt.Errorf(FailWriting, err)
	}

	return product, nil
}

func (r *repository) UpdateNamePrice(id int, nombre string, precio int) (domain.Product, error) {
	var ps []domain.Product
	if err := r.db.Read(&ps); err != nil {
		return domain.Product{}, fmt.Errorf(FailReading)
	}

	var product domain.Product
	var updated bool
	for idx, p := range ps {
		if p.Id == id {
			ps[idx].Nombre = nombre
			ps[idx].Precio = precio
			product = ps[idx]
			updated = true
		}
	}

	if !updated {
		return domain.Product{}, errors.New("Producto con id " + strconv.Itoa(id) + " no encontrado")
	}

	if err := r.db.Write(ps); err != nil {
		return domain.Product{}, fmt.Errorf(FailWriting, err)
	}

	return product, nil
}

func (r *repository) Delete(id int) (domain.Product, error) {
	var ps []domain.Product
	if err := r.db.Read(&ps); err != nil {
		return domain.Product{}, fmt.Errorf(FailReading)
	}

	var product domain.Product
	var deleted bool
	for idx, p := range ps {
		if p.Id == id {
			product = p
			ps = append(ps[:idx], ps[idx+1:]...)
			deleted = true
		}
	}

	if !deleted {
		return domain.Product{}, errors.New("Producto con id" + strconv.Itoa(id) + "no encontrado")
	}

	if err := r.db.Write(ps); err != nil {
		return domain.Product{}, fmt.Errorf(FailWriting, err)
	}

	return product, nil
}

func (r *repository) LastID() (int, error) {
	var ps []domain.Product
	if err := r.db.Read(&ps); err != nil {
		return 0, fmt.Errorf(FailReading)
	}
	if len(ps) == 0 {
		return 0, nil
	}

	return ps[len(ps)-1].Id, nil
}

/* // Obtiene la lista de productos del archivo products.json
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
} */

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
