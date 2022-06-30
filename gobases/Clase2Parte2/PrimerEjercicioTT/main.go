package main

import "fmt"

type Alumno struct {
	nombre   string
	apellido string
	dni      int
	fecha    string
}

func (a Alumno) detalles() {
	fmt.Printf("Nombre %v \n", a.nombre)
	fmt.Printf("Apellido %v \n", a.apellido)
	fmt.Printf("Dni %v \n", a.dni)
	fmt.Printf("Fecha Ingreso: %v \n", a.fecha)
}

func main() {
	alumno := Alumno{"Jose", "Gonzalez", 95959706, "26/06/2022"}
	alumno.detalles()
}
