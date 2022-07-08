package main

import "fmt"


type Usuario struct {
	Nombre string
	Apellido string
	Correo string
	Productos []Producto
}
type Producto struct {
	Nombre string
	Precio int
	Cantidad int
}

func NuevoProducto(nombre string, precio int) *Producto {
	return &Producto{
		Nombre: nombre,
		Precio: precio,
	}
}

func AgregarProducto(usuario *Usuario, producto *Producto, cantidad int) {
	producto.Cantidad = cantidad
	usuario.Productos = append(usuario.Productos, *producto)
}




func main(){
	u := &Usuario{"nahu", "monserrat", "nahuel123", make([]Producto, 0)}
	p := &Producto{"arroz", 10, 150}
	p2 := &Producto{"Calcio", 30, 10}
	
	AgregarProducto(u, p, 200)
	AgregarProducto(u, p2, 500)
	fmt.Println(u.Productos[1].Nombre)
}

/*
Una importante empresa de ventas web necesita agregar una funcionalidad para agregar productos a los usuarios. Para ello requieren que tanto los usuarios como los productos tengan la misma dirección de memoria en el main del programa como en las funciones.
Se necesitan las estructuras:
Usuario: Nombre, Apellido, Correo, Productos (array de productos).
Producto: Nombre, precio, cantidad.
Se requieren las funciones:
Nuevo producto: recibe nombre y precio, y retorna un producto.
Agregar producto: recibe usuario, producto y cantidad, no retorna nada, agrega el producto al usuario.
Borrar productos: recibe un usuario, borra los productos del usuario.

*/