package main

import "fmt"

// Ejercicio 2 - Ecommerce

// Una importante empresa de ventas web necesita agregar una funcionalidad para agregar productos
// a los usuarios. Para ello requieren que tanto los usuarios como los productos tengan la misma
// dirección de memoria en el main del programa como en las funciones.

// Se necesitan las estructuras:
//   - Usuario: Nombre, Apellido, Correo, Productos (array de productos).
//   - Producto: Nombre, precio, cantidad.
// Se requieren las funciones:
//   - Nuevo producto: recibe nombre y precio, y retorna un producto.
//   - Agregar producto: recibe usuario, producto y cantidad, no retorna nada, agrega el producto al usuario.
//   - Borrar productos: recibe un usuario, borra los productos del usuario.

// Estructura producto
type producto struct {
	nombre   string
	precio   float64
	cantidad int
}

// Estructura usuario
type usuario struct {
	nombre    string
	apellido  string
	correo    string
	productos []producto
}

// Función para crear nuevo producto
func nuevoProducto(n string, p float64) *producto {
	prod := producto{nombre: n, precio: p, cantidad: 0}
	fmt.Println(prod)
	return &prod
}

// Función para agregar productos a un usuario
func agregarProducto(u *usuario, p *producto, c int) {
	u.productos = append(u.productos, *p)
}

// Función para borrar los productos a un usuario
func borrarProductosUsuario(u *usuario) {
	u.productos = nil
}

func main() {
	fmt.Println("Ejercicio 2 - Ecommerce")
	fmt.Println("")

	// Obtenemos el producto creado
	prod := nuevoProducto("vasos", 150.0)
	fmt.Println(*prod)

	// Creamos un usuario
	user := usuario{nombre: "Nobuo", apellido: "Uematsu", correo: "nobuo.uematsu@yahoo.com.jp"}
	fmt.Println(user)

	// Agregamos un producto a un usuario
	agregarProducto(&user, prod, 10)
	fmt.Println(user)

	// Borramos los productos a un usuario
	borrarProductosUsuario(&user)
	fmt.Println(user)
}
