package main

import "fmt"

// Una importante empresa de ventas web necesita agregar una funcionalidad para agregar productos a los usuarios.
// Para ello requieren que tanto los usuarios como los productos tengan la misma dirección de memoria
// en el main del programa como en las funciones.
// Se necesitan las estructuras:
// Usuario: Nombre, Apellido, Correo, Productos (array de productos).
// Producto: Nombre, precio, cantidad.
// Se requieren las funciones:
// Nuevo producto: recibe nombre y precio, y retorna un producto.
// Agregar producto: recibe usuario, producto y cantidad, no retorna nada, agrega el producto al usuario.
// Borrar productos: recibe un usuario, borra los productos del usuario.

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

func nuevoProducto(nombre string, precio float64) *Producto {

	return &Producto{nombre, precio, 1}
}

func agregarProducto(u *Usuario, p *Producto, cantidad int) {
	p.Cantidad = cantidad
	u.Productos = append(u.Productos, *p)
}

func borrarProducto(u *Usuario) {
	u.Productos = []Producto{}
}

func main() {
	user := Usuario{
		Nombre:   "Juan",
		Apellido: "Perez",
		Correo:   "juan@perez.com",
	}

	p1 := *nuevoProducto("silla", 50000)
	p2 := *nuevoProducto("mesa", 100000)
	p3 := *nuevoProducto("monitor", 300000)
	agregarProducto(&user, &p1, 1)
	agregarProducto(&user, &p2, 5)
	agregarProducto(&user, &p3, 2)

	fmt.Println(user)

	borrarProducto(&user)

	fmt.Println(user)
}
