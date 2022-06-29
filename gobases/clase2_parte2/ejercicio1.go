package main

import "fmt"

type Alumno struct {
	Nombre   string
	Apellido string
	DNI      int
	Fecha    string
}

func main() {
	alumno1 := Alumno{"Juan", "Casas", 123, "02-12-2012"}
	alumno2 := Alumno{"Pedro", "Paredes", 456, "12-06-2022"}
	alumno3 := Alumno{"Marcos", "Zapata", 789, "21-05-2021"}
	alumno1.detalle()
	alumno2.detalle()
	alumno3.detalle()
}

func (alumno Alumno) detalle() {
	fmt.Printf("Alumno %s %s con DNI %d ingresó el %s \n", alumno.Nombre, alumno.Apellido, alumno.DNI, alumno.Fecha)
}
