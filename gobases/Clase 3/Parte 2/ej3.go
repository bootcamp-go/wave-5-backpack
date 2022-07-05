package main

import (
	"fmt"
)

type Productos struct {
	nombre   string
	precio   float64
	cantidad int
}

type Servicios struct {
	nombre             string
	precio             float64
	minutos_trabajados int
}

type Mantenimientos struct {
	nombre string
	precio float64
}

func SumarProductos(listaProductos []*Productos, res chan float64) {
	var suma float64
	for _, pr := range listaProductos {
		suma += (pr.precio * float64(pr.cantidad))
	}
	res <- suma
}

func SumarServicios(listaServicios []*Servicios, res chan float64) {
	var suma float64
	for _, ser := range listaServicios {
		mediaHorasTrabajadas := ser.minutos_trabajados / 30
		suma += (ser.precio * float64(mediaHorasTrabajadas))
		if ser.minutos_trabajados%30 != 0 {
			suma += ser.precio
		}
	}
	res <- suma
}

func SumarMantenimientos(listaMantenimientos []*Mantenimientos, res chan float64) {
	var suma float64
	for _, man := range listaMantenimientos {
		suma += man.precio
	}
	res <- suma
}

func CalcularCostoTotal(listaProductos []*Productos, listaMantenimientos []*Mantenimientos, listaServicios []*Servicios, sumaTotal chan float64) {
	chProductos := make(chan float64)
	chMantenimientos := make(chan float64)
	chServicios := make(chan float64)

	go SumarMantenimientos(listaMantenimientos, chMantenimientos)
	go SumarProductos(listaProductos, chProductos)
	go SumarServicios(listaServicios, chServicios)

	sumaTotal <- (<-chServicios + <-chMantenimientos + <-chProductos)
}

func main() {
	listaCalculoProductos := []*Productos{
		{
			nombre:   "p1",
			cantidad: 10,
			precio:   10,
		},
	}

	listaCalculoMantenimientos := []*Mantenimientos{
		{
			nombre: "m1",
			precio: 10,
		},
	}

	listaCalculoServicios := []*Servicios{
		{
			nombre:             "m1",
			precio:             10,
			minutos_trabajados: 30,
		},
	}

	chResultadoSuma := make(chan float64)

	go CalcularCostoTotal(listaCalculoProductos, listaCalculoMantenimientos, listaCalculoServicios, chResultadoSuma)
	fmt.Println(<-chResultadoSuma)
}
