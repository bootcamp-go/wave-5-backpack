package main

import "fmt"

type myCustomError struct {
	status int
	msg    string
}

func (e *myCustomError) Error() string {
	return fmt.Sprintf("%d - %v \n", e.status, e.msg)
}
func myCustomErrorTest(status int) (int, error) {
	if status <= 150000 {
		return 0, &myCustomError{
			status: status,
			msg:    "error: el salario ingresado no alcanza el mínimo imponible",
		}
	}
	return 1, nil
}
func main() {
	var salary int = 150001
	_, err := myCustomErrorTest((salary))
	if err != nil {
		fmt.Printf("%s", err)
		return
	}
	fmt.Println("Debe pagar impuesto.")
}

// En tu función “main”, define una variable llamada “salary” y asignarle un valor de tipo “int”.
// Crea un error personalizado con un struct que implemente “Error()” con el mensaje “error: el salario ingresado no alcanza el mínimo imponible"
//y lánzalo en caso de que “salary” sea menor a 150.000. Caso contrario, imprime por consola el mensaje “Debe pagar impuesto”.
