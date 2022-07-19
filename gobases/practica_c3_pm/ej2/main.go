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
	precio   float64
	cantidad int
}

func NuevoProducto(nombre string, precio float64) *Producto {
	return &Producto{
		Nombre: nombre,
		precio: precio,
	}
}

func AgregarProducto(usuario *Usuario, producto *Producto, cantidad int) {
	producto.cantidad = cantidad
	usuario.Productos = append(usuario.Productos, *producto)
}

func BorrarProductos(usuario *Usuario) {
	usuario.Productos = *new([]Producto)
}

func main() {
	var usuario1 *Usuario = &Usuario{
		Nombre:   "Matias",
		Apellido: "Carrasco",
		Correo:   "matias.carrasco@mercadolibre.cl",
	}

	var producto1 *Producto = NuevoProducto("Laptop", 1000)
	var producto2 *Producto = NuevoProducto("Mouse", 50)
	var producto3 *Producto = NuevoProducto("Teclado", 100)

	AgregarProducto(usuario1, producto1, 1)
	AgregarProducto(usuario1, producto2, 2)
	AgregarProducto(usuario1, producto3, 3)

	fmt.Println("Usuario: ", *usuario1)
	fmt.Println("Usuario direccion de memoria: ", &usuario1)

	fmt.Println("Producto1: ", *producto1)
	fmt.Println("Producto1 direccion de memoria: ", &producto1)
	apuntadorProducto1EnListaUsuario := &usuario1.Productos[0]
	fmt.Println("Producto1 direccion de memoria en productos del usuario: ", &apuntadorProducto1EnListaUsuario)

	fmt.Println("Producto2: ", *producto2)
	fmt.Println("Producto2 direccion de memoria: ", &producto2)
	apuntadorProducto2EnListaUsuario := &usuario1.Productos[1]
	fmt.Println("Producto2 direccion de memoria en productos del usuario: ", &apuntadorProducto2EnListaUsuario)

	BorrarProductos(usuario1)

	fmt.Println("Usuario: ", *usuario1)

}
