package main

import "fmt"

const (
	productoChico   = "chico"
	productoMediano = "mediano"
	productoGrande  = "grande"
)

type Producto interface {
	CalcularCosto() float64
}

type producto struct {
	tipoProducto string
	nombre       string
	precio       float64
}

func (p producto) CalcularCosto() float64 {
	switch p.tipoProducto {
	case productoChico:
		return p.precio
	case productoMediano:
		return p.precio + p.precio*0.03
	case productoGrande:
		return p.precio + p.precio*0.06 + 2500
	default:
		return -1
	}
}

type Ecommerce interface {
	Total() float64
	Agregar(p Producto)
}

type tienda struct {
	productos []Producto
}

func (t *tienda) Agregar(prod Producto) {
	t.productos = append(t.productos, prod)
}

func (t *tienda) Total() float64 {
	var aux float64 = 0
	for _, producto := range t.productos {
		aux += producto.CalcularCosto()
	}
	return aux
}

func nuevoProducto(tP string, nom string, pre float64) Producto {
	return &producto{tipoProducto: tP, nombre: nom, precio: pre}
}

func nuevaTienda() Ecommerce {
	return &tienda{}
}

func main() {
	ecomm1 := nuevaTienda()
	prod1 := nuevoProducto("Grande", "Armario", 42.12)
	prod2 := nuevoProducto("Grande", "Sillon", 51.2)
	prod3 := nuevoProducto("Pequeño", "Mesa ratona", 31.21)
	prod4 := nuevoProducto("Mediano", "Consola", 131.13)
	ecomm1.Agregar(prod1)
	ecomm1.Agregar(prod2)
	ecomm1.Agregar(prod3)
	ecomm1.Agregar(prod4)

	fmt.Println("El total de los productos de la tienda es de: ", ecomm1.Total())
}
