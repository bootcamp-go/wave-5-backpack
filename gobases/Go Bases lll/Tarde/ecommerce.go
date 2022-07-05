package main

import "fmt"

type producto struct {
	nombre   string
	precio   float64
	cantidad int
}

type usuario struct {
	nombre    string
	apellido  string
	correo    string
	Productos []producto
}

func nuevo(id string, prec float64) producto {
	productoNuevo := producto{
		nombre: id,
		precio: prec,
	}
	return productoNuevo
}

func agregar(usuario *usuario, producto *producto, cant int) {
	(*producto).cantidad = cant
	((*usuario).Productos) = append((*usuario).Productos, *producto)
}

func eliminar(user *usuario) {
	var lista []producto
	(*user).Productos = lista
}

func main() {

	usuario := usuario{
		nombre:   "Pepito",
		apellido: "Perez",
		correo:   "pepito@gmail.com",
	}
	p1 := nuevo("Xbox", 10)
	p2 := nuevo("PS5", 20)
	agregar(&usuario, &p1, 1)
	agregar(&usuario, &p2, 2)
	fmt.Println(usuario)
	eliminar(&usuario)
	fmt.Println(usuario)
}
