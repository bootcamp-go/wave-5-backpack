package main

import (
	"fmt"
)

type Student struct {
	Name     string
	LastName string
	DNI      int
	Fecha    string
}

func (s Student) detalle() {
	fmt.Printf("Nombre: %s\nApellido: %s\nDNI: %d\nFecha: %s\n", s.Name, s.LastName, s.DNI, s.Fecha)
}

func main() {
	s1 := Student{
		Name:     "Carlos",
		LastName: "Rodriguez",
		DNI:      1020,
		Fecha:    "10-07-2022",
	}
	s1.detalle()
	s2 := Student{
		Name:     "Pedro",
		LastName: "Paez",
	}
	s2.detalle()
}
