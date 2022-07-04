package main

import "fmt"

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

func main() {
	user := Usuario{}
	user.Nombre = "Jose"
	user.Apellido = "Gonzalez"
	user.Correo = "j@meli.com"
	p := nuevoProducto("gaseosa", 215)
	agregarProducto(&user, &p, 2)
	agregarProducto(&user, &p, 2)
	fmt.Print(user)
	borrarProductos(&user)
}

func nuevoProducto(nombre string, precio float64) Producto {
	return Producto{nombre, precio, 0}
}

func agregarProducto(user *Usuario, prod *Producto, cant int) {
	prod.Cantidad = cant
	user.Productos = append(user.Productos, *prod)
}

func borrarProductos(user *Usuario) {
	user.Productos = nil
}
