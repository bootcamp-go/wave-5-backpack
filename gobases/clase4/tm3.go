// Ejercicio 3 - Impuestos de salario #3
// Repite el proceso anterior, pero ahora implementando “fmt.Errorf()”, para que el mensaje de error reciba por parámetro el valor de “salary” indicando que no alcanza el mínimo imponible (el mensaje mostrado por consola deberá decir: “error: el mínimo imponible es de 150.000 y el salario ingresado es de: [salary]”, siendo [salary] el valor de tipo int pasado por parámetro).

// Bonus Track -  Impuestos de salario #4
// Vamos a hacer que nuestro programa sea un poco más complejo.
// Desarrolla las funciones necesarias para permitir a la empresa calcular:
// Salario mensual de un trabajador según la cantidad de horas trabajadas.
// La función recibirá las horas trabajadas en el mes y el valor de la hora como argumento.
// Dicha función deberá retornar más de un valor (salario calculado y error).
// En caso de que el salario mensual sea igual o superior a $150.000, se le deberá descontar el 10% en concepto de impuesto.
// En caso de que la cantidad de horas mensuales ingresadas sea menor a 80 o un número negativo, la función debe devolver un error. El mismo deberá indicar “error: el trabajador no puede haber trabajado menos de 80 hs mensuales”.
// Calcular el medio aguinaldo correspondiente al trabajador
// Fórmula de cálculo de aguinaldo:
// [mejor salario del semestre] / 12 * [meses trabajados en el semestre].
// La función que realice el cálculo deberá retornar más de un valor, incluyendo un error en caso de que se ingrese un número negativo.

// Desarrolla el código necesario para cumplir con las funcionalidades requeridas, utilizando “errors.New()”, “fmt.Errorf()” y “errors.Unwrap()”. No olvides realizar las validaciones de los retornos de error en tu función “main()”.

package main

import "fmt"

func main() {
	var wage int = 200000
	if wage < 150000 {
		err := fmt.Errorf("ERROR: es menor a 150000, el importe es %d", wage)
		fmt.Println(err)
	} else {
		fmt.Println("Si paga impuestos")
	}
}
