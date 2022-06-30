package main

import "fmt"

//Ejercicio 1 - Registro de Estudiantes

type Estudiante struct {
	nombre   string
	apellido string
	dni      int
	fecha    string
}

func (e Estudiante) detalles() {
	fmt.Println("Detalles del estudiante: ", e.nombre, e.apellido, 
e.fecha, e.dni)
}

func main() {
	e1 := Estudiante{"Christian", "Daniel", 1005815706, "16/04/2022"}
	e1.detalles()
}
