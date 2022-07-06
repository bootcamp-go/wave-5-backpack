package main

import (
	"errors"
	"fmt"
)

func main() {
	var salary int = 150001
	if salary <= 150000 {
		fmt.Println(errors.New("error: el salario ingresado no alcanza el mínimo imponible"))
		return
	}
	fmt.Println("Debe pagar impuesto.")
}

// Haz lo mismo que en el ejercicio anterior pero reformulando el código para que, en reemplazo de “Error()”,  se implemente “errors.New()”.
