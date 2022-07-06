package main

import "fmt"

func main(){
	var edad int = 22
	var isEmpleado bool = false
	var antiguedad int = 0
	var sueldo int = 100000

	if (edad >= 22 && isEmpleado == true && antiguedad > 0) {
		fmt.Printf("Felicidades fuiste seleccionado para un préstamo \n")
		if sueldo > 100000 {
			fmt.Printf("No se te cobrarán intereses \n")
		}
	} else {
		fmt.Printf("Lo sentimos, no reunes los requisitos \n")
	}
}

//Un banco quiere otorgar préstamos a sus clientes, pero no todos pueden acceder a los mismos. 
// Para ello tiene ciertas reglas para saber a qué cliente se le puede otorgar. 
// Solo le otorga préstamos a clientes cuya edad sea mayor a 22 años, se encuentren empleados y tengan más de un año de antigüedad en su trabajo. 
// Dentro de los préstamos que otorga no les cobrará interés a los que su sueldo es mejor a $100.000. 
// Es necesario realizar una aplicación que tenga estas variables y que imprima un mensaje de acuerdo a cada caso.

// Tip: tu código tiene que poder imprimir al menos 3 mensajes diferentes.
   					