package main

import "fmt"

//Ejercicio 2 - Ecommerce
type Usuario struct {
	Nombre string
	Apellido string
	Correo string
	Productos []Producto
}

type Producto struct {
	Nombre string
	Precio float64
	Cantidad int
}

func nuevoProducto(nombre *string, precio *float64) *Producto {
	return &Producto{Nombre: *nombre, Precio: *precio}
}

func (u *Usuario) agregarProducto(p *Producto, cantidad *int) {
	p.Cantidad = *cantidad
	u.Productos = append(u.Productos, *p) 
}

func (u *Usuario) borrarProductos() {
	u.Productos = []Producto{}
}

func main() {

	nombre := "Silla"
	precio := 599000.00
	cantidad := 2

	producto := nuevoProducto(&nombre, &precio)

	usuario := &Usuario{
		Nombre: "Luisa",
		Apellido: "Arango",
		Correo: "lu@mail.com",
	}
	usuario.agregarProducto(producto, &cantidad)
	fmt.Println(usuario)

	usuario.borrarProductos()
	fmt.Println(usuario)

}