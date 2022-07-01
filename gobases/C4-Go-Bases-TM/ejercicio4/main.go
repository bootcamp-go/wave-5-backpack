package main

import (
	"errors"
	"fmt"
	"strings"
)

// Bonus Track -  Impuestos de salario #4

// Vamos a hacer que nuestro programa sea un poco más complejo y útil.

// 1. Desarrolla las funciones necesarias para permitir a la empresa calcular:
//     a) Salario mensual de un trabajador según la cantidad de horas trabajadas.
//          - La función recibirá las horas trabajadas en el mes y el valor de la hora como argumento.
//          - Dicha función deberá retornar más de un valor (salario calculado y error).
//          - En caso de que el salario mensual sea igual o superior a $150.000, se le deberá descontar el 10% en concepto de impuesto.
//          - En caso de que la cantidad de horas mensuales ingresadas sea menor a 80 o un número negativo,
//            la función debe devolver un error. El mismo deberá indicar “error: el trabajador no puede haber trabajado menos de 80 hs mensuales”.

//     b) Calcular el medio aguinaldo correspondiente al trabajador
//          - Fórmula de cálculo de aguinaldo: [mejor salario del semestre] / 12 * [meses trabajados en el semestre].
//          - La función que realice el cálculo deberá retornar más de un valor, incluyendo un error en caso de que se ingrese un número negativo.
// 2. Desarrolla el código necesario para cumplir con las funcionalidades requeridas, utilizando “errors.New()”, “fmt.Errorf()” y “errors.Unwrap()”.
//    No olvides realizar las validaciones de los retornos de error en tu función “main()”.

type errorGenerico struct{}

func (e errorGenerico) Error() string {
	return "no debe ser negativo"
}

func calcularSalario(horasTrabajadas float64, valorHora float64) (float64, error) {
	if horasTrabajadas < 80 {
		return 0.0, errors.New("error: el trabajador no puede haber trabajado menos de 80 hs mensuales")
	}

	salario := horasTrabajadas * valorHora
	if salario >= 150000 {
		return salario - (salario * 0.1), nil
	}
	return salario, nil
}

// Función para dar formato a moneda
func formatearMoneda(m float64) string {
	// Formateamos la cantidad a string
	money := fmt.Sprintf("%.2f", m)
	// Separamos la cantidad de su decimal
	moneyElements := strings.Split(money, ".")
	// Invertimos la cantidad
	moneyInverted := ""
	for _, v := range moneyElements[0] {
		moneyInverted = string(v) + moneyInverted
	}
	// Reinvertimos la cantidad y agregamos las comas
	moneyValid := ""
	for i, v := range moneyInverted {
		if (i+1)%3 == 0 && (i+1) != len(moneyInverted) {
			moneyValid = "," + string(v) + moneyValid
		} else {
			moneyValid = string(v) + moneyValid
		}
	}
	// Regresamos el resultado
	return "$" + moneyValid + "." + moneyElements[1]
}

func calcularAguinaldo(salarios ...float64) (float64, error) {
	mesesTrabajados := 0
	mejorSalario := 0.0
	firstElement := true
	for _, v := range salarios {
		if v < 0.0 {
			eg := errorGenerico{}
			return 0.0, fmt.Errorf("El salario %.2f %w", v, eg)
		}
		mesesTrabajados++
		if firstElement {
			// Obtenemos el primer salario de referencia
			mejorSalario = v
			firstElement = false
		} else if v > mejorSalario {
			// Validamos los demás salarios
			mejorSalario = v
		}
	}

	// Se calcula aguinaldo
	aguinaldo := mejorSalario / (float64(12 * mesesTrabajados))
	return aguinaldo, nil
}

func main() {
	fmt.Println("Bonus Track -  Impuestos de salario #4")
	fmt.Println("")

	horasTrabajadas := 100.0

	salario, err := calcularSalario(horasTrabajadas, 1000.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("El salario es de: %s\n", formatearMoneda(salario))

	aguinaldo, err := calcularAguinaldo(100.0, 120.0, 140.0, 110.0, 130.0, 150.0, -10.0)
	if err != nil {
		fmt.Println(errors.Unwrap(err), "-> Detalles del error: ", err)
		return
	}
	fmt.Printf("El aguinaldo es de %s\n", formatearMoneda(aguinaldo))
}
