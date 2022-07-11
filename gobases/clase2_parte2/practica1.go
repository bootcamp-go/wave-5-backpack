package main

import "fmt"

type Alumnos struct {
	Nombre string
	Apellido string
	DNI int
	Fecha string
}

func (a Alumnos) detalle() {
	fmt.Println("Nombre del alumno",a.Nombre)
	fmt.Println("Apellido del alumno",a.Apellido)
	fmt.Println("DNI del alumno",a.DNI)
	fmt.Println("Fecha de ingreso del alumno",a.Fecha)
}

func main()  {
	
	alumno := Alumnos{"Vanessa","Sotomayor",12345,"12/06/2021"}
	alumno.detalle()
	//fmt.Println(detalle(alumno))
}