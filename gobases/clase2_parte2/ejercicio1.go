package main

import "fmt"

type alumnos struct {
	Nombre   string
	Apellido string
	DNI      string
	Fecha    string
}

func (a alumnos) detalles() string {
	return fmt.Sprintf("Nombre: %s\nApellido: %s\nDNI: %s\nFecha de ingreso: %s\n", a.Nombre, a.Apellido, a.DNI, a.Fecha)
}

func main() {
	alumno1 := alumnos{"Pipu", "Pepe", "12345", "12-12-12"}
	alumno2 := alumnos{"Pipu", "Pepe", "12345", "12-12-12"}

	fmt.Print(alumno1.detalles())
	fmt.Print(alumno2.detalles())
}

//Ejercicio realizado con compañeros de la sala de meets
/*
package main

import "fmt"

type Alumno struct {
	nombre   string
	apellido string
	dni      int
	fecha    string
}

func main() {
	alumno1 := Alumno{"Fulanito", "Susanito", 1, "29-06-2022"}
	alumno1.detalle()
}

func (a Alumno) detalle() {
	fmt.Println("Nombre del alumno", a.nombre)
	fmt.Println("Apellido del alumno", a.apellido)
	fmt.Println("DNI del alumno", a.dni)
	fmt.Println("Fecha de ingreso", a.fecha)
}
*/
