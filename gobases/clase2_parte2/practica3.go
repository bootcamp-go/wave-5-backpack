package main

import "fmt"

const (
	pequeno = "pequeno"
	mediano = "mediano"
	grande = "grande"
)

type Tienda struct {
	productos []producto
}

type producto struct {
	tipoProducto string
	nombre string
	precio float64
}

type Producto interface {
	CalcularCosto() float64
}

type Ecommerce interface {
	Total() float64
	Agregar()
}

func nuevoProducto(nombre string, tipo string, precio float64) producto {
	item := producto{tipo, nombre, precio}
	return item
}

func nuevaTienda() Tienda {
	nuevaT := Tienda{}
	return nuevaT
}

func (t *Tienda) Agregar(item producto) {
	t.productos = append(t.productos, item)
}

func (t Tienda) Total() float64 {
	var cont float64 = 0

	for _, producto := range t.productos {
		fmt.Println("Producto costo:", producto.CalcularCosto())

		cont += producto.CalcularCosto()
	}
	return cont
}

func (p producto) CalcularCosto() float64 {
	switch p.tipoProducto {
	case pequeno:
		return p.precio
	case mediano:
		return p.precio + (p.precio * 3 / 100)
	case grande:
		return p.precio + float64(p.precio * 6 / 100) + 2500
	default:
		fmt.Println("El tipo de producto no existe")
		return 0
	}
}

func main()  {
	t := nuevaTienda()
	item1 := nuevoProducto("cepillo", pequeno, 4000)
	item2 := nuevoProducto("air fryer", mediano, 3000)
	item3 := nuevoProducto("sofa", grande, 14000)

	t.Agregar(item1)
	t.Agregar(item2)
	t.Agregar(item3)

	fmt.Println(t)
	fmt.Println(t.Total())
}