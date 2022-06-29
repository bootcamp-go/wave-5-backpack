package main

import (
	"errors"
	"fmt"
)

const (
	MINIMUM = "minimum"
	AVERAGE = "average"
	MAXIMUM = "maximum"
)

const (
	PERRO     = 10000
	GATO      = 5000
	HAMSTER   = 250
	TARANTULA = 150
)

func impuestosDeSalario(salario float32) float32 {
	//Ejercicio 1
	var impuesto float32

	if salario > 50000 && salario < 150000 {
		impuesto = (salario * 17) / 100
	} else if salario > 150000 {
		impuesto = (salario * 10) / 100
	} else {
		impuesto = 0
	}

	return impuesto
}

func calcularPromedio(calificaciones ...int) (float64, error) {
	//Ejercicio 2
	contNotas := 0
	sumaNotas := 0
	for _, calificacion := range calificaciones {

		if calificacion < 0 {
			return 0, errors.New("La calificacion no puede ser negativa")
		}
		sumaNotas += calificacion
		contNotas++
	}

	promedio := sumaNotas / contNotas

	return float64(promedio), nil
}

func calcularSalario(mPorMes float32, categoria string) float32 {
	//Ejercicio 3
	hPorMes := mPorMes / 60
	var salario float32

	switch categoria {
	case "C":
		salario = hPorMes * 1000
	case "B":
		salario = (hPorMes * 1500) + (hPorMes*1500)*20/100
	case "A":
		salario = (hPorMes * 3000) + (hPorMes*3000)*50/100
	}

	return salario
}

func calcularEstadisticas(calculo string) {
	//Ejercicio 4
	minValue, err := estadistica(calculo)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s nota: %.2f", calculo, minValue(1, 2, 3, 4, 5))
	}
}

func estadistica(operation string) (func(notas ...int) float32, error) {
	//Ejercicio 4
	switch operation {
	case MINIMUM:
		return minFunc, nil
	case AVERAGE:
		return aveFunc, nil
	case MAXIMUM:
		return maxFunc, nil
	default:
		return nil, errors.New("Calculo no definido")
	}
}

func minFunc(notas ...int) float32 {
	//Ejercicio 4
	var min int

	for i, nota := range notas {
		if i == 0 {
			min = nota
		} else if nota < min {
			min = nota
		}
	}
	return float32(min)
}

func aveFunc(notas ...int) float32 {
	//Ejercicio 4
	countNotas := 0
	sumNotas := 0

	for _, nota := range notas {
		sumNotas += nota
		countNotas++
	}
	return float32(sumNotas / countNotas)
}

func maxFunc(notas ...int) float32 {
	//Ejercicio 4
	var max int

	for i, nota := range notas {
		if i == 0 {
			max = nota
		} else if nota > max {
			max = nota
		}
	}
	return float32(max)
}

func calcularCantidadDeAlimento(animal string) {
	//Ejercicio 5
	alimentoAnimal, err := cantidadAlimento(animal)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Cantidad alimento para %s: %.2f grs.", animal, alimentoAnimal(2))
	}
}

func cantidadAlimento(animal string) (func(cantidadAnimal int) float64, error) {
	//Ejercicio 5
	switch animal {
	case "perro":
		return alimentoPerro, nil
	case "gato":
		return alimentoGato, nil
	case "hamster":
		return alimentoHamster, nil
	case "tarantula":
		return alimentoTarantula, nil
	default:
		return nil, errors.New("El animal no existe")

	}
}

func alimentoPerro(cantidadPerros int) float64 {
	return float64(PERRO * cantidadPerros)
}

func alimentoGato(cantidadGatos int) float64 {
	return float64(GATO * cantidadGatos)
}

func alimentoHamster(cantidadHamster int) float64 {
	return float64(HAMSTER * cantidadHamster)
}

func alimentoTarantula(cantidadTarantulas int) float64 {
	return float64(TARANTULA * cantidadTarantulas)
}
