package main

import (
	"fmt"
)

func main() {
  noPaga := 100000

  err := checkImpuesto(noPaga)
  fmt.Printf("analizando: %v\n", noPaga)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println("Debe pagar impuesto")
  }

  paga := 160000

  err = checkImpuesto(paga)
  fmt.Printf("analizando: %v\n", paga)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println("Debe pagar impuesto")
  }
}

type errorImpuesto struct {
}

func (e *errorImpuesto) Error() string {
  return fmt.Sprintln("error: el salario ingresado no alcanza el mínimo imponible")
}

func checkImpuesto(salary int) error {
  if salary < 150000 {
    err := &errorImpuesto{}
    return err
  }

  return nil
}
