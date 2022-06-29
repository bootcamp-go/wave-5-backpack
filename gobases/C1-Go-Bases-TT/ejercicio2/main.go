package main

import "fmt"

// Ejercicio 2 - Préstamo

// Un banco quiere otorgar préstamos a sus clientes, pero no todos pueden acceder a los mismos. Para ello tiene ciertas reglas para saber a qué cliente se le puede otorgar. Solo le otorga préstamos a clientes cuya edad sea mayor a 22 años, se encuentren empleados y tengan más de un año de antigüedad en su trabajo. Dentro de los préstamos que otorga no les cobrará interés a los que su sueldo es mejor a $100.000.
// Es necesario realizar una aplicación que tenga estas variables y que imprima un mensaje de acuerdo a cada caso.

// Tip: tu código tiene que poder imprimir al menos 3 mensajes diferentes.

func main() {
	fmt.Println("Prestamos bancarios!")

	// Datos del posible cliente
	var edad int = 12
	var empleado bool = false
	var antiguedad int = 3 // en meses
	var sueldo float32 = 12300.00

	if edad > 22 {
		if empleado {
			if antiguedad > 12 {
				if sueldo > 100000 {
					fmt.Println("Se autorizó un préstamo a su nombre sin intereses, Felicidades!")
				} else {
					fmt.Println("Se autorizó un préstamo a su nombre con el interés del 10%")
				}
			} else {
				fmt.Println("Debe tener mas de un año de antiguedad para que se le otorgue un préstamo!")
			}
		} else {
			fmt.Println("Es indispensable ser empleado para que se le otorgue un préstamo!")
		}
	} else {
		fmt.Println("No tiene la edad necesaria para que se le otorgue un préstamo!")
	}
}
