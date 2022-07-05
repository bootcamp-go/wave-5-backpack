package main

import "fmt"

type Alumno struct {
	Nombre   string
	Apellido string
	DNI      int
	Fecha    string
}

func (alumno Alumno) detalle() {
	fmt.Printf("Nombre: %s\n", alumno.Nombre)
	fmt.Printf("Apellido: %s\n", alumno.Apellido)
	fmt.Printf("DNI: %d\n", alumno.DNI)
	fmt.Printf("Fecha: %s\n", alumno.Fecha)
}

func main() {
	alumno := Alumno{Nombre: "Matias", Apellido: "Vince", DNI: 42263580, Fecha: "21/06/2022"}
	alumno.detalle()
}
