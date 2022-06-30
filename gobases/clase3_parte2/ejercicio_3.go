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

func sumarProductos(p []Productos, c chan float64) {
	var precio_total float64
	for _, value := range p {
		precio_total += value.precio * float64(value.cantidad)
	}
	c <- precio_total
}

func sumarServicios(s []Servicios, c chan float64) {
	var precio_total float64
	for _, value := range s {
		medias_horas := math.Ceil(float64(value.minutos) / 30)
		precio_total += value.precio * medias_horas
	}
	c <- precio_total
}

func sumarMantenimiento(m []Mantenimiento, c chan float64) {
	var precio_total float64
	for _, value := range m {
		precio_total += value.precio
	}
	c <- precio_total
}

func goRoutine(c chan float64) {
	p1 := Productos{"Leche", 2000, 1}
	p2 := Productos{"Pan", 3000, 2}
	p := []Productos{p1, p2}

	s1 := Servicios{"Cocina", 3000, 30}
	s2 := Servicios{"Limpieza", 4000, 45}
	s := []Servicios{s1, s2}

	m1 := Mantenimiento{"Jardineria", 40}
	m2 := Mantenimiento{"Plomeria", 50}
	m := []Mantenimiento{m1, m2}

	go sumarProductos(p, c)
	go sumarServicios(s, c)
	go sumarMantenimiento(m, c)
}

func main() {
	c := make(chan float64)
	goRoutine(c)
	var total float64
	for i := 0; i < 3; i++ {
		valor := <-c
		fmt.Println(valor)
		total += valor
	}
	fmt.Println("Total: ", total)
}
