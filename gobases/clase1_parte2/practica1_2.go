package main

import "fmt"

func main() {
  //Ejercicio 1- Letras de una palabra
  var palabra string = "Estructura"
 
  //Imprime la cantidad de letras que tiene la palabra.
  fmt.Println(len(palabra))

  //Imprime cada una de las letras.
  for i, letra := range palabra {

    fmt.Println(i, string(letra))

  }

  //Ejercicio 2 - Prestamo
  var edad int = 23
  var empleo bool = true
  var antiguedad int = 2
  var sueldo int = 250000

  if edad > 22 {
    fmt.Println("Cumples con la edad requerida")
  }
  if empleo == true {
    fmt.Println("Cumples con el requisito de empleo")
  }
  if antiguedad > 1 {
    fmt.Println("Cumples con el requisito de antiguedad")
  }
  if sueldo > 100000 {
    fmt.Println("No se te cobrara interes")
  } else {
    fmt.Println("No cumples con alguno de los requisitos")
  }

  //Ejercicio 3 - A que mes corresponde
  var mes int = 7

  //Imprimira el nombre del mes de acuerdo al valor que se le haya asignado a la variable.
  switch mes {
  case 1:
    fmt.Println("Enero")
  case 2:
    fmt.Println("Febrero")
  case 3:
    fmt.Println("Marzo")
  case 4:
    fmt.Println("Abril")
  case 5:
    fmt.Println("Mayo")
  case 6:
    fmt.Println("Junio")
  case 7:
    fmt.Println("Julio")
  case 8:
    fmt.Println("Agosto")
  case 9:
    fmt.Println("Septiembre")
  case 10:
    fmt.Println("Octubre")
  case 11:
    fmt.Println("Noviembre")
  case 12:
    fmt.Println("Diciembre")
  }

  //Ejercicio 4 - Que edad tiene
  var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Dario": 44, "Pedro": 30}

  //Imprime la edad de Benjamin.
  fmt.Println(employees["Benjamin"])

  //Imprime los empleados mayores de 21 anos.
  for key, element := range employees {
    if element > 21 {

      fmt.Println(key, element)

    }
  }

  //Agrega un empleado nuevo.
  employees["Federico"] = 25
  fmt.Println(employees)
 
  //Elimina a Pedro del mapa.
  delete(employees, "Pedro")
  fmt.Println(employees)
}