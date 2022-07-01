package main

import "fmt"

type myCustomError struct {
	msg string
}

func (e *myCustomError) Error() string {
	return fmt.Sprintf("%s", e.msg)
}

func testSalary(salary int) (string, error) {
	if salary < 150000 {
		return "", &myCustomError{
			msg: "error :eres muy pobre",
		}

		//		"error: el salario ingresado no alcanza el mínimo imponible"
	}
	return "debe pagar impuesto", nil
}

func main() {
	var salary int
	fmt.Println("ingrese el salario")
	fmt.Scanln(&salary)

	msg, err := testSalary(salary)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s \n", msg)

}
