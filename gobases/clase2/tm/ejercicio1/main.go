package main

import "fmt"

// Una empresa de chocolates necesita calcular el impuesto de sus empleados al momento de depositar el sueldo,
// para cumplir el objetivo es necesario crear una función que devuelva el impuesto de un salario.
// Teniendo en cuenta que si la persona gana más de $50.000 se le descontará un 17% del sueldo
// y si gana más de $150.000 se le descontará además un 10%.


func main() {
  sueldo, sueldo2, sueldo3 := 40000.0,55000.0,178000.0

  desc, total := calcularImpuesto(sueldo)
  fmt.Printf("Sueldo de %v\nSe descuenta: %v\nTotal: %v\n",sueldo, desc, total)

  desc2, total2 := calcularImpuesto(sueldo2)
  fmt.Printf("Sueldo de %v\nSe descuenta: %v\nTotal: %v\n",sueldo2, desc2, total2)

  desc3, total3 := calcularImpuesto(sueldo3)
  fmt.Printf("Sueldo de %v\nSe descuenta: %v\nTotal: %v\n",sueldo3, desc3, total3)
}

// Se debe ingresar el sueldo y retorna el monto a descontar y el sueldo total (aplicando el impuesto)
func calcularImpuesto(sueldo float64) (float64, float64) {
  var impuesto, descontado, total float64

  if sueldo > 50000 {
    impuesto = 0.17
  }

  if sueldo > 150000 {
    impuesto = 0.27
  }

  descontado = sueldo * impuesto
  total = sueldo - descontado

  return descontado, total
}
