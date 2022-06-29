package main

import "fmt"

type Alumno struct {
	Nombre   string
	Apellido string
	DNI      int
	Fecha    string
}

func (a Alumno) detalle() {
	fmt.Printf("Nombre: %v\n", a.Nombre)
	fmt.Printf("Apellido: %v\n", a.Apellido)
	fmt.Printf("DNI: %v\n", a.DNI)
	fmt.Printf("Fecha: %v\n", a.Fecha)
}

func main() {
	alumno := Alumno{"Patricio", "Flood", 42225884, "29/06/2022"}
	alumno.detalle()
}
