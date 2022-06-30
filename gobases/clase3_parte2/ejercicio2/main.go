package main

import "fmt"

type Usuario struct {
	Nombre, Apellido, Correo string
	Productos                []Producto
}

type Producto struct {
	Nombre   string
	Precio   float64
	Cantidad int
}

func NuevoProducto(nombre string, precio float64) *Producto {
	return &Producto{
		Nombre:   nombre,
		Precio:   precio,
		Cantidad: 1,
	}
}

func AgregarProducto(u *Usuario, p *Producto, cantidad int) {
	p.Cantidad = cantidad
	u.Productos = append(u.Productos, *p)
}

func BorrarProductos(u *Usuario) {
	u.Productos = []Producto{}
}

func main() {
	user := Usuario{
		Nombre:   "Cristobal",
		Apellido: "Monsalve",
		Correo:   "cmonsalve@gmail.com",
	}

	p1 := *NuevoProducto("silla", 50000)
	p2 := *NuevoProducto("mesa", 100000)
	p3 := *NuevoProducto("monitor", 300000)
	AgregarProducto(&user, &p1, 1)
	AgregarProducto(&user, &p2, 5)
	AgregarProducto(&user, &p3, 2)

	fmt.Println(user)

	BorrarProductos(&user)

	fmt.Println(user)
}
