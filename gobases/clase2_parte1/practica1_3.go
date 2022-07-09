package main

import "fmt"

//Ejercicio 3 - Calcular salario
func salario(categoria string, minTrabajados float64) float64 {

  //Salario por hora
  var sueldo float64

  if categoria == "C" {

    sueldo = minTrabajados * 16.67
    return sueldo

  } else if categoria == "B" {

    sueldo = minTrabajados * 25
    sueldo += sueldo * 20 / 100
    return sueldo

  } else if categoria == "A" {

    sueldo = minTrabajados * 50
    sueldo += sueldo * 50 / 100
    return sueldo

  }

  return sueldo
}

func main() {

  resultado := salario("A", 60)
  fmt.Println(resultado)

}