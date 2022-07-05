package main

import "fmt"

type tienda struct {
	listaProductos []producto
}

type producto struct{
	Tipo string
	Nombre string
	Precio float64
}

type iProducto interface{
	CalcularCosto() float64
}

type iEcommerce interface{
	Total() float64
	Agregar()
}
func nuevoProducto(tipo string, nombre string, precio float64) producto{
	prod:=producto{
		Tipo: tipo,
		Nombre: nombre,
		Precio: precio,
	}
	return prod
}
func nuevaTienda() tienda{
	tiend:=tienda{}
	return tiend
}

func (prod *producto)CalcularCosto() float64{
	var costo=0.0
	switch prod.Tipo{
	case "Pequeño":
		costo=prod.Precio
	case "Mediano":
		costo=prod.Precio*1.03
	case "Grande":
		costo=(prod.Precio*1.06)+2500
	default:
		costo=0.0
	}
	return costo
}
func (tiend *tienda)Total() float64{
	var total=0.0
	for _,prod :=range tiend.listaProductos {
		total=total+prod.CalcularCosto();
	}
	return total
}
func (tiend *tienda)Agregar(prod producto) {
	tiend.listaProductos=append(tiend.listaProductos,prod)
}

func main(){
	fmt.Println("prueba")
}


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