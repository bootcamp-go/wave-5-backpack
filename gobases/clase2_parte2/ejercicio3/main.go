package main

import "fmt"

func main() {

	p1 := producto{
		tipoProd: "PEQUEÑO",
		nombre:   "LG",
		precio:   200,
	}

	p2 := producto{
		tipoProd: "PEQUEÑO",
		nombre:   "LG",
		precio:   200,
	}

	tienda1 := nuevaTienda()
	tienda1.agregar(p1)
	tienda1.agregar(p2)
	fmt.Println("El total es:", tienda1.total())

}

type producto struct {
	tipoProd string
	nombre   string
	precio   float64
}

type tienda struct {
	listaProds []producto
}

func nuevaTienda() Ecommer {
	return &tienda{}
}

func (t *tienda) agregar(p producto) {
	t.listaProds = append(t.listaProds, p)
}

func (t *tienda) total() float64 {
	total := 0.0
	for _, prodI := range t.listaProds {
		total += float64(prodI.calcularCosto())
	}
	return float64(total)
}

type Producto interface {
	calcularCosto() float64
}

type Ecommer interface {
	total() float64
	agregar(p producto)
}

func nuevoProducto(tp string, n string, pre float64) Producto {
	return &producto{tipoProd: tp, nombre: n, precio: pre}
}

func (p producto) calcularCosto() float64 {
	switch p.tipoProd {
	case "PEQUEÑO":
		return p.precio
	case "MEDIANO":
		return p.precio * 1.03
	case "GRANDE":
		return p.precio*1.06 + 2500.0
	default:
		return 0.0
	}

}
