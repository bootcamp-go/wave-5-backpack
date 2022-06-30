package main

import "fmt"

type Alumno struct {
	Nombre       string
	Apellido     string
	DNI          string
	FechaIngreso string
}

func (alumno Alumno) Detalle() {
	fmt.Printf("Nombre: %s\nApellido: %s\nDNI: %s\nFecha de ingreso: %s", alumno.Nombre,
		alumno.Apellido, alumno.DNI, alumno.FechaIngreso)
}
