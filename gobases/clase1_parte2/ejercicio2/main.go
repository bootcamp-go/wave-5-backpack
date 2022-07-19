package main

import "fmt"

func main() {
	var edad int = 25
	var antiguedad int = 2
	var empleado bool = true
	var sueldo float32 = 100500

	if edad > 22 && empleado && antiguedad > 1 {
		fmt.Println("Cumple con los requisitos basicos para aplicar al prestamos")
		if sueldo > 100000 {
			fmt.Println("Usted esta excento de interes, por tu alto sueldo.")
		} else {
			fmt.Println("Usted tendra que pagar una tasa de interes, por su bajo salario.")
		}
	} else {
		fmt.Println("Su prestamo ha sido rechazado, no cumple con los requsitos basicos")
	}
}
