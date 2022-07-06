package main

import "fmt"

//Estructura Alumno
type Alumno struct {
	Nombre   string
	Apellido string
	DNI      int
	Fecha    string
}

//Funcionalidad para imrpimir el detalle
func (a Alumno) Detalle() {
	fmt.Printf("Nombre: %s\nApellido: %s\nDNI:%d\nFecha:%s\n", a.Nombre, a.Apellido, a.DNI, a.Fecha)
}

func main() {
	//reemplazando valores
	alumno := Alumno{
		Nombre:   "Luz",
		Apellido: "Lucumí",
		DNI:      2626262609,
		Fecha:    "2006/26/16",
	}
	alumno.Detalle()
}