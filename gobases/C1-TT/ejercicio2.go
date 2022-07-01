package main

import "fmt"

func main() {
	edad := 21
	empleado := true
	mesesDeAntiguedad := 13
	sueldo := 100001

	if edad < 22 && empleado && mesesDeAntiguedad > 12 {
		if sueldo > 100000 {
			fmt.Println("Se te otorga el prestamo sin interes")
		} else {
			fmt.Println("Se te otorga el prestamo con interes")
		}
	} else {
		fmt.Println("No se te otorga el prestamo :(")
	}
}
