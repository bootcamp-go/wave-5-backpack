package main

import "fmt"

//Ejercicio 2 - Ecommerce

type Usuario struct {
	Nombre   string
	Apellido string
	Correo   string
	Products []Producto
}

type Producto struct {
	Nombre   string
	Precio   float64
	Cantidad int
}

func newProduct(nombre string, precio float64) Producto {
	p := Producto{Nombre: nombre, Precio: precio}
	fmt.Println("Producto Creado!")
	return p
}

func addProduct(u *Usuario, p *Producto, cantidad int) {
	p.Cantidad = cantidad
	u.Products = append(u.Products, *p)
	fmt.Println("Producto Agregado con exito!")
}

func delProduct(u *Usuario) {
	u.Products = append(u.Products[:0])
	fmt.Println("Productos eliminados con Exito!")
}

func main() {
	u := Usuario{Nombre: "Anuel", Apellido: "Bad Bunny", Correo: "bad@gmail.com"}
	p := newProduct("Manzanas", 6500.0)
	p1 := newProduct("Pasta", 2500.0)

	addProduct(&u, &p, 15)
	addProduct(&u, &p1, 5)

	fmt.Println(u.Products)
	delProduct(&u)
	fmt.Println(u.Products)
}
