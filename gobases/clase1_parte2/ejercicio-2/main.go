package main

import "fmt"

func main() {
	/*
		Un banco quiere otorgar préstamos a sus clientes,
		pero no todos pueden acceder a los mismos.
		Para ello tiene ciertas reglas para saber a qué cliente se le puede otorgar.
		Solo le otorga préstamos a clientes cuya edad sea mayor a 22 años,
		se encuentren empleados y tengan más de un año de antigüedad en su trabajo.
		Dentro de los préstamos que otorga no les cobrará interés a los que su sueldo
		es mejor a $100.000.
		Es necesario realizar una aplicación que tenga estas variables y
		que imprima un mensaje de acuerdo a cada caso.
	*/

	var edad int = 23
	var empleado bool = true
	var antiguedad int = 2
	var sueldo int = 100000

	switch {
	case edad > 22 && empleado && antiguedad > 1:
		fmt.Print("Puede otorgar préstamos")
		if sueldo > 100000 {
			fmt.Println(" y No le cobrará interés")
		} else {
			fmt.Println(" y Le cobrará interés")
		}
	default:
		fmt.Println("No puede otorgar préstamos")
	}
}
