package main

import (
	"fmt"
	"math"
)

type Productos struct {
	Nombre   string
	Precio   float64
	Cantidad int
}

type Servicios struct {
	Nombre  string
	Precio  float64
	Minutos int
}

type Mantenimiento struct {
	Nombre string
	Precio float64
}

func sumarProductos(p *[]Productos, c chan float64) {
	total := 0.0
	for _, products := range *p {
		total += products.Precio * float64(products.Cantidad)
	}
	c <- total
}

func sumarServicios(s *[]Servicios, c chan float64) {
	total := 0.0
	for _, services := range *s {
		media := 0.0
		if services.Minutos < 30 {
			media += 1
		} else {
			//Si no es menor a 30, conocemos
			//cuantas medias horas hay
			media += math.Floor(float64(services.Minutos / 30))
		}
		total += services.Precio * media
	}
	c <- total
}

func sumarMantenimiento(m *[]Mantenimiento, c chan float64) {
	total := 0.0
	for _, mant := range *m {
		total += mant.Precio
	}
	c <- total
}

func main() {
	productos := []Productos{
		{Nombre: "TV", Precio: 600000, Cantidad: 2},
		{Nombre: "Control remoto", Precio: 2000, Cantidad: 2},
		{Nombre: "PC", Precio: 3000000, Cantidad: 4},
	}

	servicios := []Servicios{
		{Nombre: "Transporte", Precio: 100000, Minutos: 15},
		{Nombre: "IT", Precio: 2000000, Minutos: 60},
		{Nombre: "Limpieza", Precio: 40000, Minutos: 30},
	}

	mantenimiento := []Mantenimiento{
		{Nombre: "PCs", Precio: 200000},
		{Nombre: "Electrodomésticos", Precio: 240000},
		{Nombre: "Maquinas", Precio: 21000},
	}

	totalProducto, totalServicio, totalMantenimiento := make(chan float64), make(chan float64), make(chan float64)

	go sumarProductos(&productos, totalProducto)
	go sumarServicios(&servicios, totalServicio)
	go sumarMantenimiento(&mantenimiento, totalMantenimiento)

	resultProduct := <-totalProducto
	resultService := <-totalServicio
	resultManteinance := <-totalMantenimiento

	fmt.Println("Suma Productos:", resultProduct)
	fmt.Println("Suma Servicios:", resultService)
	fmt.Println("Suma Mantenimientos:", resultManteinance)
	fmt.Println("Total sumas:", resultProduct+resultService+resultManteinance)
}
