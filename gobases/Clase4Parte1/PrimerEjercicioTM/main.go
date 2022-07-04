package main

import "fmt"

type errorJ struct {
	msg string
}

func (e *errorJ) Error() string {
	return fmt.Sprintf(e.msg)
}

func testError(salary float64) (float64, error) {
	if salary < 150000 {
		return 0, &errorJ{
			msg: "error: el salario ingresado no alcanza el mínimo imponible",
		}
	}
	return salary, nil
}

func main() {
	salary := 50000
	_, err := testError(float64(salary))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Debe pagar impuestos")
}
