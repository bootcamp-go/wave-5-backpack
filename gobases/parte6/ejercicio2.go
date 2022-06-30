package main

import "fmt"

type Usuario struct {
	nombre    string
	apellido  string
	correo    string
	productos []*Producto
}

type Producto struct {
	nombre   string
	precio   float64
	cantidad int
}

func nuevoProducto(nombre string, precio float64) *Producto {
	return &Producto{nombre: nombre, precio: precio}
}

func agregarProducto(usuario *Usuario, producto *Producto, cantidad int) {
	producto.cantidad = cantidad
	usuario.productos = append(usuario.productos, producto)
}

func borrarProductos(usuario *Usuario) {
	usuario.productos = nil
}

func main() {
	user := Usuario{
		nombre: "Patricio",
	}

	producto := nuevoProducto("Casa", 500.54)

	agregarProducto(&user, producto, 2)

	fmt.Println(*user.productos[0])

	producto.nombre = "Auto"

	fmt.Println(*user.productos[0])

	borrarProductos(&user)

	fmt.Println(user.productos)
}
