package main

import "fmt"

/*Ejercicio 3 - Calcular Precio

Una empresa nacional se encarga de realizar venta de productos, servicios y mantenimiento.
Para ello requieren realizar un programa que se encargue de calcular el precio total de Productos,
Servicios y Mantenimientos. Debido a la fuerte demanda y para optimizar la velocidad requieren que
el cálculo de la sumatoria se realice en paralelo mediante 3 go routines.

Se requieren 3 estructuras:
	1. Productos: nombre, precio, cantidad.
	2. Servicios: nombre, precio, minutos trabajados.
	3. Mantenimiento: nombre, precio.

Se requieren 3 funciones:
	1. Sumar Productos: recibe un array de producto y devuelve el precio total (precio * cantidad).
	2. Sumar Servicios: recibe un array de servicio y devuelve el precio total (precio * media hora trabajada,
	si no llega a trabajar 30 minutos se le cobra como si hubiese trabajado media hora).
	3. Sumar Mantenimiento: recibe un array de mantenimiento y devuelve el precio total.

Los 3 se deben ejecutar concurrentemente y al final se debe mostrar por pantalla el monto final
(sumando el total de los 3).*/

type Products struct {
	NameProduct  string
	PriceProduct float64
	Amount       int
}

type Services struct {
	NameService  string
	PriceService float64
	WorkMinutes  float64
}

type Maintenance struct {
	NameMaintenance  string
	PriceMaintenance float64
}

func AddServices(service *[]Services, c chan float64) {
	var fullService float64
	for _, value := range *service {
		if value.WorkMinutes < 30 {
			fullService += value.PriceService * 30
		} else {
			fullService += value.WorkMinutes * value.PriceService
		}
	}
	fmt.Printf("Total del costo de servicios: $%.2f\n", fullService)
	c <- fullService
	close(c)
}

func AddProducts(product *[]Products, c chan float64) {
	var fullPrice float64
	for _, value := range *product {
		fullPrice += value.PriceProduct * float64(value.Amount)
	}
	fmt.Printf("El Precio Total de los productos es: $%.2f\n", fullPrice)
	c <- fullPrice
	close(c)
}

func AddMaintenance(m *[]Maintenance, c chan float64) {
	var fullMaintenancePrice float64
	for _, value := range *m {
		fullMaintenancePrice += value.PriceMaintenance
	}
	fmt.Printf("El Precio Total de Mantenimiento es: $%.2f\n", fullMaintenancePrice)
	c <- fullMaintenancePrice
	close(c)
}

func main() {
	service := []Services{
		{"Service 1", 650, 30},
		{"Service 2", 860, 75},
		{"service 3", 300, 180},
	}

	product := []Products{
		{"Bicycle", 4500, 5},
		{"Shoes", 450, 38},
		{"Phones", 8750, 120},
	}

	maintenance := []Maintenance{
		{"Mant 1", 59},
		{"Mant 2", 36},
		{"Mant 3", 85},
	}

	c1 := make(chan float64)
	c2 := make(chan float64)
	c3 := make(chan float64)

	go AddServices(&service, c1)
	go AddProducts(&product, c2)
	go AddMaintenance(&maintenance, c3)

	t1 := <-c1
	t2 := <-c2
	t3 := <-c3

	fmt.Printf("Precio Total Final: 💷 %.2f\n", t1+t2+t3)
}
