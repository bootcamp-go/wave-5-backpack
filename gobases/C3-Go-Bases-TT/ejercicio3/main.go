package main

import (
	"fmt"
	"strings"
)

// Ejercicio 3 - Calcular Precio

// Una empresa nacional se encarga de realizar venta de productos, servicios y mantenimiento.
// Para ello requieren realizar un programa que se encargue de calcular el precio total de Productos,
// Servicios y Mantenimientos. Debido a la fuerte demanda y para optimizar la velocidad
// requieren que el cálculo de la sumatoria se realice en paralelo mediante 3 go routines.

// Se requieren 3 estructuras:
//  - Productos: nombre, precio, cantidad.
//  - Servicios: nombre, precio, minutos trabajados.
//  - Mantenimiento: nombre, precio.

// Se requieren 3 funciones:
//  - Sumar Productos: recibe un array de producto y devuelve el precio total (precio * cantidad).
//  - Sumar Servicios: recibe un array de servicio y devuelve el precio total (precio * media hora trabajada,
//    si no llega a trabajar 30 minutos se le cobra como si hubiese trabajado media hora).
//  - Sumar Mantenimiento: recibe un array de mantenimiento y devuelve el precio total.

// Los 3 se deben ejecutar concurrentemente y al final se debe mostrar por pantalla el monto final
// (sumando el total de los 3).

// Estructura producto
type producto struct {
	nombre   string
	precio   float64
	cantidad int
}

// Estructura servicio
type servicio struct {
	nombre            string
	precio            float64
	minutosTrabajados int
}

// Estructura mantenimiento
type mantenimiento struct {
	nombre string
	precio float64
}

// Función para calcular el precio del total de productos
func totalProductos(p *[]producto, c chan float64) {
	total := 0.0
	for _, v := range *p {
		total += v.precio * float64(v.cantidad)
	}
	c <- total
}

// Función para calcular el precio del total de servicios
func totalServicios(s *[]servicio, c chan float64) {
	total := 0.0
	for _, v := range *s {
		if v.minutosTrabajados < 30 {
			total += v.precio
		} else {
			mediasHorasTrabajadas := float64(v.minutosTrabajados) / 30.0
			total += v.precio * mediasHorasTrabajadas
		}
	}
	c <- total
}

// Función para calcular el precio del total de mantenimientos
func totalMantenimientos(m *[]mantenimiento, c chan float64) {
	total := 0.0
	for _, v := range *m {
		total += v.precio
	}
	//  - Sumar Mantenimiento: recibe un array de mantenimiento y devuelve el precio total.
	c <- total
}

// Función para dar formato a moneda
func formatearMoneda(m float64) string {
	// Formateamos la cantidad a string
	money := fmt.Sprintf("%.2f", m)
	// Separamos la cantidad de su decimal
	moneyElements := strings.Split(money, ".")
	// Invertimos la cantidad
	moneyInverted := ""
	for _, v := range moneyElements[0] {
		moneyInverted = string(v) + moneyInverted
	}
	// Reinvertimos la cantidad y agregamos las comas
	moneyValid := ""
	for i, v := range moneyInverted {
		if (i+1)%3 == 0 && (i+1) != len(moneyInverted) {
			moneyValid = "," + string(v) + moneyValid
		} else {
			moneyValid = string(v) + moneyValid
		}
	}
	// Regresamos el resultado
	return "$" + moneyValid + "." + moneyElements[1]
}

func main() {
	fmt.Println("Ejercicio 3 - Calcular Precio")
	fmt.Println("")

	// Creamos los productos
	productos := []producto{{nombre: "p1", precio: 567.50, cantidad: 46}, {nombre: "p2", precio: 346.50, cantidad: 245}, {nombre: "p3", precio: 4823.75, cantidad: 867}, {nombre: "p4", precio: 356678.50, cantidad: 76}, {nombre: "p5", precio: 465678.50, cantidad: 678}, {nombre: "p6", precio: 4678.75, cantidad: 45}, {nombre: "p7", precio: 2574.50, cantidad: 73}, {nombre: "p8", precio: 4679.50, cantidad: 356}, {nombre: "p9", precio: 2457.75, cantidad: 865}, {nombre: "p10", precio: 3427.00, cantidad: 34}}

	// Creamos los servicios
	servicios := []servicio{{nombre: "s1", precio: 7632.00, minutosTrabajados: 23}, {nombre: "s2", precio: 6572.50, minutosTrabajados: 635}, {nombre: "s3", precio: 2582.00, minutosTrabajados: 63}, {nombre: "s4", precio: 4692.00, minutosTrabajados: 6783}, {nombre: "s5", precio: 5798.50, minutosTrabajados: 832}, {nombre: "s6", precio: 2464.00, minutosTrabajados: 33}, {nombre: "s7", precio: 6795.50, minutosTrabajados: 7823}, {nombre: "s8", precio: 4725.00, minutosTrabajados: 463}, {nombre: "s9", precio: 4682.50, minutosTrabajados: 26}, {nombre: "s10", precio: 8672.50, minutosTrabajados: 553}}

	// Creamos los mantenimientos
	mantenimientos := []mantenimiento{{nombre: "m1", precio: 2463}, {nombre: "m2", precio: 663}, {nombre: "m3", precio: 4363}, {nombre: "m4", precio: 56763}, {nombre: "m5", precio: 63}, {nombre: "m6", precio: 46}, {nombre: "m7", precio: 56763}, {nombre: "m8", precio: 466}, {nombre: "m9", precio: 2626}, {nombre: "m10", precio: 457}}

	// Creamos los canales para cada función
	channelProductos := make(chan float64)
	channelServicios := make(chan float64)
	channelMantenimientos := make(chan float64)

	// Llamamos a las funciones de cálculo
	go totalProductos(&productos, channelProductos)
	go totalServicios(&servicios, channelServicios)
	go totalMantenimientos(&mantenimientos, channelMantenimientos)

	// Obtenemos los resultados de las funciones por medio de los canales
	totalProductos := <-channelProductos
	totalServicios := <-channelServicios
	totalMantenimientos := <-channelMantenimientos

	// Mostramos el precio total de los bienes y servicios de la empresa
	fmt.Printf("El total de los bienes y servicios de la empresa es de: %s\n", formatearMoneda(totalServicios+totalProductos+totalMantenimientos))
}
