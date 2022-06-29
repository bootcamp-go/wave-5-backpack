package main

import "fmt"

const (
	PEQUENIO = "Pequeño"
	MEDIANO  = "Mediano"
	GRANDE   = "Grande"
)

// ---------- Struct tienda ----------

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

// ---------- Struct producto ----------

type producto struct {
	tipo   string
	nombre string
	precio float64
}

func (this producto) CalcularCosto() float64 {
	switch this.tipo {
	case PEQUENIO:
		return this.precio
	case MEDIANO:
		return this.precio * 1.03
	case GRANDE:
		return this.precio * 1.06
	default:
		return 0
	}
}

// ---------- Interface Producto ----------

type Producto interface {
	CalcularCosto() float64
}

// ---------- Interface Ecommerce ----------

type Ecommerce interface {
	Total() float64
	Agregar(p Producto)
}

// --------------- Funciones ----------------

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

// ----------------- Main -------------------
func main() {

	tienda := nuevaTienda()
	tienda.Agregar(nuevoProducto(PEQUENIO, "Pizza", 10.0))
	tienda.Agregar(nuevoProducto(MEDIANO, "Hamburguesa", 21.0))
	tienda.Agregar(nuevoProducto(GRANDE, "Pollo", 31.0))

	fmt.Println("Total de la tienda 1:", tienda.Total())

	tienda.Agregar(nuevoProducto(PEQUENIO, "Pizza", 10.0))
	tienda.Agregar(nuevoProducto(MEDIANO, "Hamburguesa", 20.0))

	fmt.Println("Nuevo total de la tienda 1:", tienda.Total())

	tienda2 := nuevaTienda()
	tienda2.Agregar(nuevoProducto(PEQUENIO, "Pizza", 10.0))
	tienda2.Agregar(nuevoProducto(PEQUENIO, "Hamburguesa", 20.0))

	fmt.Println("Total de la tienda 2:", tienda2.Total())

}
