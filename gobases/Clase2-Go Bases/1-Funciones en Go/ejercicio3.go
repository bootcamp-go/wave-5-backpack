package main

import "fmt"

func main(){
	salario:=calcularSalario("C",60)
	
		fmt.Printf("El salario es %v \n",salario)
	
	 
}

func calcularSalario(categoria string, minutosTrabajados float64) (float64){
	var horasTrabajadas float64=minutosTrabajados/60
	var salario=0.0
	switch categoria{
	case "A":
		salario=horasTrabajadas*1000
	case "B":
		salario=horasTrabajadas*1500*1.2
	case "C":
		salario=horasTrabajadas*3000*1.5
	default:
		salario=0.0
	}
	
	return salario
} 


/*
Ejercicio 3 - Calcular salario
Una empresa marinera necesita calcular el salario de sus empleados basándose en la cantidad de horas trabajadas por mes y la categoría.

Si es categoría C, su salario es de $1.000 por hora
Si es categoría B, su salario es de $1.500 por hora más un %20 de su salario mensual
Si es de categoría A, su salario es de $3.000 por hora más un %50 de su salario mensual

Se solicita generar una función que reciba por parámetro la cantidad de minutos trabajados por mes y la categoría, y que devuelva su salario.

*/