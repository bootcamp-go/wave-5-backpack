// Ejercicio 3 - Calcular salario
// Una empresa marinera necesita calcular el salario de sus empleados basándose en la cantidad de horas trabajadas por mes y la categoría.

// Si es categoría C, su salario es de $1.000 por hora
// Si es categoría B, su salario es de $1.500 por hora más un %20 de su salario mensual
// Si es de categoría A, su salario es de $3.000 por hora más un %50 de su salario mensual

// Se solicita generar una función que reciba por parámetro la cantidad de minutos trabajados por mes y la categoría, y que devuelva su salario.
package main
import "fmt"

func main()  {
	var palabra string = "papanato"

	fmt.Printf("tiene esta cantidad de letras:%d\n", len(palabra))
	for i, letra := range palabra{
		fmt.Printf("posicion: %x letra: %c \n", i, letra)
	}
}