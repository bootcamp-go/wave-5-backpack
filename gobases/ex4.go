package main

import "fmt"

func ex4() {
	edad, empleado, antiguedad, sueldo := 22, true, 2, 100.000
	if edad > 21 && empleado && antiguedad > 1 {
		if sueldo > 100.000 {
			fmt.Println("Prestamo sin interes validado")
		} else {
			fmt.Println("Prestamo con interes validado")
		}
	} else {
		fmt.Println("Prestamo NO validado")
	}
}
