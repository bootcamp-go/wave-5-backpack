package main

import "fmt"

// ------- Struct Alumno ----------

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

// ------------- Main -------------

func main() {
	alumno := Alumno{
		Nombre:  "Juan",
		Apelido: "Perez",
		DNI:     "12345678",
		Fecha:   "01/01/2000",
	}
	alumno.detalle()
}
