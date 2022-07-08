package main

import (
	"fmt"
)

const (
	catA      = "A"
	catB      = "B"
	catC      = "C"
	minimo    = "minimo"
	promedio4 = "promedio"
	maximo    = "maximo"
)

func main() {

	//	Ejercicio 1
	// Una empresa de chocolates necesita calcular el impuesto de sus empleados al momento de depositar el sueldo, para cumplir el objetivo es necesario crear una función que devuelva el impuesto de un salario.
	// Teniendo en cuenta que si la persona gana más de $50.000 se le descontará un 17% del sueldo y si gana más de $150.000 se le descontará además un 10%.

	fmt.Println("**** Ejercicio 1 ****")
	fmt.Println("El descuento es:", CalcularImpuesto(160000.00))

	//Ejercicio 2
	// Un colegio necesita calcular el promedio (por alumno) de sus calificaciones.
	// Se solicita generar una función en la cual se le pueda pasar N cantidad de enteros y devuelva el promedio y un error en caso que uno de los números ingresados sea negativo

	promedio, err := CalcularPromedioNotas(4.0, -5.0, 4.8, 7.8)

	fmt.Println("**** Ejercicio 2 ****")
	if err != nil {
		fmt.Printf(err.Error())
	} else {
		fmt.Println("El promedio de notas es:", promedio)
	}

	//Ejercicio 3
	// Una empresa marinera necesita calcular el salario de sus empleados basándose en la cantidad de horas trabajadas por mes y la categoría.

	// Si es categoría C, su salario es de $1.000 por hora
	// Si es categoría B, su salario es de $1.500 por hora más un %20 de su salario mensual
	// Si es de categoría A, su salario es de $3.000 por hora más un %50 de su salario mensual

	// Se solicita generar una función que reciba por parámetro la cantidad de minutos trabajados por mes y la categoría, y que devuelva su salario.

	salario, err := CalcularSalario(50000, catC)
	fmt.Println("**** Ejercicio 3 ****")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("El salario es de:", salario)
	}

	//Ejercicio 4
	// Los profesores de una universidad de Colombia
	// necesitan calcular algunas estadísticas de calificaciones de los alumnos de un curso, requiriendo calcular los valores mínimo, máximo y promedio de sus calificaciones.

	// Se solicita generar una función que indique qué tipo de cálculo se quiere realizar (mínimo, máximo o promedio) y que devuelva otra función ( y un mensaje en caso que el cálculo no esté definido) que se le puede pasar una cantidad N de enteros y devuelva el cálculo que se indicó en la función anterior

	fmt.Println("**** Ejercicio 4 ****")
	operacionMin := Operaciones(minimo)
	operacionMax := Operaciones(maximo)
	operacionProm := Operaciones(promedio4)

	min := operacionMin(2.0, 1.0, 5.0, 9.0)
	max := operacionMax(2.0, 1.0, 5.0, 9.0)
	prom := operacionProm(2.0, 1.0, 5.0, 9.0)

	fmt.Println("El minimo es:", min)
	fmt.Println("El maximo es:", max)
	fmt.Println("El promedio es:", prom)

}

// Funciones Auxiliares

// AUX ejercicio 1
func CalcularImpuesto(sueldo float64) float64 {
	var descuento float64

	if sueldo > 50000 && sueldo < 150000 {
		descuento = sueldo * 0.17
	} else {
		descuento = sueldo * 0.27
	}

	return descuento
}

// AUX ejercicio 2
func CalcularPromedioNotas(notas ...float32) (float32, error) {
	resultado, hayNegativo := SumaNotas(notas...)
	if hayNegativo {
		return 0, fmt.Errorf("Error hay un numero negativo\n")
	}

	resultado = resultado / float32(len(notas))
	return resultado, nil
}
func SumaNotas(notas ...float32) (float32, bool) {
	var resultado float32
	var hayNegativo bool = false
	for _, nota := range notas {
		if nota < 0 {
			hayNegativo = true
		}
		resultado += nota
	}
	return resultado, hayNegativo

}

// AUX ejercicio 3
func CalcularSalario(minutosTrabajados float64, categoria string) (float64, error) {
	if minutosTrabajados == 0 {
		return 0, fmt.Errorf("El trabajador no tiene minutos asignados")
	}

	horasTrabajadas := minutosTrabajados / 60
	var salario float64
	var resultado float64
	switch categoria {
	case "A":
		resultado = horasTrabajadas * 1000
		break
	case "B":
		salario = horasTrabajadas * 1500
		resultado = salario + (salario * 0.20)
		break
	case "C":
		salario = horasTrabajadas * 3000
		resultado = salario + (salario * 0.50)
		break
	}

	return resultado, nil

}

// AUX ejercicio 4
func Operaciones(operacion string) func(notas ...float32) float32 {
	switch operacion {
	case "minimo":
		return opMinimo
	case "maximo":
		return opMaximo
	case "promedio":
		return opPromedio
	}

	return nil
}

func opMinimo(notas ...float32) float32 {
	var minimo float32 = notas[0]
	for _, nota := range notas {
		if minimo > nota {
			minimo = nota
		}
	}
	return minimo
}

func opMaximo(notas ...float32) float32 {
	var maximo float32 = notas[0]
	for _, nota := range notas {
		if maximo < nota {
			maximo = nota
		}
	}
	return maximo
}

func opPromedio(notas ...float32) float32 {
	var sumaNotas float32
	var resultado float32
	for _, nota := range notas {
		sumaNotas += nota
	}
	resultado = sumaNotas / float32(len(notas))

	return resultado
}
