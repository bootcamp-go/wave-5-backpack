package main

import "fmt"

type Alumnos struct {
	Nombre string
	Apellido string
	DNI int
	fecha string
}

func (a Alumnos) detalle(){
	fmt.Printf("Nombre: %s \nApellido: %s \nDNI: %d \nFecha de ingreso: %s \n", a.Nombre, a.Apellido, a.DNI, a.fecha)
}

func main(){
	alumno := Alumnos{"Daniela","Bedoya",1001,"04-05-2002"}
	alumno.detalle()
}