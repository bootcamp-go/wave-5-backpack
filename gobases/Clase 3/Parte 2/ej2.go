package main

import "fmt"

type Usuario struct {
	Nombre    string
	Apellido  string
	Correo    string
	Productos []*Producto
}

type Producto struct {
	Nombre   string
	Precio   float64
	Cantidad int
}

func NuevoProducto(nombre string, precio float64) *Producto {
	return &Producto{
		Nombre: nombre,
		Precio: precio,
	}
}

func AgregarProducto(u *Usuario, p *Producto, cantidad int) {
	p.Cantidad = cantidad
	u.Productos = append(u.Productos, p)
}

func BorrarProductos(u *Usuario) {
	u.Productos = []*Producto{}
}

func main() {
	user := &Usuario{
		Nombre:   "Pepe",
		Apellido: "Sierra",
		Correo:   "correo@real.com",
	}
	p1 := NuevoProducto("p1", 12)
	p2 := NuevoProducto("p2", 20)
	AgregarProducto(user, p1, 10)
	AgregarProducto(user, p2, 20)
	fmt.Println(user)
	BorrarProductos(user)
	fmt.Println(user)
}
