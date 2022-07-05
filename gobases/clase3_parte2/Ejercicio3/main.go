package main

import (
	"fmt"
)

type Productos struct {
	Nombre   string
	Precio   float64
	Cantidad int
}

type Servicios struct {
	Nombre           string
	Precio           float64
	MinutosTabajados int
}

type Mantenimiento struct {
	Nombre string
	Precio float64
}

func SumarServicios(servicios *[]Servicios, c chan float64) {
	var total float64
	for _, value := range *servicios {
		if value.MinutosTabajados < 30 {
			total += value.Precio * 30
		} else {
			total += float64(value.MinutosTabajados) * value.Precio
		}
	}

	fmt.Println("Total servicios $", total)
	c <- total
	close(c)
}

func SumarProducto(productos *[]Productos, c chan float64) {
	var total float64
	for _, value := range *productos {
		total += value.Precio * float64(value.Cantidad)
	}

	fmt.Println("Total productos $", total)
	c <- total
	close(c)
}
func SumarMantenimiento(mantenimientos *[]Mantenimiento, c chan float64) {
	var total float64
	for _, value := range *mantenimientos {
		total += value.Precio
	}

	fmt.Println("Total mantenimiento $", total)
	c <- total
	close(c)
}

func main() {
	productos := []Productos{
		{Nombre: "Producto 1", Precio: 100, Cantidad: 20},
		{Nombre: "Producto 2", Precio: 200, Cantidad: 4},
		{Nombre: "Producto 3", Precio: 300, Cantidad: 1},
		{Nombre: "Producto 4", Precio: 200, Cantidad: 200},
	}

	servicios := []Servicios{
		{Nombre: "Programación", Precio: 200, MinutosTabajados: 480},
		{Nombre: "Paqueteria", Precio: 10234, MinutosTabajados: 30},
		{Nombre: "Encomienda", Precio: 300, MinutosTabajados: 15},
		{Nombre: "Limpieza", Precio: 100, MinutosTabajados: 22},
	}

	mantenimientos := []Mantenimiento{
		{Nombre: "Mantenimiento 1", Precio: 400},
		{Nombre: "Mantenimiento 2", Precio: 2500},
		{Nombre: "Mantenimiento 3", Precio: 4100},
		{Nombre: "Mantenimiento 4", Precio: 4100},
	}

	// Usando tres canales
	c1 := make(chan float64)
	c2 := make(chan float64)
	c3 := make(chan float64)

	go SumarProducto(&productos, c1)
	go SumarServicios(&servicios, c2)
	go SumarMantenimiento(&mantenimientos, c3)

	t1 := <-c1
	t2 := <-c2
	t3 := <-c3

	fmt.Println("* Total final * $ ", t1+t2+t3)
}
