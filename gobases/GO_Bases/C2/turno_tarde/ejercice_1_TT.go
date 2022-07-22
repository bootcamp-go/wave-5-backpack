package main

import "fmt"

type Alumno struct {
	nombre        string
	apellido      string
	dni           string
	fecha_ingreso string
}

func (a Alumno) detalle() {
	fmt.Printf("Nombre: %s\n", a.nombre)
	fmt.Printf("Apellido: %s\n", a.apellido)
	fmt.Printf("DNI: %s\n", a.dni)
	fmt.Printf("Fecha de Ingreso: %s\n", a.fecha_ingreso)
}

func main() {
	alumno := Alumno{}
	alumno.nombre = "Franco"
	alumno.apellido = "Pergolini"
	alumno.dni = "37716720"
	alumno.fecha_ingreso = "21-06-2022"
	alumno.detalle()
}
