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
	var (
		nombre   string  = "Ketchup"
		precio   float64 = 2500
		cantidad int     = 100
	)

	producto := NewProducto(&nombre, &precio)
	usuario := &Usuario{
		Nombre:   "Juan",
		Apellido: "Camilo",
		Correo:   "test@gmail.com",
	}

	usuario.AgregarProducto(producto, &cantidad)

	fmt.Println("Usuario: ", usuario.Nombre, usuario.Apellido)
	fmt.Println("Correo:", usuario.Correo)
	for _, value := range usuario.Productos {
		fmt.Printf("Producto / %s / Cantidad %d / Precio %.2f / Total $%.2f\n", value.Nombre, value.Cantidad, value.Precio, (value.Precio * float64(value.Cantidad)))
	}

	fmt.Println()
	fmt.Println("Eliminando productos . . .")
	usuario.BorrarProducto()

	fmt.Println("Cantidad de productos después de la eliminación", len(usuario.Productos))
}

func NewProducto(nombre *string, precio *float64) *Producto {
	return &Producto{Nombre: *nombre, Precio: *precio}
}

func (u *Usuario) AgregarProducto(producto *Producto, cantidad *int) {
	producto.Cantidad = *cantidad
	u.Productos = append(u.Productos, *producto)
}

func (u *Usuario) BorrarProducto() {
	u.Productos = []Producto{}
}
