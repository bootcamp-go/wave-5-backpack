package main

import "fmt"

type Alumno struct {
	Nombre   string
	Apellido string
	Dni      int
	Fecha    string
}

func (a Alumno) detalle() {
	fmt.Println("Nombre:", a.Nombre, " \nApellido:", a.Apellido, " \nDNI:", a.Dni, " \nFecha Ingreso:", a.Fecha)
}

func main() {
	al := Alumno{"Sebastian", "Olivera", 123456, "04/11/2010"}

	al.detalle()
}
