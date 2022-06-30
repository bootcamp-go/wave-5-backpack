package main

import "fmt"

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

const (
	Pequeno string = "Pequeño"
	Mediano string = "Mediano"
	Grande  string = "Grande"
)

func nuevoProducto(tipo string, nombre string, precio float64) Producto {
	newProducto := &producto{tipo, nombre, precio}
	return newProducto
}

func nuevaTienda(productos ...Producto) Ecommerce {
	newTienda := &tienda{productos}
	return newTienda
}

func (p producto) CalcularCosto() float64 {
	switch p.Tipo {
	case Pequeno:
		return p.Precio
	case Mediano:
		return p.Precio * 1.03
	case Grande:
		return p.Precio*1.06 + 2500
	}
	return 0
}

func (t tienda) Total() float64 {
	var total float64
	for _, producto := range t.Productos {
		total += producto.CalcularCosto()
	}
	return total
}

func (t *tienda) Agregar(p Producto) {
	t.Productos = append(t.Productos, p)
}

func main() {
	producto1 := nuevoProducto("Pequeño", "Café", 1500)
	producto2 := nuevoProducto("Mediano", "Licuadora", 12000)
	producto3 := nuevoProducto("Grande", "Heladera", 120000)

	fmt.Println("Costo café:", producto1.CalcularCosto())
	fmt.Println("Costo licuadora", producto2.CalcularCosto())
	fmt.Println("Costo heladera", producto3.CalcularCosto())

	tiendaNueva := nuevaTienda(producto1, producto2, producto3)
	fmt.Println(tiendaNueva.Total())

	producto4 := nuevoProducto("Grande", "Smart tv", 130000)
	tiendaNueva.Agregar(producto4)
	fmt.Println("Total:", tiendaNueva.Total())
}
