package main

import "fmt"

type myError struct {
	msg string
}

func (e *myError) Error() string {
	return fmt.Sprintf("%v", e.msg)
}

func testSalary(salary int) error {
	if salary < 150000 {
		return &myError{"error: el salario ingresado no alcanza el mínimo imponible"}
	}
	return nil

}

func main() {
	salary := 5000000

	err := testSalary(salary)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Debe pagar impuestos")
}
