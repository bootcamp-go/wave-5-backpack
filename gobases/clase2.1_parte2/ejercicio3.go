package main

import (
	"fmt"
)

//Ejercicio 3 - Productos
const (
	Pequeno string = "Pequeno"
	Mediano string = "Mediano"
	Grande  string = "Grande"
)

type tienda struct {
	Productos []Producto
}

type producto struct {
	Tipo   string
	Nombre string
	Precio float64
}

type Producto interface {
	CalcularCosto() float64
}

type Ecommerce interface {
	Total() float64
	Agregar(p Producto)
}

func nuevoProducto(tipo string, nombre string, precio float64) Producto {
	pr := &producto{tipo, nombre, precio}
	return pr
}

func nuevaTienda(productos ...Producto) Ecommerce {
	ti := &tienda{productos}
	return ti
}

func (pr producto) CalcularCosto() float64 {
	switch pr.Tipo {
	case Pequeno:
		return pr.Precio
	case Mediano:
		return pr.Precio * 1.03
	case Grande:
		return pr.Precio*1.06 + 2500
	}
	return 0
}

func (ti tienda) Total() float64 {
	var total float64
	for _, pr := range ti.Productos {
		total += pr.CalcularCosto()
	}
	return total
}

func (ti *tienda) Agregar(pr Producto) {
	ti.Productos = append(ti.Productos, pr)
}

func main() {
	product1 := nuevoProducto("Grande", "Panela", 2500)
	product2 := nuevoProducto("Pequeno", "Maiz", 1500)
	product3 := nuevoProducto("Mediano", "Chocorramo", 4500)

	fmt.Println("Precio Total Producto 1: $", 
product1.CalcularCosto())
	fmt.Println("Precio Total Producto 2: $", 
product2.CalcularCosto())
	fmt.Println("Precio Total Producto 3: $", 
product3.CalcularCosto())

	tienda := nuevaTienda(product1, product2, product3)
	fmt.Println("Precio Total Articulos de la tienda: $", 
tienda.Total())
}
