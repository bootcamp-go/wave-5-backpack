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

func nuevoProducto(nombre string, precio float64) Producto {
	p1 := Producto{Nombre: nombre, Precio: precio}
	return p1
}

func agregarProducto(u *Usuario, p Producto, cantidad int) {
	p.Cantidad = cantidad
	u.Productos = append(u.Productos, p)
}

func borrarProductos(u *Usuario) {
	u.Productos = []Producto{}
}

func main() {
	p1 := nuevoProducto("leche", 2000)
	u1 := Usuario{Nombre: "Jessica", Apellido: "Escobar", Correo: "paola@hotmail.com"}
	agregarProducto(&u1, p1, 3)
	fmt.Println(u1)
	borrarProductos(&u1)
	fmt.Println(u1)
}
