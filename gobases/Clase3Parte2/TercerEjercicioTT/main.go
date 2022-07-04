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
	nombre            string
	precio            float64
	minutosTrabajados int
}

type Mantenimiento struct {
	nombre string
	precio float64
}

func main() {
	cp := make(chan float64)
	cs := make(chan float64)
	cm := make(chan float64)
	p1 := Productos{"carro", 50000, 5}
	p2 := Productos{"moto", 30000, 3}
	s1 := Servicios{"limpieza", 500, 75}
	s2 := Servicios{"pintura", 1000, 100}
	m1 := Mantenimiento{"limpieza", 1000}
	m2 := Mantenimiento{"pulitura", 2000}
	var prods []Productos
	prods = append(prods, p1)
	prods = append(prods, p2)
	var servs []Servicios
	servs = append(servs, s1)
	servs = append(servs, s2)
	var mant []Mantenimiento
	mant = append(mant, m1)
	mant = append(mant, m2)
	go sumarProductos(&prods, cp)
	go sumarServicios(&servs, cs)
	go sumarMantenimiento(&mant, cm)
	Total := <-cp + <-cs + <-cm
	fmt.Println(Total)
}

func sumarProductos(prods *[]Productos, c chan float64) {
	var acum float64
	for _, value := range *prods {
		acum += value.precio * float64(value.cantidad)
	}
	c <- acum
	close(c)
}
func sumarServicios(servs *[]Servicios, c chan float64) {
	var acum float64
	for _, value := range *servs {
		horasTrabajadas := math.Round(float64(value.minutosTrabajados) / 30)
		if horasTrabajadas < float64(value.minutosTrabajados)/30 {
			horasTrabajadas++
		}
		acum += value.precio * horasTrabajadas
	}
	c <- acum
	close(c)
}
func sumarMantenimiento(mant *[]Mantenimiento, c chan float64) {
	var acum float64
	for _, value := range *mant {
		acum += value.precio
	}
	c <- acum
	close(c)
}
