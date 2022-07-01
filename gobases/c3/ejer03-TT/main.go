package main

import "fmt"

// Una empresa nacional se encarga de realizar venta de productos, servicios y mantenimiento.
// Para ello requieren realizar un programa que se encargue de calcular el precio total de Productos, Servicios y Mantenimientos.
// Debido a la fuerte demanda y para optimizar la velocidad requieren que el cálculo de la sumatoria se realice en paralelo mediante 3 go routines.

// Se requieren 3 estructuras:
// Productos: nombre, precio, cantidad.
// Servicios: nombre, precio, minutos trabajados.
// Mantenimiento: nombre, precio.

// Se requieren 3 funciones:
// Sumar Productos: recibe un array de producto y devuelve el precio total (precio * cantidad).
// Sumar Servicios: recibe un array de servicio y devuelve el precio total (precio * media hora trabajada,
// si no llega a trabajar 30 minutos se le cobra como si hubiese trabajado media hora).
// Sumar Mantenimiento: recibe un array de mantenimiento y devuelve el precio total.

// Los 3 se deben ejecutar concurrentemente y al final se debe mostrar por pantalla el monto final (sumando el total de los 3).

type Productos struct {
	Nombre   string
	Precio   float64
	Cantidad int
}

type Servicios struct {
	Nombre        string
	Precio        float64
	MinTrabajados int
}

type Mantenimiento struct {
	Nombre string
	Precio float64
}

func SumarProductos(ap []Productos, c chan float64) {
	var precioTotal float64

	for _, prod := range ap {
		precioTotal += prod.Precio * float64(prod.Cantidad)
	}
	fmt.Println("Suma productos: ", precioTotal)
	c <- precioTotal

}

func SumarServicios(as []Servicios, c chan float64) {
	var precioTotal float64
	for _, serv := range as {
		precioTotal += serv.Precio * float64(serv.MinTrabajados)
	}
	fmt.Println("Suma servicios: ", precioTotal)
	c <- precioTotal
}

func SumarMantenimiento(am []Mantenimiento, c chan float64) {
	var precioTotal float64
	for _, mant := range am {
		precioTotal += mant.Precio
	}
	fmt.Println("Suma mantenimiento: ", precioTotal)
	c <- precioTotal
}

func main() {
	p1 := Productos{"prod1", 1, 2}
	p2 := Productos{"prod2", 3, 4}
	ap := []Productos{p1, p2}

	s1 := Servicios{"serv1", 5, 66}
	s2 := Servicios{"serv2", 8, 88}
	as := []Servicios{s1, s2}

	m1 := Mantenimiento{"mant1", 9}
	am := []Mantenimiento{m1}

	c := make(chan float64)

	go SumarProductos(ap, c)
	go SumarServicios(as, c)
	go SumarMantenimiento(am, c)

	var totalFinal float64

	for i := 0; i < 3; i++ {
		fmt.Println(i, "-Inicio")
		totalFinal += <-c
	}
	fmt.Println("Total a Pagar :", totalFinal)

}
