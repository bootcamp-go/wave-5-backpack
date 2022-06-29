package main

import (
	"fmt"
)

type Alumno struct {
	Nombre   string
	Apellido string
	Edad     int
	Fecha    string
}

func main() {
	nombre := "Agustin"
	apellido := "Nigrelli"
	edad := 32
	fecha := "23/06/2022"

	alumno := Alumno{}
	alumno.detalle(nombre, apellido, edad, fecha)

	fmt.Println(alumno)
}

func (a *Alumno) detalle(nombre string, apellido string, edad int, fecha string) {
	a.Nombre = nombre
	a.Apellido = apellido
	a.Edad = edad
	a.Fecha = fecha
}
