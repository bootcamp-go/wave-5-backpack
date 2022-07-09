package main

import "fmt"

const (
	pequeno = "pequeño"
	mediano = "mediano"
	grande  = "grande"
)

type tienda struct {
	productos []producto
}

type producto struct {
	tipo   string
	nombre string
	precio float64
}

type Producto interface {
	calcularCosto() float64
}

type Ecommerce interface {
	total() float64
	agregar(producto)
}

func main() {
	t := nuevaTienda()
	item1 := nuevoProducto("crema", pequeno, 4000)
	item2 := nuevoProducto("garrafon de agua", mediano, 3000)
	item3 := nuevoProducto("sarten", grande, 14000)

	t.agregar(item1)
	t.agregar(item2)
	t.agregar(item3)

	fmt.Println(t)

	fmt.Println(t.total())
}

func (t *tienda) agregar(item producto) {
	t.productos = append(t.productos, item)
}
func (t tienda) total() float64 {
	fmt.Println("entro en total")
	var cont float64 = 0
	// devuelve el costo total de todos los productos incluyendo costos adicionales
	for _, producto := range t.productos {
		fmt.Println("producto costo:", producto.calcularCosto())

		cont += producto.calcularCosto()
	}
	return cont
}
func nuevoProducto(nombre string, tipo string, precio float64) producto {
	item := producto{tipo, nombre, precio}
	return item
}

func nuevaTienda() tienda {
	//crea una nueva tienda
	nuevaT := tienda{}
	return nuevaT
}

func (p producto) calcularCosto() float64 {
	//calculo de un producto con sus costos adicionales (si los hay)
	switch p.tipo {
	case pequeno:
		return p.precio
	case mediano:
		return p.precio + (p.precio * 3 / 100)
	case grande:
		return p.precio + float64(p.precio*6/100) + 2500
	default:
		fmt.Println("entra en default??")
		return 0
	}
}

/*
Crear una estructura “tienda” que guarde una lista de productos.
Crear una estructura “producto” que guarde el tipo de producto, nombre y precio
Crear una interface “Producto” que tenga el método “CalcularCosto”
Crear una interface “Ecommerce” que tenga los métodos “Total” y “Agregar”.
*/
