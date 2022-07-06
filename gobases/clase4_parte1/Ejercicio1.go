package main

import (
	"fmt"
	"os"
)

type errorCustom struct {
	mensaje string
}

func (e *errorCustom) Error() string {
	return fmt.Sprintf("%v", e.mensaje)
}
func myErrorTest(valor int) (string, error) {
	if valor < 150000 {
		return "Valor inválido", &errorCustom{
			mensaje: "error: el salario ingresado no alcanza el mínimo imponible",
		}

	}
	return "Debe pagar impuestos", nil
}

func main() {
	salary := 200000
	response, err := myErrorTest(salary)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(response)
}
