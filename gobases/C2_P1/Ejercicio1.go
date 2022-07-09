package main

import (
	"fmt"
)

//Una empresa de chocolates necesita calcular el impuesto de sus empleados al momento de depositar el sueldo, para cumplir el objetivo es necesario crear una función que devuelva el impuesto de un salario.
//Teniendo en cuenta que si la persona gana más de $50.000 se le descontará un 17% del sueldo y si gana más de $150.000 se le descontará además un 10%.

func salary(salary float64) float64 {
	if salary > 50000 && salary < 150000 {
		var desc float64 = salary * 0.17
		return desc
	} else if salary > 150000 {
		var desc2 float64 = salary * 0.27
		return desc2
	}
	return 0
}

func main() {
	result := salary(155000)
	fmt.Println(result)
}
