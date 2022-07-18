package main

import "fmt"

//Ejercicio 3 - Impuesto de salario #3

func main()  {

	salary := 149000

	err := fmt.Errorf("error: el minimo imponible es de 150000 y el salario ingresado es de %d", salary)
	fmt.Println(err)
	
}