/*
Ejercicio 3 - Calcular salario
Una empresa marinera necesita calcular el salario de sus empleados basándose en la cantidad de horas trabajadas por mes y la categoría.

Si es categoría C, su salario es de $1.000 por hora
Si es categoría B, su salario es de $1.500 por hora más un %20 de su salario mensual
Si es de categoría A, su salario es de $3.000 por hora más un %50 de su salario mensual

Se solicita generar una función que reciba por parámetro la cantidad de minutos trabajados por mes y la categoría, y que devuelva su salario.


*/

/*
package main


import (
	"fmt"
)



func main() {

	fmt.Printf("El salario es: %f\n", calcularSalario(120,"A"))
	// se espera 9.000
	fmt.Printf("El salario es: %f \n" , calcularSalario(120,"B"))
	// se espera 3.600
	fmt.Printf("El salario es: %f \n", calcularSalario(120,"C"))
	// se espera 2.000
}

func calcularSalario (minutosTrabajados int, categoria string) float64 {

	var horasTrabajadas float64 = float64(minutosTrabajados) / 60
	var salario float64 = 0.0
	switch categoria {
    case "A":
		//Si es de categoría A, su salario es de $3.000 por hora más un %50 de su salario mensual
        salario = (3.000 * horasTrabajadas) * 1.5
    case "B":
		// Si es categoría B, su salario es de $1.500 por hora más un %20 de su salario mensual
		salario = (1.500 * horasTrabajadas) * 1.2
    case "C":
        salario = (1.000 * horasTrabajadas)   
    }

	return salario
}

*/