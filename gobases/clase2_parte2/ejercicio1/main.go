package main

import "fmt"

func main() {
	var a1 Alumnos
	a1.Nombre = "Yvonne"
	a1.Apellido = "Pintos"
	a1.DNI = 35004478
	a1.Fecha = "31/03/22"
	a1.detalle()
}

type Alumnos struct {
	Nombre   string
	Apellido string
	DNI      int
	Fecha    string
}

func (a Alumnos) detalle() {
	fmt.Print("Nombre:", a.Nombre, "\nApellido:", a.Apellido, "\nDNI:", a.DNI, "\nFecha:", a.Fecha, "\n")
}
