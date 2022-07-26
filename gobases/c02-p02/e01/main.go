package main

import (
	"fmt"
	"time"
)

type Estudiante struct {
	Nombre       string `json:"nombre"`
	Apellido     string `json:"apellido"`
	DNI          int    `json:"dni"`
	FechaIngreso string `json:"fecha"`
}

func (e Estudiante) detalle() string {
	return fmt.Sprintf("\nNombre y Apellido: %s %s\nDNI: %d\nFecha: %s\n", e.Nombre, e.Apellido, e.DNI, e.FechaIngreso)
}

func main() {

	e1 := &Estudiante{
		Nombre:       "Luis",
		Apellido:     "Figo",
		DNI:          12345,
		FechaIngreso: time.Now().UTC().String(),
	}

	var e2 Estudiante = Estudiante{
		Nombre:       "Tony",
		Apellido:     "Kroos",
		DNI:          12346,
		FechaIngreso: time.Now().UTC().String(),
	}

	fmt.Println(e1.detalle())
	fmt.Println(e2.detalle())

}
