package main

import (
	"fmt"
)

type Alumnos struct {
	Nombre   string
	Apellido string
	Fecha    string
	DNI      int
}

func (a Alumnos) studentDetail() {
	fmt.Printf("Nombre: %s\nApellidos: %s\nFecha: %s\nDNI: %v\n", a.Nombre, a.Apellido, a.Fecha, a.DNI)
}

func (b *Alumnos) setStudent(name, lastName, date string, dni int) {
	b.Nombre = name
	b.Apellido = lastName
	b.Fecha = date
	b.DNI = dni
}

func main() {
	student := Alumnos{}
	student.setStudent("José Luis", "Riverón Rodríguez", "16/06/1994", 94061643089)
	student.studentDetail()
}
