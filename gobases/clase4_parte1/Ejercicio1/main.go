package main

import (
	"fmt"
	"os"
)

// sólo se requiere crear un tipo que implemente el método Error()
type myCustomError struct {
	salario int
	msg     string
}

// hacemos que nuestro tipo struct implemente el método Error()
func (e *myCustomError) Error() string {
	return fmt.Sprintf("%d - %v", e.salario, e.msg)
}

func myCustomErrorTest(salario int) (int, error) {
	if salario < 150000 {
		return 500, &myCustomError{
			salario: salario,
			msg:     "error: el salario ingresado no alcanza el mínimo imponible",
		}
	}
	return 200, nil
}

func main() {
	salary := 150000
	_, err := myCustomErrorTest(salary) // 1llamamosanuestra func
	if err != nil {                     // hacemos una validación del valor de err
		fmt.Println(err) // si err no es nil,imprimimos el errory...
		os.Exit(1)       // utilizamos este método para salir del programa
	}
	fmt.Printf("La persona debe pagar impuesto \n")
}
