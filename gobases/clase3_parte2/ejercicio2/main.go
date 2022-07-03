/*Una importante empresa de ventas web necesita agregar una funcionalidad para agregar productos a los usuarios.
 Para ello requieren que tanto los usuarios como los productos tengan la misma dirección de memoria
 en el main del programa como en las funciones.
Se necesitan las estructuras:
Usuario: Nombre, Apellido, Correo, Productos (array de productos).
Producto: Nombre, precio, cantidad.
Se requieren las funciones:
Nuevo producto: recibe nombre y precio, y retorna un producto.
Agregar producto: recibe usuario, producto y cantidad, no retorna nada, agrega el producto al usuario.
Borrar productos: recibe un usuario, borra los productos del usuario.
*/

package main

import "fmt"

func main() {
	prod1 := NuevoProducto("Tele", 500)

	user1 := Usuario{
		Nombre: "Yvo",
	}

	AgregarProducto(&user1, prod1, 3)
	fmt.Print(user1)

	BorrarProductos(&user1)
	fmt.Print(user1)
}

type Producto struct {
	Nombre   string
	Precio   float64
	Cantidad int
}

type Usuario struct {
	Nombre    string
	Apellido  string
	Correo    string
	Productos []*Producto
}

func NuevoProducto(Nombre string, Precio float64) *Producto {
	return &Producto{Nombre: Nombre, Precio: Precio}
}

func AgregarProducto(user *Usuario, prod *Producto, cant int) {
	user.Productos = append(user.Productos, prod)
}

func BorrarProductos(user *Usuario) {
	user.Productos = nil
}
