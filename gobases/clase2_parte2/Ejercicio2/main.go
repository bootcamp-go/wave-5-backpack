package main

import "fmt"

var (
	edad              = 23
	empleado          = false
	antiguedadTrabajo = 2
	salario           = 110000
)

func main() {

	if edad > 22 {
		fmt.Printf("El cliente es mayor a 22 \n")
		if empleado == true {
			fmt.Printf("Cliente laborando \n")
			if antiguedadTrabajo > 1 {
				fmt.Printf("Cliente con mas de un 1 ano de antiguedad \n")
			} else {
				fmt.Printf("Cliente con menos de un 1 ano de antiguedad \n")
			}
		} else {
			fmt.Printf("El cliente no esta laborando \n")
		}
	} else {
		fmt.Printf("El cliente es menor \n")
	}
	if edad > 23 && empleado == true && antiguedadTrabajo > 1 {
		fmt.Printf("Credito aprobado \n")
		if edad > 23 && empleado == true && antiguedadTrabajo > 1 {
			fmt.Printf("Credito aprobado \n")
			if salario > 100000 {
				fmt.Printf("El cliente no paga intereses \n")
			} else {
				fmt.Printf("El cliente paga intereses \n")
			}
		}

	} else {
		fmt.Printf("Credito rechazado \n")
	}
}
