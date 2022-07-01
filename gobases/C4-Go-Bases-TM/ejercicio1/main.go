package main

import "fmt"

// Ejercicio 1 - Impuestos de salario #1

// En tu función “main”, define una variable llamada “salary” y asignarle un valor de tipo “int”.
// Crea un error personalizado con un struct que implemente “Error()” con el mensaje
// “error: el salario ingresado no alcanza el mínimo imponible" y lánzalo en caso de que “salary”
// sea menor a 150.000. Caso contrario, imprime por consola el mensaje “Debe pagar impuesto”.

type errorPersonalizado struct {
	msg string
}

func (e *errorPersonalizado) Error() string {
	return fmt.Sprintf("%s", e.msg)
}

func validarSalario(salario int) (string, error) {
	if salario < 150000 {
		return "", &errorPersonalizado{msg: "error: el salario ingresado no alcanza el mínimo imponible"}
	}
	return "Debe pagar impuesto", nil
}

func main() {
	fmt.Println("Ejercicio 1 - Impuestos de salario #1")
	fmt.Println("")

	salary := 1500

	valid, err := validarSalario(salary)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(valid)
}
