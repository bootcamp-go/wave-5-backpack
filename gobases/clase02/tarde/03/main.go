/*
Ejercicio 3 - Productos
Varias tiendas de ecommerce necesitan realizar una funcionalidad en Go para administrar productos y retornar el valor del precio total.
Las empresas tienen 3 tipos de productos: 
Pequeño, Mediano y Grande. (Se espera que sean muchos más)
Existen costos adicionales por mantener el producto en el almacén de la tienda, y costos de envío.

Sus costos adicionales son:
Pequeño: El costo del producto (sin costo adicional)
Mediano: El costo del producto + un 3% por mantenerlo en existencia en el almacén de la tienda.
Grande: El costo del producto + un 6%  por mantenimiento, y un costo adicional  por envío de $2500.

Requerimientos:
Crear una estructura “tienda” que guarde una lista de productos. 
Crear una estructura “producto” que guarde el tipo de producto, nombre y precio
Crear una interface “Producto” que tenga el método “CalcularCosto”
Crear una interface “Ecommerce” que tenga los métodos “Total” y “Agregar”.
Se requiere una función “nuevoProducto” que reciba el tipo de producto, su nombre y precio y devuelva un Producto.
Se requiere una función “nuevaTienda” que devuelva un Ecommerce.
Interface Producto:
El método “CalcularCosto” debe calcular el costo adicional según el tipo de producto.
Interface Ecommerce:
 - El método “Total” debe retornar el precio total en base al costo total de los productos y los adicionales si los hubiera.
 - El método “Agregar” debe recibir un producto y añadirlo a la lista de la tienda


*/

package main

import "fmt"

const (
	small = "Pequeño"
	medium = "Mediano"
	large = "Grande"
)

// estructura producto
type producto struct{
	productType string
	name string
	price float64
}

// estructura tienda
type tienda struct{
	lista []producto
}

// interface Producto
type Producto interface {
	CalcularCosto() float64
}

// interface Ecommerce
type Ecommerce interface {
	Total() float64
	Agregar(p producto)
}

// funcion nuevo producto
func nuevoProducto (tipo string, nombre string, precio float64) producto {
	return producto{productType: tipo, name:nombre, price:precio}
}

// funcion nueva tienda
func nuevaTienda () Ecommerce{
	return &tienda{}
}

// calcular costo total (supuestamente solo calcula el costo adicional)
func (p *producto) CalcularCosto() float64{
	switch p.productType{
	case medium:
		return p.price * 0.03
	case large:
		return p.price * 0.06 + 2500
	}
	return 0
}

func (t *tienda) Agregar(p producto) {
	t.lista = append(t.lista, p)
}

func (t *tienda) Total() float64 {
	var total float64 = 0
	for _, prod := range t.lista{
		total += prod.price + prod.CalcularCosto()
	}
	return total
}

func main() {
	ecomm1 := nuevaTienda()
	prod1 := nuevoProducto("Grande", "Tele", 42.12)
	prod2 := nuevoProducto("Grande", "Sillon", 51.2)
	prod3 := nuevoProducto("Pequeño", "Mesa", 31.21)
	prod4 := nuevoProducto("Mediano", "Play", 131.13)
	ecomm1.Agregar(prod1)
	ecomm1.Agregar(prod2)
	ecomm1.Agregar(prod3)
	ecomm1.Agregar(prod4)

	fmt.Println("El total de los productos de la tienda es de: ", ecomm1.Total())
}



