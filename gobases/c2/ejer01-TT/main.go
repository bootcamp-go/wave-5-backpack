package main

import "fmt"

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

func (a *Alumno) detalle() string {

	return fmt.Sprintf("Nombre: %s\nApellido: %s\nDNI: %d\nFecha de Ingreso: %s", a.Nombre, a.Apellido, a.DNI, a.FechaIngreso)
}

func main() {
	a := Alumno{
		"Juan", "Perez", 12345678, "11/03/86",
	}

	fmt.Println(a.detalle())
}
