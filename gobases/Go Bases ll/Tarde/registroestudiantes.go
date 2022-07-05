package main

import "fmt"

type Estudiante struct {
	nombre   string
	apellido string
	DNI      int
	fecha    string
}

func (e Estudiante) detalle() {
	fmt.Printf("nombre:\t%s\napellido:\t%s\nDNI:\t%d\nfecha\t%s\n", e.nombre, e.apellido, e.DNI, e.fecha)
}

func main() {
	e1 := Estudiante{"Camilo", "Martinez", 123456789, "10-07-2021"}
	Estudiante.detalle(e1)
}
