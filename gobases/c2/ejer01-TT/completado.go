package main

import (
	"fmt"
	"time"
)

/*Ejercicio 1 - Registro de estudiantes

Una universidad necesita registrar a los/as estudiantes y generar una funcionalidad para imprimir el
detalle de los datos de cada uno de ellos/as, de la siguiente manera:

Nombre: [Nombre del alumno]
Apellido: [Apellido del alumno]
DNI: [DNI del alumno]
Fecha: [Fecha ingreso alumno]

Los valores que están en corchetes deben ser reemplazados por los datos brindados por los alumnos/as.
Para ello es necesario generar una estructura Alumnos con las variables Nombre, Apellido, DNI,
Fecha y que tenga un método detalle.*/

type Alumno struct {
	Nombre       string `json:"nombre"`
	Apellido     string `json:"apellido"`
	DNI          uint64 `json:"dni"`
	FechaIngreso string `json:"fecha"`
}

func (a Alumno) detalle() string {
	return fmt.Sprintf("\nNombre y apellido: %s %s\nDNI: %d\nFecha: %s\n", a.Nombre, a.Apellido, a.DNI, a.FechaIngreso)
}

func main() {
	a := &Alumno{
		Nombre:       "Juan",
		Apellido:     "Perez",
		DNI:          347632182,
		FechaIngreso: time.Now().UTC().String(),
	}

	var a1 Alumno = Alumno{Nombre: "Maria", Apellido: "Martinez", DNI: 82734333, FechaIngreso: time.Now().String()}
	fmt.Println(a.detalle())
	fmt.Println(a1.detalle())
}
