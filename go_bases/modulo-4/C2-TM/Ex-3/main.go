package main

import "fmt"

type Empleado struct {
	catEmpleado    string
	minutosTrabajo int
}

func main() {
	var employee = Empleado{catEmpleado: "C", minutosTrabajo: 120}

	salario := calcSalario(employee)
	fmt.Println(salario)

}

func calcSalario(e Empleado) float32 {
	var horas float32
	var salario float32

	horas = float32(e.minutosTrabajo) / 60

	switch e.catEmpleado {
	case "A":
		salario = horas * 1000
	case "B":
		salario = horas * 1500
		salario += +(salario * .20)
	case "C":
		salario = horas * 3000
		salario += +(salario * .50)
	}
	return salario
}
