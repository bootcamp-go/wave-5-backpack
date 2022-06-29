package main

import "fmt"

/*
Una empresa de chocolates necesita calcular el impuesto de sus empleados
al momento de depositar el sueldo, para cumplir el objetivo es necesario
crear una función que devuelva el impuesto de un salario.
Teniendo en cuenta que si la persona gana más de $50.000
se le descontará un 17% del sueldo y si gana más de $150.000 se
le descontará además un 10%.
*/

const (
	IMPUESTO_BASE     = 17
	IMPUESTO_AGREGADO = 10
)

func calcularImpuesto(salario int, porcentage int) int {
	return int(float32(salario) * (float32(porcentage) / float32(100)))
}

func impuestoSalario(salario int) int {
	if salario > 150000 {
		return calcularImpuesto(salario, IMPUESTO_BASE+IMPUESTO_AGREGADO)
	}
	if salario > 50000 {
		return calcularImpuesto(salario, IMPUESTO_BASE)
	}

	return 0
}

func impuestoSalarios(salarios ...int) {
	for _, salario := range salarios {
		impuesto := impuestoSalario(salario)
		fmt.Printf("Salario base: %v, Impuesto: %v, Total: %v \n",
			salario,
			impuesto,
			salario-impuesto,
		)
	}
}

func main() {
	impuestoSalarios(200000, 100000, 40000)
}
