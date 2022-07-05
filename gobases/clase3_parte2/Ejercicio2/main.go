package main

import "fmt"

type Usuario struct {
	Nombre   string
	Apellido string
	Correo   string
	Producto []Productos
}

type Productos struct {
	Nombre   string
	Precio   float64
	Cantidad int
}

func NuevoProdutco(nombre *string, precio *float64) *Productos {
	return &Productos{Nombre: *nombre, Precio: *precio}
}

func (usuario *Usuario) AggProducto(producto *Productos, cantidad *int) {
	producto.Cantidad = *cantidad
	usuario.Producto = append(usuario.Producto, *producto)
}

func (usuario *Usuario) BorrarProducto() {
	usuario.Producto = []Productos{}
}

func main() {
	var (
		nombre   string  = "Plato"
		precio   float64 = 30000.00
		cantidad int     = 3
	)

	producto := NuevoProdutco(&nombre, &precio)
	usuario := &Usuario{
		Nombre:   "Luz",
		Apellido: "Jimenez",
		Correo:   "luz@gmail.com",
	}

	usuario.AggProducto(producto, &cantidad)
	fmt.Println("Usuario - ", usuario.Nombre, usuario.Apellido)
	fmt.Println("Correo - ", usuario.Correo)
	for _, value := range usuario.Producto {
		fmt.Printf("Producto - %s - Cantidad %d - Precio %.2f - Total $%.2f\n", value.Nombre, value.Cantidad, value.Precio, (value.Precio * float64(value.Cantidad)))
	}

	fmt.Println()
	fmt.Println("Eliminando productos . . .")
	usuario.BorrarProducto()

	fmt.Println("Cantidad de productos", len(usuario.Producto))
}
