package main

import "fmt"

//Ejercicio 1 - Impuestos de salario
func impuestoDeSalario(salario float64) float64 {

  var impuesto float64

  if salario > 50000 && salario <= 150000 {

    impuesto = salario * 17 / 100
    return impuesto

  } else if salario > 150000 {

    impuesto = salario * 27 / 100
    return impuesto

  }

  return 0

}

func main() {

  sueldoTotal := impuestoDeSalario(200000.90)
  fmt.Println(sueldoTotal)

}