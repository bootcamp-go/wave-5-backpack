/* 
	Ejercicio 2 - Préstamo

	Un banco quiere otorgar préstamos a sus clientes, pero no todos pueden acceder a los mismos. 
	Para ello tiene ciertas reglas para saber a qué cliente se le puede otorgar. 
	Solo le otorga préstamos a clientes cuya edad sea mayor a 22 años, se encuentren empleados y 
	tengan más de un año de antigüedad en su trabajo. Dentro de los préstamos que otorga no les cobrará 
	interés a los que su sueldo es mejor a $100.000. 
	Es necesario realizar una aplicación que tenga estas variables y que imprima un mensaje de acuerdo a cada caso.

	Tip: tu código tiene que poder imprimir al menos 3 mensajes diferentes.
*/

package main

import "fmt"

func main() {

	// Solución 1
	if edad:= 23; edad <= 22 {
		fmt.Println("Debe ser mayor a 22 años")
	} else if situacionLaboral := true; situacionLaboral == false {
		fmt.Println("Debe estar empleado")
	} else if antiguedadLaboral := 1; antiguedadLaboral < 1 {
		fmt.Println("Debe tener antigüedad laboral mayor a 1 año")
	} else {
		fmt.Println("¡Felicidades, puede acceder al préstamo!")
		if sueldo := 190000; sueldo >= 100000 {
			fmt.Println("¡Además, no paga interés!")
	 	}
	}
	
	/* Solución 2
	edad, antiguedadLaboral, sueldo := 22, 1, 190000
	situacionLaboral := true

	var mensaje string
	switch {
	case edad < 22:
		mensaje = "Debe ser mayor a 22 años"
	case situacionLaboral == false:
		mensaje = "Debe estar empleado"
	case antiguedadLaboral < 1:
		mensaje = "Debe tener antigüedad laboral mayor a 1 año"
	default:
		mensaje = "Felicidades, puede acceder al préstamo! "
		if sueldo >= 100000 {
			mensaje += "Además, no paga interés."
	 	}
	}
	fmt.Println(mensaje)
	*/

	/* Solución 3
	var mensajes []string
	if edad := 18; edad < 22 {
		mensajes = append(mensajes, "Debe ser mayor a 22 años")
	} 
	if situacionLaboral := true; situacionLaboral == false {
		mensajes = append(mensajes,"Debe estar empleado")
	}
	if antiguedadLaboral := 1; antiguedadLaboral < 1 {
		mensajes = append(mensajes,"Debe tener antigüedad laboral mayor a 1 año")
	} 
	
	if len(mensajes) == 0 {
		mensaje := "¡Felicidades, puede acceder al préstamo! "
	 	if sueldo := 190000; sueldo >= 100000 {
			mensaje += "Además, no paga interés."
	 	}
	 	fmt.Println(mensaje)
	} else {
		fmt.Println("Lo sentimos, no puede acceder al préstamo debido a:")
		for _, mensaje := range mensajes {
			fmt.Println("-",mensaje)
		}
	}
	*/
	
}