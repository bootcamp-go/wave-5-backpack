package main

import "fmt"

type tienda struct {
	productos []producto
}

func (t tienda) Total() float64 {
	total := .0
	for _, prod := range t.productos {
		total += prod.precio + prod.CalcularCosto()
	}
	return total
}

func (t *tienda) Agregar(prod producto) {
	t.productos = append(t.productos, prod)
}

type producto struct {
	tipo   string
	nombre string
	precio float64
}

func (p producto) CalcularCosto() float64 {
	switch p.tipo {
	case "Mediano":
		return p.precio * 0.03
	case "Grande":
		return p.precio*0.06 + 2500
	default:
		return 0
	}
}

type Producto interface {
	CalcularCosto() float64
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

	ecommerce := nuevaTienda()
	ecommerce.Agregar(nuevoProducto("Pequeño", "Mouse", 2000).(producto))
	ecommerce.Agregar(nuevoProducto("Mediano", "Notebook", 50000).(producto))
	ecommerce.Agregar(nuevoProducto("Grande", "Lavarropa", 70000).(producto))

	fmt.Printf("Total: $%v\n", ecommerce.Total())
}
