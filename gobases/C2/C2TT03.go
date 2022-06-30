package main

import "fmt"

type tienda struct {
	productos []producto
}

type producto struct {
	nombre string
	tipo   string
	precio float64
}

func (p producto) calcularCosto() float64 {

	return p.precio
}

type Producto interface {
	calcularCosto() float64
}

func (t tienda) Total() float64 {

	total := .0
	for _, prod := range t.productos {
		total += prod.precio + prod.calcularCosto()
	}
	return total
}

func (t *tienda) Agregar(prod producto) {

	t.productos = append(t.productos, prod)
}

type Ecommerce interface {
	Total() float64
	Agregar(prod producto)
}

func nuevoProducto(tipo, nombre string, precio float64) Producto {
	return producto{tipo, nombre, precio}
}

func nuevaTienda() Ecommerce {
	return &tienda{}
}

func main() {

	Ecommerce := nuevaTienda()
	Ecommerce.Agregar(nuevoProducto("reloj", "pequeno", 5000).(producto))
	Ecommerce.Agregar(nuevoProducto("cargador", "pequeno", 8000).(producto))
	fmt.Println("Total: ", Ecommerce.Total())
}
