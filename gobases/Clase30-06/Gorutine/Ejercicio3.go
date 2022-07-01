/*
Una empresa nacional se encarga de realizar venta de productos, servicios y mantenimiento.
Para ello requieren realizar un programa que se encargue de calcular el precio total de Productos, Servicios y Mantenimientos. Debido a la fuerte demanda y para optimizar la velocidad
requieren que el cálculo de la sumatoria se realice en paralelo mediante 3 go routines.

Se requieren 3 estructuras:
Productos: nombre, precio, cantidad.
Servicios: nombre, precio, minutos trabajados.
Mantenimiento: nombre, precio.

Se requieren 3 funciones:
Sumar Productos: recibe un array de producto y devuelve el precio total (precio * cantidad).
Sumar Servicios: recibe un array de servicio y devuelve el precio total (precio * media hora trabajada, si no llega a trabajar 30 minutos se le cobra como si hubiese trabajado media hora).
Sumar Mantenimiento: recibe un array de mantenimiento y devuelve el precio total.

Los 3 se deben ejecutar concurrentemente y al final se debe mostrar por pantalla el monto final (sumando el total de los 3).
*/

package main

import "fmt"

type Productos struct {
	nombre   string
	precio   int
	cantidad int
}

type Servicios struct {
	nombre             string
	precio             int
	minuotosTrabajados int
}

type Mantenimiento struct {
	nombre string
	precio int
}

func sumarProductos(produ []Productos, c chan int) {
	acumulador := 0
	for _, value := range produ {
		acumulador = acumulador + value.cantidad*value.precio
	}
	fmt.Printf("El valor de los productos es: %v\n", acumulador)
	c <- acumulador
}

func sumarServicios(servi []Servicios, c chan int) {
	acumulador := 0
	for _, value := range servi {
		if value.minuotosTrabajados < 30 {
			acumulador = acumulador + value.minuotosTrabajados*value.precio
		} else {
			j := value.minuotosTrabajados / 30
			resto := value.minuotosTrabajados % 30
			if resto != 0 {
				acumulador = acumulador + j*value.precio + 30*value.precio
			} else {
				acumulador = acumulador + j*value.precio
			}
		}
	}
	fmt.Printf("El valor de servicios es: %v\n", acumulador)
	c <- acumulador
}

func sumarMantenimiento(mante []Mantenimiento, c chan int) {
	acumulador := 0
	for _, value := range mante {
		acumulador = acumulador + value.precio
	}
	fmt.Printf("El valor de mantenimiento es: %v\n", acumulador)
	c <- acumulador
}

func main() {

	jamon := Productos{"jamon", 2, 2}
	p1 := []Productos{jamon}

	luz := Servicios{"luz", 20, 25}
	s1 := []Servicios{luz}

	plomeria := Mantenimiento{"plomeria", 30}
	m1 := []Mantenimiento{plomeria}
	c := make(chan int)

	go sumarMantenimiento(m1, c)
	go sumarProductos(p1, c)
	go sumarServicios(s1, c)
	var acumulador int = 0
	for i := 0; i < 3; i++ {
		acumulador = acumulador + <-c
	}
	fmt.Printf("El total a pagar es: %v", acumulador)
}
