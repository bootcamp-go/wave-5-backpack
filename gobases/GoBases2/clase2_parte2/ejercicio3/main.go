package main

import "fmt"

type tienda struct {
	listaProductos []Producto
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
	Agregar(p Producto)
}

func (p producto) CalcularCosto() float64 {
	costo_adicional := 0.0

	switch p.tipoProducto {
	case "pequeno":
		costo_adicional = p.precio
	case "mediano":
		costo_adicional = p.precio * (1 + 0.03)
	case "grande":
		costo_adicional = p.precio*(1+0.06) + 2500
	}
	return costo_adicional
}

func (t tienda) Total() float64 {
	res := 0.0

	for _, p := range t.listaProductos {
		res += p.CalcularCosto()
	}
	return res
}

func (t *tienda) Agregar(p Producto) {
	t.listaProductos = append(t.listaProductos, p)
}

func nuevoProducto(tipoProducto, nombre string, precio float64) Producto {
	return &producto{
		tipoProducto: tipoProducto,
		nombre:       nombre,
		precio:       precio,
	}
}

func nuevaTienda(productos ...Producto) Ecommerce {
	return &tienda{
		listaProductos: productos,
	}
}

func main() {

	p1 := nuevoProducto("pequeno", "p1", 100)
	p2 := nuevoProducto("mediano", "p2", 100)
	p3 := nuevoProducto("grande", "p2", 100)

	t1 := nuevaTienda(p1, p2, p3)
	fmt.Println(p1.CalcularCosto(), p2.CalcularCosto(), p3.CalcularCosto())
	fmt.Println(t1.Total())
	t1.Agregar(nuevoProducto("pequeno", "nuevo", 100))
	fmt.Println(t1.Total())
}
