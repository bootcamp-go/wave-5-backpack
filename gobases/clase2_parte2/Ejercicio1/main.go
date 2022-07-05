// registro de estudiantes

package main

import (
	"fmt"
)

type Alumno struct {
	Nombre       string `json:"nombre"`
	Apellido     string `json:"apellido"`
	DNI          uint64 `json:"dni"`
	FechaIngreso string `json:"fecha"`
}

func (a Alumno) detalle() string {
	return fmt.Sprintf("\nNombre y Apellido: %s %s\nDNI: %d\nFecha: %s\n", a.Nombre, a.Apellido, a.DNI, a.FechaIngreso)
}

func main() {
	a := &Alumno{
		Nombre:       "Luz",
		Apellido:     "Jimenez",
		DNI:          11234552,
		FechaIngreso: "23/2/2022",
	}

	var a1 Alumno = Alumno{Nombre: "Juan", Apellido: "Betancur", DNI: 4333245, FechaIngreso: "3/4/2022"}
	fmt.Println(a.detalle())
	fmt.Println(a1.detalle())
}
