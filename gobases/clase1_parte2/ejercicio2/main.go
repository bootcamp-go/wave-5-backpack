package main

import "fmt"

func main() {
	var edad, sueldo, antiguedad int
	var empleado bool
	var interes int = 0

	edad = 25
	sueldo = 404040
	antiguedad = 4
	empleado = false

	if edad > 22 && empleado == true && antiguedad > 1 {
		fmt.Println("Se le puede otorgar el prestamo")
		interes = 1
		//otra forma de hacer esto podria ser poner otro if para comprobar el sueldo
	} else {
		fmt.Println("Usted no cumple con las condiciones para que se le otorgue un prestamo")
	}

	if interes == 1 && sueldo > 100000 {
		fmt.Println("Y no te cobramos interes")
	} else if interes == 1 {
		fmt.Println("Y te cobramos interes")
	}
}
