package main

import "fmt"

type Usuario struct {
	nombre    string
	apellido  string
	correo    string
	productos []Producto
}

type Producto struct {
	nombre   string
	precio   float64
	cantidad int
}

func nuevoProducto(nombre *string, precio *float64) *Producto {
	return &Producto{nombre: *nombre, precio: *precio}
}

func (u *Usuario) agregarProducto(prod *Producto, cant *int) {
	prod.cantidad = *cant
	u.productos = append(u.productos, *prod)
}

func (u *Usuario) borrarProductos() {
	u.productos = []Producto{}
}

func main() {
	var (
		nombre string  = "HyperX Cloud II Wireless"
		precio float64 = 17300.00
		cant   int     = 1
	)
	prod := nuevoProducto(&nombre, &precio)

	user := Usuario{
		nombre:   "Matias",
		apellido: "Vince",
		correo:   "matiasvince9@gmail.com",
	}

	user.agregarProducto(prod, &cant)

	fmt.Println(user)

	user.borrarProductos()

	fmt.Println(user)
}
