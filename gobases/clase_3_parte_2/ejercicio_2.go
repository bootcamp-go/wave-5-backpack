package main

import "fmt"

type Producto struct {
	Nombre   string
	Precio   int
	Cantidad int
}

type Usuario struct {
	Nombre    string
	Apellido  string
	Correo    string
	Productos []*Producto
}

func NewProducto(nombre string, precio int) *Producto {
	producto := Producto{Nombre: nombre, Precio: precio}
	return &producto
}

func AgregarProducto(usuario *Usuario, producto *Producto, cantidad int) {
	producto.Cantidad = cantidad
	usuario.Productos = append(usuario.Productos, producto)
}

func BorrarProductos(usuario *Usuario) {
	usuario.Productos = []*Producto{}
}

func main() {
	producto := NewProducto("Iphone 9", 250000)
	usuario := Usuario{Nombre: "Claudio", Apellido: "Figueroa", Correo: "claudio@gmail.com"}

	AgregarProducto(&usuario, producto, 5)
	fmt.Println(usuario)
	fmt.Println(*producto)
	BorrarProductos(&usuario)
	fmt.Println(usuario)

}
