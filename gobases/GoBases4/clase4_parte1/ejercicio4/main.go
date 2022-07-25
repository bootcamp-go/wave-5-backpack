package main

import (
	"fmt"
	"os"
)

func main() {

	res, err := salarioMesHoras(80, 10000)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Su salario neto es de %0.f \n", res)

	res2, err2 := aguinaldo(100000, 1)

	if err2 != nil {
		fmt.Println(err2)
		os.Exit(1)
	}
	fmt.Printf("Su aguinaldo es %0.f \n", res2)

}

type datos struct {
	horas     float64
	valorHora float64
	msg       string
}

func (d *datos) Error() string {
	return fmt.Sprintf("%s", d.msg)
}

func salarioMesHoras(hoursWorked, valueHour float64) (float64, error) {

	var salaryMonth float64 = hoursWorked * valueHour

	if salaryMonth >= 150000 && hoursWorked < 80 {
		discount := 0
		return float64(discount), &datos{
			horas:     hoursWorked,
			valorHora: valueHour,
			msg:       "error: el trabajador no puede haber trabajado menos de 80 hs mensuales",
		}
	} else {
		if salaryMonth >= 150000 {
			discount := salaryMonth * (1 - 0.10)
			return discount, nil
		}
	}

	return salaryMonth, nil
}

func aguinaldo(bestSalary, monthWorked float64) (float64, error) {

	if bestSalary < 0 || monthWorked < 0 {
		return 0, &datos{
			msg: "ingreso un mes o salario negativo",
		}
	}

	result := bestSalary / 12 * monthWorked
	return result, nil

}
