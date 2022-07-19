package main

import "fmt"

const (
	SMALL  = "Pequeño"
	MEDIUM = "Mediano"
	LARGE  = "Grande"
)

type tienda struct {
	productos []Producto
}

func (this tienda) Total() float64 {
	var total float64 = 0.0
	for _, producto := range this.productos {
		total += producto.CalcularCosto()
	}
	return total
}

func (this *tienda) Agregar(p Producto) {
	this.productos = append(this.productos, p)
}

type producto struct {
	tipo   string
	nombre string
	precio float64
}

func (this producto) CalcularCosto() float64 {
	switch this.tipo {
	case SMALL:
		return this.precio
	case MEDIUM:
		return this.precio * 1.03
	case LARGE:
		return this.precio * 1.06
	default:
		return 0
	}
}

type Producto interface {
	CalcularCosto() float64
}

type Ecommerce interface {
	Total() float64
	Agregar(p Producto)
}

func nuevoProducto(tipo string, nombre string, precio float64) Producto {
	prod := &producto{
		tipo:   tipo,
		nombre: nombre,
		precio: precio,
	}
	return prod
}

func nuevaTienda() Ecommerce {
	return &tienda{
		productos: []Producto{},
	}
}

func main() {

	tienda := nuevaTienda()
	tienda.Agregar(nuevoProducto(SMALL, "Pizza", 5990.0))
	tienda.Agregar(nuevoProducto(MEDIUM, "Hamburguesa", 6990.0))
	tienda.Agregar(nuevoProducto(LARGE, "Pollo", 8990.0))

	fmt.Printf("Total de la tienda 1: %.2f\n", tienda.Total())

	tienda.Agregar(nuevoProducto(SMALL, "Pizza", 5990.0))
	tienda.Agregar(nuevoProducto(MEDIUM, "Hamburguesa", 6990.0))

	fmt.Printf("Nuevo total de la tienda 1: %.2f\n", tienda.Total())

	tienda2 := nuevaTienda()
	tienda2.Agregar(nuevoProducto(SMALL, "Pizza", 5990.0))
	tienda2.Agregar(nuevoProducto(SMALL, "Hamburguesa", 5990.0))

	fmt.Printf("Total de la tienda 2: %.2f\n", tienda2.Total())

}
