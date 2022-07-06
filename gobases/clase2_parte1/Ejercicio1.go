package main
import "fmt"

func main(){
	fmt.Printf("El impuesto a retener es de $%.2f \n", taxCalc(150001))
}

func taxCalc(salary float64) float64 {
	var tax float64 = 0.0
	if salary > 50000 && salary < 150000 {
		tax = salary * 0.17
	} else if salary > 150000 {
		tax = salary * 0.27
	} else {
		tax = 0
	}

	return tax
}

// Una empresa de chocolates necesita calcular el impuesto de sus empleados al momento de depositar el sueldo,
// para cumplir el objetivo es necesario crear una función que devuelva el impuesto de un salario. 
// Teniendo en cuenta que si la persona gana más de $50.000 se le descontará un 17% del sueldo y si gana más de $150.000 se le descontará
// además un 10%.
