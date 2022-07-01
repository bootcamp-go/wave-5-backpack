package main

type Usuario struct {
	Nombre    string
	Apellido  string
	Edad      int
	Correo    string
	Productos []Producto
}

type Producto struct {
	Nombre   string
	Precio   int
	Cantidad int
}

func agregarProducto(user *Usuario, producto *Producto) {
	user.Productos = append(user.Productos, *producto)
}

func borrarProductos(user *Usuario) {
	user.Productos = nil
}

func nuevoProducto(nombre string, precio int) Producto {
	return Producto{Nombre: nombre, Precio: precio, Cantidad: 1}
}

func main() {
	user1 := Usuario{
		Nombre:   "pepe",
		Apellido: "user",
		Edad:     46,
		Correo:   "dummy",
	}
	producto := nuevoProducto("lechuga", 10)
	agregarProducto(&user1, &producto)

	println(user1.Nombre, user1.Apellido, user1.Edad, user1.Correo, user1.Productos[0].Nombre)
}
