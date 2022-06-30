package main

import "fmt"

func main() {
	ecomm1 := nuevaTienda()
	prod1 := nuevoProducto("Grande", "Tele", 42.12)
	prod2 := nuevoProducto("Grande", "Sillon", 51.2)
	prod3 := nuevoProducto("Pequeño", "Mesa", 31.21)
	prod4 := nuevoProducto("Mediano", "Play", 131.13)
	ecomm1.Agregar(prod1)
	ecomm1.Agregar(prod2)
	ecomm1.Agregar(prod3)
	ecomm1.Agregar(prod4)

	fmt.Println("El total de los productos de la tienda es de: ", ecomm1.Total())
}

type tienda struct {
	prods []producto
}

type producto struct {
	tipoProducto string
	nombre       string
	precio       float64
}

type Producto interface {
	CalcularCosto() float64
}

type Ecommerce interface {
	Total() float64
	Agregar(prod producto)
}

func (t *tienda) Agregar(prod producto) {
	t.prods = append(t.prods, prod)
}

func (t *tienda) Total() float64 {
	var aux float64 = 0
	for _, pr := range t.prods {
		aux += pr.precio + pr.CalcularCosto()
	}
	return aux
}

func (p producto) CalcularCosto() float64 {
	switch p.tipoProducto {
	case "Mediano":
		return p.precio * 0.03
	case "Grande":
		return p.precio*0.06 + 1500.0
	default:
		return 0
	}
}

func nuevoProducto(tP string, nom string, pre float64) (p producto) {
	return producto{tipoProducto: tP, nombre: nom, precio: pre}
}

func nuevaTienda() Ecommerce {
	return &tienda{}
}
