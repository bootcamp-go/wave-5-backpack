package main

import (
	"fmt"
	"time"
)

type Mantenimiento struct {
	name  string
	price float64
	stock int
}

type Servicio struct {
	name    string
	price   float64
	minutes float64
}

type Producto struct {
	name  string
	price float64
	stock int
}

func sumarProductos(products []Producto) {
	var total float64 = 0

	for _, product := range products {
		total += (product.price * float64(product.stock))
	}

	fmt.Println("Total productos: $", total)

}

func sumarMantenimiento(mantos []Mantenimiento) {
	var total float64 = 0
	for _, mantto := range mantos {
		total += mantto.price
	}

	fmt.Println("Total mantenimientos: $", total)

}

func sumarServicios(servicios []Servicio) {
	var total float64 = 0

	for _, servicio := range servicios {
		var mediumHr float64 = servicio.minutes / 2
		if mediumHr < 1 {
			mediumHr = 1
		}
		total += (servicio.price * mediumHr)
	}

	fmt.Println("Total servicios: $", total)

}
func main() {

	var (
		products = []Producto{}
		services = []Servicio{}
		manttos  = []Mantenimiento{}
	)

	p1, p2, p3 := Producto{"Item1", 10, 5}, Producto{"Item2", 15, 10}, Producto{"Item3", 15, 5}
	products = append(products, p1)
	products = append(products, p2)
	products = append(products, p3)

	s1, s2, s3 := Servicio{"Service1", 15, 90}, Servicio{"Service2", 15, 30}, Servicio{"Service3", 10, 60}

	services = append(services, s1)
	services = append(services, s2)
	services = append(services, s3)

	m1, m2, m3 := Mantenimiento{"M1", 5, 1}, Mantenimiento{"M2", 1, 5}, Mantenimiento{"M3", 15, 10}
	manttos = append(manttos, m1)
	manttos = append(manttos, m2)
	manttos = append(manttos, m3)

	go sumarProductos(products)
	go sumarServicios(services)
	go sumarMantenimiento(manttos)

	time.Sleep(6000 * time.Millisecond)

}
