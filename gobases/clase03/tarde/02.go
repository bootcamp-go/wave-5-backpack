/*
Ejercicio 2 - Ecommerce
Una importante empresa de ventas web necesita agregar una funcionalidad para agregar productos a los usuarios. Para ello requieren que tanto los usuarios como los productos tengan la misma dirección de memoria en el main del programa como en las funciones.
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

type productos struct {
	Nombre string
	Precio float64
	Cantidad int
}

type usuarios struct {
	Nombre string
	Apellido string
	Correo string
	Productos []productos
}

func nuevoProducto (nombre *string, precio *float64) *productos {
	
	nuevoProd := &productos{Nombre: *nombre, Precio: *precio}
	return nuevoProd
}

func (u *usuarios) agregarProducto (p *productos, cantidad *int) {
	p.Cantidad = *cantidad

	u.Productos = append(u.Productos, *p)
}

func (u *usuarios) borrarProductos() {
	u.Productos = []productos{}
}

func main() {
	var (
		nombre   string  = "La virginia"
		precio   float64 = 1000
		cantidad int     = 10
	)

	producto := nuevoProducto(&nombre, &precio)
	usuario := &usuarios{
		Nombre:   "Marcelo",
		Apellido: "Gonzalez",
		Correo:   "test@gmail.com",
	}

	usuario.agregarProducto(producto, &cantidad)

	fmt.Println("Usuario - ", usuario.Nombre, usuario.Apellido)
	fmt.Println("Correo - ", usuario.Correo)
	for _, value := range usuario.Productos {
		fmt.Printf("Producto - %s - Cantidad %d - Precio %.2f - Total $%.2f\n", value.Nombre, value.Cantidad, value.Precio, (value.Precio * float64(value.Cantidad)))
	}

	fmt.Println()
	fmt.Println("Eliminando productos . . .")
	usuario.borrarProductos()

	fmt.Println("Cantidad de productos después de la eliminación", len(usuario.Productos))

}