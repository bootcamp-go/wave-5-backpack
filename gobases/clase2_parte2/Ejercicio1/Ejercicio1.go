package main

import (
	"fmt"
	"time"
)

type Alumno struct {
	Nombre   string
	Apellido string
	DNI      int
	Fecha    string
}

func (a Alumno) detalle() string {
	return fmt.Sprintf("Nombre y apellido: %s %s\nDNI: %d\nFecha: %s\n", a.Nombre, a.Apellido, a.DNI, a.Fecha)
}

func main() {
	a := Alumno{
		Nombre:   "Andres",
		Apellido: "Rivera",
		DNI:      5689,
		Fecha:    time.Now().UTC().Format(time.RFC3339),
	}
	fmt.Println(a.detalle())
}
