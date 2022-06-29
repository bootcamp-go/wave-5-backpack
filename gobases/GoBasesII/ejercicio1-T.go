package main

import "fmt"

type Alumnos struct {
	Nombre   string
	Apellido string
	DNI      int
	Fecha    string
}

func main() {

	a := Alumnos{Nombre: "Juan", Apellido: "Serna", DNI: 12345, Fecha: "28-06-2022"}

	a.detalle()

}

func (a Alumnos) detalle() {

	fmt.Printf("El nombre es %v ", a.Nombre)
	fmt.Printf("El apellido es %v ", a.Apellido)
	fmt.Printf("El DNI es %v ", a.DNI)
	fmt.Printf("La fecha es %v ", a.Fecha)

}
