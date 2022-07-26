package main

import "fmt"

const (
	PEQUENIO = "Pequeño"
	MEDIANO  = "Mediano"
	GRANDE   = "Grande"
)

type producto struct {
	tipo   string
	nombre string
	precio float64
}

type tienda struct {
	productos []Producto
}

type Producto interface {
	CalcularCosto() float64
}

type Ecommerce interface {
	Total() float64
	Agregar(p Producto)
}

func (t tienda) Total() float64 {
	var total float64 = 0
	for _, producto := range t.productos {
		total += producto.CalcularCosto()
	}
	return total
}

func (t *tienda) Agregar(p Producto) {
	t.productos = append(t.productos, p)
}

func (p producto) CalcularCosto() float64 {
	switch p.tipo {
	case PEQUENIO:
		return p.precio
	case MEDIANO:
		return p.precio * 1.03
	case GRANDE:
		return p.precio*1.06 + 2500
	default:
		return 0
	}
}

func nuevoProducto(tipo string, nombre string, precio float64) Producto {
	return &producto{tipo: tipo, nombre: nombre, precio: precio}
}

func nuevaTienda() Ecommerce {
	return &tienda{productos: []Producto{}}
}

func main() {

	lapiz := nuevoProducto(PEQUENIO, "Lápiz", 80)
	cuaderno := nuevoProducto(MEDIANO, "Cuaderno", 300)
	escritorio := nuevoProducto(GRANDE, "Escritorio", 4000)

	miTienda := nuevaTienda()
	miTienda.Agregar(lapiz)
	miTienda.Agregar(cuaderno)
	miTienda.Agregar(escritorio)

	fmt.Println("El total de mi tienda es: ", miTienda.Total())

}
