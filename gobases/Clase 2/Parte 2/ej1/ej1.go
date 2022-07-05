package main

import "fmt"

type Alumnos struct {
	Nombre   string
	Apellido string
	DNI      int
	Fecha    string
}

func (a Alumnos) Detalle() {
	fmt.Printf("Nombre: %s\nApellido: %s\nDNI:%d\nFecha:%s\n", a.Nombre, a.Apellido, a.DNI, a.Fecha)
}

func main() {
	al := Alumnos{
		Nombre:   "Andres",
		Apellido: "Ramos",
		DNI:      123123123,
		Fecha:    "2002/12/12",
	}
	al.Detalle()
}
