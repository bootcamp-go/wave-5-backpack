package main

import "fmt"

func prestamo() {
	//ejercicio 2

	edad := 23
	empleado := true
	antiguedad := 3
	sueldo := 200000

	if edad > 22 && empleado && antiguedad > 1 {
		if sueldo > 100000 {
			fmt.Println("Prestamo aprobado! Le informamos tambien que el mismo no tiene intereses.")
		} else {
			fmt.Println("Prestamo aprobado!")
		}
	} else {
		fmt.Println("Lo sentimos, no cumple los requisitos para acceder al prestamo")
	}

}
