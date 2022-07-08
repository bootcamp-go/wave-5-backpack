package main

import "fmt"

type producto struct {
	tipo   string
	nombre string
	precio float64
}

type tienda struct {
	productos []producto
}

type Producto interface {
	CalcularCosto() float64
}

type Ecommerce interface {
	Total() float64
	Agregar(p producto)
}

func nuevoProducto(nombre string, tipo string, precio float64) producto {
	return producto{tipo: tipo, nombre: nombre, precio: precio}
}

func nuevaTienda() Ecommerce {
	return &tienda{}
}

func (t *tienda) Agregar(p producto) {
	t.productos = append(t.productos, p)
}

func (t *tienda) Total() float64 {
	total := 0.0
	for _, e := range t.productos {
		total += e.precio + e.CalcularCosto()
	}

	return total
}

func (p producto) CalcularCosto() float64 {
	switch {
	case p.tipo == "mediano":
		return p.precio * 0.03
	case p.tipo == "grande":
		return (p.precio * 0.06) + 2500
	default:
		return 0
	}
}

func main() {
	tienda1 := nuevaTienda()
	product1 := nuevoProducto("Blem", "mediano", 100)
	product2 := nuevoProducto("Pipi", "grande", 200)
	product3 := nuevoProducto("Cuqui", "pequeno", 150)

	tienda1.Agregar(product1)
	tienda1.Agregar(product2)
	tienda1.Agregar(product3)

	fmt.Printf("El total de los productos de la tienda es: %.2f\n", tienda1.Total())
}
