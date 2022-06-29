package main

import "fmt"

type Alumno struct {
	nombre   string
	apellido string
	dni      string
	fecha    string
}

func main() {
	alumno1 := Alumno{"Fulanito", "Susanito", "123", "29-06-2022"}
	alumno1.detalle()
}

func (a Alumno) detalle() {
	fmt.Println("Nombre del alumno", a.nombre)
	fmt.Println("Apellido del alumno", a.apellido)
	fmt.Println("DNI del alumno", a.dni)
	fmt.Println("Fecha de ingreso", a.fecha)
}
