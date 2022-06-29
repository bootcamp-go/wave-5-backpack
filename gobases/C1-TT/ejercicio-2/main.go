package main

import "fmt"

func main() {
	edad := 23
	empleado := true
	antiguedad := 2
	sueldo := 100000

	if edad > 22 && empleado == true && antiguedad > 1 {
		if sueldo > 100000 {
			fmt.Println("Usted puede gozar de un prestamo sin intereses")
		} else {
			fmt.Println("Usted puede gozar de un prestamo, pero con intereses, ya que su sueldo es menor a $100.000")
		}
	} else {
		fmt.Println("Lo sentimos, pero no cumple con los requisitos para solicitar un prestamo")
	}
}
