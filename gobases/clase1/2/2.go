// Ejercicio 2 - Préstamo

// // Un banco quiere otorgar préstamos a sus clientes, pero no todos pueden acceder a los mismos. 
//Para ello tiene ciertas reglas para saber a qué cliente se le puede otorgar. 
//Solo le otorga préstamos a clientes cuya edad sea mayor a 22 años, se encuentren empleados y tengan más de un año de 
// antigüedad en su trabajo. Dentro de los préstamos que otorga no les cobrará interés a los que su sueldo es 
// mejor a $100.000. 
// // Es necesario realizar una aplicación que tenga estas variables y que imprima un mensaje de acuerdo a cada caso.

// // Tip: tu código tiene que poder imprimir al menos 3 mensajes diferentes.

package main
import "fmt"

func main()  {
	var (
		edadCliente int = 33
		estaEmpleado bool = true 
		sueldoCliente int = 900000
		antiguedad int = 2
	)
	if edadCliente > 22 && estaEmpleado && antiguedad >= 1  {
		if sueldoCliente > 100000 {
			fmt.Printf("su prestamo fue otorgado sin interes")
		}else{
			fmt.Printf("su prestamo fue aprovado con interes")
		}

	} else {
		fmt.Printf("su prestamo fue no aprovado")

	}


}
