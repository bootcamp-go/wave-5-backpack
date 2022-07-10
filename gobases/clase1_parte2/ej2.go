/*
Un banco quiere otorgar préstamos a sus clientes, pero no todos pueden acceder a los mismos. Para ello tiene ciertas reglas para saber a qué cliente se le puede otorgar.
Solo le otorga préstamos a clientes cuya edad sea mayor a 22 años,
se encuentren empleados
y tengan más de un año de antigüedad en su trabajo.
 Dentro de los préstamos que otorga no les cobrará interés a los que su sueldo es mejor a $100.000.
Es necesario realizar una aplicación que tenga estas variables y que imprima un mensaje de acuerdo a cada caso.

Tip: tu código tiene que poder imprimir al menos 3 mensajes diferentes.

*/

package main

import "fmt"

func main() {

	var edades = map[string]int{"Benjamin": 20, "Nahuel": 22, "Brenda": 26, "Darío": 44}

	var empleados = map[string]bool{"Benjamin": true, "Nahuel": true, "Brenda": true, "Darío": true}

	var antiguedad = map[string]int{"Benjamin": 0, "Nahuel": 1, "Brenda": 2, "Darío": 2}

	var sueldo = map[string]int{"Benjamin": 90000, "Nahuel": 80000, "Brenda": 101000, "Darío": 90000}

	var requisitoEdad string = ""
	var requisitoTrabajo string = ""
	var requisitoAntiguedad string = ""
	var cumplenTodosLosRequisitosConInteres string = ""
	var cumplenTodosLosRequisitosSinInteres string = ""

	for k, v := range edades {

		if v > 22 {
			//fmt.Println("index:", k)

			if empleados[k] == true {
				//fmt.Println("tiene mas de 22 y tiene trabajo", k)
				if antiguedad[k] > 1 {
					//cumplenTodosLosRequisitos +=  k +  ", "
					if sueldo[k] > 100000 {
						cumplenTodosLosRequisitosSinInteres += k + ", "
					} else {
						cumplenTodosLosRequisitosConInteres += k + ", "
					}
				} else {
					requisitoAntiguedad += k + ", "
				}

			} else {
				requisitoTrabajo += k + ", "
			}

		} else {
			//fmt.Println("No cumplen con requisito de edad:", k)
			requisitoEdad += k + ", "
		}
	}

	fmt.Println("Cumplen todos los requisitos sin interes:", cumplenTodosLosRequisitosSinInteres)
	fmt.Println("Cumplen todos los requisitos con interes:", cumplenTodosLosRequisitosConInteres)
	fmt.Println("No cumplen con requisito de Edad:", requisitoEdad)
	fmt.Println("No cumplen con requisito de Trabajo:", requisitoTrabajo)
	fmt.Println("No cumplen con requisito de Antiguedad:", requisitoAntiguedad)

}
