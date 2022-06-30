package main

import "fmt"

type Costo struct {
	Nombre string
	Precio float64
}

type Producto struct {
	Costo
	Cantidad int
}

type Servicio struct {
	Costo
	Minutos int
}

type Mantenimiento struct {
	Costo
}

func SumarProductos(p []Producto, c chan float64) {
	total := 0.0
	for _, prod := range p {
		total += prod.Precio * float64(prod.Cantidad)
	}
	fmt.Println("Total Produtos :", total)
	c <- total
}

func SumarServicios(s []Servicio, c chan float64) {
	total := 0.0

	for _, serv := range s {
		minPagar := serv.Minutos / 30
		if serv.Minutos%30 > 0 {
			minPagar++
		}
		total += serv.Precio * float64(minPagar)
	}

	fmt.Println("Total Servicios :", total)
	c <- total
}

func SumarMantenimientos(m []Mantenimiento, c chan float64) {
	total := 0.0

	for _, mant := range m {
		total += mant.Precio
	}

	fmt.Println("Total Mantenimientos :", total)
	c <- total
}

func main() {
	p1 := Producto{
		Costo:    Costo{"televisor", 20},
		Cantidad: 2,
	}
	p2 := Producto{
		Costo:    Costo{"bicicleta", 8},
		Cantidad: 1,
	}
	lp := []Producto{p1, p2}

	s1 := Servicio{
		Costo:   Costo{"internet", 2},
		Minutos: 55,
	}
	s2 := Servicio{
		Costo:   Costo{"telefonia", 1.5},
		Minutos: 61,
	}
	ls := []Servicio{s1, s2}

	m1 := Mantenimiento{
		Costo: Costo{"aire acondicionado", 2},
	}
	lm := []Mantenimiento{m1}

	c := make(chan float64)

	go SumarProductos(lp, c)
	go SumarServicios(ls, c)
	go SumarMantenimientos(lm, c)

	totalFinal := 0.0
	for i := 0; i < 3; i++ {
		fmt.Println(i, "-Recorre")
		totalFinal += <-c
	}
	fmt.Println("Total a Pagar :", totalFinal)
}
