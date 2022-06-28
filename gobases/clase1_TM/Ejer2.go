package main

import "fmt"

func main(){
	var(
		Age =25
		salary = 110000
		employeeTime = 3
	)

	fmt.Println("ingrese la edad del Cliente:")
	fmt.Scanln(&Age)
	fmt.Println("ingrese el Salario del Cliente:")
	fmt.Scanln(&salary)
	fmt.Println("ingrese la antiguedad del Cliente en años:")
	fmt.Scanln(&employeeTime)

	switch  {
	case Age>22 && employeeTime>1 && salary>100000:
		fmt.Println("puedes tomar el prestamo sin interes")
	case Age>22 && employeeTime>1 && salary<100000:
		fmt.Println("puedes tomar el credito con interes")
	case Age<=22 && employeeTime <=1:
		fmt.Println("no puedes acceder al credito , no cumples con los requerimientos")

	}
}