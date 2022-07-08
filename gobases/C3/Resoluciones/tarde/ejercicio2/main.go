package main

import "fmt"

/*Ejercicio 2 - Ecommerce
Una importante empresa de ventas web necesita agregar una funcionalidad para agregar productos
a los usuarios. Para ello requieren que tanto los usuarios como los productos tengan la misma
dirección de memoria en el main del programa como en las funciones.
Se necesitan las estructuras:
	1. Usuario: Nombre, Apellido, Correo, Productos (array de productos).
	2. Producto: Nombre, precio, cantidad.
Se requieren las funciones:
	1. Nuevo producto: recibe nombre y precio, y retorna un producto.
	2. Agregar producto: recibe usuario, producto y cantidad, no retorna nada, agrega el producto al usuario.
	3. Borrar productos: recibe un usuario, borra los productos del usuario.
*/

type Usuario struct {
	Nombre    string
	Apellido  string
	Correo    string
	Productos []Producto
}

type Producto struct {
	Nombre   string
	Precio   float64
	Cantidad int
}

func NewProducto(nombre *string, precio *float64) *Producto {
	return &Producto{Nombre: *nombre, Precio: *precio}
}

func (u *Usuario) AgregarProducto(producto *Producto, cantidad *int) {
	producto.Cantidad = *cantidad
	u.Productos = append(u.Productos, *producto)
}

func (u *Usuario) BorrarProducto() {
	u.Productos = []Producto{}
}

func main() {
	var (
		nombre   string  = "La virginia"
		precio   float64 = 1000
		cantidad int     = 10
	)

	producto := NewProducto(&nombre, &precio)
	usuario := &Usuario{
		Nombre:   "Marcelo",
		Apellido: "Gonzalez",
		Correo:   "test@gmail.com",
	}

	usuario.AgregarProducto(producto, &cantidad)

	fmt.Println("Usuario - ", usuario.Nombre, usuario.Apellido)
	fmt.Println("Correo - ", usuario.Correo)
	for _, value := range usuario.Productos {
		fmt.Printf("Producto - %s - Cantidad %d - Precio %.2f - Total $%.2f\n", value.Nombre, value.Cantidad, value.Precio, (value.Precio * float64(value.Cantidad)))
	}

	fmt.Println()
	fmt.Println("Eliminando productos . . .")
	usuario.BorrarProducto()

	fmt.Println("Cantidad de productos después de la eliminación", len(usuario.Productos))
}
