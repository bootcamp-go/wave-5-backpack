package main

import (
	"fmt"
)

//Definición de estructuras

type Producto struct {
	Nombre   string
	Precio   float64
	Cantidad float64
}

type Servicios struct {
	Nombre            string
	Precio            float64
	MinutosTrabajados float64
}

type Mantenimiento struct {
	Nombre string
	Precio float64
}

// Función Sumar Producto

func sumarProductos(c chan float64, productos ...Producto) {
	var subtotal float64
	for _, producto := range productos {
		subtotal += producto.Precio * producto.Cantidad
	}
	c <- subtotal
}

// Sumar Servicios

func sumarServicios(c chan float64, servicios ...Servicios) {
	var subtotal float64
	for _, servicio := range servicios {
		if servicio.MinutosTrabajados < 30 {
			subtotal += servicio.Precio * 30.0
		} else {
			subtotal += servicio.Precio * servicio.MinutosTrabajados
		}
	}
	c <- subtotal
}

// Sumar Matenmientos

func sumarMantenimiento(c chan float64, matenmientos ...Mantenimiento) {
	var subtotal float64
	for _, mantenimiento := range matenmientos {
		subtotal += mantenimiento.Precio
	}
	c <- subtotal
}

func main() {

	// Ejecución Canal de Productos

	producto1 := Producto{Nombre: "CELULAR", Precio: 200.99, Cantidad: 5.0}
	producto2 := Producto{Nombre: "LAPTOP", Precio: 800.99, Cantidad: 2.0}
	productosAll := []Producto{producto1, producto2}

	cSP := make(chan float64)
	go sumarProductos(cSP, productosAll...)
	subtotal1 := <-cSP
	fmt.Printf("El total de los productos es: %0.2f \n", subtotal1)

	// Ejecucición Canal de Servicios

	servicio1 := Servicios{Nombre: "VERIFICACION ESTR.", Precio: 5.99, MinutosTrabajados: 35.0}
	servicio2 := Servicios{Nombre: "DISEÑO", Precio: 8.99, MinutosTrabajados: 25.0}
	serviciosAll := []Servicios{servicio1, servicio2}

	cSS := make(chan float64)
	go sumarServicios(cSS, serviciosAll...)
	subtotal2 := <-cSS
	fmt.Printf("El total de los servicios es: %0.2f \n", subtotal2)

	// Ejecucición Canal de Servicios

	manteminiento1 := Mantenimiento{Nombre: "IMPRESORAS", Precio: 100.99}
	manteminiento2 := Mantenimiento{Nombre: "AIRE ACOND.", Precio: 150.99}
	mantenimientosAll := []Mantenimiento{manteminiento1, manteminiento2}

	cSM := make(chan float64)
	go sumarMantenimiento(cSM, mantenimientosAll...)
	subtotal3 := <-cSM
	fmt.Printf("El total de los mantenimientos es: %0.2f \n", subtotal3)

	// Suma de total

	total := subtotal1 + subtotal2 + subtotal3
	fmt.Println("El total es: ", total)

	fmt.Println("Ejecución Terminada")

}
