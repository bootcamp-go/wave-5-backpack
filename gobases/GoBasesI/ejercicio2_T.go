package main

import "fmt"

func main() {

	var edad int = 30
	var estaTrabajando string = "si"
	var antiguedad int = 1
	var sueldo float64 = 80000

	if edad > 22 {
		if estaTrabajando == "si" {
			if antiguedad > 1 {
				if sueldo > 100000 {
					fmt.Println("Puede acceder a un credito")
					fmt.Println("No se le cobrara intereses")
				} else {
					fmt.Println("Puede acceder a un credito, pero se cobra interes")
				}
			} else {
				fmt.Println("Tienes la edad, estas empleado pero no la antiguedad")
			}
		} else {
			fmt.Println("tienes la edad pero no estas trabajando")
		}
	} else {
		fmt.Println("No tienes la edad suficiente")
	}
}
