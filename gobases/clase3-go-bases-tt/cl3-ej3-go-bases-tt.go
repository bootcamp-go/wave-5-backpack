/*---------------------------------------------------------------------------------*

     Assignment:	Ejercicio #3:  Calcular Precio
         Author:	Israel Fabela
	   Language:	go1.18.3 darwin/arm64
		  Topic:	Go Bases

	Description:
		It requires 3 structures:
			- Products: name, price, quantity.
			- Services: name, price, minutes worked.
			- Maintenance: name, price.
		It requires 3 functions:
			- Sum Products: receives an array of product and returns
			  the total price (price * quantity).
			- Sum Services: receives an array of service and returns
			  the total price (price * half hour worked, if it is less
			  than hour worked, if he does not work 30 minutes he is charged as
			  if he had worked half an hour). half an hour).
			- Add Maintenance: receives a maintenance array and returns
			  the total price.
		The 3 must be executed concurrently and at the end the final amount
		must be displayed on the screen (adding the total of the the final
		amount).

	© Mercado Libre - IT Bootcamp 2022

----------------------------------------------------------------------------------*/

/* Ejercicio 3 - Calcular Precio */

package main

import (
	"fmt"
	"math"
)

type Productos struct {
	nombre   string
	precio   float64
	cantidad int
}

type Servicios struct {
	nombre  string
	precio  float64
	minutos int
}
type Mantenimiento struct {
	nombre string
	precio float64
}

func sumarProductos(lista ...Productos) float64 {
	var total float64
	for _, prod := range lista {
		total += prod.precio * float64(prod.cantidad)
	}

	return total
}

func sumarServicios(lista ...Servicios) float64 {
	var total float64
	for _, serv := range lista {
		if serv.minutos <= 30 {
			total += serv.precio
		}
		total += serv.precio * math.Round(float64(serv.minutos)/30)
	}
	return total
}

func sumarMantenimiento(lista ...Mantenimiento) float64 {
	var total float64
	for _, mant := range lista {
		total += mant.precio
	}
	return total
}

func main() {
	fmt.Println("\n\t|| Calcular Precio ||\n Concurrencia")

	var totalProd float64
	var totalServ float64
	var totalMant float64

	totalProd = sumarProductos(
		Productos{nombre: "Xiamo100T", precio: 32000, cantidad: 1},
		Productos{nombre: "Zamzun", precio: 42000, cantidad: 1},
		Productos{nombre: "Libro", precio: 140, cantidad: 5},
	)
	totalServ = sumarServicios(
		Servicios{nombre: "Medico", precio: 32000, minutos: 60},
		Servicios{nombre: "Zamzun", precio: 42000, minutos: 29},
		Servicios{nombre: "Libro", precio: 140, minutos: 5},
	)
	totalMant = sumarMantenimiento(
		Mantenimiento{nombre: "Medico", precio: 32000},
		Mantenimiento{nombre: "Zamzun", precio: 42000},
		Mantenimiento{nombre: "Libro", precio: 140},
	)

	fmt.Println("-> Total de Produ: ", totalProd)
	fmt.Println("-> Total de Serv: ", totalServ)
	fmt.Println("-> Total de Mant: ", totalMant)
	fmt.Println("-> Monto Final: ", totalProd+totalServ+totalMant)
}
