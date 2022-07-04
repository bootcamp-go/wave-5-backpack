package main

import "fmt"

type Alumnos struct {
	Nombre   string
	Apellido string
	DNI      int
	Fecha    string
}

func (a Alumnos) detalles() string {
	return (fmt.Sprintf("Nombre: %v, Apellido: %v, DNI: %d, Fecha: %v", a.Nombre, a.Apellido, a.DNI, a.Fecha))
}

func main() {
	alumno := Alumnos{"Alan", "Brito", 129421542, "29-06-22"}
	fmt.Println(alumno.detalles())
}
