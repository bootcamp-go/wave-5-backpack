package main

import "fmt"

type Alumnos struct {
	Nombre   string
	Apellido string
	DNI      int
	Fecha    string
}

func (a Alumnos) detalle() string {
	return fmt.Sprintf("El nombre y apellido son: %s %s. El DNI es %d y la fecha de ingreso fue %s ", a.Nombre, a.Apellido, a.DNI, a.Fecha)
}

func main() {
	a := Alumnos{
		Nombre:   "Juan",
		Apellido: "Salazar",
		DNI:      1144100536,
		Fecha:    "18-Enero-1998",
	}
	fmt.Println(a.detalle())
}
