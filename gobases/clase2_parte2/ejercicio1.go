package main

import "fmt"

type Alumno struct {
 Nombre string
 Apellido string
 DNI int
 Fecha string
}

func (a Alumno)detalle(){
	nombre, apellido, dni, fecha := a.Nombre, a.Apellido, a.DNI, a.Fecha
	fmt.Println(nombre)
	fmt.Println(apellido)
	fmt.Println(dni)
	fmt.Println(fecha)
	
}

func main(){
	a := Alumno{"Nahuel", "Monserrat", 44051370, "24/11/2001"}
	a.detalle()
}