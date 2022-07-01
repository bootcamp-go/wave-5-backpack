package main

import "fmt"

const (
	PEQ = "pequeño"
	MED = "mediano"
	GRA = "grande"
)

type Producto interface {
	CalcularCosto() float64
}

type Ecommerce interface {
	Total() float64
	Agregar(producto)
}

type tienda struct {
	Productos []producto
}

type producto struct {
	Tipo   string
	Nombre string
	Precio float64
}

func (p producto) CalcularCosto() float64 {
	switch p.Tipo {
	case PEQ:
		return p.Precio
	case MED:
		return calcularTotal(p.Precio, 3, 0)
	case GRA:
		return calcularTotal(p.Precio, 6, 2500)
	default:
		return 0
	}
}

func (t tienda) Total() float64 {
	total := 0.0
	for _, prod := range t.Productos {
		total += calcularTotalAll(prod)
	}
	return total
}

func (t *tienda) Agregar(prd producto) {
	t.Productos = append(t.Productos, prd)
}

func nuevoProducto(tipoProd, nombre string, precio float64) producto {
	return producto{
		Tipo:   tipoProd,
		Nombre: nombre,
		Precio: precio,
	}
}

func nuevaTienda() Ecommerce {
	return &tienda{}
}

func calcularTotal(costo, porcentajeExtra, adicional float64) float64 {
	valorExtra := costo * (porcentajeExtra / 100)
	return (costo + valorExtra + adicional)

}

func calcularTotalAll(prd Producto) float64 {
	return prd.CalcularCosto()
}

func main() {
	prd1 := nuevoProducto(PEQ, "celular", 400)
	prd2 := nuevoProducto(MED, "silla", 200)
	prd3 := nuevoProducto(GRA, "refrigerador", 600)

	tienda := nuevaTienda()
	tienda.Agregar(prd1)
	fmt.Println("Total tienda 1 prod :", tienda.Total())

	tienda.Agregar(prd2)
	fmt.Println("Total tienda 2 prod :", tienda.Total())

	tienda.Agregar(prd3)
	fmt.Println("Total tienda 3 prod :", tienda.Total())

	fmt.Println("Total producto 3 :", calcularTotalAll(&prd3))

}
