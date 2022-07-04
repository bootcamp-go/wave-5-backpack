/*
Ejercicio 1 - Impuestos de salario #1
En tu función “main”, define una variable llamada “salary” y asignarle un valor de tipo “int”.
Crea un error personalizado con un struct que implemente “Error()” con el mensaje “error: el salario ingresado no alcanza el mínimo imponible" y lánzalo en caso de que “salary” sea menor a 150.000. Caso contrario, imprime por consola el mensaje “Debe pagar impuesto”.
*/

package main

import "fmt"

type error interface{
	Error()string
}

type myCustomError struct {
	cant int
	msg string
}
func (e *myCustomError) Error() string {
	return fmt.Sprintf("%v\n", e.msg)
}

func myCustomErrorTest (cant int) error{
	if cant < 150000 {
		return &myCustomError {
		cant : cant,
		msg : "error: el salario ingresado no alcanza el mínimo imponible",
		}
	}
	return nil
}

func main(){
	var salary int = 180000
	err := myCustomErrorTest(salary)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	fmt.Println("Debe pagar impuesto")
}