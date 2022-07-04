package main

import (
	"fmt"
)

func main() {
  paga := 160000 
  noPaga := 100000

  err := checkImpuesto(paga)
  if err != nil {
  	fmt.Println(err)
  } else {
  	fmt.Println("Debe pagar impuesto")
  }

  err = checkImpuesto(noPaga)
  if err != nil {
  	fmt.Println(err)
  } else {
  	fmt.Println("Debe pagar impuesto")
  }
}

func checkImpuesto(salary int) error {
  if salary < 150000 {
  	return fmt.Errorf("error: el mínimo imponible es de 150000 y el salario ingresado es de: %v", salary)
  }

	return nil
}
