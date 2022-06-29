package main

/*
Teniendo en cuenta que si la persona gana más de $50.000 se
le descontará un 17% del sueldo y si gana más de $150.000 se
le descontará además un 10%.
*/
import "fmt"

func main() {
	var salary float32 = 0

	fmt.Println("------Ingrese el salario del empleado")
	fmt.Scanf("%v", &salary)

	var taxes int = calculateTaxes(salary)

	var totalOfTaxes = calculateTaxesbySalary(salary, taxes)

	fmt.Printf("Se aplicó un %v por ciento de taxes.\nLos impuestos a pagar por %v son %v pesos\n", taxes, salary, totalOfTaxes)
}

func calculateTaxes(salary float32) (taxes int) {
	taxes = 5
	if salary > 50000 {
		taxes += 17
	} else if salary > 150000 {
		taxes += 10
	}

	return
}

func calculateTaxesbySalary(salary float32, taxes int) (total float32) {
	total = (salary * float32(taxes)) / 100
	return
}
