package main

import "fmt"

func main() {
	var edad, sueldo, antiguedad int
	var empleado bool

	edad = 25
	sueldo = 2323232323
	antiguedad = 4
	empleado = true

	if edad > 22 && empleado == true && antiguedad > 1 {
		fmt.Println("Se le puede otorgar el prestamo")
		if sueldo > 100000 {
			fmt.Println("Y no te cobramos interes")
		} else {
			fmt.Println("Y te cobramos interes")
		}
	} else {
		fmt.Println("Usted no cumple con las condiciones para que se le otorgue un prestamo")
	}

}
