package main

import "fmt"
//import "errors"


const ( 
	C = "C"
	B = "B"
	A = "A"
)

func salary(category string, minutes int)float64  {
	switch category{
	case C:
		return float64((1000 / 60) * minutes)
	case B:
		salary := float64((1500 / 60) * minutes)
		return  salary + (salary * 0.20)
	case A:
		salary := float64((3000 / 60) * minutes)
		return  salary + (salary * 0.50)
	}
	return 0
}

func main()  {
	salary := salary(A, 10800)
	fmt.Printf("Salary: %.2f", salary)
	fmt.Println()
}