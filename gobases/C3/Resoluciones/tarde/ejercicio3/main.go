package main

import (
	"fmt"
)

/*Ejercicio 3 - Calcular Precio

Una empresa nacional se encarga de realizar venta de productos, servicios y mantenimiento.
Para ello requieren realizar un programa que se encargue de calcular el precio total de Productos,
Servicios y Mantenimientos. Debido a la fuerte demanda y para optimizar la velocidad requieren que
el cálculo de la sumatoria se realice en paralelo mediante 3 go routines.

Se requieren 3 estructuras:
	1. Productos: nombre, precio, cantidad.
	2. Servicios: nombre, precio, minutos trabajados.
	3. Mantenimiento: nombre, precio.

Se requieren 3 funciones:
	1. Sumar Productos: recibe un array de producto y devuelve el precio total (precio * cantidad).
	2. Sumar Servicios: recibe un array de servicio y devuelve el precio total (precio * media hora trabajada,
	si no llega a trabajar 30 minutos se le cobra como si hubiese trabajado media hora).
	3. Sumar Mantenimiento: recibe un array de mantenimiento y devuelve el precio total.

Los 3 se deben ejecutar concurrentemente y al final se debe mostrar por pantalla el monto final
(sumando el total de los 3).*/

type Producto struct {
	Nombre   string
	Precio   float64
	Cantidad uint64
}

type Servicio struct {
	Nombre           string
	Precio           float64
	MinutosTabajados uint64
}

type Mantenimiento struct {
	Nombre string
	Precio float64
}

func SumarServicios(servicios *[]Servicio, c chan float64) {
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

func SumarProducto(productos *[]Producto, c chan float64) {
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
	productos := []Producto{
		{Nombre: "Producto 1", Precio: 100, Cantidad: 20},
		{Nombre: "Producto 2", Precio: 200, Cantidad: 4},
		{Nombre: "Producto 3", Precio: 300, Cantidad: 1},
		{Nombre: "Producto 4", Precio: 200, Cantidad: 200},
	}

	servicios := []Servicio{
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
