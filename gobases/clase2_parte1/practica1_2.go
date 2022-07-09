package main

import (

  "fmt"
  "errors"

)

//Ejercicio 2 - Calcular promedio
func promedio(calificaciones ...int) (int, error) {

  var operacion int

  for _, valor := range calificaciones {

    if valor < 0 {

      return 0, errors.New("No se aceptan numeros negativos")

    }
    operacion += valor

  }
  operacion = operacion / len(calificaciones)
  return operacion, nil

}

func main() {

  resultado, err := promedio(4, 3, 5, 4, 5, 4, -1)

  if err != nil {

    fmt.Println(err)

  }

  fmt.Println(resultado)

}