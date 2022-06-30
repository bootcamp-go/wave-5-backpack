package main

import "fmt"

type Alumno struct {
	Nombre   string
	Apellido string
	DNI      string
	Fecha    string
}

func (a Alumno) detalle() {
	fmt.Println("===================================")
	fmt.Println("Nombre :", a.Nombre)
	fmt.Println("Apellido :", a.Apellido)
	fmt.Println("DNI :", a.DNI)
	fmt.Println("Fecha :", a.Fecha)
}

func main() {
	alumno := Alumno{
		Nombre:   "Cristobal",
		Apellido: "Monsalve",
		DNI:      "190039498",
		Fecha:    "22-06-22",
	}

	alumno.detalle()

}
