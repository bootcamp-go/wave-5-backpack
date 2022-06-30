package main

import "fmt"

type Alumno struct {
	Nombre   string
	Apellido string
	DNI      int
	Fecha    string
}

func (a Alumno) detalle() {
	fmt.Printf("Nombre: %s\nApellido: %s\nDNI: %d\nFecha: %s\n", a.Nombre, a.Apellido, a.DNI, a.Fecha)
}

func main() {
	Jessica := Alumno{"Jessica", "Escobar", 1234, "21-06-2022"}
	Jessica.detalle()
}
