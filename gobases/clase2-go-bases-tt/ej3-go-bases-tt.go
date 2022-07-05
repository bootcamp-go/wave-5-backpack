/*---------------------------------------------------------------------------------*

     Assignment:	Ejercicio #3:  Productos
         Author:	Israel Fabela
	   Language:	go1.18.3 darwin/arm64
		  Topic:	Go Bases

	Description:
		To create a functionality in Go to manage products and return products
		and return the total price value.
		Companies have 3 types of products:
			- Small, Medium and Large (expected to be many more).
		There are additional costs for keeping the product in the store's
		warehouse, and shipping costs.

		Your additional costs are:
			- Small : The cost of the product (no additional cost).
			- Medium : The cost of the product + 3% for keeping it in stock
			  in the store's warehouse.
			- Large : The cost of the product + a 6% maintenance fee, and
			  an additional shipping fee of $2500.

		Requirements :
			- Create a "store" structure that stores a list of products.
			- Create a "product" structure that stores the product type,
			  name and price.
			- Create a "Product" interface that has the "CalculateCost"
			  method.

	© Mercado Libre - IT Bootcamp 2022

----------------------------------------------------------------------------------*/

//	PACKAGE & LIBRARY
package main

import "fmt"

//	CONSTANTS
const (
	PEQUENO = "Pequeño"
	MEDIANO = "Mediano"
	GRANDE  = "Grande"
)

//	STRUCT : tienda
type tienda struct {
	lista []Producto
}

// FUNCTIONS : tienda
func nuevaTienda() Ecommerce {
	store := &tienda{
		lista: []Producto{},
	}
	return store
}

func (t *tienda) Agregar(p Producto) {
	t.lista = append(t.lista, p)
}

func (t tienda) Total() float64 {
	total := 0.0
	for _, producto := range t.lista {
		total += producto.CalcularCosto()
	}
	return total
}

//	STRUCT : producto
type producto struct {
	tipoDeProd string
	nombre     string
	precio     float64
}

//	FUNCTIONS : nuevoProducto
func nuevoProducto(tipoProducto string, nombreProducto string, precioProducto float64) Producto {
	new_prod := &producto{
		tipoDeProd: tipoProducto,
		nombre:     nombreProducto,
		precio:     precioProducto,
	}
	return new_prod
}

//	FUNCTIONS : (producto).CalcularCosto()
func (p producto) CalcularCosto() float64 {
	switch p.tipoDeProd {
	case PEQUENO:
		return p.precio
	case MEDIANO:
		return p.precio * 1.03
	case GRANDE:
		return p.precio * 1.06
	default:
		return 0
	}
}

//	INTERFACE : Producto
type Producto interface {
	CalcularCosto() float64
}

//	INTERFACE : Ecommerce
type Ecommerce interface {
	Total() float64
	Agregar(p Producto)
}

//	MAIN PROGRAM
func main() {
	fmt.Println("\n\t|| Calculo Estadistico ||")

	store1 := nuevaTienda() // Creando una nueva Tienda

	store1.Agregar(nuevoProducto(PEQUENO, "Xiami band", 1500)) // Agregando productos
	store1.Agregar(nuevoProducto(MEDIANO, "iPad", 24000))
	store1.Agregar(nuevoProducto(GRANDE, "Ryzer", 32000))

	fmt.Println("Total #1 de la 'Store1:\t", store1.Total()) // Imprimiendo Total

	store1.Agregar(nuevoProducto(PEQUENO, "Mouse", 2000)) // Agregando productos

	fmt.Println("Total #2 de la 'Store1:\t", store1.Total()) // Imprimiendo Total
}
