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

func nuevoProducto(nombre *string, precio *float64) *Producto {
	return &Producto{
		Nombre: *nombre,
		Precio: *precio,
	}
}

func (u *Usuario) agregarProducto(producto *Producto, cantidad *int) {
	producto.Cantidad = *cantidad
	u.Productos = append(u.Productos, *producto)
}

func (u *Usuario) borrarProductos() {
	u.Productos = []Producto{}
}

func main() {

	var (
		nombre   string  = "Ajiaco"
		precio   float64 = 30000
		cantidad int     = 10
	)
	producto := nuevoProducto(&nombre, &precio)

	usuario := &Usuario{
		Nombre:   "Juan Sebastian",
		Apellido: "Salazar",
		Correo:   "juansebastian.salazar@mercadolibre.com.co",
	}

	usuario.agregarProducto(producto, &cantidad)

	fmt.Println("Usuario - ", usuario.Nombre, " ", usuario.Apellido)
	fmt.Println("Correo - ", usuario.Correo)

	for _, value := range usuario.Productos {
		fmt.Printf("Producto - %s, Precio - %.0f, Cantidad - %d \n", value.Nombre, value.Precio, value.Cantidad)
	}

	fmt.Println("Eliminando productos")
	usuario.borrarProductos()
}
