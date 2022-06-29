package main

import (
	"fmt"
)

type Estudiante struct {
	Nombre   string
	Apellido string
	DNI      int
	Fecha    string
}

func (e Estudiante) detalle() {
	fmt.Println("Nombre :", e.Nombre)
	fmt.Println("Apellido :", e.Apellido)
	fmt.Println("DNI :", e.DNI)
	fmt.Println("Fecha :", e.Fecha)
}
func main() {

	e1 := Estudiante{"Diego", "Palacios", 12345678, "29/06/2022"}

	e1.detalle()

}
