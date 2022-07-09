package main

import (
	"fmt"
	"time"
)

type Alumno struct {
	Nombre string
	Apellido string
	DNI int
	Fecha string
}
func main()  {
	a := &Alumno{
		Nombre: "Nicolas",
		Apellido: "Herrera",
		DNI: 1192038594,
		Fecha: time.Now().String(),
	}

	fmt.Println(a.detalle())
}

func (a Alumno) detalle() string {
	return fmt.Sprintf("Nombre y apellido: %s %s \nDNI: %d \nFecha: %s", a.Nombre, a.Apellido, a.DNI, a.Fecha)
}
