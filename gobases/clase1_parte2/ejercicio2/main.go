package main

import "fmt"

func main() {
	edad := 22
	empleado := true
	antiguedad := 1
	sueldo := 120000

	if edad >= 22 && empleado && antiguedad >= 1 {
		fmt.Println("Credito aprobado!")
		if sueldo > 100000 {
			fmt.Println("No paga intereses")
		} else {
			fmt.Println("Paga intereses")
		}
	} else {
		fmt.Println("No cumple con los requisitos para un prestamo")
	}

	/* 	if edad >= 22 && empleado && antiguedad >= 1 {
	   		fmt.Println("Cumple con los requisitos para un prestamo")
	   		if sueldo > 100000 {
	   			fmt.Println("No se le cobrarán intereses")
	   		} else {
	   			fmt.Println("Se le cobrarán intereses")
	   		}
	   	} else {
	   		fmt.Println("NO cumple con los requisitos para un prestamo")
	   	} */
}
