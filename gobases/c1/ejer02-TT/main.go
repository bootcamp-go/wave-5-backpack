package main

import "fmt"

// Un banco quiere otorgar préstamos a sus clientes, pero no todos pueden acceder a los mismos.
// Para ello tiene ciertas reglas para saber a qué cliente se le puede otorgar.
// Solo le otorga préstamos a clientes cuya edad sea mayor a 22 años, se encuentren empleados y tengan más de un año de antigüedad en su trabajo.
// Dentro de los préstamos que otorga no les cobrará interés a los que su sueldo es mejor a $100.000.
// Es necesario realizar una aplicación que tenga estas variables y que imprima un mensaje de acuerdo a cada caso.

// Tip: tu código tiene que poder imprimir al menos 3 mensajes diferentes.

func main() {

	var edad int
	var empleo bool
	var antiguedad int
	var sueldo float32

	edad = 23
	empleo = true
	antiguedad = 1
	sueldo = 100001

	if edad > 22 && empleo == true && antiguedad > 1 && sueldo > 100000 {
		fmt.Println("Prestamo otorgado")
	} else if edad <= 22 {
		fmt.Println("No cumple con la edad")
	} else if empleo == false {
		fmt.Println("No tiene empleo")
	} else if antiguedad <= 1 {
		fmt.Println("No cumple con la antiguedad minima")
	} else if sueldo <= 100000 {
		fmt.Println("No cumple con el sueldo minimo")
	}
}
