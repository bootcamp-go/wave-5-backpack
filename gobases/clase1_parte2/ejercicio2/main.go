package main

import (
  "fmt"
)

func checkCondiciones(edad uint, empleado bool, antiguedad uint) {
  if edad < 22 {
    fmt.Println("Debe ser mayor de 22")
    return
  }

  if !empleado {
    fmt.Println("Debe estar empleado")
    return
  }

  if antiguedad < 1 {
    fmt.Println("Debe tener antiguedad de al menos 1 anio")
    return
  }

  fmt.Println("Cumple todas las condiciones")
}

func main() {
  edad, empleado, antiguedad := 22, true, 1
  checkCondiciones(uint(edad),empleado,uint(antiguedad))

  edad2, empleado2, antiguedad2 := 21, true, 1
  checkCondiciones(uint(edad2),empleado2,uint(antiguedad2))
  
  edad3, empleado3, antiguedad3 := 22, false, 1
  checkCondiciones(uint(edad3),empleado3,uint(antiguedad3))

  edad4, empleado4, antiguedad4 := 22, true, 0
  checkCondiciones(uint(edad4),empleado4,uint(antiguedad4))
}
