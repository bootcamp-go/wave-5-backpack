package main

import "fmt"


type MyCustomError struct {
	msg string
}


func main()  {
	var salary int = 2100000
	err := MyCustomErrorFunc(salary)
	if salary < 150000 {
		fmt.Println(err)
	} else {
		fmt.Println("Debe pagar impuesto")
	}

}

func (e *MyCustomError) Error() string {
	return fmt.Sprintf("%v", e.msg)
}

func MyCustomErrorFunc(salary int) error {
	return &MyCustomError{
		msg: fmt.Sprint("error: el salario ingresado: $", salary, " no alcanza el salario minimo"),
	}
}
