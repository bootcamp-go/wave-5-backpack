/* Ejercicio 1 - Impuestos de salario #1
En tu función “main”, define una variable llamada “salary” y asignarle un valor de tipo “int”.
Crea un error personalizado con un struct que implemente “Error()” con el mensaje “error: el salario ingresado no alcanza el mínimo imponible" y lánzalo en caso de que “salary” sea menor a 150.000. Caso contrario, imprime por consola el mensaje “Debe pagar impuesto”.

*/

package main

import "fmt"

type MyError struct{}

func (e *MyError) Error() string {
	return "error: no aplica impuestos"
}

func main() {
	theError := &MyError{}
	var wage int = 200000
	if wage < 150000 {
		fmt.Println(theError.Error())
	} else {
		fmt.Println("Si paga impuestos")
	}
}
