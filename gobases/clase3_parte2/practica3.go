package main

import (
	"fmt"

)

type Producto struct {
	Nombre string
	Precio float64
	Cantidad int
}

type Servicio struct {
	Nombre string
	Precio float64
	MinutosTrabajados int
}

type Mantenimiento struct {
	Nombre string
	Precio float64
}

func SumarProductos(productos *[]Producto, c chan float64) {
 	var precioTotal float64
	for _, value := range *productos {
		precioTotal += value.Precio * float64(value.Cantidad)
	}
 	fmt.Println(precioTotal)
	c <- precioTotal
	close(c)
}

func SumarServicios(servicios *[]Servicio, c chan float64) {
 	var precioTotal float64
	for _, value := range *servicios {
 		if value.MinutosTrabajados < 30 {
 			precioTotal += value.Precio * 30
		} else {
			precioTotal += float64(value.MinutosTrabajados) * value.Precio
		}
 	}
 	fmt.Println(precioTotal)
	 c <- precioTotal
	 close(c)
}

func SumarMantenimiento(mantenimientos *[]Mantenimiento, c chan float64) {
	var precioTotal float64
	for _, value := range *mantenimientos {
		precioTotal += value.Precio
	}
	fmt.Println(precioTotal)
	c <- precioTotal
	close(c)
}

func main()  {

	productos := []Producto {
		{"Llanta", 80000.00, 4},
		{"Exosto", 599000.00, 6},
	}

	servicios := []Servicio {
		{"Cambio de aceite", 39500.00, 30},
		{"Cambio de llanta", 18000.00, 120},
	}

	mantenimientos := []Mantenimiento {
		{"Mantenimiento general", 450000.00},
		{"Mantenimiento sencillo", 225000.00},
	}

	canal1 := make(chan float64)
	canal2 := make(chan float64)
	canal3 := make(chan float64)

	go SumarProductos(&productos, canal1)
	go SumarServicios(&servicios, canal2)
	go SumarMantenimiento(&mantenimientos, canal3)

	total1 := <-canal1
	total2 := <-canal2
	total3 := <-canal3

	fmt.Println("Monto final $ ", total1+total2+total3)}
