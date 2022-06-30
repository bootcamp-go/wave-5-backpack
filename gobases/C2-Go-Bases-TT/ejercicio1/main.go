package main

import "fmt"

// Ejercicio 1 - Registro de estudiantes
// Una universidad necesita registrar a los/as estudiantes y generar una funcionalidad
// para imprimir el detalle de los datos de cada uno de ellos/as, de la siguiente manera:

// Nombre: [Nombre del alumno]
// Apellido: [Apellido del alumno]
// DNI: [DNI del alumno]
// Fecha: [Fecha ingreso alumno]

// Los valores que están en corchetes deben ser reemplazados por los datos brindados por los alumnos/as.
// Para ello es necesario generar una estructura Alumnos con las variables
// Nombre, Apellido, DNI, Fecha y que tenga un método detalle

type Alumno struct {
	Nombre       string
	Apellido     string
	DNI          string
	FechaIngreso string
}

func (a Alumno) detalleAlumno() {
	fmt.Printf("Nombre: \t[%s]\nApellido: \t[%s]\nDNI: \t\t[%s]\nFecha: \t\t[%s]\n\n", a.Nombre, a.Apellido, a.DNI, a.FechaIngreso)
}

func main() {
	fmt.Println("Ejercicio 1 - Registro de estudiantes")
	fmt.Println("")

	// Se crea un Alumno
	var alumno1 Alumno

	// Se asignan los datos del alumno
	alumno1.Nombre = "Jose"
	alumno1.Apellido = "Alcantara"
	alumno1.DNI = "ALMJ920204HDFLUO01"
	alumno1.FechaIngreso = "26-01-2022"

	// Se imprimen los datos del alumno
	alumno1.detalleAlumno()
}
