package main

import "fmt"

func main() {
	var age, years_employed int
	var employed bool
	var salary float64

	age = 22
	employed = true
	years_employed = 1
	salary = 500000

	if (age < 22) {
		fmt.Println("Solo prestamos a mayores de 22 años")
	} else if (employed == false || years_employed < 1) {
		fmt.Println("Solo prestamos a empleados con al menos 1 año de antigüedad")
	} else if (salary > 100000) {
		fmt.Println("Prestamo otorgado con intereses")
	} else {
		fmt.Println("Prestamo otorgado sin intereses")
	}
}