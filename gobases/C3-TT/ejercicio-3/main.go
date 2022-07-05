package main

import "fmt"

type Producto struct {
	nombre   string
	precio   int
	cantidad int
}

type Servicio struct {
	nombre         string
	precio         int
	min_trabajados int
}

type Mantenimiento struct {
	nombre string
	precio int
}

func sumarProductos(productos []Producto, c chan int) {
	var total int

	for _, val := range productos {
		total += val.precio * val.cantidad
	}

	fmt.Printf("Total productos: %d\n", total)
	c <- total
}

func sumarServicios(servicios []Servicio, c chan int) {
	var total int

	for _, val := range servicios {
		if val.min_trabajados < 30 {
			total += val.precio
		} else {
			total += val.precio * (val.min_trabajados / 30)
		}
	}

	fmt.Printf("Total servicios: %d\n", total)
	c <- total
}

func sumarMantenimiento(mantenimientos []Mantenimiento, c chan int) {
	var total int

	for _, val := range mantenimientos {
		total += val.precio
	}

	fmt.Printf("Total mantenimientos: %d\n", total)
	c <- total
}

func main() {
	c1 := make(chan int)
	c2 := make(chan int)
	c3 := make(chan int)

	go sumarProductos([]Producto{
		{nombre: "Arroz", cantidad: 2, precio: 230},
		{nombre: "Fideos", cantidad: 5, precio: 332},
	}, c1)

	go sumarServicios([]Servicio{
		{nombre: "Transporte", precio: 4300, min_trabajados: 120},
		{nombre: "Proceso", precio: 5433, min_trabajados: 23},
	}, c2)

	go sumarMantenimiento([]Mantenimiento{
		{nombre: "Mantenimiento", precio: 1234},
		{nombre: "Reparacion", precio: 33441},
	}, c3)

	var total int
	total += <-c1
	total += <-c2
	total += <-c3

	fmt.Printf("Total: %d\n", total)

}
