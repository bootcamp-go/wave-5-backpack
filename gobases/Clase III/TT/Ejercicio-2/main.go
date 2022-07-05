package main

import "fmt"

//Definición de estructura

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

// Función Nuevo Producto

func nuevoProducto(nombre *string, precio *float64) *Producto {

	newProducto := Producto{
		Nombre: *nombre,
		Precio: *precio,
	}
	return &newProducto
}

// Agregar Producto

func (u *Usuario) addProduct(productoNuevo *Producto, cantidad *int) {

	productoNuevo.Cantidad = *cantidad
	u.Productos = append(u.Productos, *productoNuevo)

}

// Borrar productos

func (u *Usuario) deleteProducts() {

	u.Productos = []Producto{}

}

func main() {

	var (
		producto string  = "celular"
		precio   float64 = 200.99
		cantidad int     = 5
	)

	usuario := Usuario{
		Nombre:   "José",
		Apellido: "Berríos",
		Correo:   "contacto@ecommerce.com",
	}

	nuevoProducto := nuevoProducto(&producto, &precio)

	usuario.addProduct(nuevoProducto, &cantidad)
	fmt.Printf("Usuario: %s %s - %s \n", usuario.Nombre, usuario.Apellido, usuario.Correo)
	fmt.Printf("Productos Usuario: %v \n", usuario.Productos)

	usuario.deleteProducts()
	fmt.Printf("Usuario: %s %s - %s \n", usuario.Nombre, usuario.Apellido, usuario.Correo)
	fmt.Printf("Productos Usuario: %v \n", usuario.Productos)

}
