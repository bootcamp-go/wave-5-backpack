package main

import "fmt"

const (
	small  = "pequeno"
	medium = "mediano"
	big    = "grande"
)

type producto struct {
	precio float64
	nombre string
	tipo   string
}

type tienda struct {
	productos []producto
}

type Producto interface {
	CalcularCosto()
}

type Ecommerce interface {
	Agregar(p producto)
	Total() float64
}

func (t *tienda) Agregar(p producto) {
	t.productos = append(t.productos, p)
}
func (t tienda) Total() float64 {
	var total float64 = 0
	for _, producto := range t.productos {
		total += producto.CalcularCosto()
	}

	return total
}

func (p producto) CalcularCosto() float64 {
	switch p.tipo {
	case small: // Precio del producto más el costo de mantenerlo en tienda
		return p.precio
	case medium:
		almacen := (p.precio * 3) / 100
		return p.precio + almacen
	case big:
		almacen := (p.precio * 6) / 100
		return p.precio + almacen + 2500 // costo de envío
	default:
		return 0
	}
}
func nuevoProducto(tipoProducto string, nombre string, precio float64) producto {
	return producto{tipo: tipoProducto, nombre: nombre, precio: precio}
}

func nuevaTienda() Ecommerce {
	return &tienda{productos: []producto{}}
}

func main() {
	nTienda := nuevaTienda()
	fmt.Println("nueva tienda", nTienda)

	prod1 := nuevoProducto("pequeno", "sacapuntas", 1.50)
	prod2 := nuevoProducto("mediano", "cuaderno", 3.20)
	prod3 := nuevoProducto("grande", "mochila", 31.52)

	nTienda.Agregar(prod1)
	nTienda.Agregar(prod2)
	nTienda.Agregar(prod3)

	//fmt.Println("Productos tienda: ", nTienda)

	fmt.Println("Costo total de la  tienda: ", nTienda.Total())

}
