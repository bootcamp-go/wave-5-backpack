package main

import "fmt"

var (
	pequeno string = "PEQUEÑO"
	mediano string = "MEDIANO"
	grande  string = "GRANDE"
)

type Producto interface {
	Precio() float64
}

type Tienda struct {
	precio       float64
	tipoProducto string
}

func New(tipoProducto string, precio float64) Producto {
	return &Tienda{precio: precio, tipoProducto: tipoProducto}
}

func (p Tienda) Precio() float64 {
	switch p.tipoProducto {
	case pequeno: // Precio del producto más el costo de mantenerlo en tienda
		return p.precio
	case mediano:
		mantencion := (p.precio * 3) / 100
		return p.precio + mantencion
	case grande:
		mantencion := (p.precio * 6) / 100
		return p.precio + mantencion + 2500 // costo de envío
	default:
		return 0
	}
}

func main() {
	var precio float64 = 200

	tienda := New(grande, precio)
	fmt.Printf("Precio total del producto: 💰 %.2f\nEl producto es: %s \n", tienda.Precio(), grande)
}
