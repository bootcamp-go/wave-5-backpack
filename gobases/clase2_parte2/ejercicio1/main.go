package main

import "fmt"

type Estudiante struct {
	Nombre   string
	Apellido string
	DNI      int
	Fecha    string
}

func (e Estudiante) detalle() {
	fmt.Printf("Nombre: %v\nApellido: %v\nDNI: %v\nFecha: %v\n", e.Nombre, e.Apellido, e.DNI, e.Fecha)

}

func main() {
	est := Estudiante{Nombre: "Gabriel", Apellido: "Torrealba", DNI: 95670164, Fecha: "29/06/2022"}

	est.detalle()
}
