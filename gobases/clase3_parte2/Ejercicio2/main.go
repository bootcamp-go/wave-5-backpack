package main

import "fmt"

type Usuario struct {
	Nombre, Apellido, Correo string
	productos []Producto
}

type Producto struct{
	Nombre string
	precio float64
	cantidad int
}

func (p *Usuario) nuevoProducto(nombre string, precio float64, cantidad int) Producto{
	return Producto{nombre, precio, cantidad}
}

func (u *Usuario) productoUsuario(p Producto){
	u.productos = append(u.productos, p)
}

func (u *Usuario) eliminalProductos(){
	u.productos = append(u.productos[:0], u.productos[:0]...)
}

func main(){
	var user Usuario
	user.Nombre = "Daniela"
	user.Apellido = "Bedoya"
	user.Correo = "email@gmail.com"
	user.productoUsuario(user.nuevoProducto("papa",5000.0,3))
	user.productoUsuario(user.nuevoProducto("piña",2000.0,3))
	user.productoUsuario(user.nuevoProducto("yuca", 3000.0, 2))

	var user2 Usuario
	user2.Nombre = "Damaris"
	user2.Apellido = "Poveda"
	user2.Correo = "email@gmail.com"
	user2.productoUsuario(user2.nuevoProducto("yupi",5000.0,3))
	user2.productoUsuario(user2.nuevoProducto("arandanos",2000.0,3))
	user2.productoUsuario(user2.nuevoProducto("papitas", 3000.0, 2))

	

	fmt.Printf("Nombre completo: %s %s \nCorreo: %s\nProductos del cliente: %v\n", user.Nombre, user.Apellido, user.Correo,user.productos)
	fmt.Println("==============================================================================")
	fmt.Printf("Nombre completo: %s %s \nCorreo: %s\nProductos del cliente: %v\n", user2.Nombre, user2.Apellido, user2.Correo,user2.productos)

	user2.eliminalProductos()
	fmt.Println("==============================================================================")
	fmt.Printf("Nombre completo: %s %s \nCorreo: %s\nProductos del cliente: %v\n", user2.Nombre, user2.Apellido, user2.Correo,user2.productos)


}