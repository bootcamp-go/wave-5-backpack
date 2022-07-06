// Ejercicio 1 - Registro de estudiantes
// Una universidad necesita registrar a los/as estudiantes y generar una funcionalidad para imprimir el detalle de los datos de cada uno de ellos/as, de la siguiente manera:

// Nombre: [Nombre del alumno]
// Apellido: [Apellido del alumno]
// DNI: [DNI del alumno]
// Fecha: [Fecha ingreso alumno]

// Los valores que están en corchetes deben ser reemplazados por los datos brindados por los alumnos/as.
// Para ello es necesario generar una estructura Alumnos con las variables Nombre, Apellido, DNI, Fecha y que tenga un método detalle

package main

import (
	"fmt"
)

type Alumno struct {
	Nombre   string `json:"name"`
	Apellido string `json:"lastname"`
	DNI      int    `json:"dni"`
	Fecha    string `json:"date"`
}

func (al Alumno) detalle() string {

	return fmt.Sprintf("\nNombre y apellido: %s %s\nDNI: %d\nFecha: %s\n", al.Nombre, al.Apellido, al.DNI, al.Fecha)

}
func main() {

	al := Alumno{Nombre: "eimi",
		Apellido: "galvan",
		DNI:      32132132,
		Fecha:    "10 nov"}
		al1 := Alumno{Nombre: "fran",
		Apellido: "biagi",
		DNI:      33132132,
		Fecha:    "11 nov"}
	fmt.Println(al.detalle())
	fmt.Println(al1.detalle())
}
