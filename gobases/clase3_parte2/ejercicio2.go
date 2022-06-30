package main

import "fmt"

type Usuario struct {
	Nombre, Apellido, Correo string
	Productos                []Producto
}

type Producto struct {
	Nombre           string
	Precio, Cantidad int
}

func main() {
	usuario1 := Usuario{Nombre: "Camilo", Apellido: "Calderon", Correo: "camilo@correo.com"}
	fmt.Println(usuario1)
	producto1 := NuevoProducto("Papas", 2000)
	producto2 := NuevoProducto("Bebida", 5000)
	producto3 := NuevoProducto("Galleta", 1000)
	AgregarProducto(&usuario1, producto1, 5)
	AgregarProducto(&usuario1, producto2, 3)
	AgregarProducto(&usuario1, producto3, 2)
	fmt.Println(usuario1)
	BorrarProducto(&usuario1)
	fmt.Println(usuario1)
}

func NuevoProducto(Nombre string, Precio int) Producto {
	return Producto{Nombre: Nombre, Precio: Precio}
}

func AgregarProducto(Usuario *Usuario, Producto Producto, Cantidad int) {
	Producto.Cantidad = Cantidad
	Usuario.Productos = append(Usuario.Productos, Producto)
}

func BorrarProducto(Usuario *Usuario) {
	Usuario.Productos = []Producto{}
}
