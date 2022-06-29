package main

import "fmt"

type Alumno struct {
	Nombre   string
	Apellido string
	DNI      string
	Fecha    string
}

func (e *Alumno) detalle() {
	fmt.Println("Nombre:", e.Nombre)
	fmt.Println("Apellido:", e.Apellido)
	fmt.Println("DNI:", e.DNI)
	fmt.Println("Fecha:", e.Fecha)
}

func main() {
	estudiantes := []Alumno{
		{
			Nombre:   "Claudio",
			Apellido: "Figueroa",
			DNI:      "19572365-7",
			Fecha:    "22/06/2022",
		},
		{
			Nombre:   "Julian",
			Apellido: "Renteria",
			DNI:      "1873342-7",
			Fecha:    "22/06/2022",
		},
		{
			Nombre:   "Luis",
			Apellido: "Olivera",
			DNI:      "34323443-7",
			Fecha:    "22/06/2022",
		},
	}
	for _, estudiante := range estudiantes {
		estudiante.detalle()
		fmt.Println("------------------------")
	}

}
