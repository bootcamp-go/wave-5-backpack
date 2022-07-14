package main

import ("fmt"
		"errors")

func main()  {
	var salary int = 140000

	if salary < 150000 {
		fmt.Println(errors.New(fmt.Sprint("error: el salario ingresado: $", salary, " no alcanza el mínimo imponible")))
		return
	} else {
		fmt.Println("Debe pagar impuesto")
	}
}
