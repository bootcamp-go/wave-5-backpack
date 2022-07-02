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
		Apellido: "Serna",
		DNI:      12345,
		Fecha:    "21-junio-2022",
	}
	fmt.Println(a.detalle())
}
