/*Desarrolla las funciones necesarias para permitir a la empresa calcular:
Salario mensual de un trabajador según la cantidad de horas trabajadas.
La función recibirá las horas trabajadas en el mes y el valor de la hora como argumento.

Dicha función deberá retornar más de un valor (salario calculado y error).

En caso de que el salario mensual sea igual o superior a $150.000, se le deberá descontar el 10% en concepto de impuesto.
En caso de que la cantidad de horas mensuales ingresadas sea menor a 80 o un número negativo, la función debe devolver un error.
El mismo deberá indicar “error: el trabajador no puede haber trabajado menos de 80 hs mensuales”.

Calcular el medio aguinaldo correspondiente al trabajador

Fórmula de cálculo de aguinaldo:
[mejor salario del semestre] / 12 * [meses trabajados en el semestre].
La función que realice el cálculo deberá retornar más de un valor, incluyendo un error en caso de que se ingrese un número
negativo.

Desarrolla el código necesario para cumplir con las funcionalidades requeridas, utilizando “errors.New()”, “fmt.Errorf()” y
 “errors.Unwrap()”. No olvides realizar las validaciones de los retornos de error en tu función “main()”.

*/

package main

import (
	"errors"
	"fmt"
)

func main() {
	t1 := Trabajador{
		horasTrabajadas: 10,
		valorHora:       500,
	}

	salario, error := t1.SalarioMensual(t1.horasTrabajadas, t1.valorHora)

	if error != nil {
		fmt.Println(error)
	}
	fmt.Println(salario)

	error1 := errorEstructura{}
	otroError := fmt.Errorf("My error %w", error1)

	aguinaldo, error2 := t1.CalculoAguinaldo(500000.00, -3)

	if error2 != nil {
		fmt.Println(errors.Unwrap(otroError))
	}
	fmt.Println(aguinaldo)
}

type errorEstructura struct{}

func (e errorEstructura) Error() string {
	return "valores incorrectos"
}

type Trabajador struct {
	horasTrabajadas int
	valorHora       int
}

func (t1 *Trabajador) SalarioMensual(hs int, vH int) (float64, error) {
	totalSalario := hs * vH
	var totalDescuentos float64

	if hs < 80 || hs < 0 {
		err1 := errors.New("error el trabajador no puede haber trabajado menos de 80 hs mensuales")
		return 0.0, err1
	} else if totalSalario >= 150000 {
		totalDescuentos = float64(totalSalario) * 0.90
		return totalDescuentos, nil
	} else {
		return float64(totalSalario), nil
	}
}

func (t1 *Trabajador) CalculoAguinaldo(salario float64, meses int) (float64, error) {
	if meses > 0 {
		aguinaldo := salario / 12.0 * float64(meses)
		return aguinaldo, nil
	}
	return 0.0, fmt.Errorf("Se ingreso un valor negativo en los meses del aguinaldo")

}
