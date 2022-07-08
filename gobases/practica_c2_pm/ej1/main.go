package main

import "fmt"

type Alumno struct {
	Nombre  string
	Apelido string
	DNI     string
	Fecha   string
}

func (this *Alumno) detalle() {
	fmt.Println("Nombre:", this.Nombre)
	fmt.Println("Apelido:", this.Apelido)
	fmt.Println("DNI:", this.DNI)
	fmt.Println("Fecha:", this.Fecha)
}

func main() {
	alumno := Alumno{
		Nombre:  "Matias",
		Apelido: "Carrasco",
		DNI:     "159600783",
		Fecha:   "09-12-1984",
	}
	alumno.detalle()
}
