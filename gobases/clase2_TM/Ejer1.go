package main

import "fmt"

func main(){
		var salary float64
		fmt.Println("introduzca el salario del empleado, para calcular el monto de impuestos  $xx.x ")
		fmt.Scanln(&salary)

		fmt.Printf("los impuestos son %v \n",taxesEmployee(salary))
}

func taxesEmployee(salary float64) float64{
	var tax float64 
	if salary<=500000{
		tax = salary*0.20
	}else if salary>50000 && salary<=150000{
		tax = salary*0.17
	}else if salary>150000{
		tax = salary*0.10
	}
	return tax
}