package main

import "fmt"

const (
	PEQUENIO = "Pequeño"
	MEDIANO  = "Mediano"
	GRANDE   = "Grande"
)

// ********* STRUCTS *********

type tienda struct {
	productos []Producto
}

type producto struct {
	tipo   string
	nombre string
	precio float64
}


// ********* INTERFACES *********

type Producto interface {
	CalcularCosto() float64
}

type Ecommerce interface {
	Total() float64
	Agregar(p Producto)
}

// ********* FUNCTIONS *********

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

func CalcularCostoTotal(costo, porcentaje, costoAdicional float64)  float64{
	calc := costo * (porcentaje / 100)
	return (costo + calc + costoAdicional)
}

func (p producto) CalcularCosto() float64 {
	switch p.tipo {
	case PEQUENIO:
		return p.precio
	case MEDIANO:
		return CalcularCostoTotal(p.precio, 3, 0)
	case GRANDE:
		return CalcularCostoTotal(p.precio, 6, 2500)
	default:
		return 0
	}
}

func nuevoProducto(tipo, nombre string, precio float64) Producto {
	return &producto{tipo,nombre,precio,}
}

func nuevaTienda() Ecommerce {
	return &tienda{
		productos: []Producto{},
	}
}
func main() {
	newTienda := nuevaTienda()
	newTienda.Agregar(nuevoProducto(PEQUENIO, "papas", 5000))
	newTienda.Agregar(nuevoProducto(MEDIANO, "yuca", 6000))
	newTienda.Agregar(nuevoProducto(GRANDE, "ñame", 9000))

	fmt.Printf("El total es: $%.f \n",  newTienda.Total())
}
