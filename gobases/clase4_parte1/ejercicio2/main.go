package main

import (
	"errors"
	"fmt"
)


func calculoSalario(i int) (string, error)  {
	if i <  150000{
		return "", errors.New("error: el salario ingresado no alcanza el mínimo imponible")
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