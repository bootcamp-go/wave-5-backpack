package main

import "fmt"

func main() {
	fmt.Println(impuestosAlSalario(55000))
	fmt.Println(impuestosAlSalario(165000))
}

func impuestosAlSalario(salario float64) (salarioRestante, deducciones float64) {
	salarioRestante = salario
	deducciones = 0
	if salario <= 50000 {
		return salarioRestante, deducciones
	} else {
		nivel1 := salario - 50000
		if nivel1 >= 100000 {
			nivel1 = 100000
			nivel2 := (salario - 150000) * 0.27
			salarioRestante -= nivel2
			deducciones += nivel2
		}
		nivel1 *= 0.17
		salarioRestante -= nivel1
		deducciones += nivel1
	}
	return salarioRestante, deducciones
}
