package main

import "fmt"

type Alumno struct {
	Nombre   string
	Apellido string
	DNI      int
	fecha    string
}

func (a Alumno) detalle() {
	fmt.Println(a)
}

func main() {
	alumno1 := Alumno{"Juan", "Perez", 40123123, "2 de marzo"}

	alumno1.detalle()

}
