package main

import "fmt"

var (
	peque   string = "pequeño"
	mediano string = "mediano"
	grande  string = "grande"
)

type tienda struct {
	prods []Producto
}

type producto struct {
	tipo_producto string
	nombre        string
	precio        float64
}

type Ecommerce interface {
	Total() float64
	Agregar(producto)
}

type Producto interface {
	CalcularCosto() float64
}

func nuevoProducto(tipo_producto string, nombre string, precio float64) Producto {
	return &producto{tipo_producto: tipo_producto, nombre: nombre, precio: precio}
}

func nuevaTienda(productos ...Producto) Ecommerce {
	return &tienda{prods: productos}
}

func (prod producto) CalcularCosto() float64 {
	switch prod.tipo_producto {
	case peque:
		return prod.precio
	case mediano:
		return prod.precio * 1.3
	case grande:
		return prod.precio*1.6 + 2500
	default:
		fmt.Println("No existe el tipo de producto indicado")
		return prod.precio
	}
}

func (t tienda) Agregar(prod producto) {
	t.prods = append(t.prods, prod)
}

func (t tienda) Total() float64 {
	var total float64
	for _, val := range t.prods {
		total += val.CalcularCosto()
	}
	return total
}

func main() {
	var tiend Ecommerce
	tiend = nuevaTienda(
		nuevoProducto(peque, "HyperX Cloud II Wireless", 17300.77),
		nuevoProducto(grande, "Sony Wh1000xm4", 62333.77),
		nuevoProducto(peque, "StellSeries RX22", 44331.22),
	)

	fmt.Println(tiend.Total())
}
