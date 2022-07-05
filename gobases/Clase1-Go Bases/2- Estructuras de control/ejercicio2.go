package main

import "fmt"

func main(){
	clienteEdad := 22
	clienteOcupado := false
	clienteAntiguedad := 5
	clienteSueldo := 100001
	switch {
		case (clienteEdad < 22):
			fmt.Print("No cumple la edad\n")
			fallthrough
		case (!clienteOcupado):
			fmt.Print("No está ocupado\n")
			fallthrough
		case (clienteAntiguedad < 1):
			fmt.Print("No tiene antiguedad suficiente\n")
			
		default:
			if(clienteSueldo > 100000){
				fmt.Print("Se le otorga prestamo sin intereses\n")
			}else{
				fmt.Printf("Se le otorga prestamo con intereses\n")
			}	


	}
}

//Ejercicio 2 - Préstamo

//Un banco quiere otorgar préstamos a sus clientes, pero no todos pueden acceder a los mismos. 
//Para ello tiene ciertas reglas para saber a qué cliente se le puede otorgar. 
//Solo le otorga préstamos a clientes cuya edad sea mayor a 22 años, se encuentren empleados y tengan más de un año de antigüedad en su trabajo. 
//Dentro de los préstamos que otorga no les cobrará interés a los que su sueldo es mejor a $100.000. 
//Es necesario realizar una aplicación que tenga estas variables y que imprima un mensaje de acuerdo a cada caso.

//Tip: tu código tiene que poder imprimir al menos 3 mensajes diferentes.
