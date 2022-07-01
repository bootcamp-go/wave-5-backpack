package main

import (
	"errors"
	"fmt"
)

// Haz lo mismo que en el ejercicio anterior pero reformulando el código para que,
// en reemplazo de “Error()”,  se implemente “errors.New()”.

func main() {

	salary := 100000

	if salary < 150000 {
		fmt.Println(errors.New("Salario no alcanzado por el impuesto"))
		return
	}
	fmt.Println("Debe pagar impuesto")
}
