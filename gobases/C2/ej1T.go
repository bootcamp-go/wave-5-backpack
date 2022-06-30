package main

import (
	"fmt"
)

type Alumno struct {
	Nombre   string
	Apellido string
	DNI      int
	Fecha    string
}

func (a Alumno) detalle() {
	nombre, apellido, dni, fecha := a.Nombre, a.Apellido, a.DNI, a.Fecha

	fmt.Println("Nombre: " + nombre)
	fmt.Println("Apellido: " + apellido)
	fmt.Printf("DNI: %d \n", dni)
	fmt.Println("Fecha: " + fecha)

}

func main() {
	joaquin := Alumno{
		Nombre:   "Joaquin",
		Apellido: "Alvarez",
		DNI:      51060297,
		Fecha:    "04/08/2001",
	}
	joaquin.detalle()

}
