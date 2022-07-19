package main

import "fmt"

/*
Una empresa nacional se encarga de realizar venta de productos, servicios y mantenimiento.
Para ello requieren realizar un programa que se encargue de calcular el precio total de Productos, Servicios y Mantenimientos.
Debido a la fuerte demanda y para optimizar la velocidad requieren que el cálculo de la sumatoria se realice en paralelo mediante 3 go routines.

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

func main() {
	/* Slices de Actividades */
	var productos []Productos
	var servicios []Servicios
	var mantenimiento []Mantenimiento

	/* Detalle de Actividades */
	macbook := Productos{nombre: "MacBook", precio: 35000, cantidad: 2}
	acer := Productos{nombre: "Acer", precio: 10000, cantidad: 1}
	thinkpad := Productos{nombre: "Thinkpad", precio: 35000, cantidad: 1}

	instalacion := Servicios{nombre: "Instalacion", precio: 1000, minutos: 120}
	reparacion := Servicios{nombre: "Reparación", precio: 1200, minutos: 250}

	preventivo := Mantenimiento{nombre: "Mantto Preventivo", precio: 2000}
	correctivo := Mantenimiento{nombre: "Mantto Correctivo", precio: 3500}

	fmt.Println(productos, servicios, mantenimiento)

	productos = append(productos, macbook, acer, thinkpad)
	servicios = append(servicios, instalacion, reparacion)
	mantenimiento = append(mantenimiento, preventivo, correctivo)

	c1 := make(chan float64)
	c2 := make(chan float64)
	c3 := make(chan float64)

	go sumProductos(&productos, c1)
	go sumServicios(&servicios, c2)
	go sumMantenimiento(&mantenimiento, c3)

	t1 := <-c1
	t2 := <-c2
	t3 := <-c3

	fmt.Println("* Total final * $ ", t1+t2+t3)

}

func sumProductos(p *[]Productos, c chan float64) {
	var total float64
	for _, v := range *p {
		total += float64(v.precio * v.cantidad)
	}

	fmt.Println("Total productos", total)
	c <- total
	close(c)
}

func sumServicios(s *[]Servicios, c chan float64) {
	var total float64
	for _, v := range *s {
		total += float64(v.precio * (v.minutos / 30)) //ojo con el residuo
	}
	fmt.Println("Total servicios", total)
	c <- total
	close(c)
}

func sumMantenimiento(m *[]Mantenimiento, c chan float64) {
	var total float64
	for _, v := range *m {
		total += float64(v.precio)
	}

	fmt.Println("Total mantenimiento", total)
	c <- total
	close(c)
}
