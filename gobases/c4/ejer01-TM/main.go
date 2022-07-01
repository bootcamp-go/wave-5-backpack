package main

import (
	"fmt"
	"os"
)

// En tu función “main”, define una variable llamada “salary” y asignarle un valor de tipo “int”.
// Crea un error personalizado con un struct que implemente “Error()” con el mensaje
// “error: el salario ingresado no alcanza el mínimo imponible"
// y lánzalo en caso de que “salary” sea menor a 150.000. Caso contrario, imprime por consola el mensaje “Debe pagar impuesto”.

type errorPersonalizado struct {
	mensaje string
}

func (c *errorPersonalizado) Error() string {
	return fmt.Sprintf(c.mensaje)
}

func validarSalario(salary int) error {
	if salary < 150000 {
		return &errorPersonalizado{
			"error: el salario ingresado no alcanza el minimo no imponible",
		}
	}
	return nil
}

func main() {

	salary := 100000

	err := validarSalario(salary)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Debe pagar impuesto")
}
