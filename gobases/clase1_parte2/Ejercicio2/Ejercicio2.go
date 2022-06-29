package main

import "fmt"

func main() {
	var (
		edad       = 28
		sueldo     = 100001
		antiguedad = 4
		empleado   = true
	)

	if edad > 22 && empleado && antiguedad > 1 {
		fmt.Println("El empleado es elegible")
		if sueldo > 100000 {
			fmt.Println("Y no se le cobrará interés")
		} else {
			fmt.Println("Se le cobrará interés")
		}
	} else {
		fmt.Println("El empleado no cumple con las condiciones para que se le otorgue un préstamo")
	}
}
