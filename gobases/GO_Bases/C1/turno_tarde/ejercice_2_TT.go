package main

import "fmt"

func main() {
	client_age := 23
	client_employed := true
	client_workedAge := 0.2

	if client_age <= 22 {
		fmt.Println("No tiene la edad suficiente para conseguir el prestamo")
	} else if !client_employed {
		fmt.Println("No se le concede el prestamo porque esta desempleado")
	} else if client_workedAge < 1 {
		fmt.Println("No se le concede el prestamo porque no tiene tiempo suficiente en su trabajo")
	} else {
		fmt.Println("Se le puede conceder el prestamo")
	}
}
