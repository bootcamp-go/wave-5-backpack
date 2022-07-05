package main

import "fmt"

type Productos struct {
	nombre   string
	precio   int
	cantidad int
}
type Servicios struct {
	nombre  string
	precio  int
	minutos int
}
type Mantenimiento struct {
	nombre string
	precio int
}

func sumarProductos(productos []Productos, c chan int) {
	suma := 0
	for _, valor := range productos {
		suma += valor.precio * valor.cantidad
	}
	c <- suma
	close(c)
}
func sumarServicios(servicios []Servicios, c chan int) {
	suma := 0
	for _, valor := range servicios {
		suma += valor.precio * valor.minutos
	}
	c <- suma
	close(c)
}
func sumarMantenimientos(mantenimiento []Mantenimiento, c chan int) {
	suma := 0
	for _, valor := range mantenimiento {
		suma += valor.precio
	}
	c <- suma
	close(c)
}

func main() {
	productos := []Productos{
		{nombre: "Pan", precio: 2000, cantidad: 30},
		{nombre: "Leche", precio: 2500, cantidad: 10},
		{nombre: "Huevos", precio: 6000, cantidad: 60},
	}

	servicios := []Servicios{
		{nombre: "Internet", precio: 70000, minutos: 360},
		{nombre: "Agua", precio: 30000, minutos: 200},
		{nombre: "Luz", precio: 50000, minutos: 300},
		{nombre: "Gas", precio: 2000, minutos: 10},
	}

	mantenimientos := []Mantenimiento{
		{nombre: "Mantenimiento Pc", precio: 500000},
		{nombre: "Mantenimiento Carro", precio: 1000000},
		{nombre: "Mantenimiento Luz", precio: 200000},
		{nombre: "Mantenimiento Estufa", precio: 30000},
	}
	c1 := make(chan int)
	c2 := make(chan int)
	c3 := make(chan int)

	go sumarProductos(productos, c1)
	go sumarServicios(servicios, c2)
	go sumarMantenimientos(mantenimientos, c3)

	t1 := <-c1
	t2 := <-c2
	t3 := <-c3

	total := t1 + t2 + t3

	fmt.Println("Total:", total)

}
