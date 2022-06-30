package main

import (
	"fmt"
	"math"
)

type Producto struct {
	Nombre           string
	Precio, Cantidad int
}

type Servicio struct {
	Nombre          string
	Precio, Minutos int
}

type Mantenimiento struct {
	Nombre string
	Precio int
}

func main() {
	c := make(chan int)
	runConcurrence(c)
	total := 0
	for i := 0; i < 3; i++ {
		valor := <-c
		total += valor
		fmt.Printf("Valor %d: %d\n", i+1, valor)
	}
	fmt.Printf("El total es %d\n", total)
}

func SumarProductos(productos []Producto, c chan int) {
	precioTotal := 0
	for _, producto := range productos {
		precioTotal += producto.Precio * producto.Cantidad
	}
	c <- precioTotal
}

func SumarServicios(servicios []Servicio, c chan int) {
	precioTotal := 0
	for _, servicio := range servicios {
		mediaHora := math.Ceil(float64(servicio.Minutos) / 30)
		precioTotal += servicio.Precio * int(mediaHora)
	}
	c <- precioTotal
}

func SumarMantenimientos(mantenimientos []Mantenimiento, c chan int) {
	precioTotal := 0
	for _, mantenimiento := range mantenimientos {
		precioTotal += mantenimiento.Precio
	}
	c <- precioTotal
}

func runConcurrence(c chan int) {

	producto1 := Producto{"Papas", 2000, 5}
	producto2 := Producto{"Bebida", 5000, 2}
	producto3 := Producto{"Galleta", 1000, 3}
	productos := []Producto{producto1, producto2, producto3}

	servicio1 := Servicio{"Limpieza", 5000, 120}
	servicio2 := Servicio{"Programacion", 50000, 35}
	servicio3 := Servicio{"Planeacion", 20000, 68}
	servicios := []Servicio{servicio1, servicio2, servicio3}

	mantenimiento1 := Mantenimiento{"PC", 20000}
	mantenimiento2 := Mantenimiento{"Teclado", 15000}
	mantenimiento3 := Mantenimiento{"Televisor", 10000}
	mantenimientos := []Mantenimiento{mantenimiento1, mantenimiento2, mantenimiento3}

	go SumarProductos(productos, c)
	go SumarServicios(servicios, c)
	go SumarMantenimientos(mantenimientos, c)
}
