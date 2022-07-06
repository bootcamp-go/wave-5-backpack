// Ejercicio 2 - Ecommerce
// Una importante empresa de ventas web necesita agregar una funcionalidad para agregar productos a los usuarios. Para ello requieren que tanto los usuarios como los productos tengan la misma dirección de memoria en el main del programa como en las funciones.
// Se necesitan las estructuras:
// Usuario: Nombre, Apellido, Correo, Productos (array de productos).
// Producto: Nombre, precio, cantidad.
// Se requieren las funciones:
// Nuevo producto: recibe nombre y precio, y retorna un producto.
// Agregar producto: recibe usuario, producto y cantidad, no retorna nada, agrega el producto al usuario.
// Borrar productos: recibe un usuario, borra los productos del usuario.

package main

import "fmt"

type Producto struct {
	Nombre   string
	Precio   float64
	Cantidad int
}

func (u *Usuario) DelProd() {
	u.Productos = []Producto{}
}
func NewProd(nombre *string, precio *float64) *Producto {
	return &Producto{Nombre: *nombre, Precio: *precio}
}

func (u *Usuario) AddProd(producto *Producto, cantidad *int) {
	producto.Cantidad = *cantidad
	u.Productos = append(u.Productos, *producto)
}

type Usuario struct {
	Name    string
	LastName  string
	Email    string
	Productos []Producto
}

func main() {
	var (
		nombre   string  = "PAPAS FRITAS"
		precio   float64 = 8000
		cantidad int     = 20
	)

	prod := NewProd(&nombre, &precio)
	user := &Usuario{
		Name:   "pedro",
		LastName: "perez",
		Email:   "peperez@mail.com",
	}

	user.AddProd(prod, &cantidad)

	fmt.Println("Usuario: ", user.Name, user.LastName)
	fmt.Println("Correo: ", user.Email)
	for _, value := range user.Productos {
		fmt.Printf("prod %s, cantidad  %d, precio %.2f, total $%.2f\n", value.Nombre, value.Cantidad, value.Precio, (value.Precio * float64(value.Cantidad)))
	}

	user.DelProd()

	fmt.Println("Productos que quedan ", len(user.Productos))
}
