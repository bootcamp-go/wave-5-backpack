package main

import "fmt"

//Estructuras
type Usuario struct {
	Nombre    string
	Apellido  string
	Correo    string
	Productos []*Producto
}

type Producto struct {
	Nombre   string
	Precio   float64
	Cantidad int
}

//Funciones requeridas
func nuevoProducto(nombre string, precio float64) *Producto {
	product := Producto{Nombre: nombre, Precio: precio, Cantidad: 1}
	return &product
}

func agregarProducto(usuario *Usuario, producto *Producto, cantidad int) {
	producto.Cantidad = cantidad
	usuario.Productos = append(usuario.Productos, producto)
}

func borrarProductos(usuario *Usuario) {
	usuario.Productos = []*Producto{}
}

func imprimir(prod Producto) {
	fmt.Println("Nombre:", prod.Nombre, "\nPrecio:", prod.Precio, "\nCantidad:", prod.Cantidad)
}

func main() {
	productExample := Producto{
		Nombre:   "Lámpara",
		Precio:   230000,
		Cantidad: 6,
	}

	//Probando funciones
	product := nuevoProducto("MacBook Pro", 7490000)

	products := []*Producto{product}

	user := Usuario{
		Nombre:    "Luz",
		Apellido:  "Lucumí Hernández",
		Correo:    "luz@hotmail.com",
		Productos: products,
	}

	fmt.Println("----------\nProducto nuevo\n----------")
	fmt.Println("Nombre:", nuevoProducto("Escritorio", 400000).Nombre)
	fmt.Println("Precio:", nuevoProducto("Escritorio", 400000).Precio)

	//Producto
	fmt.Println("\nNUEVO PRODUCTO")
	imprimir(productExample)

	//Agregando profuctos al usuario
	agregarProducto(&user, &productExample, 6)
	fmt.Println("\nProductos del usuario", user.Productos)

	//Borrando productos del usuario
	borrarProductos(&user)
	fmt.Println("Productos eliminados del usuario", user.Productos)
}
