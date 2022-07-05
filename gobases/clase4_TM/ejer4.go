package main

import (
	"errors"
	"fmt"
)

func testSalary(salary int) (string, error) {
	if salary < 150000 {
		return "", fmt.Errorf("error: eres muy pobre")
		//		"error: el salario ingresado no alcanza el mínimo imponible"
	}
	return "debe pagar impuesto", nil
}

type employeeMonth struct {
	hoursWorked int
	price       float64
}

func main() {

	worksemester := []employeeMonth{{100, 12.3}, {200, 15.0}}

	// var hours int
	// fmt.Println("ingrese el salario")
	// fmt.Scanln(&hours)

	bonus, err := calcBonusPerSemester(worksemester)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%f \n", bonus)

}

func calcSalary(hours int, priceHour float64) (float64, error) {

	if hours < 80 {
		return 0, errors.New("error : no puede trabajar menos de 80 horas")
	}
	salary := float64(hours) * priceHour
	if salary > 150000.0 {
		salary *= 0.8
	}

	return salary, nil
}

func calcBonusPerSemester(s []employeeMonth) (float64, error) {
	var maxSalary float64
	for i, _ := range s { // i assume that's the slice contain less tha 6 months , if theres more information should evaluate len<6 of slice
		monthSalary, err := calcSalary(s[i].hoursWorked, s[i].price)
		if err != nil {
			return 0.0, err
		}
		if monthSalary > maxSalary {
			maxSalary = monthSalary
		}
	}

	return maxSalary / 12, nil
}
