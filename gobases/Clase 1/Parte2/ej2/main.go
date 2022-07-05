package main

import "fmt"

func main() {
	var edadCliente uint = 22
	var estado string = "empleado"
	var antiguedad uint = 1
	var sueldo uint = 1000000

	if edadCliente < 22 {
		fmt.Println("No cumple con los requisitos de edad")
	} else if estado != "empleado" {
		fmt.Println("No cumple con los requisitos de empleo")
	} else if antiguedad < 1 {
		fmt.Println("No cumple con los requisitos de antiguedad")
	} else if sueldo < 100000 {
		fmt.Println("Cumple con los requisitos, pero genera interes")
	} else {
		fmt.Println("Cumple con los requisitos, y NO genera interes")
	}
}
