/*
	Ejercicio 1 - Impuestos de salario #1
	En tu función “main”, define una variable llamada “salary” y asignarle un valor de tipo “int”.

	Crea un error personalizado con un struct que implemente “Error()” con el mensaje 
	“error: el salario ingresado no alcanza el mínimo imponible" y lánzalo en caso de que 
	“salary” sea menor a 150.000. Caso contrario, imprime por consola el mensaje 
	“Debe pagar impuesto”.
*/
package main

import "fmt"

type errorImpuesto struct {
	msg string
}

func (e *errorImpuesto) Error() string {
	return fmt.Sprintf("%v", e.msg)
}

func verificaImpuesto(salary int) (string, error) {
	if salary < 150000 {
		return "", &errorImpuesto {
			msg: "error: el salario ingresado no alcanza el mínimo imponible",
		}
	}
	return "Debe pagar impuesto", nil
}

func main() {
	var salary int = 200000
	mensaje, err := verificaImpuesto(salary)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(mensaje)
	}
}