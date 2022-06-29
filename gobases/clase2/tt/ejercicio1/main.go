package main

import "fmt"

func main() {
	alumno := Alumno{
		Nombre:   "Cristian",
		Apellido: "Lopez",
		DNI:      12345678,
		Fecha:    "29-06-2022",
	}

	alumno.Detalle()
}

type Alumno struct {
	Nombre   string
	Apellido string
	DNI      uint
	Fecha    string
}

func (a Alumno) Detalle() {
	fmt.Printf("Nombre: %s\nApellido: %s\nDNI: %v\nFecha: %s\n", a.Nombre, a.Apellido, a.DNI, a.Fecha)
}
