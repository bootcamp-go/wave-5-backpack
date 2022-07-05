/*---------------------------------------------------------------------------------*

     Assignment:	Ejercicio #2:  Ecommerce
         Author:	Israel Fabela
	   Language:	go1.18.3 darwin/arm64
		  Topic:	Go Bases

	Description:
		A major web sales company needs to add functionality to add products to
		users. To do this they require that both users and products have the same
		memory address in the program main as in the functions.

		The following structures are needed:
			- User: First name, Last name, Email, Products (array of products).
			- Product: Name, Price, Quantity.

		The functions are required:
			- New product: receives name and price, and returns a product.
			- Delete products: receives a user, deletes the user's products.
			- Add product: receives user, product and quantity, returns nothing,
			  adds the product to the user.

	© Mercado Libre - IT Bootcamp 2022

----------------------------------------------------------------------------------*/

//	PACKAGE & LIBRARY
package main

import (
	"fmt"
)

//	STRUCT : Usuario
type Usuario struct {
	Nombre    string
	Apellido  string
	Correo    string
	Productos []producto
}

//	STRUCT : producto
type producto struct {
	NombreProd string
	Precio     float64
	Cantidad   int
}

//	FUNCTIONS
func nuevoProducto(nombre string, precio float64) producto {
	newProd := &producto{
		NombreProd: nombre,
		Precio:     precio,
	}
	return *newProd
}

func agregarProducto(user *Usuario, p *producto, cantidad int) {
	p.Cantidad = cantidad
	user.Productos = append(user.Productos, *p)
}

func borrarProductos(user *Usuario) {
	user.Productos = []producto{}
}

//	MAIN PROGRAM
func main() {
	fmt.Println("\n\t|| Ecommerce ||")

	usr1 := Usuario{
		Nombre:   "Lalo",
		Apellido: "Dorado",
		Correo:   "la.dorado@correo.com",
	}

	prd1 := nuevoProducto("tablet-X", 12000)

	fmt.Println("> Usuario:")

	agregarProducto(&usr1, &prd1, 3)

	fmt.Println("\tNombre:\t\t", usr1.Nombre)
	fmt.Println("\tApellido:\t", usr1.Apellido)
	fmt.Println("\tCorreo:\t\t", usr1.Correo)
	fmt.Println("\tProductos:\t", usr1.Productos)

	borrarProductos(&usr1)

	fmt.Println("\n> Datos de productos borrados:")

	fmt.Println("\tNombre:\t\t", usr1.Nombre)
	fmt.Println("\tApellido:\t", usr1.Apellido)
	fmt.Println("\tCorreo:\t\t", usr1.Correo)
	fmt.Println("\tProductos:\t", usr1.Productos)
}
