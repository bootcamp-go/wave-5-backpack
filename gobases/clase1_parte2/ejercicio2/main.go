package main

import "fmt"

type Persona struct {
	nombre     string
	edad       int
	empleado   bool
	antiguedad int
	sueldo     int
}

func prestamo(p Persona) {
	switch {
	case p.edad > 22, p.empleado == true, p.antiguedad > 12:
		if p.sueldo > 100000 {
			fmt.Println("El prestamo es aprobado y no se cobraran intereses")
		} else {
			fmt.Println("El prestamo es aprobado")
		}
	default:
		fmt.Println("No cumple con los requisitos para el prestamo")
	}
}

func main() {

	var persona1 Persona = Persona{nombre: "Juan", edad: 24, empleado: true, antiguedad: 18, sueldo: 320000}
	var persona2 Persona = Persona{nombre: "Pedro", edad: 34, empleado: true, antiguedad: 13, sueldo: 90000}
	var persona3 Persona = Persona{nombre: "Adrian", edad: 19, empleado: false, antiguedad: 0, sueldo: 320000}

	prestamo(persona1)
	prestamo(persona2)
	prestamo(persona3)

}
