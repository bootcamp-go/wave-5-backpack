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
	usuario := Usuario{"Cosme", "Fulanito", "cosme.fulanito@gmail.com", []Producto{}}
	fmt.Println("Usuario antes de agregar productos: ", usuario)

	caramelo := nuevoProducto("Caramelos", 1)
	agregarProducto(&usuario, caramelo, 5)

	fmt.Println("Usuario despues de agregar un producto: ", usuario)

	borrarProductos(&usuario)

	fmt.Println("Usuario despues de borrar productos: ", usuario)
}

func nuevoProducto(nombre string, precio float64) Producto {
	return Producto{nombre, precio, 0}
}

func agregarProducto(usuario *Usuario, producto Producto, cantidad int) {
	producto.Cantidad = cantidad
	usuario.Productos = append(usuario.Productos, producto)
}

func borrarProductos(usuario *Usuario) {
	usuario.Productos = []Producto{}
}
