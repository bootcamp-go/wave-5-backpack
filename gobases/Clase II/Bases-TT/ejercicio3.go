package main

import "fmt"

const (
	Pequeno = "chico"
	Mediano = "mediano"
	Grande  = "grande"
)

type tienda struct {
	productos []Producto
}

type producto struct {
	typeProduct  string
	nameProduct  string
	priceProduct float64
}

type Producto interface {
	CalcularCosto() float64
}

func (p producto) CalcularCosto() float64 {
	switch p.typeProduct {
	case Pequeno:
		return p.priceProduct
	case Mediano:
		return p.priceProduct * 1.03
	case Grande:
		return (p.priceProduct*1.06 + 2500)
	}
	return 0.0
}

type Ecommerce interface {
	Total() float64
	Agregar(Producto)
}

func (t tienda) Total() float64 {
	var total float64
	for _, producto := range t.productos {
		total += producto.CalcularCosto()
	}
	return total
}

func (t *tienda) Agregar(aProducto Producto) {

	t.productos = append(t.productos, aProducto)
}

func nuevoProducto(value1, value2 string, price float64) Producto {
	nProducto := producto{typeProduct: value1, nameProduct: value2, priceProduct: price}
	return nProducto
}

func nuevaTienda() Ecommerce {
	return &tienda{}

}

func main() {

	produc1 := nuevoProducto(Grande, "Mesa", 500)
	produc2 := nuevoProducto(Mediano, "Laptop", 500)
	produc3 := nuevoProducto(Pequeno, "Celular", 200)

	nt := nuevaTienda()
	nt.Agregar(produc1)
	nt.Agregar(produc2)
	nt.Agregar(produc3)
	fmt.Println(produc1.CalcularCosto())
	fmt.Println(nt.Total())

}
