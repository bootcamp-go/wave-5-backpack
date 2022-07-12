package main

import "fmt"

type Alumno struct {
	Nombre string
	Apellido string
	DNI string
	Fecha string
}
func (alum Alumno) Detalle(){
	fmt.Printf("Nombre: %s\n",alum.Nombre)
	fmt.Printf("Apellido: %s\n",alum.Apellido)
	fmt.Printf("DNI: %s\n",alum.DNI)
	fmt.Printf("Fecha: %s\n",alum.Fecha)
}

func main(){
	alumno := Alumno{"Luis","Tapia","00001","18/07/1997"}
	alumno.Detalle();
}