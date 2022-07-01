package main

import "fmt"

func main() {
  //Ejercicio 1 - Imprime tu nombre
  var name string = "Vanessa Sotomayor Ampudia"
  var address string = "Barrio Caney, Torre 9 Apto 531"

  fmt.Println(name, address)

  //Ejercicio 2 - Clima
  var temperatura float64 = 29
  var humedad float64 = 72
  var presionAtmosferica float64 = 1013

  fmt.Println(temperatura, humedad, presionAtmosferica)

  //Ejercicio 3 - Declaracion de variables
  //var 1nombre string -> Incorrecta, el nombre de la variable no debe empezar con numero.
  //var apellido string -> Correcta
  //var int edad -> Incorrecta, el tipo de dato debe de ir de ultimo en la declaracion de la variable.
  //1apellido := 6 -> Incorrecta, el nombre de la variable no debe empezar con numero.
  //var licencia_de_conducir = true -> Incorrecta, el nombre de la variable debe ir en camelCase y se debe escribir el tipo de dato en la declaracion de variable larga.
  //var estatura de la persona int -> Incorrecta, el nombre de la variable no debe ir con espacios.
  //cantidadDeHijos := 2 -> Correcta

  //Ejercicio 4 - Tipos de datos
  //var apellido string = "Gomez" -> Correcta
  //var edad int = "35" -> Incorrecta, el valor asignado a una variable de tipo int no va entre comillas.
  //boolean := "false"; -> Incorrecta, las lineas de codigo no llevan ; al final.
  //var sueldo string = 45857.90 -> Incorrecta, una variable de tipo string no guarda valores de tipo int
  //var nombre string = "Julián" -> Correcta

}