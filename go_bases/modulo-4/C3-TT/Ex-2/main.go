package main

import "fmt"

type Usuario struct {
	nombre    string
	apellido  string
	email     string
	productos []Producto
}

type Producto struct {
	nombre   string
	precio   int
	cantidad int
}

func main() {
	/*
		Una importante empresa de ventas web necesita agregar una funcionalidad para agregar productos a los usuarios.
		Para ello requieren que tanto los usuarios como los productos tengan la misma dirección de memoria en el main del programa como en las funciones.

		Se necesitan las estructuras:
			+ Usuario: Nombre, Apellido, Correo, Productos (array de productos).
			+ Producto: Nombre, precio, cantidad.
		Se requieren las funciones:
			+ Nuevo producto: recibe nombre y precio, y retorna un producto.
			+ Agregar producto: recibe usuario, producto y cantidad, no retorna nada, agrega el producto al usuario.
			Borrar productos: recibe un usuario, borra los productos del usuario.

	*/

	usuario := Usuario{nombre: "Roko", apellido: "Doggy", email: "roko@dogmail.com"}

	itemCroquetas := nuevoProducto("DogChow", 700)

	agregarProducto(&usuario, &itemCroquetas, 2)
	fmt.Println("Carrito de", usuario.nombre, "🛍: ", "| productos:", usuario.productos[0].nombre, "| precio:", usuario.productos[0].precio, "| cant:", usuario.productos[0].cantidad)

	borrarProductos(&usuario)

	fmt.Println("Carrito de", usuario.nombre, "🛍: ", "| productos:", usuario.productos[0].nombre, "| precio:", usuario.productos[0].precio, "| cant:", usuario.productos[0].cantidad)

}

func nuevoProducto(n string, p int) Producto {
	product := Producto{nombre: n, precio: p}
	return product
}

func agregarProducto(u *Usuario, p *Producto, c int) {
	fmt.Println("Hola", *&u.nombre, "😎", ", has añadido el producto:", *p, "cantidad:", c)

	*&p.cantidad = c
	*&u.productos = append(*&u.productos, *p)
}
func borrarProductos(u *Usuario) {
	fmt.Println("Hola", *&u.nombre, "estamos vaciando tu carrito...", "🗑")

	for k := range u.productos {
		*&u.productos[k].nombre = ""
		*&u.productos[k].precio = 0
		*&u.productos[k].cantidad = 0
	}

}
