package main

import "fmt"

type tienda struct {
	productos []Producto
}

type producto struct {
	tipo   string
	nombre string
	precio float64
}

type Producto interface {
	CalcularCosto() float64
}

type Ecommerce interface {
	Total() float64
	Agregar(p Producto)
}

func nuevoProducto(tipo, nombre string, precio float64) Producto {
	p1 := &producto{tipo, nombre, precio}

	return p1
}

const (
	Pequeno = "Pequeño"
	Mediano = "Mediano"
	Grande  = "Grande"
)

func (p producto) CalcularCosto() float64 {
	switch p.tipo {
	case Pequeno:
		return p.precio
	case Mediano:
		return p.precio * 1.03
	case Grande:
		return p.precio*1.06 + 2500
	}
	return 0
}

func nuevaTienda(productos ...Producto) Ecommerce {
	e1 := &tienda{productos}
	return e1
}

func (t *tienda) Agregar(p Producto) {
	t.productos = append(t.productos, p)
}

func (t tienda) Total() float64 {
	var costo_total float64
	for _, value := range t.productos {
		costo_total += value.CalcularCosto()
	}
	return costo_total
}

func main() {
	p1 := nuevoProducto(Pequeno, "leche", 1000)
	p2 := nuevoProducto(Mediano, "pan", 2000)
	p3 := nuevoProducto(Grande, "chocolate", 4000)

	e1 := nuevaTienda(p1, p2, p3)

	p4 := nuevoProducto(Pequeno, "dulce", 500)
	e1.Agregar(p4)

	fmt.Println(e1.Total())
}
