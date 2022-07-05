package main

import "fmt"

var ErrorSalario = &myError{msg: "error: el salario ingresado no alcanza el mínimo imponible"}

type myError struct{
msg string
}

func (e *myError) Error() string {
	return e.msg
}

func calculoSalario(i int) (string, error)  {
	if i <  150000{
		return "", ErrorSalario
	}
	return "debe pagar impuesto", nil
}

func main(){

	sueldo := 149000

	res, err := calculoSalario(sueldo)

	if err != nil {
		fmt.Println(err)
		return
		
	}

	fmt.Printf(res)

}