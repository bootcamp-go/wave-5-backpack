package main

import "fmt"

type producto struct{
	Nombre string
	precio float64
	cantidad int
}

func NuevoProducto(id string, prec float64) producto{
	prod:=producto{
		Nombre: id,
		precio: prec,
	}
	return prod
}

func AgregarProducto(user *usuario,prod *producto, cant int){
	(*prod).cantidad=cant
	((*user).Productos)=append((*user).Productos,*prod)
}

func BorrarProductos(user *usuario){
	var lista []producto
	(*user).Productos=lista
}

type usuario struct{
	Nombre string
	Apellido string
	correo string
	Productos[]producto
}

func main(){
	user:=usuario{
	Nombre:"Diego",
	Apellido:"Jota",
	correo: "dj@gmail.com",
	}	

	fmt.Println(user)

	prod1:=NuevoProducto("Triciclo",3500)
	prod2:=NuevoProducto("Bicicleta",5000)
	AgregarProducto(&user,&prod1,5)
	AgregarProducto(&user,&prod2,10)
	fmt.Println(user)
	BorrarProductos(&user)
	fmt.Println(user)
}


/*
Ejercicio 2 - Ecommerce
Una importante empresa de ventas web necesita agregar una funcionalidad para agregar productos a los usuarios. 
Para ello requieren que tanto los usuarios como los productos tengan la misma dirección de memoria en el main del 
programa como en las funciones.
Se necesitan las estructuras:
Usuario: Nombre, Apellido, Correo, Productos (array de productos).
Producto: Nombre, precio, cantidad.
Se requieren las funciones:
Nuevo producto: recibe nombre y precio, y retorna un producto.
Agregar producto: recibe usuario, producto y cantidad, no retorna nada, agrega el producto al usuario.
Borrar productos: recibe un usuario, borra los productos del usuario.

*/

