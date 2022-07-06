// Ejercicio 2 - Impuestos de salario #2

// Haz lo mismo que en el ejercicio anterior pero reformulando el código para que, en reemplazo de “Error()”,  se implemente “errors.New()”.

package main

import (
	"errors"
	"fmt"
)

func main() {
	var wage int = 100000
	if wage < 150000 {
		fmt.Println(errors.New("error: no aplica impuestos"))
	} else {
		fmt.Println("Si paga impuestos")
	}
}