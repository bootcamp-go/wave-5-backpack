package main

import "fmt"

//Ejercicio 3 - Calcular precio

const (
	MED float64 = 1000
	HOR float64 = 1500
)

type Productos struct {
	nombre   string
	precio   float64
	cantidad int
}

type Servicios struct {
	nombre  string
	precio  float64
	minutes int
}

type Mantenimiento struct {
	nombre string
	precio float64
}

func sumProductos(productos []Productos, c chan float64) {
	var sum float64 = 0
	for _, prod := range productos {
		sum += float64(prod.cantidad) * prod.precio
	}
	c <- sum
}

func sumServicios(servicio []Servicios, c chan float64) {
	var sum float64 = 0
	for _, serv := range servicio {
		if serv.minutes > 30 {
			serv.precio = MED
			sum += serv.precio
		} else {
			serv.precio = HOR
			sum += serv.precio
		}
	}
	c <- sum
}

func sumMantenimiento(mante []Mantenimiento, c chan float64) {
	var sum float64 = 0
	for _, mant := range mante {
		sum += float64(mant.precio)
	}
	c <- sum
}

func main() {
	p := Productos{nombre: "Maiz", precio: 1500, cantidad: 12}
	p1 := Productos{nombre: "Atun", precio: 4500, cantidad: 15}

	produclist := []Productos{p, p1}
	fmt.Println(produclist)

	s := Servicios{nombre: "Trasporte", minutes: 15}
	s1 := Servicios{nombre: "Trasporte", minutes: 39}

	servicelist := []Servicios{s, s1}
	fmt.Println(servicelist)

	m := Mantenimiento{nombre: "Electrico", precio: 8500}
	m1 := Mantenimiento{nombre: "Computadoras", precio: 4500}

	mantenlist := []Mantenimiento{m, m1}
	fmt.Println(mantenlist)

	c1 := make(chan float64)
	c2 := make(chan float64)
	c3 := make(chan float64)

	go sumProductos(produclist, c1)
	go sumServicios(servicelist, c2)
	go sumMantenimiento(mantenlist, c3)

	v1 := <-c1
	v2 := <-c2
	v3 := <-c3

	fmt.Println("Resultado total $", v1+v2+v3)
}
